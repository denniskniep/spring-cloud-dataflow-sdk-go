package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StepExecution_failureExceptions 
type StepExecution_failureExceptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cause property
    cause StepExecution_failureExceptions_causeable
    // The localizedMessage property
    localizedMessage *string
    // The message property
    message *string
    // The stackTrace property
    stackTrace []StepExecution_failureExceptions_stackTraceable
}
// NewStepExecution_failureExceptions instantiates a new StepExecution_failureExceptions and sets the default values.
func NewStepExecution_failureExceptions()(*StepExecution_failureExceptions) {
    m := &StepExecution_failureExceptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStepExecution_failureExceptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStepExecution_failureExceptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStepExecution_failureExceptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StepExecution_failureExceptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCause gets the cause property value. The cause property
func (m *StepExecution_failureExceptions) GetCause()(StepExecution_failureExceptions_causeable) {
    return m.cause
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StepExecution_failureExceptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cause"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStepExecution_failureExceptions_causeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCause(val.(StepExecution_failureExceptions_causeable))
        }
        return nil
    }
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
        val, err := n.GetCollectionOfObjectValues(CreateStepExecution_failureExceptions_stackTraceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StepExecution_failureExceptions_stackTraceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StepExecution_failureExceptions_stackTraceable)
                }
            }
            m.SetStackTrace(res)
        }
        return nil
    }
    return res
}
// GetLocalizedMessage gets the localizedMessage property value. The localizedMessage property
func (m *StepExecution_failureExceptions) GetLocalizedMessage()(*string) {
    return m.localizedMessage
}
// GetMessage gets the message property value. The message property
func (m *StepExecution_failureExceptions) GetMessage()(*string) {
    return m.message
}
// GetStackTrace gets the stackTrace property value. The stackTrace property
func (m *StepExecution_failureExceptions) GetStackTrace()([]StepExecution_failureExceptions_stackTraceable) {
    return m.stackTrace
}
// Serialize serializes information the current object
func (m *StepExecution_failureExceptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cause", m.GetCause())
        if err != nil {
            return err
        }
    }
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
func (m *StepExecution_failureExceptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCause sets the cause property value. The cause property
func (m *StepExecution_failureExceptions) SetCause(value StepExecution_failureExceptions_causeable)() {
    m.cause = value
}
// SetLocalizedMessage sets the localizedMessage property value. The localizedMessage property
func (m *StepExecution_failureExceptions) SetLocalizedMessage(value *string)() {
    m.localizedMessage = value
}
// SetMessage sets the message property value. The message property
func (m *StepExecution_failureExceptions) SetMessage(value *string)() {
    m.message = value
}
// SetStackTrace sets the stackTrace property value. The stackTrace property
func (m *StepExecution_failureExceptions) SetStackTrace(value []StepExecution_failureExceptions_stackTraceable)() {
    m.stackTrace = value
}
// StepExecution_failureExceptionsable 
type StepExecution_failureExceptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCause()(StepExecution_failureExceptions_causeable)
    GetLocalizedMessage()(*string)
    GetMessage()(*string)
    GetStackTrace()([]StepExecution_failureExceptions_stackTraceable)
    SetCause(value StepExecution_failureExceptions_causeable)()
    SetLocalizedMessage(value *string)()
    SetMessage(value *string)()
    SetStackTrace(value []StepExecution_failureExceptions_stackTraceable)()
}
