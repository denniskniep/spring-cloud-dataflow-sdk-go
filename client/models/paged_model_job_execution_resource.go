package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PagedModelJobExecutionResource 
type PagedModelJobExecutionResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content property
    content []JobExecutionResourceable
    // The links property
    links []Linkable
    // The page property
    page PageMetadataable
}
// NewPagedModelJobExecutionResource instantiates a new PagedModelJobExecutionResource and sets the default values.
func NewPagedModelJobExecutionResource()(*PagedModelJobExecutionResource) {
    m := &PagedModelJobExecutionResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePagedModelJobExecutionResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePagedModelJobExecutionResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPagedModelJobExecutionResource(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PagedModelJobExecutionResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
func (m *PagedModelJobExecutionResource) GetContent()([]JobExecutionResourceable) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PagedModelJobExecutionResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobExecutionResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobExecutionResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JobExecutionResourceable)
                }
            }
            m.SetContent(res)
        }
        return nil
    }
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
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePageMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val.(PageMetadataable))
        }
        return nil
    }
    return res
}
// GetLinks gets the links property value. The links property
func (m *PagedModelJobExecutionResource) GetLinks()([]Linkable) {
    return m.links
}
// GetPage gets the page property value. The page property
func (m *PagedModelJobExecutionResource) GetPage()(PageMetadataable) {
    return m.page
}
// Serialize serializes information the current object
func (m *PagedModelJobExecutionResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContent() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContent()))
        for i, v := range m.GetContent() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("content", cast)
        if err != nil {
            return err
        }
    }
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
    {
        err := writer.WriteObjectValue("page", m.GetPage())
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
func (m *PagedModelJobExecutionResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *PagedModelJobExecutionResource) SetContent(value []JobExecutionResourceable)() {
    m.content = value
}
// SetLinks sets the links property value. The links property
func (m *PagedModelJobExecutionResource) SetLinks(value []Linkable)() {
    m.links = value
}
// SetPage sets the page property value. The page property
func (m *PagedModelJobExecutionResource) SetPage(value PageMetadataable)() {
    m.page = value
}
// PagedModelJobExecutionResourceable 
type PagedModelJobExecutionResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]JobExecutionResourceable)
    GetLinks()([]Linkable)
    GetPage()(PageMetadataable)
    SetContent(value []JobExecutionResourceable)()
    SetLinks(value []Linkable)()
    SetPage(value PageMetadataable)()
}
