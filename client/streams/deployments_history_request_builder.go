package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsHistoryRequestBuilder builds and executes requests for operations under \streams\deployments\history
type DeploymentsHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.streams.deployments.history.item collection
func (m *DeploymentsHistoryRequestBuilder) ByName(name string)(*DeploymentsHistoryWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewDeploymentsHistoryWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsHistoryRequestBuilderInternal instantiates a new HistoryRequestBuilder and sets the default values.
func NewDeploymentsHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsHistoryRequestBuilder) {
    m := &DeploymentsHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/history", pathParameters),
    }
    return m
}
// NewDeploymentsHistoryRequestBuilder instantiates a new HistoryRequestBuilder and sets the default values.
func NewDeploymentsHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
