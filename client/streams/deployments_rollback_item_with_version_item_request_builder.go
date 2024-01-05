package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsRollbackItemWithVersionItemRequestBuilder builds and executes requests for operations under \streams\deployments\rollback\{name}\{version}
type DeploymentsRollbackItemWithVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeploymentsRollbackItemWithVersionItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeploymentsRollbackItemWithVersionItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeploymentsRollbackItemWithVersionItemRequestBuilderInternal instantiates a new WithVersionItemRequestBuilder and sets the default values.
func NewDeploymentsRollbackItemWithVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRollbackItemWithVersionItemRequestBuilder) {
    m := &DeploymentsRollbackItemWithVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/rollback/{name}/{version}", pathParameters),
    }
    return m
}
// NewDeploymentsRollbackItemWithVersionItemRequestBuilder instantiates a new WithVersionItemRequestBuilder and sets the default values.
func NewDeploymentsRollbackItemWithVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRollbackItemWithVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsRollbackItemWithVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsRollbackItemWithVersionItemRequestBuilder) Post(ctx context.Context, requestConfiguration *DeploymentsRollbackItemWithVersionItemRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *DeploymentsRollbackItemWithVersionItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeploymentsRollbackItemWithVersionItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeploymentsRollbackItemWithVersionItemRequestBuilder) WithUrl(rawUrl string)(*DeploymentsRollbackItemWithVersionItemRequestBuilder) {
    return NewDeploymentsRollbackItemWithVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
