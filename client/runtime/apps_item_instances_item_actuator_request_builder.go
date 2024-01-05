package runtime

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/models"
)

// AppsItemInstancesItemActuatorRequestBuilder builds and executes requests for operations under \runtime\apps\{appId}\instances\{instanceId}\actuator
type AppsItemInstancesItemActuatorRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppsItemInstancesItemActuatorRequestBuilderGetQueryParameters 
type AppsItemInstancesItemActuatorRequestBuilderGetQueryParameters struct {
    // 
    Endpoint *string `uriparametername:"endpoint"`
}
// AppsItemInstancesItemActuatorRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppsItemInstancesItemActuatorRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppsItemInstancesItemActuatorRequestBuilderGetQueryParameters
}
// AppsItemInstancesItemActuatorRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppsItemInstancesItemActuatorRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppsItemInstancesItemActuatorRequestBuilderInternal instantiates a new ActuatorRequestBuilder and sets the default values.
func NewAppsItemInstancesItemActuatorRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppsItemInstancesItemActuatorRequestBuilder) {
    m := &AppsItemInstancesItemActuatorRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/runtime/apps/{appId}/instances/{instanceId}/actuator{?endpoint*}", pathParameters),
    }
    return m
}
// NewAppsItemInstancesItemActuatorRequestBuilder instantiates a new ActuatorRequestBuilder and sets the default values.
func NewAppsItemInstancesItemActuatorRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppsItemInstancesItemActuatorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppsItemInstancesItemActuatorRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AppsItemInstancesItemActuatorRequestBuilder) Get(ctx context.Context, requestConfiguration *AppsItemInstancesItemActuatorRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *AppsItemInstancesItemActuatorRequestBuilder) Post(ctx context.Context, body iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.ActuatorPostRequestable, requestConfiguration *AppsItemInstancesItemActuatorRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *AppsItemInstancesItemActuatorRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppsItemInstancesItemActuatorRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppsItemInstancesItemActuatorRequestBuilder) ToPostRequestInformation(ctx context.Context, body iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.ActuatorPostRequestable, requestConfiguration *AppsItemInstancesItemActuatorRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AppsItemInstancesItemActuatorRequestBuilder) WithUrl(rawUrl string)(*AppsItemInstancesItemActuatorRequestBuilder) {
    return NewAppsItemInstancesItemActuatorRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
