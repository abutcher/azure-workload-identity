package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FederatedIdentityCredentialCollectionResponse 
type FederatedIdentityCredentialCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewFederatedIdentityCredentialCollectionResponse instantiates a new FederatedIdentityCredentialCollectionResponse and sets the default values.
func NewFederatedIdentityCredentialCollectionResponse()(*FederatedIdentityCredentialCollectionResponse) {
    m := &FederatedIdentityCredentialCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateFederatedIdentityCredentialCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFederatedIdentityCredentialCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFederatedIdentityCredentialCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FederatedIdentityCredentialCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFederatedIdentityCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FederatedIdentityCredentialable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FederatedIdentityCredentialable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *FederatedIdentityCredentialCollectionResponse) GetValue()([]FederatedIdentityCredentialable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FederatedIdentityCredentialable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FederatedIdentityCredentialCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *FederatedIdentityCredentialCollectionResponse) SetValue(value []FederatedIdentityCredentialable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// FederatedIdentityCredentialCollectionResponseable 
type FederatedIdentityCredentialCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]FederatedIdentityCredentialable)
    SetValue(value []FederatedIdentityCredentialable)()
}
