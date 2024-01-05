package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExecutionsExternalWithExternalExecutionItemRequestBuilder builds and executes requests for operations under \tasks\executions\external\{externalExecutionId}
type ExecutionsExternalWithExternalExecutionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsExternalWithExternalExecutionItemRequestBuilderGetQueryParameters 
type ExecutionsExternalWithExternalExecutionItemRequestBuilderGetQueryParameters struct {
    // 
    Platform *string `uriparametername:"platform"`
}
// ExecutionsExternalWithExternalExecutionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsExternalWithExternalExecutionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsExternalWithExternalExecutionItemRequestBuilderGetQueryParameters
}
// NewExecutionsExternalWithExternalExecutionItemRequestBuilderInternal instantiates a new WithExternalExecutionItemRequestBuilder and sets the default values.
func NewExecutionsExternalWithExternalExecutionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsExternalWithExternalExecutionItemRequestBuilder) {
    m := &ExecutionsExternalWithExternalExecutionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/executions/external/{externalExecutionId}{?platform*}", pathParameters),
    }
    return m
}
// NewExecutionsExternalWithExternalExecutionItemRequestBuilder instantiates a new WithExternalExecutionItemRequestBuilder and sets the default values.
func NewExecutionsExternalWithExternalExecutionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsExternalWithExternalExecutionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsExternalWithExternalExecutionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsExternalWithExternalExecutionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExecutionsExternalWithExternalExecutionItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
func (m *ExecutionsExternalWithExternalExecutionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExecutionsExternalWithExternalExecutionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ExecutionsExternalWithExternalExecutionItemRequestBuilder) WithUrl(rawUrl string)(*ExecutionsExternalWithExternalExecutionItemRequestBuilder) {
    return NewExecutionsExternalWithExternalExecutionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
