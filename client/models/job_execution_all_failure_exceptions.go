package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobExecution_allFailureExceptions 
type JobExecution_allFailureExceptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cause property
    cause JobExecution_allFailureExceptions_causeable
    // The localizedMessage property
    localizedMessage *string
    // The message property
    message *string
    // The stackTrace property
    stackTrace []JobExecution_allFailureExceptions_stackTraceable
}
// NewJobExecution_allFailureExceptions instantiates a new JobExecution_allFailureExceptions and sets the default values.
func NewJobExecution_allFailureExceptions()(*JobExecution_allFailureExceptions) {
    m := &JobExecution_allFailureExceptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobExecution_allFailureExceptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobExecution_allFailureExceptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobExecution_allFailureExceptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecution_allFailureExceptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCause gets the cause property value. The cause property
func (m *JobExecution_allFailureExceptions) GetCause()(JobExecution_allFailureExceptions_causeable) {
    return m.cause
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobExecution_allFailureExceptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cause"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobExecution_allFailureExceptions_causeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCause(val.(JobExecution_allFailureExceptions_causeable))
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
        val, err := n.GetCollectionOfObjectValues(CreateJobExecution_allFailureExceptions_stackTraceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobExecution_allFailureExceptions_stackTraceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobExecution_allFailureExceptions_stackTraceable)
                }
            }
            m.SetStackTrace(res)
        }
        return nil
    }
    return res
}
// GetLocalizedMessage gets the localizedMessage property value. The localizedMessage property
func (m *JobExecution_allFailureExceptions) GetLocalizedMessage()(*string) {
    return m.localizedMessage
}
// GetMessage gets the message property value. The message property
func (m *JobExecution_allFailureExceptions) GetMessage()(*string) {
    return m.message
}
// GetStackTrace gets the stackTrace property value. The stackTrace property
func (m *JobExecution_allFailureExceptions) GetStackTrace()([]JobExecution_allFailureExceptions_stackTraceable) {
    return m.stackTrace
}
// Serialize serializes information the current object
func (m *JobExecution_allFailureExceptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *JobExecution_allFailureExceptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCause sets the cause property value. The cause property
func (m *JobExecution_allFailureExceptions) SetCause(value JobExecution_allFailureExceptions_causeable)() {
    m.cause = value
}
// SetLocalizedMessage sets the localizedMessage property value. The localizedMessage property
func (m *JobExecution_allFailureExceptions) SetLocalizedMessage(value *string)() {
    m.localizedMessage = value
}
// SetMessage sets the message property value. The message property
func (m *JobExecution_allFailureExceptions) SetMessage(value *string)() {
    m.message = value
}
// SetStackTrace sets the stackTrace property value. The stackTrace property
func (m *JobExecution_allFailureExceptions) SetStackTrace(value []JobExecution_allFailureExceptions_stackTraceable)() {
    m.stackTrace = value
}
// JobExecution_allFailureExceptionsable 
type JobExecution_allFailureExceptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCause()(JobExecution_allFailureExceptions_causeable)
    GetLocalizedMessage()(*string)
    GetMessage()(*string)
    GetStackTrace()([]JobExecution_allFailureExceptions_stackTraceable)
    SetCause(value JobExecution_allFailureExceptions_causeable)()
    SetLocalizedMessage(value *string)()
    SetMessage(value *string)()
    SetStackTrace(value []JobExecution_allFailureExceptions_stackTraceable)()
}
