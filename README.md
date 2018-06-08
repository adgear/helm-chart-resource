# Helm chart resource

Tracks the helm releases in a [helm](https://helm.sh/) repository.

## Source Configuration

* `chart_name`: *Required.* The chart name

* `repository_name`: *Optional.* The repository where the chart resides.

* `repos`: A list of repos to check for `chart_name`

### Example

Resource configuration for incubator repository

``` yaml
resources:
- name: helm-etcd
  type: helm-chart-resource
  source:
    chart_name: etcd
    repository_name: incubator
    repos:
    - name: incubator
      url: https://kubernetes-charts-incubator.storage.googleapis.com/
```

Resource configuration for private repository

``` yaml
resources:
- name: helm-some_app
  type: helm-chart-resource
  source:
    chart_name: some_app
    repository_name: adgear-helm
    repos:
    - name: adgear-helm
      url: https://adgear-charts.example.com/
      username: ((username))
      password: ((password))
```

Triggering on new release of chart `some_app`

``` yaml
- get: helm-some_app
  trigger: true
```

## Behavior

### `check`: Check for new helm chart versions

Search for the latest version of `source.chart_name`.

### `in`: Get the latest version ref.

Output the latest `ref` to file `.version`

### `out`: Package and push to a helm repository.

Package `params.path` using the version in `{params.path}/Chart.yaml` and push it to an helm repo.

#### Parameters

* `path`: *Required.* The path of the previously `get` resource containing your helm chart code.

* `type`: *Required.* The type of helm repository to push to.
Only supports `artifactory` for now.

* `api_url`: *Optional.* The artifactory api url.

## Development

### Prerequisites

* Common sense.

### Running the tests

To be implemented.

### Contributing

TBD.