package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Graph 
type Graph struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The links property
    links []Linkable
    // The nodes property
    nodes []Nodeable
}
// NewGraph instantiates a new Graph and sets the default values.
func NewGraph()(*Graph) {
    m := &Graph{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGraphFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGraphFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGraph(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Graph) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Graph) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Linkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Linkable)
                }
            }
            m.SetLinks(res)
        }
        return nil
    }
    res["nodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Nodeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Nodeable)
                }
            }
            m.SetNodes(res)
        }
        return nil
    }
    return res
}
// GetLinks gets the links property value. The links property
func (m *Graph) GetLinks()([]Linkable) {
    return m.links
}
// GetNodes gets the nodes property value. The nodes property
func (m *Graph) GetNodes()([]Nodeable) {
    return m.nodes
}
// Serialize serializes information the current object
func (m *Graph) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNodes()))
        for i, v := range m.GetNodes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("nodes", cast)
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
func (m *Graph) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLinks sets the links property value. The links property
func (m *Graph) SetLinks(value []Linkable)() {
    m.links = value
}
// SetNodes sets the nodes property value. The nodes property
func (m *Graph) SetNodes(value []Nodeable)() {
    m.nodes = value
}
// Graphable 
type Graphable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLinks()([]Linkable)
    GetNodes()([]Nodeable)
    SetLinks(value []Linkable)()
    SetNodes(value []Nodeable)()
}
