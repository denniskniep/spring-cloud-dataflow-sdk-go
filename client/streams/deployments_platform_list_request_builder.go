package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsPlatformListRequestBuilder builds and executes requests for operations under \streams\deployments\platform\list
type DeploymentsPlatformListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeploymentsPlatformListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeploymentsPlatformListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeploymentsPlatformListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewDeploymentsPlatformListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsPlatformListRequestBuilder) {
    m := &DeploymentsPlatformListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/platform/list", pathParameters),
    }
    return m
}
// NewDeploymentsPlatformListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewDeploymentsPlatformListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsPlatformListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsPlatformListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsPlatformListRequestBuilder) Get(ctx context.Context, requestConfiguration *DeploymentsPlatformListRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *DeploymentsPlatformListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeploymentsPlatformListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeploymentsPlatformListRequestBuilder) WithUrl(rawUrl string)(*DeploymentsPlatformListRequestBuilder) {
    return NewDeploymentsPlatformListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
