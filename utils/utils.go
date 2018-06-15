package utils

import (
	"bufio"
	"encoding/json"
	"os/exec"
	"strings"

	resty "gopkg.in/resty.v1"
)

// FilterSearch takes helm search output and only returns 1 or all the package lines
func FilterSearch(searchResults string, matchSelector string, first bool) []string {
	var resultsList []string
	scanner := bufio.NewScanner(strings.NewReader(searchResults))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, matchSelector) {
			resultsList = append(resultsList, line)
			if first {
				return resultsList
			}
		}
	}

	return resultsList
}

// GetLineInfo take one line from helm search and extract all the fields
func GetLineInfo(line string) (map[string]interface{}, error) {
	var info map[string]interface{}
	info = make(map[string]interface{})
	infoSlice := strings.Fields(line)

	info["name"] = infoSlice[0]
	info["chart_version"] = infoSlice[1]
	info["app_version"] = infoSlice[2]
	info["description"] = ""

	for i := 3; i < len(infoSlice); i++ {
		info["description"] = info["description"].(string) + " " + infoSlice[i]
	}

	return info, nil
}

// InstallHelmRepo setup helm repo for install locally
func InstallHelmRepo(repos []Repo) error {
	cmdName := "helm"
	var (
		cmdArgs []string
		// cmdOut  []byte
		err error
	)

	for _, repo := range repos {
		cmdArgs = []string{
			"repo",
			"add",
			repo.Name,
			repo.URL,
		}

		if repo.Username != "" {
			cmdArgs = append(cmdArgs, "--username "+repo.Username)
		}

		if repo.Password != "" {
			cmdArgs = append(cmdArgs, "--password "+repo.Password)
		}

		if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			return err
		}

		// cmdOutput := string(cmdOut)

		// fmt.Println(cmdOutput)
	}
	return nil
}

// BuildHelmChart builds the chart from requirements.lock
func BuildHelmChart(source string, path string) error {
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
func PackageHelmChart(source string, path string, tmpdir string) error {
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

// UploadArtifactoryChart takes the .tgz and POST to artifactory
func UploadArtifactoryChart(source Source, params Params, version map[string]string, tmpdir string) (string, error) {
	for _, repo := range source.Repos {
		if repo.Name == source.RepositoryName {
			resp, err := resty.R().
				SetBasicAuth(repo.Username, repo.Password).
				SetFormData(map[string]string{
					"username": repo.Username,
				}).
				Post(params.APIURL + "/security/token")

			if err != nil {
				return "", err
			}

			if resp.StatusCode() > 201 {
				panic(string(resp.Body()))
			}

			var respMap map[string]interface{}
			err = json.Unmarshal(resp.Body(), &respMap)
			if err != nil {
				return "", err
			}

			resp, err = resty.R().
				SetFile(source.ChartName+"-"+version["ref"]+".tgz", tmpdir+"/"+source.ChartName+"-"+version["ref"]+".tgz").
				SetAuthToken(respMap["access_token"].(string)).
				Put(repo.URL + "/" + source.ChartName + "-" + version["ref"] + ".tgz")

			if err != nil {
				return "", err
			}

			return string(resp.Body()), nil
		}
	}

	return "", nil
}
