package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsUpdateRequestBuilder builds and executes requests for operations under \streams\deployments\update
type DeploymentsUpdateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.update.item collection
func (m *DeploymentsUpdateRequestBuilder) ByName(name string)(*DeploymentsUpdateWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewDeploymentsUpdateWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsUpdateRequestBuilderInternal instantiates a new UpdateRequestBuilder and sets the default values.
func NewDeploymentsUpdateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsUpdateRequestBuilder) {
    m := &DeploymentsUpdateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/update", pathParameters),
    }
    return m
}
// NewDeploymentsUpdateRequestBuilder instantiates a new UpdateRequestBuilder and sets the default values.
func NewDeploymentsUpdateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsUpdateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsUpdateRequestBuilderInternal(urlParams, requestAdapter)
}
