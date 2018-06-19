package actions_test

import (
	"errors"
	"os"
	"testing"

	"github.com/adgear/helm-chart-resource/actions"
	"github.com/adgear/helm-chart-resource/utils"
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/adgear/helm-chart-resource/mocks"
	"github.com/golang/mock/gomock"
)

var (
	helmMock          *mocks.MockHelm
	artifactoryMock   *mocks.MockArtifactory
	checkResourceMock *mocks.MockCheckResource
)

func setup(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	helmMock = mocks.NewMockHelm(mockCtrl)
	artifactoryMock = mocks.NewMockArtifactory(mockCtrl)
	checkResourceMock = mocks.NewMockCheckResource(mockCtrl)
}

func TestNoChartName(t *testing.T) {
	setup(t)

	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	input := utils.Input{
		Source: utils.Source{},
	}

	cr, _ := actions.NewCheckResource(helmMock)

	assert.PanicsWithValue(t, "os.Exit called", func() { cr.Execute(input.Source) }, "os.Exit was not called")
}

// func TestWrongInput(t *testing.T) {
// 	// So far it's always valid...
// 	setup(t)

// 	input := utils.Input{
// 		Source: utils.Source{
// 			ChartName:      "etcd",
// 			RepositoryName: "incubator",
// 		},
// 	}
// 	searchResults := 6

// 	repo := "incubator/etcd"

// 	helmMock.EXPECT().InstallHelmRepo(nil).Return(nil).Times(1)
// 	helmMock.EXPECT().Search(repo).Return(searchResults, nil).Times(1)

// 	cr, _ := actions.NewCheckResource(helmMock)

// 	_, err := cr.Execute(input.Source)

// 	assert.Error(t, err)
// }

func TestPublicRepoFound(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
		},
	}
	searchResults := `NAME          	CHART VERSION	APP VERSION	DESCRIPTION                                       
	incubator/etcd	0.4.0        	2.2.5      	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.9        	2.2.5      	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.8        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.7        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.6        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.5        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.4        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.3        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.3.2        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.2.2        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.2.1        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.2.0        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.1.3        	           	Distributed reliable key-value store for the mo...
	incubator/etcd	0.1.2        	           	Etcd Helm chart for Kubernetes.                   
	incubator/etcd	0.1.1        	           	Etcd Helm chart for Kubernetes.                   
	incubator/etcd	0.1.0        	           	Etcd Helm chart for Kubernetes.                   `
	repo := "incubator/etcd"
	output := "[{\"ref\":\"0.4.0\"}]"

	helmMock.EXPECT().InstallHelmRepo(nil).Return(nil).Times(1)
	helmMock.EXPECT().Search(repo).Return(searchResults, nil).Times(1)

	cr, _ := actions.NewCheckResource(helmMock)

	o, err := cr.Execute(input.Source)

	assert.Equal(t, output, o, "they should be equal")

	assert.NoError(t, err)
}

func TestPublicRepoNotFound(t *testing.T) {
	setup(t)

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "etcd",
			RepositoryName: "incubator",
		},
	}
	searchResults := `No results found
	`
	repo := "incubator/etcd"

	helmMock.EXPECT().InstallHelmRepo(nil).Return(nil).Times(1)
	helmMock.EXPECT().Search(repo).Return(searchResults, nil).Times(1)

	cr, _ := actions.NewCheckResource(helmMock)

	o, err := cr.Execute(input.Source)

	assert.Empty(t, o)

	assert.Error(t, err)
}

func TestPrivateRepoFound(t *testing.T) {
	setup(t)

	var repos []utils.Repo

	repos = append(repos, utils.Repo{
		Name:     "adgear-helm",
		URL:      "https://someurl.com",
		Username: "potatoes",
		Password: "tomatoes",
	})

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "netbox",
			RepositoryName: "adgear-helm",
			Repos:          repos,
		},
	}
	searchResults := `NAME              	CHART VERSION	APP VERSION	DESCRIPTION                         
	adgear-helm/netbox	0.1.3        	1.0        	A Helm chart for digitalocean netbox
	adgear-helm/netbox	0.1.2        	1.0        	A Helm chart for digitalocean netbox
	adgear-helm/netbox	0.1.1        	1.0        	A Helm chart for digitalocean netbox
	adgear-helm/netbox	0.1.0        	1.0        	A Helm chart for digitalocean netbox
	`

	repo := "adgear-helm/netbox"
	output := "[{\"ref\":\"0.1.3\"}]"

	helmMock.EXPECT().InstallHelmRepo(repos).Return(nil).Times(1)
	helmMock.EXPECT().Search(repo).Return(searchResults, nil).Times(1)

	cr, _ := actions.NewCheckResource(helmMock)

	o, err := cr.Execute(input.Source)

	assert.Equal(t, output, o)

	assert.NoError(t, err)
}

func TestBadRepo(t *testing.T) {
	setup(t)

	var repos []utils.Repo

	repos = append(repos, utils.Repo{
		Name:     "adgear-helm",
		URL:      "https://someurl.com",
		Username: "potatoes",
		Password: "tomatoes",
	})

	input := utils.Input{
		Source: utils.Source{
			ChartName:      "netbox",
			RepositoryName: "adgear-helm",
			Repos:          repos,
		},
	}
	searchResults := `NAME              	CHART VERSION	APP VERSION	DESCRIPTION                         
	adgear-helm/netbox	0.1.3        	1.0        	A Helm chart for digitalocean netbox
	adgear-helm/netbox	0.1.2        	1.0        	A Helm chart for digitalocean netbox
	adgear-helm/netbox	0.1.1        	1.0        	A Helm chart for digitalocean netbox
	adgear-helm/netbox	0.1.0        	1.0        	A Helm chart for digitalocean netbox
	`

	repo := "adgear-helm/netbox"
	// output := "[{\"ref\":\"0.1.3\"}]"

	helmMock.EXPECT().InstallHelmRepo(repos).Return(errors.New("DIEEE")).Times(1)
	helmMock.EXPECT().Search(repo).Return(searchResults, nil).Times(1)

	cr, _ := actions.NewCheckResource(helmMock)

	_, err := cr.Execute(input.Source)

	assert.Error(t, err)
}
