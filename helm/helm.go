package helm

import (
	"log"
	"os"
	"os/exec"
)

// Helm interface
type Helm interface {
	DepUpdate(path string) (string, error)
	Search(repo string) (string, error)
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
