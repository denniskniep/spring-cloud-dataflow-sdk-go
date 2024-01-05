package schema

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchemaRequestBuilder builds and executes requests for operations under \schema
type SchemaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewSchemaRequestBuilderInternal instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaRequestBuilder) {
    m := &SchemaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/schema", pathParameters),
    }
    return m
}
// NewSchemaRequestBuilder instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// Targets the targets property
func (m *SchemaRequestBuilder) Targets()(*TargetsRequestBuilder) {
    return NewTargetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions the versions property
func (m *SchemaRequestBuilder) Versions()(*VersionsRequestBuilder) {
    return NewVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
