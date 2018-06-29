package actions

//go:generate mockgen -destination=../mocks/mock_check.go -package=mocks github.com/adgear/helm-chart-resource CheckResource

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/adgear/helm-chart-resource/helm"
	"github.com/adgear/helm-chart-resource/utils"
)

// CheckResource interface
type CheckResource interface {
	Execute(source utils.Source) (string, error)
}

type checkResource struct {
	helm helm.Helm
}

// NewCheckResource returns a new instance
func NewCheckResource(helm helm.Helm) (CheckResource, error) {
	return checkResource{
		helm: helm,
	}, nil
}

// Execute the check resource
func (cr checkResource) Execute(source utils.Source) (string, error) {
	if source.ChartName == "" {
		log.Fatal("Chart Name can't be empty.")
		os.Exit(1)
	}

	var refs []map[string]string
	err := cr.helm.InstallHelmRepo(source.Repos)

	if err != nil {
		return "", err
	}

	cr.helm.RepoUpdate()

	cmdOutput, err := cr.helm.Search(source.RepositoryName + "/" + source.ChartName)
	if cmdOutput == `No results found
	` {
		return "", errors.New(cmdOutput)
	}

	lines := utils.FilterSearch(cmdOutput, source.RepositoryName+"/"+source.ChartName, true)

	for _, line := range lines {
		info, _ := utils.GetLineInfo(line)
		refs = append(refs, map[string]string{"ref": info["chart_version"].(string)})
	}

	output, _ := json.Marshal(refs)

	return string(output), nil
}
