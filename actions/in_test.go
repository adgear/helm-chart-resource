package actions_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/adgear/helm-chart-resource/actions"
	"github.com/adgear/helm-chart-resource/utils"
	"github.com/stretchr/testify/assert"
)

func TestCannotWriteFile(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
			Repos:          []utils.Repo{},
		},
		Version: map[string]string{
			"ref": "0.4.0",
		},
	}
	destination := "/jaldfsji/"

	cr, _ := actions.NewInResource(helmMock)

	helmMock.EXPECT().InstallHelmRepo([]utils.Repo{}).Return(nil).Times(1)
	helmMock.EXPECT().RepoUpdate().Times(1)
	helmMock.EXPECT().BuildHelmChart(destination, input.Params.Path).Return(nil).Times(1)
	helmMock.EXPECT().PackageHelmChart(destination, input.Params.Path, "/tmp").Return(nil).Times(1)

	_, err := cr.Execute(input.Source, destination, input.Version["ref"])
	assert.Error(t, err)
}

func TestWriteToFile(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
			Repos:          []utils.Repo{},
		},
		Version: map[string]string{
			"ref": "0.4.0",
		},
	}
	destination, err := ioutil.TempDir("/tmp", "gomock")
	if err != nil {
		panic(err)
	}
	output := "{version: {ref: \"0.4.0\"}}"
	cr, _ := actions.NewInResource(helmMock)

	helmMock.EXPECT().InstallHelmRepo([]utils.Repo{}).Return(nil).Times(1)
	helmMock.EXPECT().RepoUpdate().Times(1)
	helmMock.EXPECT().BuildHelmChart(destination, input.Params.Path).Return(nil).Times(1)
	helmMock.EXPECT().PackageHelmChart(destination, input.Params.Path, "/tmp").Return(nil).Times(1)

	o, err := cr.Execute(input.Source, destination, input.Version["ref"])

	assert.Equal(t, output, o)
	assert.NoError(t, err)
	assert.FileExists(t, destination+"/.version")
	defer os.RemoveAll(destination)
}
