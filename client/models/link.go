package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Link 
type Link struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deprecation property
    deprecation *string
    // The href property
    href *string
    // The hreflang property
    hreflang *string
    // The media property
    media *string
    // The name property
    name *string
    // The profile property
    profile *string
    // The rel property
    rel *string
    // The title property
    title *string
    // The type property
    typeEscaped *string
}
// NewLink instantiates a new Link and sets the default values.
func NewLink()(*Link) {
    m := &Link{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLink(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Link) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeprecation gets the deprecation property value. The deprecation property
func (m *Link) GetDeprecation()(*string) {
    return m.deprecation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Link) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deprecation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprecation(val)
        }
        return nil
    }
    res["href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHref(val)
        }
        return nil
    }
    res["hreflang"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHreflang(val)
        }
        return nil
    }
    res["media"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedia(val)
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
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val)
        }
        return nil
    }
    res["rel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRel(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetHref gets the href property value. The href property
func (m *Link) GetHref()(*string) {
    return m.href
}
// GetHreflang gets the hreflang property value. The hreflang property
func (m *Link) GetHreflang()(*string) {
    return m.hreflang
}
// GetMedia gets the media property value. The media property
func (m *Link) GetMedia()(*string) {
    return m.media
}
// GetName gets the name property value. The name property
func (m *Link) GetName()(*string) {
    return m.name
}
// GetProfile gets the profile property value. The profile property
func (m *Link) GetProfile()(*string) {
    return m.profile
}
// GetRel gets the rel property value. The rel property
func (m *Link) GetRel()(*string) {
    return m.rel
}
// GetTitle gets the title property value. The title property
func (m *Link) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. The type property
func (m *Link) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Link) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deprecation", m.GetDeprecation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("href", m.GetHref())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hreflang", m.GetHreflang())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("media", m.GetMedia())
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
        err := writer.WriteStringValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rel", m.GetRel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *Link) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeprecation sets the deprecation property value. The deprecation property
func (m *Link) SetDeprecation(value *string)() {
    m.deprecation = value
}
// SetHref sets the href property value. The href property
func (m *Link) SetHref(value *string)() {
    m.href = value
}
// SetHreflang sets the hreflang property value. The hreflang property
func (m *Link) SetHreflang(value *string)() {
    m.hreflang = value
}
// SetMedia sets the media property value. The media property
func (m *Link) SetMedia(value *string)() {
    m.media = value
}
// SetName sets the name property value. The name property
func (m *Link) SetName(value *string)() {
    m.name = value
}
// SetProfile sets the profile property value. The profile property
func (m *Link) SetProfile(value *string)() {
    m.profile = value
}
// SetRel sets the rel property value. The rel property
func (m *Link) SetRel(value *string)() {
    m.rel = value
}
// SetTitle sets the title property value. The title property
func (m *Link) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Link) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// Linkable 
type Linkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeprecation()(*string)
    GetHref()(*string)
    GetHreflang()(*string)
    GetMedia()(*string)
    GetName()(*string)
    GetProfile()(*string)
    GetRel()(*string)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    SetDeprecation(value *string)()
    SetHref(value *string)()
    SetHreflang(value *string)()
    SetMedia(value *string)()
    SetName(value *string)()
    SetProfile(value *string)()
    SetRel(value *string)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
}
