package actions

//go:generate mockgen -destination=../mocks/mock_in.go -package=mocks github.com/adgear/helm-chart-resource InResource

import (
	"fmt"
	"io/ioutil"

	"github.com/adgear/helm-chart-resource/helm"
	"github.com/adgear/helm-chart-resource/utils"
)

// InResource interface
type InResource interface {
	Execute(source utils.Source, destination string, version string) (string, error)
}

type inResource struct {
	helm helm.Helm
}

// NewInResource returns a new instance
func NewInResource(helm helm.Helm) (InResource, error) {
	return inResource{
		helm: helm,
	}, nil
}

// Execute the in resource
func (ir inResource) Execute(source utils.Source, destination string, version string) (string, error) {
	utils.InstallHelmRepo(source.Repos)

	var output string

	err := ioutil.WriteFile(destination+"/.version", []byte(version), 0644)
	if err != nil {
		return "", err
	}

	fmt.Println()

	output = "{version: {ref: \"" + version + "\"}}"

	return output, nil
}
