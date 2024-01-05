package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateStreamRequest 
type UpdateStreamRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The appNames property
    appNames []string
    // The force property
    force *bool
    // The packageIdentifier property
    packageIdentifier PackageIdentifierable
    // The releaseName property
    releaseName *string
    // The updateProperties property
    updateProperties UpdateStreamRequest_updatePropertiesable
}
// NewUpdateStreamRequest instantiates a new UpdateStreamRequest and sets the default values.
func NewUpdateStreamRequest()(*UpdateStreamRequest) {
    m := &UpdateStreamRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdateStreamRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateStreamRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateStreamRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateStreamRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppNames gets the appNames property value. The appNames property
func (m *UpdateStreamRequest) GetAppNames()([]string) {
    return m.appNames
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateStreamRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAppNames(res)
        }
        return nil
    }
    res["force"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForce(val)
        }
        return nil
    }
    res["packageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageIdentifier(val.(PackageIdentifierable))
        }
        return nil
    }
    res["releaseName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseName(val)
        }
        return nil
    }
    res["updateProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdateStreamRequest_updatePropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateProperties(val.(UpdateStreamRequest_updatePropertiesable))
        }
        return nil
    }
    return res
}
// GetForce gets the force property value. The force property
func (m *UpdateStreamRequest) GetForce()(*bool) {
    return m.force
}
// GetPackageIdentifier gets the packageIdentifier property value. The packageIdentifier property
func (m *UpdateStreamRequest) GetPackageIdentifier()(PackageIdentifierable) {
    return m.packageIdentifier
}
// GetReleaseName gets the releaseName property value. The releaseName property
func (m *UpdateStreamRequest) GetReleaseName()(*string) {
    return m.releaseName
}
// GetUpdateProperties gets the updateProperties property value. The updateProperties property
func (m *UpdateStreamRequest) GetUpdateProperties()(UpdateStreamRequest_updatePropertiesable) {
    return m.updateProperties
}
// Serialize serializes information the current object
func (m *UpdateStreamRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppNames() != nil {
        err := writer.WriteCollectionOfStringValues("appNames", m.GetAppNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("force", m.GetForce())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("packageIdentifier", m.GetPackageIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("releaseName", m.GetReleaseName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("updateProperties", m.GetUpdateProperties())
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
func (m *UpdateStreamRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppNames sets the appNames property value. The appNames property
func (m *UpdateStreamRequest) SetAppNames(value []string)() {
    m.appNames = value
}
// SetForce sets the force property value. The force property
func (m *UpdateStreamRequest) SetForce(value *bool)() {
    m.force = value
}
// SetPackageIdentifier sets the packageIdentifier property value. The packageIdentifier property
func (m *UpdateStreamRequest) SetPackageIdentifier(value PackageIdentifierable)() {
    m.packageIdentifier = value
}
// SetReleaseName sets the releaseName property value. The releaseName property
func (m *UpdateStreamRequest) SetReleaseName(value *string)() {
    m.releaseName = value
}
// SetUpdateProperties sets the updateProperties property value. The updateProperties property
func (m *UpdateStreamRequest) SetUpdateProperties(value UpdateStreamRequest_updatePropertiesable)() {
    m.updateProperties = value
}
// UpdateStreamRequestable 
type UpdateStreamRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppNames()([]string)
    GetForce()(*bool)
    GetPackageIdentifier()(PackageIdentifierable)
    GetReleaseName()(*string)
    GetUpdateProperties()(UpdateStreamRequest_updatePropertiesable)
    SetAppNames(value []string)()
    SetForce(value *bool)()
    SetPackageIdentifier(value PackageIdentifierable)()
    SetReleaseName(value *string)()
    SetUpdateProperties(value UpdateStreamRequest_updatePropertiesable)()
}
