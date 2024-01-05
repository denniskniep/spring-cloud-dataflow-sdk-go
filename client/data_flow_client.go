package client

import (
    "context"
    i014b946f8dad9d8ec144a17118dae574626d8a719fe4a60bafd10ff8d19f4b0d "dataflow/client/schema"
    i0e32efe3c55f6fbaa0a2c9b7455c84521faf0b458182924cffb7b091fae89599 "dataflow/client/security"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i437045e3100d51af2c928b60d9de692fea26495060c3dd1a6f60a71445818f82 "dataflow/client/streams"
    i49a1b207c844bca4071a1dd5ac03d692a2fa71fd85a36b1676e7250cb29adc4d "dataflow/client/tools"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i550ab9f0040019acf07a47c09c9b287dc54a668ff63e2fc2e907a3ce87be3103 "dataflow/client/auditrecords"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i6c9d14e3ee8ce9e9424f33dec6e18f5747f119b313f3a413998da0997ec704bc "dataflow/client/apps"
    i6caec42106ad331f292224fa28b6c03fe43751f46903823ba09a19ea6be056cb "dataflow/client/about"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i85d984cb264483fb9d50d64a3583cd1bc8972029b831e007633f4b388b7f20be "dataflow/client/tasks"
    ia3be9830440430a3684c579378118576201942e81e600a785865508b83c68bb9 "dataflow/client/runtime"
    ib302e7238b92266e421d105fc2f6f5191923e0493caf7fa41886db090b9f5f13 "dataflow/client/completions"
    ib645fed3c2110ebccf7329f4acd512b5ab8d12233fe3d114862596835d462ad0 "dataflow/client/jobs"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
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
func (m *DataFlowClient) About()(*i6caec42106ad331f292224fa28b6c03fe43751f46903823ba09a19ea6be056cb.AboutRequestBuilder) {
    return i6caec42106ad331f292224fa28b6c03fe43751f46903823ba09a19ea6be056cb.NewAboutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
func (m *DataFlowClient) Apps()(*i6c9d14e3ee8ce9e9424f33dec6e18f5747f119b313f3a413998da0997ec704bc.AppsRequestBuilder) {
    return i6c9d14e3ee8ce9e9424f33dec6e18f5747f119b313f3a413998da0997ec704bc.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditRecords the auditRecords property
func (m *DataFlowClient) AuditRecords()(*i550ab9f0040019acf07a47c09c9b287dc54a668ff63e2fc2e907a3ce87be3103.AuditRecordsRequestBuilder) {
    return i550ab9f0040019acf07a47c09c9b287dc54a668ff63e2fc2e907a3ce87be3103.NewAuditRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Completions the completions property
func (m *DataFlowClient) Completions()(*ib302e7238b92266e421d105fc2f6f5191923e0493caf7fa41886db090b9f5f13.CompletionsRequestBuilder) {
    return ib302e7238b92266e421d105fc2f6f5191923e0493caf7fa41886db090b9f5f13.NewCompletionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *DataFlowClient) Jobs()(*ib645fed3c2110ebccf7329f4acd512b5ab8d12233fe3d114862596835d462ad0.JobsRequestBuilder) {
    return ib645fed3c2110ebccf7329f4acd512b5ab8d12233fe3d114862596835d462ad0.NewJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *DataFlowClient) Runtime()(*ia3be9830440430a3684c579378118576201942e81e600a785865508b83c68bb9.RuntimeRequestBuilder) {
    return ia3be9830440430a3684c579378118576201942e81e600a785865508b83c68bb9.NewRuntimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schema the schema property
func (m *DataFlowClient) Schema()(*i014b946f8dad9d8ec144a17118dae574626d8a719fe4a60bafd10ff8d19f4b0d.SchemaRequestBuilder) {
    return i014b946f8dad9d8ec144a17118dae574626d8a719fe4a60bafd10ff8d19f4b0d.NewSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Security the security property
func (m *DataFlowClient) Security()(*i0e32efe3c55f6fbaa0a2c9b7455c84521faf0b458182924cffb7b091fae89599.SecurityRequestBuilder) {
    return i0e32efe3c55f6fbaa0a2c9b7455c84521faf0b458182924cffb7b091fae89599.NewSecurityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Streams the streams property
func (m *DataFlowClient) Streams()(*i437045e3100d51af2c928b60d9de692fea26495060c3dd1a6f60a71445818f82.StreamsRequestBuilder) {
    return i437045e3100d51af2c928b60d9de692fea26495060c3dd1a6f60a71445818f82.NewStreamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks the tasks property
func (m *DataFlowClient) Tasks()(*i85d984cb264483fb9d50d64a3583cd1bc8972029b831e007633f4b388b7f20be.TasksRequestBuilder) {
    return i85d984cb264483fb9d50d64a3583cd1bc8972029b831e007633f4b388b7f20be.NewTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *DataFlowClient) Tools()(*i49a1b207c844bca4071a1dd5ac03d692a2fa71fd85a36b1676e7250cb29adc4d.ToolsRequestBuilder) {
    return i49a1b207c844bca4071a1dd5ac03d692a2fa71fd85a36b1676e7250cb29adc4d.NewToolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
