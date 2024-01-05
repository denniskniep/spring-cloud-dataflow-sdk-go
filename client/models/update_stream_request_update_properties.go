package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateStreamRequest_updateProperties 
type UpdateStreamRequest_updateProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewUpdateStreamRequest_updateProperties instantiates a new UpdateStreamRequest_updateProperties and sets the default values.
func NewUpdateStreamRequest_updateProperties()(*UpdateStreamRequest_updateProperties) {
    m := &UpdateStreamRequest_updateProperties{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdateStreamRequest_updatePropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateStreamRequest_updatePropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateStreamRequest_updateProperties(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateStreamRequest_updateProperties) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateStreamRequest_updateProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *UpdateStreamRequest_updateProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateStreamRequest_updateProperties) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// UpdateStreamRequest_updatePropertiesable 
type UpdateStreamRequest_updatePropertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
