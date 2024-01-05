package auditrecords

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuditRecordsRequestBuilder builds and executes requests for operations under \audit-records
type AuditRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuditRecordsRequestBuilderGetQueryParameters 
type AuditRecordsRequestBuilderGetQueryParameters struct {
    // 
    Actions []string `uriparametername:"actions"`
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    FromDate *string `uriparametername:"fromDate"`
    // 
    Operations []string `uriparametername:"operations"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    ToDate *string `uriparametername:"toDate"`
}
// AuditRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuditRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuditRecordsRequestBuilderGetQueryParameters
}
// AuditActionTypes the auditActionTypes property
func (m *AuditRecordsRequestBuilder) AuditActionTypes()(*AuditActionTypesRequestBuilder) {
    return NewAuditActionTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditOperationTypes the auditOperationTypes property
func (m *AuditRecordsRequestBuilder) AuditOperationTypes()(*AuditOperationTypesRequestBuilder) {
    return NewAuditOperationTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ById gets an item from the dataflow/client.auditRecords.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *AuditRecordsRequestBuilder) ById(id string)(*AuditRecordsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewAuditRecordsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInt64 gets an item from the dataflow/client.auditRecords.item collection
func (m *AuditRecordsRequestBuilder) ByIdInt64(id int64)(*AuditRecordsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewAuditRecordsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuditRecordsRequestBuilderInternal instantiates a new AuditRecordsRequestBuilder and sets the default values.
func NewAuditRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditRecordsRequestBuilder) {
    m := &AuditRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/audit-records{?pageable*,actions*,operations*,fromDate*,toDate*,assembler*}", pathParameters),
    }
    return m
}
// NewAuditRecordsRequestBuilder instantiates a new AuditRecordsRequestBuilder and sets the default values.
func NewAuditRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AuditRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuditRecordsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *AuditRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuditRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AuditRecordsRequestBuilder) WithUrl(rawUrl string)(*AuditRecordsRequestBuilder) {
    return NewAuditRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
