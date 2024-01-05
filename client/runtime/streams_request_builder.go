package runtime

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StreamsRequestBuilder builds and executes requests for operations under \runtime\streams
type StreamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StreamsRequestBuilderGetQueryParameters 
type StreamsRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Names []string `uriparametername:"names"`
    // 
    Pageable *string `uriparametername:"pageable"`
}
// StreamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StreamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *StreamsRequestBuilderGetQueryParameters
}
// ByStreamNames gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.runtime.streams.item collection
func (m *StreamsRequestBuilder) ByStreamNames(streamNames string)(*StreamsWithStreamNamesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if streamNames != "" {
        urlTplParams["streamNames"] = streamNames
    }
    return NewStreamsWithStreamNamesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewStreamsRequestBuilderInternal instantiates a new StreamsRequestBuilder and sets the default values.
func NewStreamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamsRequestBuilder) {
    m := &StreamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/runtime/streams{?names*,pageable*,assembler*}", pathParameters),
    }
    return m
}
// NewStreamsRequestBuilder instantiates a new StreamsRequestBuilder and sets the default values.
func NewStreamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStreamsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *StreamsRequestBuilder) Get(ctx context.Context, requestConfiguration *StreamsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *StreamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *StreamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *StreamsRequestBuilder) WithUrl(rawUrl string)(*StreamsRequestBuilder) {
    return NewStreamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
