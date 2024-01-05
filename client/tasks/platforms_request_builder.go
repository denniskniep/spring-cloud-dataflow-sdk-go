package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PlatformsRequestBuilder builds and executes requests for operations under \tasks\platforms
type PlatformsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PlatformsRequestBuilderGetQueryParameters 
type PlatformsRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    SchedulesEnabled *string `uriparametername:"schedulesEnabled"`
}
// PlatformsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlatformsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlatformsRequestBuilderGetQueryParameters
}
// NewPlatformsRequestBuilderInternal instantiates a new PlatformsRequestBuilder and sets the default values.
func NewPlatformsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlatformsRequestBuilder) {
    m := &PlatformsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/platforms{?pageable*,schedulesEnabled*,assembler*}", pathParameters),
    }
    return m
}
// NewPlatformsRequestBuilder instantiates a new PlatformsRequestBuilder and sets the default values.
func NewPlatformsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlatformsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlatformsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PlatformsRequestBuilder) Get(ctx context.Context, requestConfiguration *PlatformsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *PlatformsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PlatformsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PlatformsRequestBuilder) WithUrl(rawUrl string)(*PlatformsRequestBuilder) {
    return NewPlatformsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
