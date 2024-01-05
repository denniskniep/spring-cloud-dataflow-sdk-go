package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsPlatformRequestBuilder builds and executes requests for operations under \streams\deployments\platform
type DeploymentsPlatformRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewDeploymentsPlatformRequestBuilderInternal instantiates a new PlatformRequestBuilder and sets the default values.
func NewDeploymentsPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsPlatformRequestBuilder) {
    m := &DeploymentsPlatformRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/platform", pathParameters),
    }
    return m
}
// NewDeploymentsPlatformRequestBuilder instantiates a new PlatformRequestBuilder and sets the default values.
func NewDeploymentsPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsPlatformRequestBuilderInternal(urlParams, requestAdapter)
}
// List the list property
func (m *DeploymentsPlatformRequestBuilder) List()(*DeploymentsPlatformListRequestBuilder) {
    return NewDeploymentsPlatformListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
