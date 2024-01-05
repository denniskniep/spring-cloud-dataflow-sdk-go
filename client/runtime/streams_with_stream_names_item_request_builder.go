package runtime

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StreamsWithStreamNamesItemRequestBuilder builds and executes requests for operations under \runtime\streams\{streamNames}
type StreamsWithStreamNamesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StreamsWithStreamNamesItemRequestBuilderGetQueryParameters 
type StreamsWithStreamNamesItemRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Pageable *string `uriparametername:"pageable"`
}
// StreamsWithStreamNamesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StreamsWithStreamNamesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *StreamsWithStreamNamesItemRequestBuilderGetQueryParameters
}
// NewStreamsWithStreamNamesItemRequestBuilderInternal instantiates a new WithStreamNamesItemRequestBuilder and sets the default values.
func NewStreamsWithStreamNamesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamsWithStreamNamesItemRequestBuilder) {
    m := &StreamsWithStreamNamesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/runtime/streams/{streamNames}{?pageable*,assembler*}", pathParameters),
    }
    return m
}
// NewStreamsWithStreamNamesItemRequestBuilder instantiates a new WithStreamNamesItemRequestBuilder and sets the default values.
func NewStreamsWithStreamNamesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamsWithStreamNamesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStreamsWithStreamNamesItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *StreamsWithStreamNamesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *StreamsWithStreamNamesItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *StreamsWithStreamNamesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *StreamsWithStreamNamesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *StreamsWithStreamNamesItemRequestBuilder) WithUrl(rawUrl string)(*StreamsWithStreamNamesItemRequestBuilder) {
    return NewStreamsWithStreamNamesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
