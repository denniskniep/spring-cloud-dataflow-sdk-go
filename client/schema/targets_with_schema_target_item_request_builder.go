package schema

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TargetsWithSchemaTargetItemRequestBuilder builds and executes requests for operations under \schema\targets\{schemaTarget}
type TargetsWithSchemaTargetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TargetsWithSchemaTargetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetsWithSchemaTargetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTargetsWithSchemaTargetItemRequestBuilderInternal instantiates a new WithSchemaTargetItemRequestBuilder and sets the default values.
func NewTargetsWithSchemaTargetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetsWithSchemaTargetItemRequestBuilder) {
    m := &TargetsWithSchemaTargetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/schema/targets/{schemaTarget}", pathParameters),
    }
    return m
}
// NewTargetsWithSchemaTargetItemRequestBuilder instantiates a new WithSchemaTargetItemRequestBuilder and sets the default values.
func NewTargetsWithSchemaTargetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetsWithSchemaTargetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetsWithSchemaTargetItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetsWithSchemaTargetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TargetsWithSchemaTargetItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
func (m *TargetsWithSchemaTargetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TargetsWithSchemaTargetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *TargetsWithSchemaTargetItemRequestBuilder) WithUrl(rawUrl string)(*TargetsWithSchemaTargetItemRequestBuilder) {
    return NewTargetsWithSchemaTargetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
