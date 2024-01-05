package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsScaleRequestBuilder builds and executes requests for operations under \streams\deployments\scale
type DeploymentsScaleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByStreamName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.streams.deployments.scale.item collection
func (m *DeploymentsScaleRequestBuilder) ByStreamName(streamName string)(*DeploymentsScaleWithStreamNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if streamName != "" {
        urlTplParams["streamName"] = streamName
    }
    return NewDeploymentsScaleWithStreamNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsScaleRequestBuilderInternal instantiates a new ScaleRequestBuilder and sets the default values.
func NewDeploymentsScaleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleRequestBuilder) {
    m := &DeploymentsScaleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/scale", pathParameters),
    }
    return m
}
// NewDeploymentsScaleRequestBuilder instantiates a new ScaleRequestBuilder and sets the default values.
func NewDeploymentsScaleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsScaleRequestBuilderInternal(urlParams, requestAdapter)
}
