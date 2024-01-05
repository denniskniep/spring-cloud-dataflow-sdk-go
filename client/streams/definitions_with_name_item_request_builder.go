package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DefinitionsWithNameItemRequestBuilder builds and executes requests for operations under \streams\definitions\{name}
type DefinitionsWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DefinitionsWithNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefinitionsWithNameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DefinitionsWithNameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefinitionsWithNameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Applications the applications property
func (m *DefinitionsWithNameItemRequestBuilder) Applications()(*DefinitionsItemApplicationsRequestBuilder) {
    return NewDefinitionsItemApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDefinitionsWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDefinitionsWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefinitionsWithNameItemRequestBuilder) {
    m := &DefinitionsWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/definitions/{name}", pathParameters),
    }
    return m
}
// NewDefinitionsWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDefinitionsWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefinitionsWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefinitionsWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DefinitionsWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DefinitionsWithNameItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
func (m *DefinitionsWithNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DefinitionsWithNameItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Related the related property
func (m *DefinitionsWithNameItemRequestBuilder) Related()(*DefinitionsItemRelatedRequestBuilder) {
    return NewDefinitionsItemRelatedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *DefinitionsWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DefinitionsWithNameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
func (m *DefinitionsWithNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefinitionsWithNameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DefinitionsWithNameItemRequestBuilder) WithUrl(rawUrl string)(*DefinitionsWithNameItemRequestBuilder) {
    return NewDefinitionsWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
