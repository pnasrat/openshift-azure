package kubeclient

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/util/retry"

	"github.com/openshift/openshift-azure/pkg/util/wait"
)

func (u *Kubeclientset) DrainAndDeleteWorker(ctx context.Context, hostname string) error {
	err := u.setUnschedulable(hostname, true)
	switch {
	case err == nil:
	case kerrors.IsNotFound(err):
		u.Log.Info("drain: node not found, skipping")
		return nil
	default:
		return err
	}

	err = u.deletePods(ctx, hostname)
	if err != nil {
		return err
	}

	return u.Client.CoreV1().Nodes().Delete(hostname, &metav1.DeleteOptions{})
}

func (u *Kubeclientset) DeleteMaster(hostname string) error {
	err := u.Client.CoreV1().Nodes().Delete(hostname, &metav1.DeleteOptions{})
	if kerrors.IsNotFound(err) {
		err = nil
	}
	return err
}

func (u *Kubeclientset) setUnschedulable(hostname string, unschedulable bool) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		node, err := u.Client.CoreV1().Nodes().Get(hostname, metav1.GetOptions{})
		if err != nil {
			return err
		}

		node.Spec.Unschedulable = unschedulable
		_, err = u.Client.CoreV1().Nodes().Update(node)
		return err
	})
}

func getControllerRef(pod *v1.Pod) *metav1.OwnerReference {
	for _, ref := range pod.OwnerReferences {
		if ref.Controller != nil && *ref.Controller {
			return &ref
		}
	}
	return nil
}

func max(i, j time.Duration) time.Duration {
	if i > j {
		return i
	}
	return j
}

func (u *Kubeclientset) deletePods(ctx context.Context, hostname string) error {
	podList, err := u.Client.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{
		FieldSelector: fields.SelectorFromSet(fields.Set{"spec.nodeName": hostname}).String(),
	})
	if err != nil {
		return err
	}

	pods := map[*v1.Pod]struct{}{}
	duration := time.Duration(0)
	for i, pod := range podList.Items {
		if _, isMirror := pod.Annotations[v1.MirrorPodAnnotationKey]; isMirror {
			continue
		}

		controllerRef := getControllerRef(&pod)
		if controllerRef != nil && controllerRef.Kind == "DaemonSet" {
			continue
		}

		err = u.Client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{})
		switch {
		case err == nil:
			d := 30 * time.Second
			if pod.Spec.TerminationGracePeriodSeconds != nil {
				d = 3 * time.Duration(*pod.Spec.TerminationGracePeriodSeconds+2) * time.Second
			}
			duration = max(duration, d)

		case apierrors.IsNotFound(err):
		default:
			// TODO: handle 429

			return err
		}

		pods[&podList.Items[i]] = struct{}{}
	}

	t := time.NewTimer(duration)
	defer t.Stop()
	return wait.PollImmediateUntil(time.Second, func() (bool, error) {
		for pod := range pods {
			p, err := u.Client.CoreV1().Pods(pod.Namespace).Get(pod.Name, metav1.GetOptions{})
			switch {
			case apierrors.IsNotFound(err) || (p != nil && p.ObjectMeta.UID != pod.ObjectMeta.UID):
				delete(pods, pod)
			case err == nil:
			default:
				return false, err
			}
		}
		if len(pods) == 0 {
			return true, nil
		}
		select {
		case <-t.C:
			return true, nil
		default:
			return false, nil
		}
	}, ctx.Done())
}
