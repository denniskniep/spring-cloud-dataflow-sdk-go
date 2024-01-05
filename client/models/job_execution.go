package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobExecution 
type JobExecution struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allFailureExceptions property
    allFailureExceptions []JobExecution_allFailureExceptionsable
    // The createTime property
    createTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The endTime property
    endTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The executionContext property
    executionContext ExecutionContextable
    // The exitStatus property
    exitStatus ExitStatusable
    // The failureExceptions property
    failureExceptions []JobExecution_failureExceptionsable
    // The id property
    id *int64
    // The jobConfigurationName property
    jobConfigurationName *string
    // The jobId property
    jobId *int64
    // The jobInstance property
    jobInstance JobInstanceable
    // The jobParameters property
    jobParameters JobParametersable
    // The lastUpdated property
    lastUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The running property
    running *bool
    // The startTime property
    startTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *JobExecution_status
    // The stepExecutions property
    stepExecutions []StepExecutionable
    // The stopping property
    stopping *bool
    // The version property
    version *int32
}
// NewJobExecution instantiates a new JobExecution and sets the default values.
func NewJobExecution()(*JobExecution) {
    m := &JobExecution{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobExecutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobExecutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobExecution(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecution) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllFailureExceptions gets the allFailureExceptions property value. The allFailureExceptions property
func (m *JobExecution) GetAllFailureExceptions()([]JobExecution_allFailureExceptionsable) {
    return m.allFailureExceptions
}
// GetCreateTime gets the createTime property value. The createTime property
func (m *JobExecution) GetCreateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createTime
}
// GetEndTime gets the endTime property value. The endTime property
func (m *JobExecution) GetEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endTime
}
// GetExecutionContext gets the executionContext property value. The executionContext property
func (m *JobExecution) GetExecutionContext()(ExecutionContextable) {
    return m.executionContext
}
// GetExitStatus gets the exitStatus property value. The exitStatus property
func (m *JobExecution) GetExitStatus()(ExitStatusable) {
    return m.exitStatus
}
// GetFailureExceptions gets the failureExceptions property value. The failureExceptions property
func (m *JobExecution) GetFailureExceptions()([]JobExecution_failureExceptionsable) {
    return m.failureExceptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobExecution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allFailureExceptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobExecution_allFailureExceptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobExecution_allFailureExceptionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobExecution_allFailureExceptionsable)
                }
            }
            m.SetAllFailureExceptions(res)
        }
        return nil
    }
    res["createTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateTime(val)
        }
        return nil
    }
    res["endTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["executionContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExecutionContextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionContext(val.(ExecutionContextable))
        }
        return nil
    }
    res["exitStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExitStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExitStatus(val.(ExitStatusable))
        }
        return nil
    }
    res["failureExceptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobExecution_failureExceptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobExecution_failureExceptionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobExecution_failureExceptionsable)
                }
            }
            m.SetFailureExceptions(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["jobConfigurationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobConfigurationName(val)
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
    res["jobInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobInstance(val.(JobInstanceable))
        }
        return nil
    }
    res["jobParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobParameters(val.(JobParametersable))
        }
        return nil
    }
    res["lastUpdated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdated(val)
        }
        return nil
    }
    res["running"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunning(val)
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseJobExecution_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*JobExecution_status))
        }
        return nil
    }
    res["stepExecutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStepExecutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StepExecutionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StepExecutionable)
                }
            }
            m.SetStepExecutions(res)
        }
        return nil
    }
    res["stopping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStopping(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *JobExecution) GetId()(*int64) {
    return m.id
}
// GetJobConfigurationName gets the jobConfigurationName property value. The jobConfigurationName property
func (m *JobExecution) GetJobConfigurationName()(*string) {
    return m.jobConfigurationName
}
// GetJobId gets the jobId property value. The jobId property
func (m *JobExecution) GetJobId()(*int64) {
    return m.jobId
}
// GetJobInstance gets the jobInstance property value. The jobInstance property
func (m *JobExecution) GetJobInstance()(JobInstanceable) {
    return m.jobInstance
}
// GetJobParameters gets the jobParameters property value. The jobParameters property
func (m *JobExecution) GetJobParameters()(JobParametersable) {
    return m.jobParameters
}
// GetLastUpdated gets the lastUpdated property value. The lastUpdated property
func (m *JobExecution) GetLastUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdated
}
// GetRunning gets the running property value. The running property
func (m *JobExecution) GetRunning()(*bool) {
    return m.running
}
// GetStartTime gets the startTime property value. The startTime property
func (m *JobExecution) GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startTime
}
// GetStatus gets the status property value. The status property
func (m *JobExecution) GetStatus()(*JobExecution_status) {
    return m.status
}
// GetStepExecutions gets the stepExecutions property value. The stepExecutions property
func (m *JobExecution) GetStepExecutions()([]StepExecutionable) {
    return m.stepExecutions
}
// GetStopping gets the stopping property value. The stopping property
func (m *JobExecution) GetStopping()(*bool) {
    return m.stopping
}
// GetVersion gets the version property value. The version property
func (m *JobExecution) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *JobExecution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllFailureExceptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllFailureExceptions()))
        for i, v := range m.GetAllFailureExceptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("allFailureExceptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createTime", m.GetCreateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("executionContext", m.GetExecutionContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("exitStatus", m.GetExitStatus())
        if err != nil {
            return err
        }
    }
    if m.GetFailureExceptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFailureExceptions()))
        for i, v := range m.GetFailureExceptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("failureExceptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jobConfigurationName", m.GetJobConfigurationName())
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
        err := writer.WriteObjectValue("jobInstance", m.GetJobInstance())
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
        err := writer.WriteTimeValue("lastUpdated", m.GetLastUpdated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("running", m.GetRunning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startTime", m.GetStartTime())
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
    if m.GetStepExecutions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStepExecutions()))
        for i, v := range m.GetStepExecutions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("stepExecutions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stopping", m.GetStopping())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("version", m.GetVersion())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecution) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllFailureExceptions sets the allFailureExceptions property value. The allFailureExceptions property
