package streams

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsManifestWithNameItemRequestBuilder builds and executes requests for operations under \streams\deployments\manifest\{name}
type DeploymentsManifestWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByVersion gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.manifest.item.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *DeploymentsManifestWithNameItemRequestBuilder) ByVersion(version string)(*DeploymentsManifestItemWithVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if version != "" {
        urlTplParams["version"] = version
    }
    return NewDeploymentsManifestItemWithVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByVersionInteger gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.manifest.item.item collection
func (m *DeploymentsManifestWithNameItemRequestBuilder) ByVersionInteger(version int32)(*DeploymentsManifestItemWithVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["version"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(version), 10)
    return NewDeploymentsManifestItemWithVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsManifestWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDeploymentsManifestWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsManifestWithNameItemRequestBuilder) {
    m := &DeploymentsManifestWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/manifest/{name}", pathParameters),
    }
    return m
}
// NewDeploymentsManifestWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDeploymentsManifestWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsManifestWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsManifestWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
