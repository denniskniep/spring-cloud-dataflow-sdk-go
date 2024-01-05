package tools

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ToolsRequestBuilder builds and executes requests for operations under \tools
type ToolsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewToolsRequestBuilderInternal instantiates a new ToolsRequestBuilder and sets the default values.
func NewToolsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ToolsRequestBuilder) {
    m := &ToolsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tools", pathParameters),
    }
    return m
}
// NewToolsRequestBuilder instantiates a new ToolsRequestBuilder and sets the default values.
func NewToolsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ToolsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewToolsRequestBuilderInternal(urlParams, requestAdapter)
}
// ConvertTaskGraphToText the convertTaskGraphToText property
func (m *ToolsRequestBuilder) ConvertTaskGraphToText()(*ConvertTaskGraphToTextRequestBuilder) {
    return NewConvertTaskGraphToTextRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParseTaskTextToGraph the parseTaskTextToGraph property
func (m *ToolsRequestBuilder) ParseTaskTextToGraph()(*ParseTaskTextToGraphRequestBuilder) {
    return NewParseTaskTextToGraphRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