func (m *JobExecution) SetAllFailureExceptions(value []JobExecution_allFailureExceptionsable)() {
    m.allFailureExceptions = value
}
// SetCreateTime sets the createTime property value. The createTime property
func (m *JobExecution) SetCreateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createTime = value
}
// SetEndTime sets the endTime property value. The endTime property
func (m *JobExecution) SetEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endTime = value
}
// SetExecutionContext sets the executionContext property value. The executionContext property
func (m *JobExecution) SetExecutionContext(value ExecutionContextable)() {
    m.executionContext = value
}
// SetExitStatus sets the exitStatus property value. The exitStatus property
func (m *JobExecution) SetExitStatus(value ExitStatusable)() {
    m.exitStatus = value
}
// SetFailureExceptions sets the failureExceptions property value. The failureExceptions property
func (m *JobExecution) SetFailureExceptions(value []JobExecution_failureExceptionsable)() {
    m.failureExceptions = value
}
// SetId sets the id property value. The id property
func (m *JobExecution) SetId(value *int64)() {
    m.id = value
}
// SetJobConfigurationName sets the jobConfigurationName property value. The jobConfigurationName property
func (m *JobExecution) SetJobConfigurationName(value *string)() {
    m.jobConfigurationName = value
}
// SetJobId sets the jobId property value. The jobId property
func (m *JobExecution) SetJobId(value *int64)() {
    m.jobId = value
}
// SetJobInstance sets the jobInstance property value. The jobInstance property
func (m *JobExecution) SetJobInstance(value JobInstanceable)() {
    m.jobInstance = value
}
// SetJobParameters sets the jobParameters property value. The jobParameters property
func (m *JobExecution) SetJobParameters(value JobParametersable)() {
    m.jobParameters = value
}
// SetLastUpdated sets the lastUpdated property value. The lastUpdated property
func (m *JobExecution) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdated = value
}
// SetRunning sets the running property value. The running property
func (m *JobExecution) SetRunning(value *bool)() {
    m.running = value
}
// SetStartTime sets the startTime property value. The startTime property
func (m *JobExecution) SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startTime = value
}
// SetStatus sets the status property value. The status property
func (m *JobExecution) SetStatus(value *JobExecution_status)() {
    m.status = value
}
// SetStepExecutions sets the stepExecutions property value. The stepExecutions property
func (m *JobExecution) SetStepExecutions(value []StepExecutionable)() {
    m.stepExecutions = value
}
// SetStopping sets the stopping property value. The stopping property
func (m *JobExecution) SetStopping(value *bool)() {
    m.stopping = value
}
// SetVersion sets the version property value. The version property
func (m *JobExecution) SetVersion(value *int32)() {
    m.version = value
}
// JobExecutionable 
type JobExecutionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllFailureExceptions()([]JobExecution_allFailureExceptionsable)
    GetCreateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExecutionContext()(ExecutionContextable)
    GetExitStatus()(ExitStatusable)
    GetFailureExceptions()([]JobExecution_failureExceptionsable)
    GetId()(*int64)
    GetJobConfigurationName()(*string)
    GetJobId()(*int64)
    GetJobInstance()(JobInstanceable)
    GetJobParameters()(JobParametersable)
    GetLastUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRunning()(*bool)
    GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*JobExecution_status)
    GetStepExecutions()([]StepExecutionable)
    GetStopping()(*bool)
    GetVersion()(*int32)
    SetAllFailureExceptions(value []JobExecution_allFailureExceptionsable)()
    SetCreateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExecutionContext(value ExecutionContextable)()
    SetExitStatus(value ExitStatusable)()
    SetFailureExceptions(value []JobExecution_failureExceptionsable)()
    SetId(value *int64)()
    SetJobConfigurationName(value *string)()
    SetJobId(value *int64)()
    SetJobInstance(value JobInstanceable)()
    SetJobParameters(value JobParametersable)()
    SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRunning(value *bool)()
    SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *JobExecution_status)()
    SetStepExecutions(value []StepExecutionable)()
    SetStopping(value *bool)()
    SetVersion(value *int32)()
}
