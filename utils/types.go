package utils

// Source structure
type Source struct {
	ChartName      string `json:"chart_name"`
	RepositoryName string `json:"repository_name"`
	Repos          []Repo
}

// Check structure
type Check struct {
	Source Source
}

// Input struct
type Input struct {
	Source  Source
	Version map[string]string
	Params  Params
}

// Params struct
type Params struct {
	Type   string
	APIURL string `json:"api_url"`
	Path   string
}

// MetadataItem ref metadata
type MetadataItem struct {
	Name  string
	Value string
}

// Repo struct
type Repo struct {
	Name     string
	URL      string
	Username string
	Password string
}
