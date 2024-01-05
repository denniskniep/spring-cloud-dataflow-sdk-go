package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder builds and executes requests for operations under \streams\deployments\scale\{streamName}\{appName}\instances\{count}
type DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeploymentsScaleItemItemInstancesWithCountItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeploymentsScaleItemItemInstancesWithCountItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilderInternal instantiates a new WithCountItemRequestBuilder and sets the default values.
func NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) {
    m := &DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/scale/{streamName}/{appName}/instances/{count}", pathParameters),
    }
    return m
}
// NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilder instantiates a new WithCountItemRequestBuilder and sets the default values.
func NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) Post(ctx context.Context, requestConfiguration *DeploymentsScaleItemItemInstancesWithCountItemRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
func (m *DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeploymentsScaleItemItemInstancesWithCountItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) WithUrl(rawUrl string)(*DeploymentsScaleItemItemInstancesWithCountItemRequestBuilder) {
    return NewDeploymentsScaleItemItemInstancesWithCountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
