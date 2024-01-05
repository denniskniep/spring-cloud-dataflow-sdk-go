package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstancesInstancesItemRequestBuilder builds and executes requests for operations under \jobs\instances\{id}
type InstancesInstancesItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InstancesInstancesItemRequestBuilderGetQueryParameters 
type InstancesInstancesItemRequestBuilderGetQueryParameters struct {
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// InstancesInstancesItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstancesInstancesItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstancesInstancesItemRequestBuilderGetQueryParameters
}
// NewInstancesInstancesItemRequestBuilderInternal instantiates a new InstancesItemRequestBuilder and sets the default values.
func NewInstancesInstancesItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesInstancesItemRequestBuilder) {
    m := &InstancesInstancesItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/instances/{id}{?schemaTarget*}", pathParameters),
    }
    return m
}
// NewInstancesInstancesItemRequestBuilder instantiates a new InstancesItemRequestBuilder and sets the default values.
func NewInstancesInstancesItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesInstancesItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstancesInstancesItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *InstancesInstancesItemRequestBuilder) Get(ctx context.Context, requestConfiguration *InstancesInstancesItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *InstancesInstancesItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InstancesInstancesItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *InstancesInstancesItemRequestBuilder) WithUrl(rawUrl string)(*InstancesInstancesItemRequestBuilder) {
    return NewInstancesInstancesItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
