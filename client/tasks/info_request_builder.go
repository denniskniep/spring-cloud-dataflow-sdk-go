package tasks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InfoRequestBuilder builds and executes requests for operations under \tasks\info
type InfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewInfoRequestBuilderInternal instantiates a new InfoRequestBuilder and sets the default values.
func NewInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InfoRequestBuilder) {
    m := &InfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/info", pathParameters),
    }
    return m
}
// NewInfoRequestBuilder instantiates a new InfoRequestBuilder and sets the default values.
func NewInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Executions the executions property
func (m *InfoRequestBuilder) Executions()(*InfoExecutionsRequestBuilder) {
    return NewInfoExecutionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
