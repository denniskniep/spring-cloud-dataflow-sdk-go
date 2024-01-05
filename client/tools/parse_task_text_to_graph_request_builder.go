package tools

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ParseTaskTextToGraphRequestBuilder builds and executes requests for operations under \tools\parseTaskTextToGraph
type ParseTaskTextToGraphRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ParseTaskTextToGraphRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ParseTaskTextToGraphRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewParseTaskTextToGraphRequestBuilderInternal instantiates a new ParseTaskTextToGraphRequestBuilder and sets the default values.
func NewParseTaskTextToGraphRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ParseTaskTextToGraphRequestBuilder) {
    m := &ParseTaskTextToGraphRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tools/parseTaskTextToGraph", pathParameters),
    }
    return m
}
// NewParseTaskTextToGraphRequestBuilder instantiates a new ParseTaskTextToGraphRequestBuilder and sets the default values.
func NewParseTaskTextToGraphRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ParseTaskTextToGraphRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParseTaskTextToGraphRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParseTaskTextToGraphRequestBuilder) Post(ctx context.Context, body ParseTaskTextToGraphPostRequestBodyable, requestConfiguration *ParseTaskTextToGraphRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *ParseTaskTextToGraphRequestBuilder) ToPostRequestInformation(ctx context.Context, body ParseTaskTextToGraphPostRequestBodyable, requestConfiguration *ParseTaskTextToGraphRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ParseTaskTextToGraphRequestBuilder) WithUrl(rawUrl string)(*ParseTaskTextToGraphRequestBuilder) {
    return NewParseTaskTextToGraphRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
