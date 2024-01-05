package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StepExecution_failureExceptions_cause 
type StepExecution_failureExceptions_cause struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The localizedMessage property
    localizedMessage *string
    // The message property
    message *string
    // The stackTrace property
    stackTrace []StepExecution_failureExceptions_cause_stackTraceable
}
// NewStepExecution_failureExceptions_cause instantiates a new StepExecution_failureExceptions_cause and sets the default values.
func NewStepExecution_failureExceptions_cause()(*StepExecution_failureExceptions_cause) {
    m := &StepExecution_failureExceptions_cause{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStepExecution_failureExceptions_causeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStepExecution_failureExceptions_causeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStepExecution_failureExceptions_cause(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StepExecution_failureExceptions_cause) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StepExecution_failureExceptions_cause) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["localizedMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalizedMessage(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["stackTrace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStepExecution_failureExceptions_cause_stackTraceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StepExecution_failureExceptions_cause_stackTraceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StepExecution_failureExceptions_cause_stackTraceable)
                }
            }
            m.SetStackTrace(res)
        }
        return nil
    }
    return res
}
// GetLocalizedMessage gets the localizedMessage property value. The localizedMessage property
func (m *StepExecution_failureExceptions_cause) GetLocalizedMessage()(*string) {
    return m.localizedMessage
}
// GetMessage gets the message property value. The message property
func (m *StepExecution_failureExceptions_cause) GetMessage()(*string) {
    return m.message
}
// GetStackTrace gets the stackTrace property value. The stackTrace property
func (m *StepExecution_failureExceptions_cause) GetStackTrace()([]StepExecution_failureExceptions_cause_stackTraceable) {
    return m.stackTrace
}
// Serialize serializes information the current object
func (m *StepExecution_failureExceptions_cause) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("localizedMessage", m.GetLocalizedMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetStackTrace() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStackTrace()))
        for i, v := range m.GetStackTrace() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("stackTrace", cast)
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
func (m *StepExecution_failureExceptions_cause) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocalizedMessage sets the localizedMessage property value. The localizedMessage property
func (m *StepExecution_failureExceptions_cause) SetLocalizedMessage(value *string)() {
    m.localizedMessage = value
}
// SetMessage sets the message property value. The message property
func (m *StepExecution_failureExceptions_cause) SetMessage(value *string)() {
    m.message = value
}
// SetStackTrace sets the stackTrace property value. The stackTrace property
func (m *StepExecution_failureExceptions_cause) SetStackTrace(value []StepExecution_failureExceptions_cause_stackTraceable)() {
    m.stackTrace = value
}
// StepExecution_failureExceptions_causeable 
type StepExecution_failureExceptions_causeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocalizedMessage()(*string)
    GetMessage()(*string)
    GetStackTrace()([]StepExecution_failureExceptions_cause_stackTraceable)
    SetLocalizedMessage(value *string)()
    SetMessage(value *string)()
    SetStackTrace(value []StepExecution_failureExceptions_cause_stackTraceable)()
}
