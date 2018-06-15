package actions_test

import (
	"testing"

	"github.com/adgear/helm-chart-resource/actions"
	"github.com/adgear/helm-chart-resource/utils"
	"github.com/stretchr/testify/assert"
)

// NEED TO MOCK ARTIFACTORY
// func TestArtifactory(t *testing.T) {
// 	setup(t)

// 	input := utils.Input{
// 		Source: utils.Source{
// 			ChartName:      "etcd",
// 			RepositoryName: "incubator",
// 		},
// 		Params: utils.Params{
// 			APIURL: "http://someurl/",
// 			Path:   "somepath/",
// 			Type:   "artifactory",
// 		},
// 		Version: map[string]string{
// 			"ref": "0.4.0",
// 		},
// 	}
// 	destination := "/potatoes/"
// 	output := "{version: {ref: 0.4.0}}"

// 	cr, _ := actions.NewOutResource(helmMock)

// 	o, err := cr.Execute(input, destination, "/tmp")

// 	assert.Error(t, err)
// 	assert.Equal(t, output, o)
// }

func TestNotArtifactory(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
		},
		Params: utils.Params{
			APIURL: "http://someurl/",
			Path:   "ci/assets/dummy-chart",
			Type:   "s3",
		},
		Version: map[string]string{
			"ref": "0.4.0",
		},
	}
	destination := "/potatoes/"

	cr, _ := actions.NewOutResource(helmMock)

	_, err := cr.Execute(input, destination, "/tmp")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "unsupported")
}
