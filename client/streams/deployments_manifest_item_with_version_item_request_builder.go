package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsManifestItemWithVersionItemRequestBuilder builds and executes requests for operations under \streams\deployments\manifest\{name}\{version}
type DeploymentsManifestItemWithVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeploymentsManifestItemWithVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeploymentsManifestItemWithVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeploymentsManifestItemWithVersionItemRequestBuilderInternal instantiates a new WithVersionItemRequestBuilder and sets the default values.
func NewDeploymentsManifestItemWithVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsManifestItemWithVersionItemRequestBuilder) {
    m := &DeploymentsManifestItemWithVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/manifest/{name}/{version}", pathParameters),
    }
    return m
}
// NewDeploymentsManifestItemWithVersionItemRequestBuilder instantiates a new WithVersionItemRequestBuilder and sets the default values.
func NewDeploymentsManifestItemWithVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsManifestItemWithVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsManifestItemWithVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsManifestItemWithVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeploymentsManifestItemWithVersionItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *DeploymentsManifestItemWithVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeploymentsManifestItemWithVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeploymentsManifestItemWithVersionItemRequestBuilder) WithUrl(rawUrl string)(*DeploymentsManifestItemWithVersionItemRequestBuilder) {
    return NewDeploymentsManifestItemWithVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
