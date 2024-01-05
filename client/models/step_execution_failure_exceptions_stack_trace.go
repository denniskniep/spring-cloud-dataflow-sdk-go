package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StepExecution_failureExceptions_stackTrace 
type StepExecution_failureExceptions_stackTrace struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The classLoaderName property
    classLoaderName *string
    // The className property
    className *string
    // The fileName property
    fileName *string
    // The lineNumber property
    lineNumber *int32
    // The methodName property
    methodName *string
    // The moduleName property
    moduleName *string
    // The moduleVersion property
    moduleVersion *string
    // The nativeMethod property
    nativeMethod *bool
}
// NewStepExecution_failureExceptions_stackTrace instantiates a new StepExecution_failureExceptions_stackTrace and sets the default values.
func NewStepExecution_failureExceptions_stackTrace()(*StepExecution_failureExceptions_stackTrace) {
    m := &StepExecution_failureExceptions_stackTrace{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStepExecution_failureExceptions_stackTraceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStepExecution_failureExceptions_stackTraceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStepExecution_failureExceptions_stackTrace(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StepExecution_failureExceptions_stackTrace) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClassLoaderName gets the classLoaderName property value. The classLoaderName property
func (m *StepExecution_failureExceptions_stackTrace) GetClassLoaderName()(*string) {
    return m.classLoaderName
}
// GetClassName gets the className property value. The className property
func (m *StepExecution_failureExceptions_stackTrace) GetClassName()(*string) {
    return m.className
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StepExecution_failureExceptions_stackTrace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classLoaderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassLoaderName(val)
        }
        return nil
    }
    res["className"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassName(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["lineNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineNumber(val)
        }
        return nil
    }
    res["methodName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethodName(val)
        }
        return nil
    }
    res["moduleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModuleName(val)
        }
        return nil
    }
    res["moduleVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModuleVersion(val)
        }
        return nil
    }
    res["nativeMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNativeMethod(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The fileName property
func (m *StepExecution_failureExceptions_stackTrace) GetFileName()(*string) {
    return m.fileName
}
// GetLineNumber gets the lineNumber property value. The lineNumber property
func (m *StepExecution_failureExceptions_stackTrace) GetLineNumber()(*int32) {
    return m.lineNumber
}
// GetMethodName gets the methodName property value. The methodName property
func (m *StepExecution_failureExceptions_stackTrace) GetMethodName()(*string) {
    return m.methodName
}
// GetModuleName gets the moduleName property value. The moduleName property
func (m *StepExecution_failureExceptions_stackTrace) GetModuleName()(*string) {
    return m.moduleName
}
// GetModuleVersion gets the moduleVersion property value. The moduleVersion property
func (m *StepExecution_failureExceptions_stackTrace) GetModuleVersion()(*string) {
    return m.moduleVersion
}
// GetNativeMethod gets the nativeMethod property value. The nativeMethod property
func (m *StepExecution_failureExceptions_stackTrace) GetNativeMethod()(*bool) {
    return m.nativeMethod
}
// Serialize serializes information the current object
func (m *StepExecution_failureExceptions_stackTrace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("classLoaderName", m.GetClassLoaderName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("className", m.GetClassName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("lineNumber", m.GetLineNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("methodName", m.GetMethodName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moduleName", m.GetModuleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moduleVersion", m.GetModuleVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("nativeMethod", m.GetNativeMethod())
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
func (m *StepExecution_failureExceptions_stackTrace) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClassLoaderName sets the classLoaderName property value. The classLoaderName property
func (m *StepExecution_failureExceptions_stackTrace) SetClassLoaderName(value *string)() {
    m.classLoaderName = value
}
// SetClassName sets the className property value. The className property
func (m *StepExecution_failureExceptions_stackTrace) SetClassName(value *string)() {
    m.className = value
}
// SetFileName sets the fileName property value. The fileName property
func (m *StepExecution_failureExceptions_stackTrace) SetFileName(value *string)() {
    m.fileName = value
}
// SetLineNumber sets the lineNumber property value. The lineNumber property
func (m *StepExecution_failureExceptions_stackTrace) SetLineNumber(value *int32)() {
    m.lineNumber = value
}
// SetMethodName sets the methodName property value. The methodName property
func (m *StepExecution_failureExceptions_stackTrace) SetMethodName(value *string)() {
    m.methodName = value
}
// SetModuleName sets the moduleName property value. The moduleName property
func (m *StepExecution_failureExceptions_stackTrace) SetModuleName(value *string)() {
    m.moduleName = value
}
// SetModuleVersion sets the moduleVersion property value. The moduleVersion property
func (m *StepExecution_failureExceptions_stackTrace) SetModuleVersion(value *string)() {
    m.moduleVersion = value
}
// SetNativeMethod sets the nativeMethod property value. The nativeMethod property
func (m *StepExecution_failureExceptions_stackTrace) SetNativeMethod(value *bool)() {
    m.nativeMethod = value
}
// StepExecution_failureExceptions_stackTraceable 
type StepExecution_failureExceptions_stackTraceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassLoaderName()(*string)
    GetClassName()(*string)
    GetFileName()(*string)
    GetLineNumber()(*int32)
    GetMethodName()(*string)
    GetModuleName()(*string)
    GetModuleVersion()(*string)
    GetNativeMethod()(*bool)
    SetClassLoaderName(value *string)()
    SetClassName(value *string)()
    SetFileName(value *string)()
    SetLineNumber(value *int32)()
    SetMethodName(value *string)()
    SetModuleName(value *string)()
    SetModuleVersion(value *string)()
    SetNativeMethod(value *bool)()
}
