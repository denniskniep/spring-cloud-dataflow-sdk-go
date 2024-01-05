package auditrecords

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuditOperationTypesRequestBuilder builds and executes requests for operations under \audit-records\audit-operation-types
type AuditOperationTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuditOperationTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditOperationTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuditOperationTypesRequestBuilderInternal instantiates a new AuditOperationTypesRequestBuilder and sets the default values.
func NewAuditOperationTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditOperationTypesRequestBuilder) {
    m := &AuditOperationTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/audit-records/audit-operation-types", pathParameters),
    }
    return m
}
// NewAuditOperationTypesRequestBuilder instantiates a new AuditOperationTypesRequestBuilder and sets the default values.
func NewAuditOperationTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditOperationTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditOperationTypesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AuditOperationTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditOperationTypesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *AuditOperationTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditOperationTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AuditOperationTypesRequestBuilder) WithUrl(rawUrl string)(*AuditOperationTypesRequestBuilder) {
    return NewAuditOperationTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
