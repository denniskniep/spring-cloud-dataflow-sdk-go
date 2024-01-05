package streams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsRequestBuilder builds and executes requests for operations under \streams\deployments
type DeploymentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeploymentsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeploymentsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByName gets an item from the dataflow/client.streams.deployments.item collection
func (m *DeploymentsRequestBuilder) ByName(name string)(*DeploymentsWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewDeploymentsWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsRequestBuilderInternal instantiates a new DeploymentsRequestBuilder and sets the default values.
func NewDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRequestBuilder) {
    m := &DeploymentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments", pathParameters),
    }
    return m
}
// NewDeploymentsRequestBuilder instantiates a new DeploymentsRequestBuilder and sets the default values.
func NewDeploymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeploymentsRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeploymentsRequestBuilderDeleteRequestConfiguration)([]byte, error) {
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
// History the history property
func (m *DeploymentsRequestBuilder) History()(*DeploymentsHistoryRequestBuilder) {
    return NewDeploymentsHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Manifest the manifest property
func (m *DeploymentsRequestBuilder) Manifest()(*DeploymentsManifestRequestBuilder) {
    return NewDeploymentsManifestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Platform the platform property
func (m *DeploymentsRequestBuilder) Platform()(*DeploymentsPlatformRequestBuilder) {
    return NewDeploymentsPlatformRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rollback the rollback property
func (m *DeploymentsRequestBuilder) Rollback()(*DeploymentsRollbackRequestBuilder) {
    return NewDeploymentsRollbackRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Scale the scale property
func (m *DeploymentsRequestBuilder) Scale()(*DeploymentsScaleRequestBuilder) {
    return NewDeploymentsScaleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *DeploymentsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeploymentsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Update the update property
func (m *DeploymentsRequestBuilder) Update()(*DeploymentsUpdateRequestBuilder) {
    return NewDeploymentsUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeploymentsRequestBuilder) WithUrl(rawUrl string)(*DeploymentsRequestBuilder) {
    return NewDeploymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
