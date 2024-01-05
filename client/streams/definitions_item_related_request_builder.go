package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DefinitionsItemRelatedRequestBuilder builds and executes requests for operations under \streams\definitions\{name}\related
type DefinitionsItemRelatedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DefinitionsItemRelatedRequestBuilderGetQueryParameters 
type DefinitionsItemRelatedRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Nested *bool `uriparametername:"nested"`
    // 
    Pageable *string `uriparametername:"pageable"`
}
// DefinitionsItemRelatedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefinitionsItemRelatedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DefinitionsItemRelatedRequestBuilderGetQueryParameters
}
// NewDefinitionsItemRelatedRequestBuilderInternal instantiates a new RelatedRequestBuilder and sets the default values.
func NewDefinitionsItemRelatedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefinitionsItemRelatedRequestBuilder) {
    m := &DefinitionsItemRelatedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/definitions/{name}/related{?pageable*,nested*,assembler*}", pathParameters),
    }
    return m
}
// NewDefinitionsItemRelatedRequestBuilder instantiates a new RelatedRequestBuilder and sets the default values.
func NewDefinitionsItemRelatedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefinitionsItemRelatedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefinitionsItemRelatedRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DefinitionsItemRelatedRequestBuilder) Get(ctx context.Context, requestConfiguration *DefinitionsItemRelatedRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *DefinitionsItemRelatedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefinitionsItemRelatedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DefinitionsItemRelatedRequestBuilder) WithUrl(rawUrl string)(*DefinitionsItemRelatedRequestBuilder) {
    return NewDefinitionsItemRelatedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
