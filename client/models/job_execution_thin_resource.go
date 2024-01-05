package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobExecutionThinResource 
type JobExecutionThinResource struct {
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
    // The instanceId property
    instanceId *int64
    // The jobId property
    jobId *int64
    // The jobParameters property
    jobParameters JobExecutionThinResource_jobParametersable
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
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The startTime property
    startTime *string
    // The status property
    status *JobExecutionThinResource_status
    // The stepExecutionCount property
    stepExecutionCount *int32
    // The stoppable property
    stoppable *bool
    // The taskExecutionId property
    taskExecutionId *int64
    // The timeZone property
    timeZone JobExecutionThinResource_timeZoneable
}
// NewJobExecutionThinResource instantiates a new JobExecutionThinResource and sets the default values.
func NewJobExecutionThinResource()(*JobExecutionThinResource) {
    m := &JobExecutionThinResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobExecutionThinResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobExecutionThinResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobExecutionThinResource(), nil
}
// GetAbandonable gets the abandonable property value. The abandonable property
func (m *JobExecutionThinResource) GetAbandonable()(*bool) {
    return m.abandonable
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecutionThinResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefined gets the defined property value. The defined property
func (m *JobExecutionThinResource) GetDefined()(*bool) {
    return m.defined
}
// GetDuration gets the duration property value. The duration property
func (m *JobExecutionThinResource) GetDuration()(*string) {
    return m.duration
}
// GetExecutionId gets the executionId property value. The executionId property
func (m *JobExecutionThinResource) GetExecutionId()(*int64) {
    return m.executionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobExecutionThinResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["instanceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceId(val)
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
        val, err := n.GetObjectValue(CreateJobExecutionThinResource_jobParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobParameters(val.(JobExecutionThinResource_jobParametersable))
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseJobExecutionThinResource_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*JobExecutionThinResource_status))
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
        val, err := n.GetObjectValue(CreateJobExecutionThinResource_timeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val.(JobExecutionThinResource_timeZoneable))
        }
        return nil
    }
    return res
}
// GetInstanceId gets the instanceId property value. The instanceId property
func (m *JobExecutionThinResource) GetInstanceId()(*int64) {
    return m.instanceId
}
// GetJobId gets the jobId property value. The jobId property
func (m *JobExecutionThinResource) GetJobId()(*int64) {
    return m.jobId
}
// GetJobParameters gets the jobParameters property value. The jobParameters property
func (m *JobExecutionThinResource) GetJobParameters()(JobExecutionThinResource_jobParametersable) {
    return m.jobParameters
}
// GetJobParametersString gets the jobParametersString property value. The jobParametersString property
func (m *JobExecutionThinResource) GetJobParametersString()(*string) {
    return m.jobParametersString
}
// GetLinks gets the links property value. The links property
func (m *JobExecutionThinResource) GetLinks()([]Linkable) {
    return m.links
}
// GetName gets the name property value. The name property
func (m *JobExecutionThinResource) GetName()(*string) {
    return m.name
}
// GetRestartable gets the restartable property value. The restartable property
func (m *JobExecutionThinResource) GetRestartable()(*bool) {
    return m.restartable
}
// GetSchemaTarget gets the schemaTarget property value. The schemaTarget property
func (m *JobExecutionThinResource) GetSchemaTarget()(*string) {
    return m.schemaTarget
}
// GetStartDate gets the startDate property value. The startDate property
func (m *JobExecutionThinResource) GetStartDate()(*string) {
    return m.startDate
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *JobExecutionThinResource) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStartTime gets the startTime property value. The startTime property
func (m *JobExecutionThinResource) GetStartTime()(*string) {
    return m.startTime
}
// GetStatus gets the status property value. The status property
func (m *JobExecutionThinResource) GetStatus()(*JobExecutionThinResource_status) {
    return m.status
}
// GetStepExecutionCount gets the stepExecutionCount property value. The stepExecutionCount property
func (m *JobExecutionThinResource) GetStepExecutionCount()(*int32) {
    return m.stepExecutionCount
}
// GetStoppable gets the stoppable property value. The stoppable property
func (m *JobExecutionThinResource) GetStoppable()(*bool) {
    return m.stoppable
}
// GetTaskExecutionId gets the taskExecutionId property value. The taskExecutionId property
func (m *JobExecutionThinResource) GetTaskExecutionId()(*int64) {
    return m.taskExecutionId
}
// GetTimeZone gets the timeZone property value. The timeZone property
func (m *JobExecutionThinResource) GetTimeZone()(JobExecutionThinResource_timeZoneable) {
    return m.timeZone
}
// Serialize serializes information the current object
func (m *JobExecutionThinResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt64Value("instanceId", m.GetInstanceId())
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
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *JobExecutionThinResource) SetAbandonable(value *bool)() {
    m.abandonable = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecutionThinResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefined sets the defined property value. The defined property
func (m *JobExecutionThinResource) SetDefined(value *bool)() {
    m.defined = value
}
// SetDuration sets the duration property value. The duration property
func (m *JobExecutionThinResource) SetDuration(value *string)() {
    m.duration = value
}
// SetExecutionId sets the executionId property value. The executionId property
func (m *JobExecutionThinResource) SetExecutionId(value *int64)() {
    m.executionId = value
}
// SetInstanceId sets the instanceId property value. The instanceId property
func (m *JobExecutionThinResource) SetInstanceId(value *int64)() {
    m.instanceId = value
}
// SetJobId sets the jobId property value. The jobId property
func (m *JobExecutionThinResource) SetJobId(value *int64)() {
    m.jobId = value
}
// SetJobParameters sets the jobParameters property value. The jobParameters property
func (m *JobExecutionThinResource) SetJobParameters(value JobExecutionThinResource_jobParametersable)() {
    m.jobParameters = value
}
// SetJobParametersString sets the jobParametersString property value. The jobParametersString property
func (m *JobExecutionThinResource) SetJobParametersString(value *string)() {
    m.jobParametersString = value
}
// SetLinks sets the links property value. The links property
func (m *JobExecutionThinResource) SetLinks(value []Linkable)() {
    m.links = value
}
// SetName sets the name property value. The name property
func (m *JobExecutionThinResource) SetName(value *string)() {
    m.name = value
}
// SetRestartable sets the restartable property value. The restartable property
func (m *JobExecutionThinResource) SetRestartable(value *bool)() {
    m.restartable = value
}
// SetSchemaTarget sets the schemaTarget property value. The schemaTarget property
func (m *JobExecutionThinResource) SetSchemaTarget(value *string)() {
    m.schemaTarget = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *JobExecutionThinResource) SetStartDate(value *string)() {
    m.startDate = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *JobExecutionThinResource) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStartTime sets the startTime property value. The startTime property
func (m *JobExecutionThinResource) SetStartTime(value *string)() {
    m.startTime = value
}
// SetStatus sets the status property value. The status property
func (m *JobExecutionThinResource) SetStatus(value *JobExecutionThinResource_status)() {
    m.status = value
}
// SetStepExecutionCount sets the stepExecutionCount property value. The stepExecutionCount property
func (m *JobExecutionThinResource) SetStepExecutionCount(value *int32)() {
    m.stepExecutionCount = value
}
// SetStoppable sets the stoppable property value. The stoppable property
func (m *JobExecutionThinResource) SetStoppable(value *bool)() {
    m.stoppable = value
}
// SetTaskExecutionId sets the taskExecutionId property value. The taskExecutionId property
func (m *JobExecutionThinResource) SetTaskExecutionId(value *int64)() {
    m.taskExecutionId = value
}
// SetTimeZone sets the timeZone property value. The timeZone property
func (m *JobExecutionThinResource) SetTimeZone(value JobExecutionThinResource_timeZoneable)() {
    m.timeZone = value
}
// JobExecutionThinResourceable 
type JobExecutionThinResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAbandonable()(*bool)
    GetDefined()(*bool)
    GetDuration()(*string)
    GetExecutionId()(*int64)
    GetInstanceId()(*int64)
    GetJobId()(*int64)
    GetJobParameters()(JobExecutionThinResource_jobParametersable)
    GetJobParametersString()(*string)
    GetLinks()([]Linkable)
    GetName()(*string)
    GetRestartable()(*bool)
    GetSchemaTarget()(*string)
    GetStartDate()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStartTime()(*string)
    GetStatus()(*JobExecutionThinResource_status)
    GetStepExecutionCount()(*int32)
    GetStoppable()(*bool)
    GetTaskExecutionId()(*int64)
    GetTimeZone()(JobExecutionThinResource_timeZoneable)
    SetAbandonable(value *bool)()
    SetDefined(value *bool)()
    SetDuration(value *string)()
    SetExecutionId(value *int64)()
    SetInstanceId(value *int64)()
    SetJobId(value *int64)()
    SetJobParameters(value JobExecutionThinResource_jobParametersable)()
    SetJobParametersString(value *string)()
    SetLinks(value []Linkable)()
    SetName(value *string)()
    SetRestartable(value *bool)()
    SetSchemaTarget(value *string)()
    SetStartDate(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStartTime(value *string)()
    SetStatus(value *JobExecutionThinResource_status)()
    SetStepExecutionCount(value *int32)()
    SetStoppable(value *bool)()
    SetTaskExecutionId(value *int64)()
    SetTimeZone(value JobExecutionThinResource_timeZoneable)()
}
