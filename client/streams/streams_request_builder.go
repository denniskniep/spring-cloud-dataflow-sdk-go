package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StreamsRequestBuilder builds and executes requests for operations under \streams
type StreamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewStreamsRequestBuilderInternal instantiates a new StreamsRequestBuilder and sets the default values.
func NewStreamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamsRequestBuilder) {
    m := &StreamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams", pathParameters),
    }
    return m
}
// NewStreamsRequestBuilder instantiates a new StreamsRequestBuilder and sets the default values.
func NewStreamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStreamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Definitions the definitions property
func (m *StreamsRequestBuilder) Definitions()(*DefinitionsRequestBuilder) {
    return NewDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deployments the deployments property
func (m *StreamsRequestBuilder) Deployments()(*DeploymentsRequestBuilder) {
    return NewDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logs the logs property
func (m *StreamsRequestBuilder) Logs()(*LogsRequestBuilder) {
    return NewLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Validation the validation property
func (m *StreamsRequestBuilder) Validation()(*ValidationRequestBuilder) {
    return NewValidationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
