// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package obx

import (
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type deviceService_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var DeviceServiceBinding = deviceService_EntityInfo{
	Entity: objectbox.Entity{
		Id: 3,
	},
	Uid: 8592729144137394465,
}

// DeviceService_ contains type-based Property helpers to facilitate some common operations such as Queries.
var DeviceService_ = struct {
	Timestamps_Created  *objectbox.PropertyInt64
	Timestamps_Modified *objectbox.PropertyInt64
	Timestamps_Origin   *objectbox.PropertyInt64
	Description         *objectbox.PropertyString
	Id                  *objectbox.PropertyUint64
	Name                *objectbox.PropertyString
	LastConnected       *objectbox.PropertyInt64
	LastReported        *objectbox.PropertyInt64
	OperatingState      *objectbox.PropertyString
	Labels              *objectbox.PropertyStringVector
	Addressable         *objectbox.RelationToOne
	AdminState          *objectbox.PropertyString
}{
	Timestamps_Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Timestamps_Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Timestamps_Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	LastConnected: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	LastReported: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	OperatingState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Labels: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
	Addressable: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     11,
			Entity: &DeviceServiceBinding.Entity,
		},
		Target: &AddressableBinding.Entity,
	},
	AdminState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &DeviceServiceBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (deviceService_EntityInfo) GeneratorVersion() int {
	return 2
}

// AddToModel is called by ObjectBox during model build
func (deviceService_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("DeviceService", 3, 8592729144137394465)
	model.Property("Timestamps_Created", objectbox.PropertyType_Long, 1, 8937887139535500921)
	model.Property("Timestamps_Modified", objectbox.PropertyType_Long, 2, 3372180570590395340)
	model.Property("Timestamps_Origin", objectbox.PropertyType_Long, 3, 5493121179345774275)
	model.Property("Description", objectbox.PropertyType_String, 4, 8372447212547125896)
	model.Property("Id", objectbox.PropertyType_Long, 5, 1139277417242868914)
	model.PropertyFlags(objectbox.PropertyFlags_ID | objectbox.PropertyFlags_UNSIGNED)
	model.Property("Name", objectbox.PropertyType_String, 6, 6755402022460936928)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(2, 5111743529621501669)
	model.Property("LastConnected", objectbox.PropertyType_Long, 7, 3352512121287165304)
	model.Property("LastReported", objectbox.PropertyType_Long, 8, 3101703780598732288)
	model.Property("OperatingState", objectbox.PropertyType_String, 9, 7157873871038944207)
	model.Property("Labels", objectbox.PropertyType_StringVector, 10, 2994616260915039620)
	model.Property("Addressable", objectbox.PropertyType_Relation, 11, 2640177732493661289)
	model.PropertyFlags(objectbox.PropertyFlags_UNSIGNED)
	model.PropertyRelation("Addressable", 3, 4107923026042948129)
	model.Property("AdminState", objectbox.PropertyType_String, 12, 2249464337832106382)
	model.EntityLastPropertyId(12, 2249464337832106382)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (deviceService_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*DeviceService); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Service.Id), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(DeviceService).Service.Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (deviceService_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*DeviceService); ok {
		obj.Service.Id = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(DeviceService).Service.Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (deviceService_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	if rel := &object.(*DeviceService).Addressable; rel != nil {
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
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (deviceService_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *DeviceService
	if objPtr, ok := object.(*DeviceService); ok {
		obj = objPtr
	} else {
		objVal := object.(DeviceService)
		obj = &objVal
	}

	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.Service.DescribedObject.Description)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Service.Name)
	var offsetOperatingState = fbutils.CreateStringOffset(fbb, string(obj.Service.OperatingState))
	var offsetLabels = fbutils.CreateStringVectorOffset(fbb, obj.Service.Labels)
	var offsetAdminState = fbutils.CreateStringOffset(fbb, string(obj.AdminState))

	var rIdAddressable uint64
	if rel := &obj.Addressable; rel != nil {
		if rId, err := AddressableBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdAddressable = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(12)
	fbutils.SetInt64Slot(fbb, 0, obj.Service.DescribedObject.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Service.DescribedObject.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Service.DescribedObject.Timestamps.Origin)
	fbutils.SetUOffsetTSlot(fbb, 3, offsetDescription)
	fbutils.SetUint64Slot(fbb, 4, id)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetName)
	fbutils.SetInt64Slot(fbb, 6, obj.Service.LastConnected)
	fbutils.SetInt64Slot(fbb, 7, obj.Service.LastReported)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetOperatingState)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetLabels)
	fbutils.SetUint64Slot(fbb, 10, rIdAddressable)
	fbutils.SetUOffsetTSlot(fbb, 11, offsetAdminState)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (deviceService_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(12, 0)

	var relAddressable *Addressable
	if rId := table.GetUint64Slot(24, 0); rId > 0 {
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

	return &DeviceService{
		Service: Service{
			DescribedObject: DescribedObject{
				Timestamps: Timestamps{
					Created:  table.GetInt64Slot(4, 0),
					Modified: table.GetInt64Slot(6, 0),
					Origin:   table.GetInt64Slot(8, 0),
				},
				Description: fbutils.GetStringSlot(table, 10),
			},
			Id:             objectbox.StringIdConvertToEntityProperty(id),
			Name:           fbutils.GetStringSlot(table, 14),
			LastConnected:  table.GetInt64Slot(16, 0),
			LastReported:   table.GetInt64Slot(18, 0),
			OperatingState: models.OperatingState(fbutils.GetStringSlot(table, 20)),
			Labels:         fbutils.GetStringVectorSlot(table, 22),
			Addressable:    *relAddressable,
		},
		AdminState: models.AdminState(fbutils.GetStringSlot(table, 26)),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (deviceService_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]DeviceService, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (deviceService_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]DeviceService), *object.(*DeviceService))
}

