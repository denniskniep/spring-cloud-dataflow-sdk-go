package streams

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeploymentsRollbackWithNameItemRequestBuilder builds and executes requests for operations under \streams\deployments\rollback\{name}
type DeploymentsRollbackWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByVersion gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.rollback.item.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *DeploymentsRollbackWithNameItemRequestBuilder) ByVersion(version string)(*DeploymentsRollbackItemWithVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if version != "" {
        urlTplParams["version"] = version
    }
    return NewDeploymentsRollbackItemWithVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByVersionInteger gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.streams.deployments.rollback.item.item collection
func (m *DeploymentsRollbackWithNameItemRequestBuilder) ByVersionInteger(version int32)(*DeploymentsRollbackItemWithVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["version"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(version), 10)
    return NewDeploymentsRollbackItemWithVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeploymentsRollbackWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDeploymentsRollbackWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRollbackWithNameItemRequestBuilder) {
    m := &DeploymentsRollbackWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/deployments/rollback/{name}", pathParameters),
    }
    return m
}
// NewDeploymentsRollbackWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewDeploymentsRollbackWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeploymentsRollbackWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsRollbackWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
