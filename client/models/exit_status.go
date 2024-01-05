package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExitStatus 
type ExitStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The exitCode property
    exitCode *string
    // The exitDescription property
    exitDescription *string
    // The running property
    running *bool
}
// NewExitStatus instantiates a new ExitStatus and sets the default values.
func NewExitStatus()(*ExitStatus) {
    m := &ExitStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExitStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExitStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExitStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExitStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExitCode gets the exitCode property value. The exitCode property
func (m *ExitStatus) GetExitCode()(*string) {
    return m.exitCode
}
// GetExitDescription gets the exitDescription property value. The exitDescription property
func (m *ExitStatus) GetExitDescription()(*string) {
    return m.exitDescription
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExitStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exitCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExitCode(val)
        }
        return nil
    }
    res["exitDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExitDescription(val)
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
    return res
}
// GetRunning gets the running property value. The running property
func (m *ExitStatus) GetRunning()(*bool) {
    return m.running
}
// Serialize serializes information the current object
func (m *ExitStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("exitCode", m.GetExitCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("exitDescription", m.GetExitDescription())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExitStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExitCode sets the exitCode property value. The exitCode property
func (m *ExitStatus) SetExitCode(value *string)() {
    m.exitCode = value
}
// SetExitDescription sets the exitDescription property value. The exitDescription property
func (m *ExitStatus) SetExitDescription(value *string)() {
    m.exitDescription = value
}
// SetRunning sets the running property value. The running property
func (m *ExitStatus) SetRunning(value *bool)() {
    m.running = value
}
// ExitStatusable 
type ExitStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExitCode()(*string)
    GetExitDescription()(*string)
    GetRunning()(*bool)
    SetExitCode(value *string)()
    SetExitDescription(value *string)()
    SetRunning(value *bool)()
}
