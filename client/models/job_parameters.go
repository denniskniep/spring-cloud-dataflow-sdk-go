package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobParameters 
type JobParameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The empty property
    empty *bool
    // The parameters property
    parameters JobParameters_parametersable
}
// NewJobParameters instantiates a new JobParameters and sets the default values.
func NewJobParameters()(*JobParameters) {
    m := &JobParameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobParameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobParameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmpty gets the empty property value. The empty property
func (m *JobParameters) GetEmpty()(*bool) {
    return m.empty
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["empty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmpty(val)
        }
        return nil
    }
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJobParameters_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(JobParameters_parametersable))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
func (m *JobParameters) GetParameters()(JobParameters_parametersable) {
    return m.parameters
}
// Serialize serializes information the current object
func (m *JobParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("empty", m.GetEmpty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parameters", m.GetParameters())
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
func (m *JobParameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmpty sets the empty property value. The empty property
func (m *JobParameters) SetEmpty(value *bool)() {
    m.empty = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *JobParameters) SetParameters(value JobParameters_parametersable)() {
    m.parameters = value
}
// JobParametersable 
type JobParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmpty()(*bool)
    GetParameters()(JobParameters_parametersable)
    SetEmpty(value *bool)()
    SetParameters(value JobParameters_parametersable)()
}
