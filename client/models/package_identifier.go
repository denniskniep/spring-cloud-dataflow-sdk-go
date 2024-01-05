package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PackageIdentifier 
type PackageIdentifier struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The packageName property
    packageName *string
    // The packageVersion property
    packageVersion *string
    // The repositoryName property
    repositoryName *string
}
// NewPackageIdentifier instantiates a new PackageIdentifier and sets the default values.
func NewPackageIdentifier()(*PackageIdentifier) {
    m := &PackageIdentifier{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePackageIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageIdentifier(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PackageIdentifier) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PackageIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["packageName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageName(val)
        }
        return nil
    }
    res["packageVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageVersion(val)
        }
        return nil
    }
    res["repositoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryName(val)
        }
        return nil
    }
    return res
}
// GetPackageName gets the packageName property value. The packageName property
func (m *PackageIdentifier) GetPackageName()(*string) {
    return m.packageName
}
// GetPackageVersion gets the packageVersion property value. The packageVersion property
func (m *PackageIdentifier) GetPackageVersion()(*string) {
    return m.packageVersion
}
// GetRepositoryName gets the repositoryName property value. The repositoryName property
func (m *PackageIdentifier) GetRepositoryName()(*string) {
    return m.repositoryName
}
// Serialize serializes information the current object
func (m *PackageIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("packageName", m.GetPackageName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("packageVersion", m.GetPackageVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositoryName", m.GetRepositoryName())
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
func (m *PackageIdentifier) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPackageName sets the packageName property value. The packageName property
func (m *PackageIdentifier) SetPackageName(value *string)() {
    m.packageName = value
}
// SetPackageVersion sets the packageVersion property value. The packageVersion property
func (m *PackageIdentifier) SetPackageVersion(value *string)() {
    m.packageVersion = value
}
// SetRepositoryName sets the repositoryName property value. The repositoryName property
func (m *PackageIdentifier) SetRepositoryName(value *string)() {
    m.repositoryName = value
}
// PackageIdentifierable 
type PackageIdentifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPackageName()(*string)
    GetPackageVersion()(*string)
    GetRepositoryName()(*string)
    SetPackageName(value *string)()
    SetPackageVersion(value *string)()
    SetRepositoryName(value *string)()
}
