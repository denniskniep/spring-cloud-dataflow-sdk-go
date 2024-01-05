package completions

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompletionsRequestBuilder builds and executes requests for operations under \completions
type CompletionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewCompletionsRequestBuilderInternal instantiates a new CompletionsRequestBuilder and sets the default values.
func NewCompletionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompletionsRequestBuilder) {
    m := &CompletionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/completions", pathParameters),
    }
    return m
}
// NewCompletionsRequestBuilder instantiates a new CompletionsRequestBuilder and sets the default values.
func NewCompletionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompletionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompletionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Stream the stream property
func (m *CompletionsRequestBuilder) Stream()(*StreamRequestBuilder) {
    return NewStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Task the task property
func (m *CompletionsRequestBuilder) Task()(*TaskRequestBuilder) {
    return NewTaskRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
