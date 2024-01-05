package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActuatorPostRequest 
type ActuatorPostRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The body property
    body ActuatorPostRequest_bodyable
    // The endpoint property
    endpoint *string
}
// NewActuatorPostRequest instantiates a new ActuatorPostRequest and sets the default values.
func NewActuatorPostRequest()(*ActuatorPostRequest) {
    m := &ActuatorPostRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActuatorPostRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActuatorPostRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActuatorPostRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActuatorPostRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The body property
func (m *ActuatorPostRequest) GetBody()(ActuatorPostRequest_bodyable) {
    return m.body
}
// GetEndpoint gets the endpoint property value. The endpoint property
func (m *ActuatorPostRequest) GetEndpoint()(*string) {
    return m.endpoint
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActuatorPostRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActuatorPostRequest_bodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ActuatorPostRequest_bodyable))
        }
        return nil
    }
    res["endpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpoint(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ActuatorPostRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endpoint", m.GetEndpoint())
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
func (m *ActuatorPostRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The body property
func (m *ActuatorPostRequest) SetBody(value ActuatorPostRequest_bodyable)() {
    m.body = value
}
// SetEndpoint sets the endpoint property value. The endpoint property
func (m *ActuatorPostRequest) SetEndpoint(value *string)() {
    m.endpoint = value
}
// ActuatorPostRequestable 
type ActuatorPostRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(ActuatorPostRequest_bodyable)
    GetEndpoint()(*string)
    SetBody(value ActuatorPostRequest_bodyable)()
    SetEndpoint(value *string)()
}
