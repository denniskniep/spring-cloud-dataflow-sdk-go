package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StepExecution 
type StepExecution struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commitCount property
    commitCount *int32
    // The endTime property
    endTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The executionContext property
    executionContext ExecutionContextable
    // The exitStatus property
    exitStatus ExitStatusable
    // The failureExceptions property
    failureExceptions []StepExecution_failureExceptionsable
    // The filterCount property
    filterCount *int32
    // The id property
    id *int64
    // The jobExecution property
    jobExecution JobExecutionable
    // The jobExecutionId property
    jobExecutionId *int64
    // The jobParameters property
    jobParameters JobParametersable
    // The lastUpdated property
    lastUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The processSkipCount property
    processSkipCount *int32
    // The readCount property
    readCount *int32
    // The readSkipCount property
    readSkipCount *int32
    // The rollbackCount property
    rollbackCount *int32
    // The skipCount property
    skipCount *int32
    // The startTime property
    startTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *StepExecution_status
    // The stepName property
    stepName *string
    // The summary property
    summary *string
    // The terminateOnly property
    terminateOnly *bool
    // The version property
    version *int32
    // The writeCount property
    writeCount *int32
    // The writeSkipCount property
    writeSkipCount *int32
}
// NewStepExecution instantiates a new StepExecution and sets the default values.
func NewStepExecution()(*StepExecution) {
    m := &StepExecution{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStepExecutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStepExecutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStepExecution(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StepExecution) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommitCount gets the commitCount property value. The commitCount property
func (m *StepExecution) GetCommitCount()(*int32) {
    return m.commitCount
}
// GetEndTime gets the endTime property value. The endTime property
func (m *StepExecution) GetEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endTime
}
// GetExecutionContext gets the executionContext property value. The executionContext property
func (m *StepExecution) GetExecutionContext()(ExecutionContextable) {
    return m.executionContext
}
// GetExitStatus gets the exitStatus property value. The exitStatus property
func (m *StepExecution) GetExitStatus()(ExitStatusable) {
    return m.exitStatus
}
// GetFailureExceptions gets the failureExceptions property value. The failureExceptions property
func (m *StepExecution) GetFailureExceptions()([]StepExecution_failureExceptionsable) {
    return m.failureExceptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StepExecution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commitCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitCount(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateStepExecution_failureExceptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StepExecution_failureExceptionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StepExecution_failureExceptionsable)
                }
            }
            m.SetFailureExceptions(res)
        }
        return nil
    }
    res["filterCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterCount(val)
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
    res["jobExecutionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobExecutionId(val)
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
    res["processSkipCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessSkipCount(val)
        }
        return nil
    }
    res["readCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadCount(val)
        }
        return nil
    }
    res["readSkipCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadSkipCount(val)
        }
        return nil
    }
    res["rollbackCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRollbackCount(val)
        }
        return nil
    }
    res["skipCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipCount(val)
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
        val, err := n.GetEnumValue(ParseStepExecution_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*StepExecution_status))
        }
        return nil
    }
    res["stepName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStepName(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    res["terminateOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTerminateOnly(val)
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
    res["writeCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWriteCount(val)
        }
        return nil
    }
    res["writeSkipCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWriteSkipCount(val)
        }
        return nil
    }
    return res
}
// GetFilterCount gets the filterCount property value. The filterCount property
func (m *StepExecution) GetFilterCount()(*int32) {
    return m.filterCount
}
// GetId gets the id property value. The id property
func (m *StepExecution) GetId()(*int64) {
    return m.id
}
// GetJobExecution gets the jobExecution property value. The jobExecution property
func (m *StepExecution) GetJobExecution()(JobExecutionable) {
    return m.jobExecution
}
// GetJobExecutionId gets the jobExecutionId property value. The jobExecutionId property
func (m *StepExecution) GetJobExecutionId()(*int64) {
    return m.jobExecutionId
}
// GetJobParameters gets the jobParameters property value. The jobParameters property
func (m *StepExecution) GetJobParameters()(JobParametersable) {
    return m.jobParameters
}
// GetLastUpdated gets the lastUpdated property value. The lastUpdated property
func (m *StepExecution) GetLastUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdated
}
// GetProcessSkipCount gets the processSkipCount property value. The processSkipCount property
func (m *StepExecution) GetProcessSkipCount()(*int32) {
    return m.processSkipCount
}
// GetReadCount gets the readCount property value. The readCount property
func (m *StepExecution) GetReadCount()(*int32) {
    return m.readCount
}
// GetReadSkipCount gets the readSkipCount property value. The readSkipCount property
func (m *StepExecution) GetReadSkipCount()(*int32) {
    return m.readSkipCount
}
// GetRollbackCount gets the rollbackCount property value. The rollbackCount property
func (m *StepExecution) GetRollbackCount()(*int32) {
    return m.rollbackCount
}
// GetSkipCount gets the skipCount property value. The skipCount property
func (m *StepExecution) GetSkipCount()(*int32) {
    return m.skipCount
}
// GetStartTime gets the startTime property value. The startTime property
func (m *StepExecution) GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startTime
}
// GetStatus gets the status property value. The status property
func (m *StepExecution) GetStatus()(*StepExecution_status) {
    return m.status
}
// GetStepName gets the stepName property value. The stepName property
func (m *StepExecution) GetStepName()(*string) {
    return m.stepName
}
// GetSummary gets the summary property value. The summary property
func (m *StepExecution) GetSummary()(*string) {
    return m.summary
}
// GetTerminateOnly gets the terminateOnly property value. The terminateOnly property
func (m *StepExecution) GetTerminateOnly()(*bool) {
    return m.terminateOnly
}
// GetVersion gets the version property value. The version property
func (m *StepExecution) GetVersion()(*int32) {
    return m.version
}
// GetWriteCount gets the writeCount property value. The writeCount property
func (m *StepExecution) GetWriteCount()(*int32) {
    return m.writeCount
}
// GetWriteSkipCount gets the writeSkipCount property value. The writeSkipCount property
func (m *StepExecution) GetWriteSkipCount()(*int32) {
    return m.writeSkipCount
}
// Serialize serializes information the current object
func (m *StepExecution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("commitCount", m.GetCommitCount())
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
        err := writer.WriteInt32Value("filterCount", m.GetFilterCount())
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
        err := writer.WriteObjectValue("jobExecution", m.GetJobExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("jobExecutionId", m.GetJobExecutionId())
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
        err := writer.WriteInt32Value("processSkipCount", m.GetProcessSkipCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("readCount", m.GetReadCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("readSkipCount", m.GetReadSkipCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rollbackCount", m.GetRollbackCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skipCount", m.GetSkipCount())
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
    {
        err := writer.WriteStringValue("stepName", m.GetStepName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("terminateOnly", m.GetTerminateOnly())
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
        err := writer.WriteInt32Value("writeCount", m.GetWriteCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("writeSkipCount", m.GetWriteSkipCount())
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
func (m *StepExecution) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommitCount sets the commitCount property value. The commitCount property
func (m *StepExecution) SetCommitCount(value *int32)() {
    m.commitCount = value
}
// SetEndTime sets the endTime property value. The endTime property
func (m *StepExecution) SetEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endTime = value
}
// SetExecutionContext sets the executionContext property value. The executionContext property
func (m *StepExecution) SetExecutionContext(value ExecutionContextable)() {
    m.executionContext = value
}
// SetExitStatus sets the exitStatus property value. The exitStatus property
func (m *StepExecution) SetExitStatus(value ExitStatusable)() {
    m.exitStatus = value
}
// SetFailureExceptions sets the failureExceptions property value. The failureExceptions property
func (m *StepExecution) SetFailureExceptions(value []StepExecution_failureExceptionsable)() {
    m.failureExceptions = value
}
// SetFilterCount sets the filterCount property value. The filterCount property
func (m *StepExecution) SetFilterCount(value *int32)() {
    m.filterCount = value
}
// SetId sets the id property value. The id property
func (m *StepExecution) SetId(value *int64)() {
    m.id = value
}
// SetJobExecution sets the jobExecution property value. The jobExecution property
func (m *StepExecution) SetJobExecution(value JobExecutionable)() {
    m.jobExecution = value
}
// SetJobExecutionId sets the jobExecutionId property value. The jobExecutionId property
func (m *StepExecution) SetJobExecutionId(value *int64)() {
    m.jobExecutionId = value
}
// SetJobParameters sets the jobParameters property value. The jobParameters property
func (m *StepExecution) SetJobParameters(value JobParametersable)() {
    m.jobParameters = value
}
// SetLastUpdated sets the lastUpdated property value. The lastUpdated property
func (m *StepExecution) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdated = value
}
// SetProcessSkipCount sets the processSkipCount property value. The processSkipCount property
func (m *StepExecution) SetProcessSkipCount(value *int32)() {
    m.processSkipCount = value
}
// SetReadCount sets the readCount property value. The readCount property
func (m *StepExecution) SetReadCount(value *int32)() {
    m.readCount = value
}
// SetReadSkipCount sets the readSkipCount property value. The readSkipCount property
func (m *StepExecution) SetReadSkipCount(value *int32)() {
    m.readSkipCount = value
}
// SetRollbackCount sets the rollbackCount property value. The rollbackCount property
func (m *StepExecution) SetRollbackCount(value *int32)() {
    m.rollbackCount = value
}
// SetSkipCount sets the skipCount property value. The skipCount property
func (m *StepExecution) SetSkipCount(value *int32)() {
    m.skipCount = value
}
// SetStartTime sets the startTime property value. The startTime property
func (m *StepExecution) SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startTime = value
}
// SetStatus sets the status property value. The status property
func (m *StepExecution) SetStatus(value *StepExecution_status)() {
    m.status = value
}
// SetStepName sets the stepName property value. The stepName property
func (m *StepExecution) SetStepName(value *string)() {
    m.stepName = value
}
// SetSummary sets the summary property value. The summary property
func (m *StepExecution) SetSummary(value *string)() {
    m.summary = value
}
// SetTerminateOnly sets the terminateOnly property value. The terminateOnly property
func (m *StepExecution) SetTerminateOnly(value *bool)() {
    m.terminateOnly = value
}
// SetVersion sets the version property value. The version property
func (m *StepExecution) SetVersion(value *int32)() {
    m.version = value
}
// SetWriteCount sets the writeCount property value. The writeCount property
func (m *StepExecution) SetWriteCount(value *int32)() {
    m.writeCount = value
}
// SetWriteSkipCount sets the writeSkipCount property value. The writeSkipCount property
func (m *StepExecution) SetWriteSkipCount(value *int32)() {
    m.writeSkipCount = value
}
// StepExecutionable 
type StepExecutionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommitCount()(*int32)
    GetEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExecutionContext()(ExecutionContextable)
    GetExitStatus()(ExitStatusable)
    GetFailureExceptions()([]StepExecution_failureExceptionsable)
    GetFilterCount()(*int32)
    GetId()(*int64)
    GetJobExecution()(JobExecutionable)
    GetJobExecutionId()(*int64)
    GetJobParameters()(JobParametersable)
    GetLastUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProcessSkipCount()(*int32)
    GetReadCount()(*int32)
    GetReadSkipCount()(*int32)
    GetRollbackCount()(*int32)
    GetSkipCount()(*int32)
    GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*StepExecution_status)
    GetStepName()(*string)
    GetSummary()(*string)
    GetTerminateOnly()(*bool)
    GetVersion()(*int32)
    GetWriteCount()(*int32)
    GetWriteSkipCount()(*int32)
    SetCommitCount(value *int32)()
    SetEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExecutionContext(value ExecutionContextable)()
    SetExitStatus(value ExitStatusable)()
    SetFailureExceptions(value []StepExecution_failureExceptionsable)()
    SetFilterCount(value *int32)()
    SetId(value *int64)()
    SetJobExecution(value JobExecutionable)()
    SetJobExecutionId(value *int64)()
    SetJobParameters(value JobParametersable)()
    SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProcessSkipCount(value *int32)()
    SetReadCount(value *int32)()
    SetReadSkipCount(value *int32)()
    SetRollbackCount(value *int32)()
    SetSkipCount(value *int32)()
    SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *StepExecution_status)()
    SetStepName(value *string)()
    SetSummary(value *string)()
    SetTerminateOnly(value *bool)()
    SetVersion(value *int32)()
    SetWriteCount(value *int32)()
    SetWriteSkipCount(value *int32)()
}
