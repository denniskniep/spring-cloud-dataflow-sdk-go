package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobInstance 
type JobInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int64
    // The instanceId property
    instanceId *int64
    // The jobName property
    jobName *string
    // The version property
    version *int32
}
// NewJobInstance instantiates a new JobInstance and sets the default values.
func NewJobInstance()(*JobInstance) {
    m := &JobInstance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobInstance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobInstance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["jobName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobName(val)
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
func (m *JobInstance) GetId()(*int64) {
    return m.id
}
// GetInstanceId gets the instanceId property value. The instanceId property
func (m *JobInstance) GetInstanceId()(*int64) {
    return m.instanceId
}
// GetJobName gets the jobName property value. The jobName property
func (m *JobInstance) GetJobName()(*string) {
    return m.jobName
}
// GetVersion gets the version property value. The version property
func (m *JobInstance) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *JobInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteStringValue("jobName", m.GetJobName())
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
func (m *JobInstance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *JobInstance) SetId(value *int64)() {
    m.id = value
}
// SetInstanceId sets the instanceId property value. The instanceId property
func (m *JobInstance) SetInstanceId(value *int64)() {
    m.instanceId = value
}
// SetJobName sets the jobName property value. The jobName property
func (m *JobInstance) SetJobName(value *string)() {
    m.jobName = value
}
// SetVersion sets the version property value. The version property
func (m *JobInstance) SetVersion(value *int32)() {
    m.version = value
}
// JobInstanceable 
type JobInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int64)
    GetInstanceId()(*int64)
    GetJobName()(*string)
    GetVersion()(*int32)
    SetId(value *int64)()
    SetInstanceId(value *int64)()
    SetJobName(value *string)()
    SetVersion(value *int32)()
}
