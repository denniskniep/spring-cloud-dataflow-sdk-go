package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobExecutionResource_timeZone 
type JobExecutionResource_timeZone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The displayName property
    displayName *string
    // The dstsavings property
    dstsavings *int32
    // The id property
    id *string
    // The rawOffset property
    rawOffset *int32
}
// NewJobExecutionResource_timeZone instantiates a new JobExecutionResource_timeZone and sets the default values.
func NewJobExecutionResource_timeZone()(*JobExecutionResource_timeZone) {
    m := &JobExecutionResource_timeZone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateJobExecutionResource_timeZoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobExecutionResource_timeZoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJobExecutionResource_timeZone(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *JobExecutionResource_timeZone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *JobExecutionResource_timeZone) GetDisplayName()(*string) {
    return m.displayName
}
// GetDstsavings gets the dstsavings property value. The dstsavings property
func (m *JobExecutionResource_timeZone) GetDstsavings()(*int32) {
    return m.dstsavings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobExecutionResource_timeZone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["dstsavings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDstsavings(val)
        }
        return nil
    }
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
    res["rawOffset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawOffset(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *JobExecutionResource_timeZone) GetId()(*string) {
    return m.id
}
// GetRawOffset gets the rawOffset property value. The rawOffset property
func (m *JobExecutionResource_timeZone) GetRawOffset()(*int32) {
    return m.rawOffset
}
// Serialize serializes information the current object
func (m *JobExecutionResource_timeZone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dstsavings", m.GetDstsavings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rawOffset", m.GetRawOffset())
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
func (m *JobExecutionResource_timeZone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *JobExecutionResource_timeZone) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDstsavings sets the dstsavings property value. The dstsavings property
func (m *JobExecutionResource_timeZone) SetDstsavings(value *int32)() {
    m.dstsavings = value
}
// SetId sets the id property value. The id property
func (m *JobExecutionResource_timeZone) SetId(value *string)() {
    m.id = value
}
// SetRawOffset sets the rawOffset property value. The rawOffset property
func (m *JobExecutionResource_timeZone) SetRawOffset(value *int32)() {
    m.rawOffset = value
}
// JobExecutionResource_timeZoneable 
type JobExecutionResource_timeZoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetDstsavings()(*int32)
    GetId()(*string)
    GetRawOffset()(*int32)
    SetDisplayName(value *string)()
    SetDstsavings(value *int32)()
    SetId(value *string)()
    SetRawOffset(value *int32)()
}
