package tasks

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExecutionsLaunchRequestBuilder builds and executes requests for operations under \tasks\executions\launch
type ExecutionsLaunchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsLaunchRequestBuilderPostQueryParameters 
type ExecutionsLaunchRequestBuilderPostQueryParameters struct {
    // 
    Arguments *string `uriparametername:"arguments"`
    // 
    Name *string `uriparametername:"name"`
    // 
    Properties *string `uriparametername:"properties"`
}
// ExecutionsLaunchRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsLaunchRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsLaunchRequestBuilderPostQueryParameters
}
// NewExecutionsLaunchRequestBuilderInternal instantiates a new LaunchRequestBuilder and sets the default values.
func NewExecutionsLaunchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsLaunchRequestBuilder) {
    m := &ExecutionsLaunchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/executions/launch{?name*,properties*,arguments*}", pathParameters),
    }
    return m
}
// NewExecutionsLaunchRequestBuilder instantiates a new LaunchRequestBuilder and sets the default values.
func NewExecutionsLaunchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsLaunchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsLaunchRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsLaunchRequestBuilder) Post(ctx context.Context, requestConfiguration *ExecutionsLaunchRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ExecutionsLaunchRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ExecutionsLaunchRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExecutionsLaunchRequestBuilder) WithUrl(rawUrl string)(*ExecutionsLaunchRequestBuilder) {
    return NewExecutionsLaunchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
