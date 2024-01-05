package streams

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ValidationRequestBuilder builds and executes requests for operations under \streams\validation
type ValidationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.streams.validation.item collection
func (m *ValidationRequestBuilder) ByName(name string)(*ValidationWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewValidationWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewValidationRequestBuilderInternal instantiates a new ValidationRequestBuilder and sets the default values.
func NewValidationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ValidationRequestBuilder) {
    m := &ValidationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/streams/validation", pathParameters),
    }
    return m
}
// NewValidationRequestBuilder instantiates a new ValidationRequestBuilder and sets the default values.
func NewValidationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ValidationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewValidationRequestBuilderInternal(urlParams, requestAdapter)
}
