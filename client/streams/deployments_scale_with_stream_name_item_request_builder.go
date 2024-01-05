package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsScaleWithStreamNameItemRequestBuilder builds and executes requests for operations under \streams\deployments\scale\{streamName}
type DeploymentsScaleWithStreamNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByAppName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.scale.item.item collection
func (m *DeploymentsScaleWithStreamNameItemRequestBuilder) ByAppName(appName string)(*DeploymentsScaleItemWithAppNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if appName != "" {
        urlTplParams["appName"] = appName
    }
    return NewDeploymentsScaleItemWithAppNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsScaleWithStreamNameItemRequestBuilderInternal instantiates a new WithStreamNameItemRequestBuilder and sets the default values.
func NewDeploymentsScaleWithStreamNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleWithStreamNameItemRequestBuilder) {
    m := &DeploymentsScaleWithStreamNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/scale/{streamName}", pathParameters),
    }
    return m
}
// NewDeploymentsScaleWithStreamNameItemRequestBuilder instantiates a new WithStreamNameItemRequestBuilder and sets the default values.
func NewDeploymentsScaleWithStreamNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleWithStreamNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsScaleWithStreamNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
