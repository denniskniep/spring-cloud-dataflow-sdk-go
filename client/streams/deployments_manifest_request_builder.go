package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsManifestRequestBuilder builds and executes requests for operations under \streams\deployments\manifest
type DeploymentsManifestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.streams.deployments.manifest.item collection
func (m *DeploymentsManifestRequestBuilder) ByName(name string)(*DeploymentsManifestWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewDeploymentsManifestWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsManifestRequestBuilderInternal instantiates a new ManifestRequestBuilder and sets the default values.
func NewDeploymentsManifestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsManifestRequestBuilder) {
    m := &DeploymentsManifestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/manifest", pathParameters),
    }
    return m
}
// NewDeploymentsManifestRequestBuilder instantiates a new ManifestRequestBuilder and sets the default values.
func NewDeploymentsManifestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsManifestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsManifestRequestBuilderInternal(urlParams, requestAdapter)
}
