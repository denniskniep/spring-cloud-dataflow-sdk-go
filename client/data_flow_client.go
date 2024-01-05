package client

import (
    "context"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i34b8e2fd66099266db79b74496d3c3e30853ae055a475f89902b7f096e86f569 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/completions"
    i433d5dac3040ebc08b4bd12b2de02b126d91cadd6980a9aeef355cb63826d2ee "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/streams"
    i67934d0263c4586557b7c0ac1f884137bbd3b32ac2b948c626329772767c4319 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/about"
    i7638ff68664772d482f8dc66c218408bc63d9f7cd492c2aac4a4dcd01182a0da "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/runtime"
    i768ac0037ddaab39dabfdf36fb1e47c151a3840faaf879f661c819335f83c38f "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/auditrecords"
    i866027ee37135123c0585c891ec25e9426e4445edf81f525db97ef95414bccf5 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/schema"
    i8c8b06b621e92b0c16a6e2ba094614a310f4137ec10df70708b307bdcda39dae "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/apps"
    i8e01bb5c870c8814a57f4218191f03102cf349429197bc489647fe16cd826967 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/security"
    ic7df6e6b923e174f81f252602787c09e4b873c83caf7076419b134f689cd7ccf "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/jobs"
    idb23fff012d069ba3f0f450171203c6db046b80f728a31ff1fb60291a92e225e "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/tasks"
    if8eaade36db4af52e4eb64165ed82f1064a844854820a23f730905e404c0aa3a "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/tools"
)

// DataFlowClient the main entry point of the SDK, exposes the configuration and the fluent API.
type DataFlowClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataFlowClientDataFlowClientDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataFlowClientDataFlowClientGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataFlowClientDataFlowClientHeadRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientHeadRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataFlowClientDataFlowClientOptionsRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientOptionsRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataFlowClientDataFlowClientPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataFlowClientDataFlowClientPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataFlowClientDataFlowClientPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataFlowClientDataFlowClientPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// About the about property
func (m *DataFlowClient) About()(*i67934d0263c4586557b7c0ac1f884137bbd3b32ac2b948c626329772767c4319.AboutRequestBuilder) {
    return i67934d0263c4586557b7c0ac1f884137bbd3b32ac2b948c626329772767c4319.NewAboutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
func (m *DataFlowClient) Apps()(*i8c8b06b621e92b0c16a6e2ba094614a310f4137ec10df70708b307bdcda39dae.AppsRequestBuilder) {
    return i8c8b06b621e92b0c16a6e2ba094614a310f4137ec10df70708b307bdcda39dae.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditRecords the auditRecords property
func (m *DataFlowClient) AuditRecords()(*i768ac0037ddaab39dabfdf36fb1e47c151a3840faaf879f661c819335f83c38f.AuditRecordsRequestBuilder) {
    return i768ac0037ddaab39dabfdf36fb1e47c151a3840faaf879f661c819335f83c38f.NewAuditRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Completions the completions property
func (m *DataFlowClient) Completions()(*i34b8e2fd66099266db79b74496d3c3e30853ae055a475f89902b7f096e86f569.CompletionsRequestBuilder) {
    return i34b8e2fd66099266db79b74496d3c3e30853ae055a475f89902b7f096e86f569.NewCompletionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDataFlowClient instantiates a new DataFlowClient and sets the default values.
func NewDataFlowClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataFlowClient) {
    m := &DataFlowClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("http://localhost:9393")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
func (m *DataFlowClient) Delete(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
func (m *DataFlowClient) Get(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientGetRequestConfiguration)([]byte, error) {
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
func (m *DataFlowClient) Head(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientHeadRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToHeadRequestInformation(ctx, requestConfiguration);
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
// Jobs the jobs property
func (m *DataFlowClient) Jobs()(*ic7df6e6b923e174f81f252602787c09e4b873c83caf7076419b134f689cd7ccf.JobsRequestBuilder) {
    return ic7df6e6b923e174f81f252602787c09e4b873c83caf7076419b134f689cd7ccf.NewJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *DataFlowClient) Options(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientOptionsRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToOptionsRequestInformation(ctx, requestConfiguration);
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
func (m *DataFlowClient) Patch(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientPatchRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, requestConfiguration);
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
func (m *DataFlowClient) Post(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
func (m *DataFlowClient) Put(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
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
// Runtime the runtime property
func (m *DataFlowClient) Runtime()(*i7638ff68664772d482f8dc66c218408bc63d9f7cd492c2aac4a4dcd01182a0da.RuntimeRequestBuilder) {
    return i7638ff68664772d482f8dc66c218408bc63d9f7cd492c2aac4a4dcd01182a0da.NewRuntimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schema the schema property
func (m *DataFlowClient) Schema()(*i866027ee37135123c0585c891ec25e9426e4445edf81f525db97ef95414bccf5.SchemaRequestBuilder) {
    return i866027ee37135123c0585c891ec25e9426e4445edf81f525db97ef95414bccf5.NewSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Security the security property
func (m *DataFlowClient) Security()(*i8e01bb5c870c8814a57f4218191f03102cf349429197bc489647fe16cd826967.SecurityRequestBuilder) {
    return i8e01bb5c870c8814a57f4218191f03102cf349429197bc489647fe16cd826967.NewSecurityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Streams the streams property
func (m *DataFlowClient) Streams()(*i433d5dac3040ebc08b4bd12b2de02b126d91cadd6980a9aeef355cb63826d2ee.StreamsRequestBuilder) {
    return i433d5dac3040ebc08b4bd12b2de02b126d91cadd6980a9aeef355cb63826d2ee.NewStreamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks the tasks property
func (m *DataFlowClient) Tasks()(*idb23fff012d069ba3f0f450171203c6db046b80f728a31ff1fb60291a92e225e.TasksRequestBuilder) {
    return idb23fff012d069ba3f0f450171203c6db046b80f728a31ff1fb60291a92e225e.NewTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *DataFlowClient) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
func (m *DataFlowClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
func (m *DataFlowClient) ToHeadRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientHeadRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.HEAD, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// Tools the tools property
func (m *DataFlowClient) Tools()(*if8eaade36db4af52e4eb64165ed82f1064a844854820a23f730905e404c0aa3a.ToolsRequestBuilder) {
    return if8eaade36db4af52e4eb64165ed82f1064a844854820a23f730905e404c0aa3a.NewToolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *DataFlowClient) ToOptionsRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientOptionsRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.OPTIONS, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
func (m *DataFlowClient) ToPatchRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
func (m *DataFlowClient) ToPostRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
func (m *DataFlowClient) ToPutRequestInformation(ctx context.Context, requestConfiguration *DataFlowClientDataFlowClientPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
