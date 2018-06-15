package actions_test

import (
	"testing"

	"github.com/adgear/helm-chart-resource/actions"
	"github.com/adgear/helm-chart-resource/utils"
	"github.com/stretchr/testify/assert"
)

func TestArtifactory(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
			Repos:          []utils.Repo{},
		},
		Params: utils.Params{
			APIURL: "http://someurl/",
			Path:   "ci/assets/dummy-chart",
			Type:   "artifactory",
		},
		Version: map[string]string{
			"ref": "0.4.0",
		},
	}
	destination := "/potatoes/"
	output := "{version: {ref: \"0.4.0\"}}"

	helmMock.EXPECT().InstallHelmRepo([]utils.Repo{}).Return(nil).Times(1)
	helmMock.EXPECT().BuildHelmChart(destination, input.Params.Path).Return(nil).Times(1)
	helmMock.EXPECT().PackageHelmChart(destination, input.Params.Path, "/tmp").Return(nil).Times(1)
	artifactoryMock.EXPECT().UploadArtifactoryChart(input.Source, input.Params, input.Version, "/tmp").Return(nil)

	cr, _ := actions.NewOutResource(helmMock, artifactoryMock)

	o, err := cr.Execute(input, destination, "/tmp")

	assert.NoError(t, err)
	assert.Equal(t, output, o)
}

func TestNotArtifactory(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
			Repos:          []utils.Repo{},
		},
		Params: utils.Params{
			APIURL: "http://someurl/",
			Path:   "ci/assets/dummy-chart",
			Type:   "s3",
		},
		Version: map[string]string{
			"ref": "0.4.0",
		},
	}
	destination := "/potatoes/"

	helmMock.EXPECT().InstallHelmRepo([]utils.Repo{}).Return(nil).Times(1)
	helmMock.EXPECT().BuildHelmChart(destination, input.Params.Path).Return(nil).Times(1)
	helmMock.EXPECT().PackageHelmChart(destination, input.Params.Path, "/tmp").Return(nil).Times(1)

	cr, _ := actions.NewOutResource(helmMock, artifactoryMock)

	_, err := cr.Execute(input, destination, "/tmp")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "unsupported chart repository")
}
