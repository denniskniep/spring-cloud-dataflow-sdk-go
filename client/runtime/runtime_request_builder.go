package runtime

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RuntimeRequestBuilder builds and executes requests for operations under \runtime
type RuntimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Apps the apps property
func (m *RuntimeRequestBuilder) Apps()(*AppsRequestBuilder) {
    return NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRuntimeRequestBuilderInternal instantiates a new RuntimeRequestBuilder and sets the default values.
func NewRuntimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RuntimeRequestBuilder) {
    m := &RuntimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/runtime", pathParameters),
    }
    return m
}
// NewRuntimeRequestBuilder instantiates a new RuntimeRequestBuilder and sets the default values.
func NewRuntimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RuntimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRuntimeRequestBuilderInternal(urlParams, requestAdapter)
}
// Streams the streams property
func (m *RuntimeRequestBuilder) Streams()(*StreamsRequestBuilder) {
    return NewStreamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
