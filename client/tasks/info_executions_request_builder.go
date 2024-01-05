package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InfoExecutionsRequestBuilder builds and executes requests for operations under \tasks\info\executions
type InfoExecutionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InfoExecutionsRequestBuilderGetQueryParameters 
type InfoExecutionsRequestBuilderGetQueryParameters struct {
    // 
    Completed *string `uriparametername:"completed"`
    // 
    Days *int32 `uriparametername:"days"`
    // 
    Name *string `uriparametername:"name"`
}
// InfoExecutionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InfoExecutionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InfoExecutionsRequestBuilderGetQueryParameters
}
// NewInfoExecutionsRequestBuilderInternal instantiates a new ExecutionsRequestBuilder and sets the default values.
func NewInfoExecutionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InfoExecutionsRequestBuilder) {
    m := &InfoExecutionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/info/executions{?completed*,name*,days*}", pathParameters),
    }
    return m
}
// NewInfoExecutionsRequestBuilder instantiates a new ExecutionsRequestBuilder and sets the default values.
func NewInfoExecutionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InfoExecutionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInfoExecutionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *InfoExecutionsRequestBuilder) Get(ctx context.Context, requestConfiguration *InfoExecutionsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *InfoExecutionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InfoExecutionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *InfoExecutionsRequestBuilder) WithUrl(rawUrl string)(*InfoExecutionsRequestBuilder) {
    return NewInfoExecutionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
