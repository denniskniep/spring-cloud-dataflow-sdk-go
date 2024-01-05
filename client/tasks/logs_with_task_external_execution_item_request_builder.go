package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LogsWithTaskExternalExecutionItemRequestBuilder builds and executes requests for operations under \tasks\logs\{taskExternalExecutionId}
type LogsWithTaskExternalExecutionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LogsWithTaskExternalExecutionItemRequestBuilderGetQueryParameters 
type LogsWithTaskExternalExecutionItemRequestBuilderGetQueryParameters struct {
    // 
    PlatformName *string `uriparametername:"platformName"`
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// LogsWithTaskExternalExecutionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LogsWithTaskExternalExecutionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LogsWithTaskExternalExecutionItemRequestBuilderGetQueryParameters
}
// NewLogsWithTaskExternalExecutionItemRequestBuilderInternal instantiates a new WithTaskExternalExecutionItemRequestBuilder and sets the default values.
func NewLogsWithTaskExternalExecutionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsWithTaskExternalExecutionItemRequestBuilder) {
    m := &LogsWithTaskExternalExecutionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/logs/{taskExternalExecutionId}{?platformName*,schemaTarget*}", pathParameters),
    }
    return m
}
// NewLogsWithTaskExternalExecutionItemRequestBuilder instantiates a new WithTaskExternalExecutionItemRequestBuilder and sets the default values.
func NewLogsWithTaskExternalExecutionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsWithTaskExternalExecutionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLogsWithTaskExternalExecutionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *LogsWithTaskExternalExecutionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LogsWithTaskExternalExecutionItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *LogsWithTaskExternalExecutionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LogsWithTaskExternalExecutionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LogsWithTaskExternalExecutionItemRequestBuilder) WithUrl(rawUrl string)(*LogsWithTaskExternalExecutionItemRequestBuilder) {
    return NewLogsWithTaskExternalExecutionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
