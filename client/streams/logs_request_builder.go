package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LogsRequestBuilder builds and executes requests for operations under \streams\logs
type LogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByStreamName gets an item from the dataflow/client.streams.logs.item collection
func (m *LogsRequestBuilder) ByStreamName(streamName string)(*LogsWithStreamNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if streamName != "" {
        urlTplParams["streamName"] = streamName
    }
    return NewLogsWithStreamNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLogsRequestBuilderInternal instantiates a new LogsRequestBuilder and sets the default values.
func NewLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsRequestBuilder) {
    m := &LogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/logs", pathParameters),
    }
    return m
}
// NewLogsRequestBuilder instantiates a new LogsRequestBuilder and sets the default values.
func NewLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLogsRequestBuilderInternal(urlParams, requestAdapter)
}
