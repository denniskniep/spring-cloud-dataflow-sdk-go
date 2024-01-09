package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchedulesRequestBuilder builds and executes requests for operations under \tasks\schedules
type SchedulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchedulesRequestBuilderGetQueryParameters 
type SchedulesRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesRequestBuilderGetQueryParameters
}
// SchedulesRequestBuilderPostQueryParameters 
type SchedulesRequestBuilderPostQueryParameters struct {
    // 
    Arguments *string `uriparametername:"arguments"`
    // 
    Platform *string `uriparametername:"platform"`
    // 
    Properties *string `uriparametername:"properties"`
    // 
    ScheduleName *string `uriparametername:"scheduleName"`
    // 
    TaskDefinitionName *string `uriparametername:"taskDefinitionName"`
}
// SchedulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesRequestBuilderPostQueryParameters
}
// BySchedulesId gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.tasks.schedules.item collection
func (m *SchedulesRequestBuilder) BySchedulesId(schedulesId string)(*SchedulesSchedulesItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if schedulesId != "" {
        urlTplParams["schedules%2Did"] = schedulesId
    }
    return NewSchedulesSchedulesItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSchedulesRequestBuilderInternal instantiates a new SchedulesRequestBuilder and sets the default values.
func NewSchedulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchedulesRequestBuilder) {
    m := &SchedulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/schedules{?pageable*,platform*,assembler*,scheduleName*,taskDefinitionName*,properties*,arguments*}", pathParameters),
    }
    return m
}
// NewSchedulesRequestBuilder instantiates a new SchedulesRequestBuilder and sets the default values.
func NewSchedulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchedulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchedulesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SchedulesRequestBuilder) Get(ctx context.Context, requestConfiguration *SchedulesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Instances the instances property
func (m *SchedulesRequestBuilder) Instances()(*SchedulesInstancesRequestBuilder) {
    return NewSchedulesInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *SchedulesRequestBuilder) Post(ctx context.Context, requestConfiguration *SchedulesRequestBuilderPostRequestConfiguration)(error) {
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
func (m *SchedulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchedulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchedulesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SchedulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SchedulesRequestBuilder) WithUrl(rawUrl string)(*SchedulesRequestBuilder) {
    return NewSchedulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
