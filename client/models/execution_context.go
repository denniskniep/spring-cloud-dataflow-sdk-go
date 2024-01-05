package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExecutionContext 
type ExecutionContext struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dirty property
    dirty *bool
    // The empty property
    empty *bool
}
// NewExecutionContext instantiates a new ExecutionContext and sets the default values.
func NewExecutionContext()(*ExecutionContext) {
    m := &ExecutionContext{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExecutionContextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExecutionContextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExecutionContext(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExecutionContext) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDirty gets the dirty property value. The dirty property
func (m *ExecutionContext) GetDirty()(*bool) {
    return m.dirty
}
// GetEmpty gets the empty property value. The empty property
func (m *ExecutionContext) GetEmpty()(*bool) {
    return m.empty
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExecutionContext) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dirty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirty(val)
        }
        return nil
    }
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
    return res
}
// Serialize serializes information the current object
func (m *ExecutionContext) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("dirty", m.GetDirty())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("empty", m.GetEmpty())
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
func (m *ExecutionContext) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDirty sets the dirty property value. The dirty property
func (m *ExecutionContext) SetDirty(value *bool)() {
    m.dirty = value
}
// SetEmpty sets the empty property value. The empty property
func (m *ExecutionContext) SetEmpty(value *bool)() {
    m.empty = value
}
// ExecutionContextable 
type ExecutionContextable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDirty()(*bool)
    GetEmpty()(*bool)
    SetDirty(value *bool)()
    SetEmpty(value *bool)()
}
