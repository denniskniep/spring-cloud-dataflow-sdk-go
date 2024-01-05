package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Node 
type Node struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *string
    // The metadata property
    metadata Node_metadataable
    // The name property
    name *string
    // The properties property
    properties Node_propertiesable
}
// NewNode instantiates a new Node and sets the default values.
func NewNode()(*Node) {
    m := &Node{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNode(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Node) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Node) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNode_metadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(Node_metadataable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNode_propertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(Node_propertiesable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *Node) GetId()(*string) {
    return m.id
}
// GetMetadata gets the metadata property value. The metadata property
func (m *Node) GetMetadata()(Node_metadataable) {
    return m.metadata
}
// GetName gets the name property value. The name property
func (m *Node) GetName()(*string) {
    return m.name
}
// GetProperties gets the properties property value. The properties property
func (m *Node) GetProperties()(Node_propertiesable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *Node) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("properties", m.GetProperties())
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
func (m *Node) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *Node) SetId(value *string)() {
    m.id = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *Node) SetMetadata(value Node_metadataable)() {
    m.metadata = value
}
// SetName sets the name property value. The name property
func (m *Node) SetName(value *string)() {
    m.name = value
}
// SetProperties sets the properties property value. The properties property
func (m *Node) SetProperties(value Node_propertiesable)() {
    m.properties = value
}
// Nodeable 
type Nodeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetMetadata()(Node_metadataable)
    GetName()(*string)
    GetProperties()(Node_propertiesable)
    SetId(value *string)()
    SetMetadata(value Node_metadataable)()
    SetName(value *string)()
    SetProperties(value Node_propertiesable)()
}
