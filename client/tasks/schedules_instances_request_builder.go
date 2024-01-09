package tasks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SchedulesInstancesRequestBuilder builds and executes requests for operations under \tasks\schedules\instances
type SchedulesInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTaskDefinitionName gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.tasks.schedules.instances.item collection
func (m *SchedulesInstancesRequestBuilder) ByTaskDefinitionName(taskDefinitionName string)(*SchedulesInstancesWithTaskDefinitionNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if taskDefinitionName != "" {
        urlTplParams["taskDefinitionName"] = taskDefinitionName
    }
    return NewSchedulesInstancesWithTaskDefinitionNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSchedulesInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewSchedulesInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchedulesInstancesRequestBuilder) {
    m := &SchedulesInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/schedules/instances", pathParameters),
    }
    return m
}
// NewSchedulesInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewSchedulesInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SchedulesInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchedulesInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
