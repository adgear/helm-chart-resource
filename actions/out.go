package actions

//go:generate mockgen -destination=../mocks/mock_out.go -package=mocks github.com/adgear/helm-chart-resource OutResource

import (
	"errors"

	"github.com/adgear/helm-chart-resource/artifactory"
	"github.com/adgear/helm-chart-resource/helm"
	"github.com/adgear/helm-chart-resource/utils"
)

// OutResource interface
type OutResource interface {
	Execute(input utils.Input, sourcePath string, tmpdir string) (string, error)
}

type outResource struct {
	helm        helm.Helm
	artifactory artifactory.Artifactory
}

// NewOutResource returns a new instance
func NewOutResource(helm helm.Helm, artifactory artifactory.Artifactory) (OutResource, error) {
	return outResource{
		helm:        helm,
		artifactory: artifactory,
	}, nil
}

// Execute the out resource
func (or outResource) Execute(input utils.Input, sourcePath string, tmpdir string) (string, error) {
	version, err := or.helm.ExtractChartVersion(sourcePath, input.Params.Path)
	if err != nil {
		return "", err
	}
	or.helm.InstallHelmRepo(input.Source.Repos)
	or.helm.BuildHelmChart(sourcePath, input.Params.Path)
	or.helm.PackageHelmChart(sourcePath, input.Params.Path, tmpdir)

	if input.Params.Type == "artifactory" {
		err := or.artifactory.UploadArtifactoryChart(input.Source, input.Params, version, tmpdir)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("unsupported chart repository")
	}

	var output string

	output = "{\"version\": {\"ref\": \"" + version + "\"}}"

	return output, nil
}
