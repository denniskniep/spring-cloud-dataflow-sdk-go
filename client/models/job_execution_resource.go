package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobExecutionResource 
type JobExecutionResource struct {
    // The abandonable property
    abandonable *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The defined property
    defined *bool
    // The duration property
    duration *string
    // The executionId property
    executionId *int64
    // The jobExecution property
    jobExecution JobExecutionable
    // The jobId property
    jobId *int64
    // The jobParameters property
    jobParameters JobExecutionResource_jobParametersable
    // The jobParametersString property
    jobParametersString *string
    // The links property
    links []Linkable
    // The name property
    name *string
    // The restartable property
    restartable *bool
    // The schemaTarget property
    schemaTarget *string
    // The startDate property
    startDate *string
    // The startTime property
    startTime *string
    // The stepExecutionCount property
    stepExecutionCount *int32
    // The stoppable property
    stoppable *bool
    // The taskExecutionId property
    taskExecutionId *int64
    // The timeZone property
    timeZone JobExecutionResource_timeZoneable
}
// NewJobExecutionResource instantiates a new JobExecutionResource and sets the default values.
func NewJobExecutionResource()(*JobExecutionResource) {
    m := &JobExecutionResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobExecutionResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobExecutionResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobExecutionResource(), nil
}
// GetAbandonable gets the abandonable property value. The abandonable property
func (m *JobExecutionResource) GetAbandonable()(*bool) {
    return m.abandonable
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecutionResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefined gets the defined property value. The defined property
func (m *JobExecutionResource) GetDefined()(*bool) {
    return m.defined
}
// GetDuration gets the duration property value. The duration property
func (m *JobExecutionResource) GetDuration()(*string) {
    return m.duration
}
// GetExecutionId gets the executionId property value. The executionId property
func (m *JobExecutionResource) GetExecutionId()(*int64) {
    return m.executionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobExecutionResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["abandonable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAbandonable(val)
        }
        return nil
    }
    res["defined"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefined(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["executionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionId(val)
        }
        return nil
    }
    res["jobExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobExecutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobExecution(val.(JobExecutionable))
        }
        return nil
    }
    res["jobId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobId(val)
        }
        return nil
    }
    res["jobParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobExecutionResource_jobParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobParameters(val.(JobExecutionResource_jobParametersable))
        }
        return nil
    }
    res["jobParametersString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobParametersString(val)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Linkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Linkable)
                }
            }
            m.SetLinks(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["restartable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartable(val)
        }
        return nil
    }
    res["schemaTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaTarget(val)
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["stepExecutionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStepExecutionCount(val)
        }
        return nil
    }
    res["stoppable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStoppable(val)
        }
        return nil
    }
    res["taskExecutionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskExecutionId(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobExecutionResource_timeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val.(JobExecutionResource_timeZoneable))
        }
        return nil
    }
    return res
}
// GetJobExecution gets the jobExecution property value. The jobExecution property
func (m *JobExecutionResource) GetJobExecution()(JobExecutionable) {
    return m.jobExecution
}
// GetJobId gets the jobId property value. The jobId property
func (m *JobExecutionResource) GetJobId()(*int64) {
    return m.jobId
}
// GetJobParameters gets the jobParameters property value. The jobParameters property
func (m *JobExecutionResource) GetJobParameters()(JobExecutionResource_jobParametersable) {
    return m.jobParameters
}
// GetJobParametersString gets the jobParametersString property value. The jobParametersString property
func (m *JobExecutionResource) GetJobParametersString()(*string) {
    return m.jobParametersString
}
// GetLinks gets the links property value. The links property
func (m *JobExecutionResource) GetLinks()([]Linkable) {
    return m.links
}
// GetName gets the name property value. The name property
func (m *JobExecutionResource) GetName()(*string) {
    return m.name
}
// GetRestartable gets the restartable property value. The restartable property
func (m *JobExecutionResource) GetRestartable()(*bool) {
    return m.restartable
}
// GetSchemaTarget gets the schemaTarget property value. The schemaTarget property
func (m *JobExecutionResource) GetSchemaTarget()(*string) {
    return m.schemaTarget
}
// GetStartDate gets the startDate property value. The startDate property
func (m *JobExecutionResource) GetStartDate()(*string) {
    return m.startDate
}
// GetStartTime gets the startTime property value. The startTime property
func (m *JobExecutionResource) GetStartTime()(*string) {
    return m.startTime
}
// GetStepExecutionCount gets the stepExecutionCount property value. The stepExecutionCount property
func (m *JobExecutionResource) GetStepExecutionCount()(*int32) {
    return m.stepExecutionCount
}
// GetStoppable gets the stoppable property value. The stoppable property
func (m *JobExecutionResource) GetStoppable()(*bool) {
    return m.stoppable
}
// GetTaskExecutionId gets the taskExecutionId property value. The taskExecutionId property
func (m *JobExecutionResource) GetTaskExecutionId()(*int64) {
    return m.taskExecutionId
}
// GetTimeZone gets the timeZone property value. The timeZone property
func (m *JobExecutionResource) GetTimeZone()(JobExecutionResource_timeZoneable) {
    return m.timeZone
}
// Serialize serializes information the current object
func (m *JobExecutionResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("abandonable", m.GetAbandonable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("defined", m.GetDefined())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("executionId", m.GetExecutionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("jobExecution", m.GetJobExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("jobId", m.GetJobId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("jobParameters", m.GetJobParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobParametersString", m.GetJobParametersString())
        if err != nil {
            return err
        }
    }
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("restartable", m.GetRestartable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaTarget", m.GetSchemaTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startTime", m.GetStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("stepExecutionCount", m.GetStepExecutionCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stoppable", m.GetStoppable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("taskExecutionId", m.GetTaskExecutionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAbandonable sets the abandonable property value. The abandonable property
func (m *JobExecutionResource) SetAbandonable(value *bool)() {
    m.abandonable = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecutionResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefined sets the defined property value. The defined property
func (m *JobExecutionResource) SetDefined(value *bool)() {
    m.defined = value
}
// SetDuration sets the duration property value. The duration property
func (m *JobExecutionResource) SetDuration(value *string)() {
    m.duration = value
}
// SetExecutionId sets the executionId property value. The executionId property
func (m *JobExecutionResource) SetExecutionId(value *int64)() {
    m.executionId = value
}
// SetJobExecution sets the jobExecution property value. The jobExecution property
func (m *JobExecutionResource) SetJobExecution(value JobExecutionable)() {
    m.jobExecution = value
}
// SetJobId sets the jobId property value. The jobId property
func (m *JobExecutionResource) SetJobId(value *int64)() {
    m.jobId = value
}
// SetJobParameters sets the jobParameters property value. The jobParameters property
func (m *JobExecutionResource) SetJobParameters(value JobExecutionResource_jobParametersable)() {
    m.jobParameters = value
}
// SetJobParametersString sets the jobParametersString property value. The jobParametersString property
func (m *JobExecutionResource) SetJobParametersString(value *string)() {
    m.jobParametersString = value
}
// SetLinks sets the links property value. The links property
func (m *JobExecutionResource) SetLinks(value []Linkable)() {
    m.links = value
}
// SetName sets the name property value. The name property
func (m *JobExecutionResource) SetName(value *string)() {
    m.name = value
}
// SetRestartable sets the restartable property value. The restartable property
func (m *JobExecutionResource) SetRestartable(value *bool)() {
    m.restartable = value
}
// SetSchemaTarget sets the schemaTarget property value. The schemaTarget property
func (m *JobExecutionResource) SetSchemaTarget(value *string)() {
    m.schemaTarget = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *JobExecutionResource) SetStartDate(value *string)() {
    m.startDate = value
}
// SetStartTime sets the startTime property value. The startTime property
func (m *JobExecutionResource) SetStartTime(value *string)() {
    m.startTime = value
}
// SetStepExecutionCount sets the stepExecutionCount property value. The stepExecutionCount property
func (m *JobExecutionResource) SetStepExecutionCount(value *int32)() {
    m.stepExecutionCount = value
}
// SetStoppable sets the stoppable property value. The stoppable property
func (m *JobExecutionResource) SetStoppable(value *bool)() {
    m.stoppable = value
}
// SetTaskExecutionId sets the taskExecutionId property value. The taskExecutionId property
func (m *JobExecutionResource) SetTaskExecutionId(value *int64)() {
    m.taskExecutionId = value
}
// SetTimeZone sets the timeZone property value. The timeZone property
func (m *JobExecutionResource) SetTimeZone(value JobExecutionResource_timeZoneable)() {
    m.timeZone = value
}
// JobExecutionResourceable 
type JobExecutionResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAbandonable()(*bool)
    GetDefined()(*bool)
    GetDuration()(*string)
    GetExecutionId()(*int64)
    GetJobExecution()(JobExecutionable)
    GetJobId()(*int64)
    GetJobParameters()(JobExecutionResource_jobParametersable)
    GetJobParametersString()(*string)
    GetLinks()([]Linkable)
    GetName()(*string)
    GetRestartable()(*bool)
    GetSchemaTarget()(*string)
    GetStartDate()(*string)
    GetStartTime()(*string)
    GetStepExecutionCount()(*int32)
    GetStoppable()(*bool)
    GetTaskExecutionId()(*int64)
    GetTimeZone()(JobExecutionResource_timeZoneable)
    SetAbandonable(value *bool)()
    SetDefined(value *bool)()
    SetDuration(value *string)()
    SetExecutionId(value *int64)()
    SetJobExecution(value JobExecutionable)()
    SetJobId(value *int64)()
    SetJobParameters(value JobExecutionResource_jobParametersable)()
    SetJobParametersString(value *string)()
    SetLinks(value []Linkable)()
    SetName(value *string)()
    SetRestartable(value *bool)()
    SetSchemaTarget(value *string)()
    SetStartDate(value *string)()
    SetStartTime(value *string)()
    SetStepExecutionCount(value *int32)()
    SetStoppable(value *bool)()
    SetTaskExecutionId(value *int64)()
    SetTimeZone(value JobExecutionResource_timeZoneable)()
}
