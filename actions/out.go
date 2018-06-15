package actions

//go:generate mockgen -destination=../mocks/mock_out.go -package=mocks github.com/adgear/helm-chart-resource OutResource

import (
	"errors"

	"github.com/adgear/helm-chart-resource/helm"
	"github.com/adgear/helm-chart-resource/utils"
)

// OutResource interface
type OutResource interface {
	Execute(input utils.Input, sourcePath string, tmpdir string) (string, error)
}

type outResource struct {
	helm helm.Helm
}

// NewOutResource returns a new instance
func NewOutResource(helm helm.Helm) (OutResource, error) {
	return outResource{
		helm: helm,
	}, nil
}

// Execute the in resource
func (or outResource) Execute(input utils.Input, sourcePath string, tmpdir string) (string, error) {
	utils.InstallHelmRepo(input.Source.Repos)

	utils.BuildHelmChart(sourcePath, input.Params.Path)

	utils.PackageHelmChart(sourcePath, input.Params.Path, tmpdir)

	if input.Params.Type == "artifactory" {
		_, err := utils.UploadArtifactoryChart(input.Source, input.Params, input.Version, tmpdir)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("unsupported")
	}

	var output string

	output = "{version: {ref: \"" + input.Version["ref"] + "\"}}"

	return output, nil
}
