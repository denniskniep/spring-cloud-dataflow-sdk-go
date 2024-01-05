package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageMetadata 
type PageMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number property
    number *int64
    // The size property
    size *int64
    // The totalElements property
    totalElements *int64
    // The totalPages property
    totalPages *int64
}
// NewPageMetadata instantiates a new PageMetadata and sets the default values.
func NewPageMetadata()(*PageMetadata) {
    m := &PageMetadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePageMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePageMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPageMetadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PageMetadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PageMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["totalElements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalElements(val)
        }
        return nil
    }
    res["totalPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPages(val)
        }
        return nil
    }
    return res
}
// GetNumber gets the number property value. The number property
func (m *PageMetadata) GetNumber()(*int64) {
    return m.number
}
// GetSize gets the size property value. The size property
func (m *PageMetadata) GetSize()(*int64) {
    return m.size
}
// GetTotalElements gets the totalElements property value. The totalElements property
func (m *PageMetadata) GetTotalElements()(*int64) {
    return m.totalElements
}
// GetTotalPages gets the totalPages property value. The totalPages property
func (m *PageMetadata) GetTotalPages()(*int64) {
    return m.totalPages
}
// Serialize serializes information the current object
func (m *PageMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalElements", m.GetTotalElements())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalPages", m.GetTotalPages())
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
func (m *PageMetadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumber sets the number property value. The number property
func (m *PageMetadata) SetNumber(value *int64)() {
    m.number = value
}
// SetSize sets the size property value. The size property
func (m *PageMetadata) SetSize(value *int64)() {
    m.size = value
}
// SetTotalElements sets the totalElements property value. The totalElements property
func (m *PageMetadata) SetTotalElements(value *int64)() {
    m.totalElements = value
}
// SetTotalPages sets the totalPages property value. The totalPages property
func (m *PageMetadata) SetTotalPages(value *int64)() {
    m.totalPages = value
}
// PageMetadataable 
type PageMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumber()(*int64)
    GetSize()(*int64)
    GetTotalElements()(*int64)
    GetTotalPages()(*int64)
    SetNumber(value *int64)()
    SetSize(value *int64)()
    SetTotalElements(value *int64)()
    SetTotalPages(value *int64)()
}
