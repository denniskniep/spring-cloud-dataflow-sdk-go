package tools

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/models"
)

// ConvertTaskGraphToTextRequestBuilder builds and executes requests for operations under \tools\convertTaskGraphToText
type ConvertTaskGraphToTextRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConvertTaskGraphToTextRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConvertTaskGraphToTextRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConvertTaskGraphToTextRequestBuilderInternal instantiates a new ConvertTaskGraphToTextRequestBuilder and sets the default values.
func NewConvertTaskGraphToTextRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConvertTaskGraphToTextRequestBuilder) {
    m := &ConvertTaskGraphToTextRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tools/convertTaskGraphToText", pathParameters),
    }
    return m
}
// NewConvertTaskGraphToTextRequestBuilder instantiates a new ConvertTaskGraphToTextRequestBuilder and sets the default values.
func NewConvertTaskGraphToTextRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConvertTaskGraphToTextRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConvertTaskGraphToTextRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ConvertTaskGraphToTextRequestBuilder) Post(ctx context.Context, body iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.Graphable, requestConfiguration *ConvertTaskGraphToTextRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *ConvertTaskGraphToTextRequestBuilder) ToPostRequestInformation(ctx context.Context, body iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.Graphable, requestConfiguration *ConvertTaskGraphToTextRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ConvertTaskGraphToTextRequestBuilder) WithUrl(rawUrl string)(*ConvertTaskGraphToTextRequestBuilder) {
    return NewConvertTaskGraphToTextRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
