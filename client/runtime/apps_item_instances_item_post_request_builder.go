package runtime

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AppsItemInstancesItemPostRequestBuilder builds and executes requests for operations under \runtime\apps\{appId}\instances\{instanceId}\post
type AppsItemInstancesItemPostRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppsItemInstancesItemPostRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppsItemInstancesItemPostRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppsItemInstancesItemPostRequestBuilderInternal instantiates a new PostRequestBuilder and sets the default values.
func NewAppsItemInstancesItemPostRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppsItemInstancesItemPostRequestBuilder) {
    m := &AppsItemInstancesItemPostRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/runtime/apps/{appId}/instances/{instanceId}/post", pathParameters),
    }
    return m
}
// NewAppsItemInstancesItemPostRequestBuilder instantiates a new PostRequestBuilder and sets the default values.
func NewAppsItemInstancesItemPostRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppsItemInstancesItemPostRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppsItemInstancesItemPostRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AppsItemInstancesItemPostRequestBuilder) Post(ctx context.Context, body *string, requestConfiguration *AppsItemInstancesItemPostRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *AppsItemInstancesItemPostRequestBuilder) ToPostRequestInformation(ctx context.Context, body *string, requestConfiguration *AppsItemInstancesItemPostRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    requestInfo.SetContentFromScalar(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AppsItemInstancesItemPostRequestBuilder) WithUrl(rawUrl string)(*AppsItemInstancesItemPostRequestBuilder) {
    return NewAppsItemInstancesItemPostRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
