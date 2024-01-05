package tasks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CtrRequestBuilder builds and executes requests for operations under \tasks\ctr
type CtrRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewCtrRequestBuilderInternal instantiates a new CtrRequestBuilder and sets the default values.
func NewCtrRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CtrRequestBuilder) {
    m := &CtrRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/ctr", pathParameters),
    }
    return m
}
// NewCtrRequestBuilder instantiates a new CtrRequestBuilder and sets the default values.
func NewCtrRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CtrRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCtrRequestBuilderInternal(urlParams, requestAdapter)
}
// OptionsPath the optionsPath property
func (m *CtrRequestBuilder) OptionsPath()(*CtrOptionsRequestBuilder) {
    return NewCtrOptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
