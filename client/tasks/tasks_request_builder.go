package tasks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TasksRequestBuilder builds and executes requests for operations under \tasks
type TasksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewTasksRequestBuilderInternal instantiates a new TasksRequestBuilder and sets the default values.
func NewTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TasksRequestBuilder) {
    m := &TasksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks", pathParameters),
    }
    return m
}
// NewTasksRequestBuilder instantiates a new TasksRequestBuilder and sets the default values.
func NewTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TasksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTasksRequestBuilderInternal(urlParams, requestAdapter)
}
// Ctr the ctr property
func (m *TasksRequestBuilder) Ctr()(*CtrRequestBuilder) {
    return NewCtrRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Definitions the definitions property
func (m *TasksRequestBuilder) Definitions()(*DefinitionsRequestBuilder) {
    return NewDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Executions the executions property
func (m *TasksRequestBuilder) Executions()(*ExecutionsRequestBuilder) {
    return NewExecutionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Info the info property
func (m *TasksRequestBuilder) Info()(*InfoRequestBuilder) {
    return NewInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logs the logs property
func (m *TasksRequestBuilder) Logs()(*LogsRequestBuilder) {
    return NewLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Platforms the platforms property
func (m *TasksRequestBuilder) Platforms()(*PlatformsRequestBuilder) {
    return NewPlatformsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Validation the validation property
func (m *TasksRequestBuilder) Validation()(*ValidationRequestBuilder) {
    return NewValidationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
