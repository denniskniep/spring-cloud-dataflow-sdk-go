package jobs

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/models"
)

// ThinexecutionsRequestBuilder builds and executes requests for operations under \jobs\thinexecutions
type ThinexecutionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThinexecutionsRequestBuilderGetQueryParameters 
type ThinexecutionsRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    FromDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"fromDate"`
    // 
    JobInstanceId *int32 `uriparametername:"jobInstanceId"`
    // 
    Name *string `uriparametername:"name"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
    // 
    TaskExecutionId *int32 `uriparametername:"taskExecutionId"`
    // 
    ToDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"toDate"`
}
// ThinexecutionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThinexecutionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThinexecutionsRequestBuilderGetQueryParameters
}
// NewThinexecutionsRequestBuilderInternal instantiates a new ThinexecutionsRequestBuilder and sets the default values.
func NewThinexecutionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThinexecutionsRequestBuilder) {
    m := &ThinexecutionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/thinexecutions{?pageable*,assembler*,taskExecutionId*,schemaTarget*,name*,jobInstanceId*,fromDate*,toDate*}", pathParameters),
    }
    return m
}
// NewThinexecutionsRequestBuilder instantiates a new ThinexecutionsRequestBuilder and sets the default values.
func NewThinexecutionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThinexecutionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThinexecutionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ThinexecutionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThinexecutionsRequestBuilderGetRequestConfiguration)(iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.PagedModelJobExecutionThinResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.CreatePagedModelJobExecutionThinResourceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.PagedModelJobExecutionThinResourceable), nil
}
func (m *ThinexecutionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThinexecutionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ThinexecutionsRequestBuilder) WithUrl(rawUrl string)(*ThinexecutionsRequestBuilder) {
    return NewThinexecutionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
