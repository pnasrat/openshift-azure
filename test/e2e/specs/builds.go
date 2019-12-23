package specs

import (
	"context"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	buildv1 "github.com/openshift/api/build/v1"
	buildclientmanual "github.com/openshift/origin/pkg/build/client/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/test/sanity"
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

		Expect(err).To(BeNil())
		// defer func() {
		// 	By("deleting project")
		// 	_ = sanity.Checker.DeleteProject(ctx, namespace)
		// }()

		// buildConfig := &buildv1.BuildConfig{}
		// buildConfig.Name = fmt.Sprintf("%v-build", namespace)

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

		// Spec: buildv1.BuildSpec{
		// 	CommonSpec: buildv1.CommonSpec{
		// 		Source: buildv1.BuildSource{

		// buildv1.BuildSourceBinary{}
		By("creating a binary BuildConfig")
		_, buildErr := sanity.Checker.Client.EndUser.BuildV1.BuildConfigs(namespace).Create(bc)
		Expect(buildErr).To(BeNil())

		By("validating starting a build for a binary BuildConfig")

		buildRequest := &buildv1.BinaryBuildRequestOptions{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			AsFile: "Dockerfile",
		}

		instantiateClient := buildclientmanual.NewBuildInstantiateBinaryClient(sanity.Checker.Client.EndUser.BuildV1.RESTClient(), namespace)

		// b, err := sanity.Checker.Client.EndUser.BuildV1.BuildConfigs(namespace).Instantiate(name, buildRequest)
		//b, err := sanity.Checker.Client.EndUser.BuildV1.
		r := strings.NewReader("FROM scratch")
		b, err := instantiateClient.InstantiateBinary(name, buildRequest, r)

		// errs = sanity.Checker.ValidateTestApp(ctx, namespace)
		Expect(err).To(BeNil())
		Expect(b).To(BeNil())
	})
})
