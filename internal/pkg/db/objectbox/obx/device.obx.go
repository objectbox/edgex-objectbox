// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	"github.com/edgexfoundry/edgex-go/pkg/models"
	. "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type device_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var DeviceBinding = device_EntityInfo{
	Entity: objectbox.Entity{
		Id: 3,
	},
	Uid: 8277139392933333295,
}

// Device_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Device_ = struct {
	BaseObject_Created  *objectbox.PropertyInt64
	BaseObject_Modified *objectbox.PropertyInt64
	BaseObject_Origin   *objectbox.PropertyInt64
	Description         *objectbox.PropertyString
	Id                  *objectbox.PropertyUint64
	Name                *objectbox.PropertyString
	AdminState          *objectbox.PropertyString
	OperatingState      *objectbox.PropertyString
	Addressable         *objectbox.RelationOneToMany
	LastConnected       *objectbox.PropertyInt64
	LastReported        *objectbox.PropertyInt64
	Labels              *objectbox.PropertyStringVector
	Location            *objectbox.PropertyByteVector
	Service             *objectbox.RelationOneToMany
	Profile             *objectbox.RelationOneToMany
}{
	BaseObject_Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &DeviceBinding.Entity,
		},
	},
	BaseObject_Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &DeviceBinding.Entity,
		},
	},
	BaseObject_Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &DeviceBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &DeviceBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &DeviceBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &DeviceBinding.Entity,
		},
	},
	AdminState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &DeviceBinding.Entity,
		},
	},
	OperatingState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &DeviceBinding.Entity,
		},
	},
	Addressable: &objectbox.RelationOneToMany{
		Property: &objectbox.BaseProperty{
			Id:     9,
			Entity: &DeviceBinding.Entity,
		},
		Target: &AddressableBinding.Entity,
	},
	LastConnected: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &DeviceBinding.Entity,
		},
	},
	LastReported: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &DeviceBinding.Entity,
		},
	},
	Labels: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &DeviceBinding.Entity,
		},
	},
	Location: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &DeviceBinding.Entity,
		},
	},
	Service: &objectbox.RelationOneToMany{
		Property: &objectbox.BaseProperty{
			Id:     14,
			Entity: &DeviceBinding.Entity,
		},
		Target: &DeviceServiceBinding.Entity,
	},
	Profile: &objectbox.RelationOneToMany{
		Property: &objectbox.BaseProperty{
			Id:     15,
			Entity: &DeviceBinding.Entity,
		},
		Target: &DeviceProfileBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (device_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by ObjectBox during model build
func (device_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Device", 3, 8277139392933333295)
	model.Property("BaseObject_Created", objectbox.PropertyType_Long, 1, 1426902891719146417)
	model.Property("BaseObject_Modified", objectbox.PropertyType_Long, 2, 5883868846286861351)
	model.Property("BaseObject_Origin", objectbox.PropertyType_Long, 3, 4095520008548384869)
	model.Property("Description", objectbox.PropertyType_String, 4, 8000536243793605912)
	model.Property("Id", objectbox.PropertyType_Long, 5, 4515089217959618723)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Name", objectbox.PropertyType_String, 6, 716857237269823862)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(2, 167828898366445354)
	model.Property("AdminState", objectbox.PropertyType_String, 7, 8163911656701609670)
	model.Property("OperatingState", objectbox.PropertyType_String, 8, 145868754249953301)
	model.Property("Addressable", objectbox.PropertyType_Relation, 9, 5293395039385098111)
	model.PropertyRelation("Addressable", 3, 8824925290223488279)
	model.Property("LastConnected", objectbox.PropertyType_Long, 10, 783154740742750654)
	model.Property("LastReported", objectbox.PropertyType_Long, 11, 9043009336819819958)
	model.Property("Labels", objectbox.PropertyType_StringVector, 12, 1085892164679946752)
	model.Property("Location", objectbox.PropertyType_ByteVector, 13, 6218275966482027583)
	model.Property("Service", objectbox.PropertyType_Relation, 14, 6082578288181241795)
	model.PropertyRelation("DeviceService", 4, 6949046153132189499)
	model.Property("Profile", objectbox.PropertyType_Relation, 15, 6924709779874779070)
	model.PropertyRelation("DeviceProfile", 5, 1047090894033747089)
	model.EntityLastPropertyId(15, 6924709779874779070)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (device_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Device); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Device).Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (device_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*Device); ok {
		obj.Id = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Device).Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (device_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	if rel := &object.(*Device).Addressable; rel != nil {
		rId, err := AddressableBinding.GetId(rel)
		if err != nil {
			return err
		} else if rId == 0 {
			if err := txn.RunWithCursor(AddressableBinding.Id, func(targetCursor *objectbox.Cursor) error {
				_, err := targetCursor.Put(rel) // NOTE Put/PutAsync() has a side-effect of setting the rel.ID
				return err
			}); err != nil {
				return err
			}
		}
	}
	if rel := &object.(*Device).Service; rel != nil {
		rId, err := DeviceServiceBinding.GetId(rel)
		if err != nil {
			return err
		} else if rId == 0 {
			if err := txn.RunWithCursor(DeviceServiceBinding.Id, func(targetCursor *objectbox.Cursor) error {
				_, err := targetCursor.Put(rel) // NOTE Put/PutAsync() has a side-effect of setting the rel.ID
				return err
			}); err != nil {
				return err
			}
		}
	}
	if rel := &object.(*Device).Profile; rel != nil {
		rId, err := DeviceProfileBinding.GetId(rel)
		if err != nil {
			return err
		} else if rId == 0 {
			if err := txn.RunWithCursor(DeviceProfileBinding.Id, func(targetCursor *objectbox.Cursor) error {
				_, err := targetCursor.Put(rel) // NOTE Put/PutAsync() has a side-effect of setting the rel.ID
				return err
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (device_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Device)
	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.DescribedObject.Description)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetAdminState = fbutils.CreateStringOffset(fbb, string(obj.AdminState))
	var offsetOperatingState = fbutils.CreateStringOffset(fbb, string(obj.OperatingState))
	var offsetLabels = fbutils.CreateStringVectorOffset(fbb, obj.Labels)
	var offsetLocation = fbutils.CreateByteVectorOffset(fbb, interfaceJsonToDatabaseValue(obj.Location))

	var rIdAddressable uint64
	if rel := &obj.Addressable; rel != nil {
		if rId, err := AddressableBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdAddressable = rId
		}
	}

	var rIdService uint64
	if rel := &obj.Service; rel != nil {
		if rId, err := DeviceServiceBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdService = rId
		}
	}

	var rIdProfile uint64
	if rel := &obj.Profile; rel != nil {
		if rId, err := DeviceProfileBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdProfile = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(15)
	fbutils.SetInt64Slot(fbb, 0, obj.DescribedObject.BaseObject.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.DescribedObject.BaseObject.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.DescribedObject.BaseObject.Origin)
	fbutils.SetUOffsetTSlot(fbb, 3, offsetDescription)
	fbutils.SetUint64Slot(fbb, 4, id)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetAdminState)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetOperatingState)
	fbutils.SetUint64Slot(fbb, 8, rIdAddressable)
	fbutils.SetInt64Slot(fbb, 9, obj.LastConnected)
	fbutils.SetInt64Slot(fbb, 10, obj.LastReported)
	fbutils.SetUOffsetTSlot(fbb, 11, offsetLabels)
	fbutils.SetUOffsetTSlot(fbb, 12, offsetLocation)
	fbutils.SetUint64Slot(fbb, 13, rIdService)
	fbutils.SetUint64Slot(fbb, 14, rIdProfile)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (device_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(12, 0)

	var relAddressable *Addressable
	if rId := table.GetUint64Slot(20, 0); rId > 0 {
		if err := txn.RunWithCursor(AddressableBinding.Id, func(targetCursor *objectbox.Cursor) error {
			if relObject, err := targetCursor.Get(rId); err != nil {
				return err
			} else if relObj, ok := relObject.(*Addressable); ok {
				relAddressable = relObj
			} else {
				var relObj = relObject.(Addressable)
				relAddressable = &relObj
			}
			return nil
		}); err != nil {
			return nil, err
		}
	} else {
		relAddressable = &Addressable{}
	}

	var relService *DeviceService
	if rId := table.GetUint64Slot(30, 0); rId > 0 {
		if err := txn.RunWithCursor(DeviceServiceBinding.Id, func(targetCursor *objectbox.Cursor) error {
			if relObject, err := targetCursor.Get(rId); err != nil {
				return err
			} else if relObj, ok := relObject.(*DeviceService); ok {
				relService = relObj
			} else {
				var relObj = relObject.(DeviceService)
				relService = &relObj
			}
			return nil
		}); err != nil {
			return nil, err
		}
	} else {
		relService = &DeviceService{}
	}

	var relProfile *DeviceProfile
	if rId := table.GetUint64Slot(32, 0); rId > 0 {
		if err := txn.RunWithCursor(DeviceProfileBinding.Id, func(targetCursor *objectbox.Cursor) error {
			if relObject, err := targetCursor.Get(rId); err != nil {
				return err
			} else if relObj, ok := relObject.(*DeviceProfile); ok {
				relProfile = relObj
			} else {
				var relObj = relObject.(DeviceProfile)
				relProfile = &relObj
			}
			return nil
		}); err != nil {
			return nil, err
		}
	} else {
		relProfile = &DeviceProfile{}
	}

	return &Device{
		DescribedObject: models.DescribedObject{
			BaseObject: BaseObject{
				Created:  table.GetInt64Slot(4, 0),
				Modified: table.GetInt64Slot(6, 0),
				Origin:   table.GetInt64Slot(8, 0),
			},
			Description: fbutils.GetStringSlot(table, 10),
		},
		Id:             objectbox.StringIdConvertToEntityProperty(id),
		Name:           fbutils.GetStringSlot(table, 14),
		AdminState:     models.AdminState(fbutils.GetStringSlot(table, 16)),
		OperatingState: models.OperatingState(fbutils.GetStringSlot(table, 18)),
		Addressable:    *relAddressable,
		LastConnected:  table.GetInt64Slot(22, 0),
		LastReported:   table.GetInt64Slot(24, 0),
		Labels:         fbutils.GetStringVectorSlot(table, 26),
		Location:       interfaceJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 28)),
		Service:        *relService,
		Profile:        *relProfile,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (device_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Device, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (device_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]Device), *object.(*Device))
}

// Box provides CRUD access to Device objects
type DeviceBox struct {
	*objectbox.Box
}

// BoxForDevice opens a box of Device objects
func BoxForDevice(ob *objectbox.ObjectBox) *DeviceBox {
	return &DeviceBox{
		Box: ob.InternalBox(3),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Device.Id property on the passed object will be assigned the new ID as well.
func (box *DeviceBox) Put(object *Device) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Device.Id property on the passed object will be assigned the new ID as well.
//
// It's executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "Put & Forget:" you gain faster puts as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
//
// In situations with (extremely) high async load, this method may be throttled (~1ms) or delayed (<1s).
// In the unlikely event that the object could not be enqueued after delaying, an error will be returned.
//
// Note that this method does not give you hard durability guarantees like the synchronous Put provides.
// There is a small time window (typically 3 ms) in which the data may not have been committed durably yet.
func (box *DeviceBox) PutAsync(object *Device) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Device.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Device.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *DeviceBox) PutAll(objects []Device) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *DeviceBox) Get(id uint64) (*Device, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Device), nil
}

// Get reads all stored objects
func (box *DeviceBox) GetAll() ([]Device, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Device), nil
}

// Remove deletes a single object
func (box *DeviceBox) Remove(object *Device) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.Id))
}

// Creates a query with the given conditions. Use the fields of the Device_ struct to create conditions.
// Keep the *DeviceQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *DeviceBox) Query(conditions ...objectbox.Condition) *DeviceQuery {
	return &DeviceQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Device_ struct to create conditions.
// Keep the *DeviceQuery if you intend to execute the query multiple times.
func (box *DeviceBox) QueryOrError(conditions ...objectbox.Condition) (*DeviceQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &DeviceQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Device which Id is either 42 or 47:
// 		box.Query(Device_.Id.In(42, 47)).Find()
type DeviceQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *DeviceQuery) Find() ([]Device, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Device), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *DeviceQuery) Offset(offset uint64) *DeviceQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *DeviceQuery) Limit(limit uint64) *DeviceQuery {
	query.Query.Limit(limit)
	return query
}
