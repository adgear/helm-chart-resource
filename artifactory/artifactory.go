package artifactory

import (
	"encoding/json"
	"errors"

	"github.com/adgear/helm-chart-resource/utils"
	resty "gopkg.in/resty.v1"
)

// Artifactory interface
type Artifactory interface {
	UploadArtifactoryChart(source utils.Source, params utils.Params, version map[string]string, tmpdir string) error
}

type artifactory struct{}

// NewArtifactory instance
func NewArtifactory() Artifactory {
	return artifactory{}
}

// UploadArtifactoryChart takes the .tgz and POST to artifactory
func (a artifactory) UploadArtifactoryChart(source utils.Source, params utils.Params, version map[string]string, tmpdir string) error {
	for _, repo := range source.Repos {
		if repo.Name == source.RepositoryName {
			resp, err := resty.R().
				SetBasicAuth(repo.Username, repo.Password).
				SetFormData(map[string]string{
					"username": repo.Username,
				}).
				Post(params.APIURL + "/security/token")

			if err != nil {
				return err
			}

			if resp.StatusCode() > 201 {
				return errors.New(string(resp.Body()))
			}

			var respMap map[string]interface{}
			err = json.Unmarshal(resp.Body(), &respMap)
			if err != nil {
				return err
			}

			resp, err = resty.R().
				SetFile(source.ChartName+"-"+version["ref"]+".tgz", tmpdir+"/"+source.ChartName+"-"+version["ref"]+".tgz").
				SetAuthToken(respMap["access_token"].(string)).
				Put(repo.URL + "/" + source.ChartName + "-" + version["ref"] + ".tgz")

			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}
