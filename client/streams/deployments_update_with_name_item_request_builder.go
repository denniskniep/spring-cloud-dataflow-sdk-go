package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f "dataflow/client/models"
)

// DeploymentsUpdateWithNameItemRequestBuilder builds and executes requests for operations under \streams\deployments\update\{name}
type DeploymentsUpdateWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeploymentsUpdateWithNameItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeploymentsUpdateWithNameItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeploymentsUpdateWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDeploymentsUpdateWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsUpdateWithNameItemRequestBuilder) {
    m := &DeploymentsUpdateWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/update/{name}", pathParameters),
    }
    return m
}
// NewDeploymentsUpdateWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDeploymentsUpdateWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsUpdateWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsUpdateWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsUpdateWithNameItemRequestBuilder) Post(ctx context.Context, body i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f.UpdateStreamRequestable, requestConfiguration *DeploymentsUpdateWithNameItemRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
func (m *DeploymentsUpdateWithNameItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f.UpdateStreamRequestable, requestConfiguration *DeploymentsUpdateWithNameItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeploymentsUpdateWithNameItemRequestBuilder) WithUrl(rawUrl string)(*DeploymentsUpdateWithNameItemRequestBuilder) {
    return NewDeploymentsUpdateWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
