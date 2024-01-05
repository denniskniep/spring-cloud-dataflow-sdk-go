package apps

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppsRequestBuilder builds and executes requests for operations under \apps
type AppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppsRequestBuilderGetQueryParameters 
type AppsRequestBuilderGetQueryParameters struct {
    // 
    DefaultVersion *bool `uriparametername:"defaultVersion"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    PagedResourcesAssembler *string `uriparametername:"pagedResourcesAssembler"`
    // 
    Search *string `uriparametername:"search"`
    // 
    // Deprecated: This property is deprecated, use typeAsGetTypeQueryParameterType instead
    Type *string `uriparametername:"type"`
    // 
    TypeAsGetTypeQueryParameterType *GetTypeQueryParameterType `uriparametername:"type"`
    // 
    Version *string `uriparametername:"version"`
}
// AppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppsRequestBuilderGetQueryParameters
}
// AppsRequestBuilderPostQueryParameters 
type AppsRequestBuilderPostQueryParameters struct {
    // 
    Apps *string `uriparametername:"apps"`
    // 
    Force *bool `uriparametername:"force"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    PagedResourcesAssembler *string `uriparametername:"pagedResourcesAssembler"`
    // 
    Uri *string `uriparametername:"uri"`
}
// AppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppsRequestBuilderPostQueryParameters
}
// ByType gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.apps.item collection
func (m *AppsRequestBuilder) ByType(type string)(*WithTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if type != "" {
        urlTplParams["type"] = type
    }
    return NewWithTypeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppsRequestBuilderInternal instantiates a new AppsRequestBuilder and sets the default values.
func NewAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppsRequestBuilder) {
    m := &AppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/apps{?pageable*,pagedResourcesAssembler*,type*,search*,version*,defaultVersion*,uri*,apps*,force*}", pathParameters),
    }
    return m
}
// NewAppsRequestBuilder instantiates a new AppsRequestBuilder and sets the default values.
func NewAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AppsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppsRequestBuilderDeleteRequestConfiguration)([]byte, error) {
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
func (m *AppsRequestBuilder) Get(ctx context.Context, requestConfiguration *AppsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *AppsRequestBuilder) Post(ctx context.Context, requestConfiguration *AppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AppsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
func (m *AppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *AppsRequestBuilder) WithUrl(rawUrl string)(*AppsRequestBuilder) {
    return NewAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
