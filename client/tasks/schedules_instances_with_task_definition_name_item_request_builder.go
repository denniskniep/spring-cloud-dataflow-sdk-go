package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder builds and executes requests for operations under \tasks\schedules\instances\{taskDefinitionName}
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteQueryParameters
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetQueryParameters
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadQueryParameters
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsQueryParameters
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchQueryParameters
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostQueryParameters
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutQueryParameters 
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Platform *string `uriparametername:"platform"`
}
// SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutQueryParameters
}
// NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilderInternal instantiates a new WithTaskDefinitionNameItemRequestBuilder and sets the default values.
func NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) {
    m := &SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/schedules/instances/{taskDefinitionName}{?platform*,assembler*}", pathParameters),
    }
    return m
}
// NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilder instantiates a new WithTaskDefinitionNameItemRequestBuilder and sets the default values.
func NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Head(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToHeadRequestInformation(ctx, requestConfiguration);
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Options(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToOptionsRequestInformation(ctx, requestConfiguration);
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Patch(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, requestConfiguration);
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Post(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) Put(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToHeadRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderHeadRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.HEAD, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToOptionsRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderOptionsRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.OPTIONS, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) WithUrl(rawUrl string)(*SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) {
    return NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
