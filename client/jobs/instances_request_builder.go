package jobs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstancesRequestBuilder builds and executes requests for operations under \jobs\instances
type InstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InstancesRequestBuilderGetQueryParameters 
type InstancesRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Name *string `uriparametername:"name"`
    // 
    Pageable *string `uriparametername:"pageable"`
}
// InstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstancesRequestBuilderGetQueryParameters
}
// ById gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.jobs.instances.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *InstancesRequestBuilder) ById(id string)(*InstancesInstancesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewInstancesInstancesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInt64 gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.jobs.instances.item collection
func (m *InstancesRequestBuilder) ByIdInt64(id int64)(*InstancesInstancesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewInstancesInstancesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    m := &InstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/instances{?name*,pageable*,assembler*}", pathParameters),
    }
    return m
}
// NewInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *InstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *InstancesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *InstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *InstancesRequestBuilder) WithUrl(rawUrl string)(*InstancesRequestBuilder) {
    return NewInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
