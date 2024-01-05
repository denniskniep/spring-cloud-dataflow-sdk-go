package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsScaleItemWithAppNameItemRequestBuilder builds and executes requests for operations under \streams\deployments\scale\{streamName}\{appName}
type DeploymentsScaleItemWithAppNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewDeploymentsScaleItemWithAppNameItemRequestBuilderInternal instantiates a new WithAppNameItemRequestBuilder and sets the default values.
func NewDeploymentsScaleItemWithAppNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleItemWithAppNameItemRequestBuilder) {
    m := &DeploymentsScaleItemWithAppNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/scale/{streamName}/{appName}", pathParameters),
    }
    return m
}
// NewDeploymentsScaleItemWithAppNameItemRequestBuilder instantiates a new WithAppNameItemRequestBuilder and sets the default values.
func NewDeploymentsScaleItemWithAppNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleItemWithAppNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsScaleItemWithAppNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Instances the instances property
func (m *DeploymentsScaleItemWithAppNameItemRequestBuilder) Instances()(*DeploymentsScaleItemItemInstancesRequestBuilder) {
    return NewDeploymentsScaleItemItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
