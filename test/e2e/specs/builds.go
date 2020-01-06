package specs

import (
	"context"
	"encoding/json"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	buildv1 "github.com/openshift/api/build/v1"
	buildclientmanual "github.com/openshift/origin/pkg/build/client/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/test/sanity"
	"github.com/openshift/openshift-azure/test/util/log"
)

var _ = Describe("Openshift on Azure end user e2e tests [EndUser][EveryPR]", func() {

	// It("should be able to list builds [EndUser][Builds]", func() {
	// 	ctx := context.Background()
	// 	namespace, err := sanity.Checker.CreateProject(ctx)

	// 	Expect(err).To(BeNil())
	// 	defer func() {
	// 		By("deleting test project")
	// 		_ = sanity.Checker.DeleteProject(ctx, namespace)
	// 	}()

	// 	builds, buildErrs := sanity.Checker.Client.EndUser.BuildV1.Builds(fmt.Sprintf("%v", namespace)).List(metav1.ListOptions{})

	// 	By("validating test build")
	// 	// errs = sanity.Checker.ValidateTestApp(ctx, namespace)
	// 	Expect(builds.Items).To(BeEmpty())
	// 	Expect(buildErrs).To(BeEmpty())
	// })

	It("should be able to create builds [EndUser][Builds]", func() {
		ctx := context.Background()
		namespace, err := sanity.Checker.CreateProject(ctx)
		name := namespace + "-build"
		testlogger := log.GetTestLogger()

		// Expect(err).To(BeNil())
		// defer func() {
		// 	By("deleting project")
		// 	_ = sanity.Checker.DeleteProject(ctx, namespace)
		// }()

		By("creating a binary BuildConfig")
		bc := &buildv1.BuildConfig{
			ObjectMeta: metav1.ObjectMeta{
				Name:   name,
				Labels: map[string]string{"build": name},
			},
			Spec: buildv1.BuildConfigSpec{
				CommonSpec: buildv1.CommonSpec{
					Source: buildv1.BuildSource{
						Type: buildv1.BuildSourceBinary,
					},
					Strategy: buildv1.BuildStrategy{
						DockerStrategy: &buildv1.DockerBuildStrategy{},
					},
				},
			},
		}
		build, buildErr := sanity.Checker.Client.EndUser.BuildV1.BuildConfigs(namespace).Create(bc)
		Expect(buildErr).To(BeNil())

		By("validating starting a build for a binary BuildConfig")

		buildClient := sanity.Checker.Client.EndUser.BuildV1.RESTClient()

		buildRequestOptions := &buildv1.BinaryBuildRequestOptions{
			ObjectMeta: metav1.ObjectMeta{
				Name:      build.Name,
				Namespace: build.Namespace,
			},
			AsFile: "Dockerfile",
		}
		bytes, err := json.Marshal(buildRequestOptions)
		testlogger.Error(string(bytes))

		instantiateClient := buildclientmanual.NewBuildInstantiateBinaryClient(buildClient, namespace)

		r := strings.NewReader("FROM scratch")
		_, err = instantiateClient.InstantiateBinary("foo", buildRequestOptions, r)

		// errs = sanity.Checker.ValidateTestApp(ctx, namespace)
		Expect(err).To(BeNil())
		//Expect(json.Marshal(b)).To(BeNil())
	})
})
