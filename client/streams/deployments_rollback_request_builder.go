package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsRollbackRequestBuilder builds and executes requests for operations under \streams\deployments\rollback
type DeploymentsRollbackRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.rollback.item collection
func (m *DeploymentsRollbackRequestBuilder) ByName(name string)(*DeploymentsRollbackWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewDeploymentsRollbackWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsRollbackRequestBuilderInternal instantiates a new RollbackRequestBuilder and sets the default values.
func NewDeploymentsRollbackRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRollbackRequestBuilder) {
    m := &DeploymentsRollbackRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/rollback", pathParameters),
    }
    return m
}
// NewDeploymentsRollbackRequestBuilder instantiates a new RollbackRequestBuilder and sets the default values.
func NewDeploymentsRollbackRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRollbackRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsRollbackRequestBuilderInternal(urlParams, requestAdapter)
}
