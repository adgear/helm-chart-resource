package helm

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/adgear/helm-chart-resource/utils"
	"gopkg.in/yaml.v2"
)

// Helm interface
type Helm interface {
	DepUpdate(path string) (string, error)
	Search(repo string) (string, error)
	InstallHelmRepo(repos []utils.Repo) error
	BuildHelmChart(source string, path string) error
	PackageHelmChart(source string, path string, tmpdir string) error
	ExtractChartVersion(source string, path string) (string, error)
}

type helm struct{}

// NewHelm instance
func NewHelm() Helm {
	return helm{}
}

func (h helm) DepUpdate(path string) (string, error) {
	return "", nil
}

func (h helm) Search(repo string) (string, error) {
	var (
		cmdOut []byte
		err    error
	)

	cmdName := "helm"
	cmdArgs := []string{
		"search",
		repo,
		"-l",
	}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	return string(cmdOut), nil
}

// InstallHelmRepo setup helm repo for install locally
func (h helm) InstallHelmRepo(repos []utils.Repo) error {
	cmdName := "helm"
	var (
		cmdArgs []string
	)

	for _, repo := range repos {
		cmdArgs = []string{
			"repo",
			"add",
			repo.Name,
			repo.URL,
		}

		if repo.Username != "" {
			cmdArgs = append(cmdArgs, "--username")
			cmdArgs = append(cmdArgs, repo.Username)
		}

		if repo.Password != "" {
			cmdArgs = append(cmdArgs, "--password")
			cmdArgs = append(cmdArgs, repo.Password)
		}

		if _, err := exec.Command(cmdName, cmdArgs...).CombinedOutput(); err != nil {
			return err
		}
	}

	return nil
}

// BuildHelmChart builds the chart from requirements.lock
func (h helm) BuildHelmChart(source string, path string) error {
	cmdName := "helm"
	cmdArgs := []string{
		"dep",
		"build",
		source + "/" + path,
	}

	var (
		err error
	)

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		return err
	}

	return nil
}

// PackageHelmChart packages the chart to .tgz file
func (h helm) PackageHelmChart(source string, path string, tmpdir string) error {
	cmdName := "helm"
	cmdArgs := []string{
		"package",
		source + "/" + path,
		"-d",
		tmpdir + "/",
	}

	var (
		err error
	)

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		return err
	}

	return nil
}

func (h helm) ExtractChartVersion(source string, path string) (string, error) {
	var chartYaml map[string]string
	chartYamlByte, err := ioutil.ReadFile(source + "/" + path + "/Chart.yaml")

	if err != nil {
		return "null", err
	}

	err = yaml.Unmarshal(chartYamlByte, &chartYaml)

	if err != nil {
		return "null", err
	}

	return chartYaml["version"], nil
}