// Box provides CRUD access to DeviceService objects
type DeviceServiceBox struct {
	*objectbox.Box
}

// BoxForDeviceService opens a box of DeviceService objects
func BoxForDeviceService(ob *objectbox.ObjectBox) *DeviceServiceBox {
	return &DeviceServiceBox{
		Box: ob.InternalBox(3),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Service.Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the DeviceService.Service.Id property on the passed object will be assigned the new ID as well.
func (box *DeviceServiceBox) Put(object *DeviceService) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the DeviceService.Service.Id property on the passed object will be assigned the new ID as well.
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
func (box *DeviceServiceBox) PutAsync(object *DeviceService) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Service.Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the DeviceService.Service.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the DeviceService.Service.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *DeviceServiceBox) PutAll(objects []DeviceService) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *DeviceServiceBox) Get(id uint64) (*DeviceService, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*DeviceService), nil
}

// Get reads all stored objects
func (box *DeviceServiceBox) GetAll() ([]DeviceService, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceService), nil
}

// Remove deletes a single object
func (box *DeviceServiceBox) Remove(object *DeviceService) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.Service.Id))
}

// Creates a query with the given conditions. Use the fields of the DeviceService_ struct to create conditions.
// Keep the *DeviceServiceQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *DeviceServiceBox) Query(conditions ...objectbox.Condition) *DeviceServiceQuery {
	return &DeviceServiceQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the DeviceService_ struct to create conditions.
// Keep the *DeviceServiceQuery if you intend to execute the query multiple times.
func (box *DeviceServiceBox) QueryOrError(conditions ...objectbox.Condition) (*DeviceServiceQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &DeviceServiceQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all DeviceService which Id is either 42 or 47:
// 		box.Query(DeviceService_.Id.In(42, 47)).Find()
type DeviceServiceQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *DeviceServiceQuery) Find() ([]DeviceService, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceService), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *DeviceServiceQuery) Offset(offset uint64) *DeviceServiceQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *DeviceServiceQuery) Limit(limit uint64) *DeviceServiceQuery {
	query.Query.Limit(limit)
	return query
}