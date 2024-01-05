package apps

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithTypeItemRequestBuilder builds and executes requests for operations under \apps\{type}
type WithTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the dataflow/client.apps.item.item collection
// Deprecated: 
func (m *WithTypeItemRequestBuilder) ByName(name string)(*ItemWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewItemWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithTypeItemRequestBuilderInternal instantiates a new WithTypeItemRequestBuilder and sets the default values.
func NewWithTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithTypeItemRequestBuilder) {
    m := &WithTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/apps/{type}", pathParameters),
    }
    return m
}
// NewWithTypeItemRequestBuilder instantiates a new WithTypeItemRequestBuilder and sets the default values.
func NewWithTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
