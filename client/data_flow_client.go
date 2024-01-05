package client

import (
    "context"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i1097e231fe4d6d044e0b4832829147ef85fc75d1b96551da1c704e06892a3f5c "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/auditrecords"
    i579d79eb321ef938fbd487e854e61eaea08c2ae4d31fd0dffc3f41283cfb5250 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/tasks"
    i57c28cfda18dcf5bbb2202ea2355c12eb642c3f5494cb5d5a9882c49c97055ca "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/jobs"
    i5ce27d5b249357bb9efc5d8e6538d85ee7dd872223cb12208da98e76d81a5118 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/runtime"
    i653664a6775d7e19eaef36749f9b5c3419d9800079e2328ba53d2b195492847b "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/schema"
    i77c622c0bccd95729eb4344c8d1a1d844ae3ae68b0a5935dfeca1255bd00c3ac "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/apps"
    i96607ccd2027d970b5c219f59e693b1f8ad1ca62be0954e1bc12259aa9a982b1 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/about"
    ib1261bf55cac3763431209246ce343a8dbab7c0aa86ca4621d13a01ad57c23e0 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/security"
    ib2a070990fff214db7736622ba849386240da42ae4d0ae4a902d0edd100e391e "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/completions"
    ic6f7762fc45e264230ef6aaea5d9f07e68db12149f8fb73f6c33e10288591e4a "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/streams"
    ifec7978daa52ce7ad7c7c0c6679a9bdb9c60c48bdc067538b01ea7d21a0734f7 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client/tools"
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
func (m *DataFlowClient) About()(*i96607ccd2027d970b5c219f59e693b1f8ad1ca62be0954e1bc12259aa9a982b1.AboutRequestBuilder) {
    return i96607ccd2027d970b5c219f59e693b1f8ad1ca62be0954e1bc12259aa9a982b1.NewAboutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
func (m *DataFlowClient) Apps()(*i77c622c0bccd95729eb4344c8d1a1d844ae3ae68b0a5935dfeca1255bd00c3ac.AppsRequestBuilder) {
    return i77c622c0bccd95729eb4344c8d1a1d844ae3ae68b0a5935dfeca1255bd00c3ac.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditRecords the auditRecords property
func (m *DataFlowClient) AuditRecords()(*i1097e231fe4d6d044e0b4832829147ef85fc75d1b96551da1c704e06892a3f5c.AuditRecordsRequestBuilder) {
    return i1097e231fe4d6d044e0b4832829147ef85fc75d1b96551da1c704e06892a3f5c.NewAuditRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Completions the completions property
func (m *DataFlowClient) Completions()(*ib2a070990fff214db7736622ba849386240da42ae4d0ae4a902d0edd100e391e.CompletionsRequestBuilder) {
    return ib2a070990fff214db7736622ba849386240da42ae4d0ae4a902d0edd100e391e.NewCompletionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *DataFlowClient) Jobs()(*i57c28cfda18dcf5bbb2202ea2355c12eb642c3f5494cb5d5a9882c49c97055ca.JobsRequestBuilder) {
    return i57c28cfda18dcf5bbb2202ea2355c12eb642c3f5494cb5d5a9882c49c97055ca.NewJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *DataFlowClient) Runtime()(*i5ce27d5b249357bb9efc5d8e6538d85ee7dd872223cb12208da98e76d81a5118.RuntimeRequestBuilder) {
    return i5ce27d5b249357bb9efc5d8e6538d85ee7dd872223cb12208da98e76d81a5118.NewRuntimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schema the schema property
func (m *DataFlowClient) Schema()(*i653664a6775d7e19eaef36749f9b5c3419d9800079e2328ba53d2b195492847b.SchemaRequestBuilder) {
    return i653664a6775d7e19eaef36749f9b5c3419d9800079e2328ba53d2b195492847b.NewSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Security the security property
func (m *DataFlowClient) Security()(*ib1261bf55cac3763431209246ce343a8dbab7c0aa86ca4621d13a01ad57c23e0.SecurityRequestBuilder) {
    return ib1261bf55cac3763431209246ce343a8dbab7c0aa86ca4621d13a01ad57c23e0.NewSecurityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Streams the streams property
func (m *DataFlowClient) Streams()(*ic6f7762fc45e264230ef6aaea5d9f07e68db12149f8fb73f6c33e10288591e4a.StreamsRequestBuilder) {
    return ic6f7762fc45e264230ef6aaea5d9f07e68db12149f8fb73f6c33e10288591e4a.NewStreamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks the tasks property
func (m *DataFlowClient) Tasks()(*i579d79eb321ef938fbd487e854e61eaea08c2ae4d31fd0dffc3f41283cfb5250.TasksRequestBuilder) {
    return i579d79eb321ef938fbd487e854e61eaea08c2ae4d31fd0dffc3f41283cfb5250.NewTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *DataFlowClient) Tools()(*ifec7978daa52ce7ad7c7c0c6679a9bdb9c60c48bdc067538b01ea7d21a0734f7.ToolsRequestBuilder) {
    return ifec7978daa52ce7ad7c7c0c6679a9bdb9c60c48bdc067538b01ea7d21a0734f7.NewToolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
