package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adgear/helm-chart-resource/actions"
	"github.com/adgear/helm-chart-resource/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	artifactoryLib "github.com/adgear/helm-chart-resource/artifactory"
	helmLib "github.com/adgear/helm-chart-resource/helm"
)

var version = "local"

var (
	rootCmd = &cobra.Command{}

	// Flags
	action      string
	destination string
	source      string
	tmpdir      string
	verbose     bool
	showVersion bool
	help        bool

	helm        helmLib.Helm
	artifactory artifactoryLib.Artifactory
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)

	rootCmd.PersistentFlags().StringVarP(&action, "action", "a", "check", "Concourse resource action, defaults to check. Can be check, in or out.")
	rootCmd.PersistentFlags().StringVarP(&destination, "destination", "d", "./", "Destination of input resource")
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "./", "Source of output resource")
	rootCmd.PersistentFlags().StringVarP(&tmpdir, "tmpdir", "t", "/tmp/", "Temporary directory for packaging helm chart.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode.")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "V", false, "Prints the version and exit.")
	rootCmd.PersistentFlags().BoolVarP(&help, "help", "h", false, "Show help.")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	if help {
		rootCmd.Usage()
		os.Exit(0)
	}

	if showVersion {
		log.Info("Version: " + version)
		os.Exit(0)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var inputMap utils.Input

	err = json.Unmarshal(input, &inputMap)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	helm = helmLib.NewHelm()
	artifactory = artifactoryLib.NewArtifactory()

	switch action {
	case "check":
		checkResource, err := actions.NewCheckResource(helm)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		output, err := checkResource.Execute(inputMap.Source)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println(output)
		os.Exit(0)
	case "in":
		inResource, err := actions.NewInResource(helm)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		output, err := inResource.Execute(inputMap.Source, destination, inputMap.Version["ref"])
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println(output)
		os.Exit(0)
	case "out":
		outResource, err := actions.NewOutResource(helm, artifactory)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		output, err := outResource.Execute(inputMap, source, tmpdir)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println(output)
		os.Exit(0)
	default:
		fmt.Println("Nope")
		os.Exit(1)
	}
}
