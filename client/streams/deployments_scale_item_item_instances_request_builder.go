package streams

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsScaleItemItemInstancesRequestBuilder builds and executes requests for operations under \streams\deployments\scale\{streamName}\{appName}\instances
type DeploymentsScaleItemItemInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCount gets an item from the dataflow/client.streams.deployments.scale.item.item.instances.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *DeploymentsScaleItemItemInstancesRequestBuilder) ByCount(count string)(*DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if count != "" {
        urlTplParams["count"] = count
    }
    return NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCountInteger gets an item from the dataflow/client.streams.deployments.scale.item.item.instances.item collection
func (m *DeploymentsScaleItemItemInstancesRequestBuilder) ByCountInteger(count int32)(*DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["count"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(count), 10)
    return NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsScaleItemItemInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewDeploymentsScaleItemItemInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleItemItemInstancesRequestBuilder) {
    m := &DeploymentsScaleItemItemInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/scale/{streamName}/{appName}/instances", pathParameters),
    }
    return m
}
// NewDeploymentsScaleItemItemInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewDeploymentsScaleItemItemInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleItemItemInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsScaleItemItemInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
