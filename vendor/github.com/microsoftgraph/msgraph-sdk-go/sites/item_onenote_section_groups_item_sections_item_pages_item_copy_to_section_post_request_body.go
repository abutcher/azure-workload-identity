package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody 
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody instantiates a new ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) {
    m := &ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
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
    res["siteCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetId gets the id property value. The id property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetSiteCollectionId()(*string) {
    val, err := m.GetBackingStore().Get("siteCollectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSiteId gets the siteId property value. The siteId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) GetSiteId()(*string) {
    val, err := m.GetBackingStore().Get("siteId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
        err := writer.WriteStringValue("siteCollectionId", m.GetSiteCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteId", m.GetSiteId())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. The id property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) SetSiteCollectionId(value *string)() {
    err := m.GetBackingStore().Set("siteCollectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSiteId sets the siteId property value. The siteId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBody) SetSiteId(value *string)() {
    err := m.GetBackingStore().Set("siteId", value)
    if err != nil {
        panic(err)
    }
}
// ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBodyable 
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetGroupId()(*string)
    GetId()(*string)
    GetSiteCollectionId()(*string)
    GetSiteId()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetGroupId(value *string)()
    SetId(value *string)()
    SetSiteCollectionId(value *string)()
    SetSiteId(value *string)()
}
