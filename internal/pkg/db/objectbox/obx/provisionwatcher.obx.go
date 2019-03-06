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

type provisionWatcher_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var ProvisionWatcherBinding = provisionWatcher_EntityInfo{
	Entity: objectbox.Entity{
		Id: 12,
	},
	Uid: 7116913470590795583,
}

// ProvisionWatcher_ contains type-based Property helpers to facilitate some common operations such as Queries.
var ProvisionWatcher_ = struct {
	Created        *objectbox.PropertyInt64
	Modified       *objectbox.PropertyInt64
	Origin         *objectbox.PropertyInt64
	Id             *objectbox.PropertyUint64
	Name           *objectbox.PropertyString
	Identifiers    *objectbox.PropertyByteVector
	Profile        *objectbox.RelationToOne
	Service        *objectbox.RelationToOne
	OperatingState *objectbox.PropertyString
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Identifiers: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Profile: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     7,
			Entity: &ProvisionWatcherBinding.Entity,
		},
		Target: &DeviceProfileBinding.Entity,
	},
	Service: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     8,
			Entity: &ProvisionWatcherBinding.Entity,
		},
		Target: &DeviceServiceBinding.Entity,
	},
	OperatingState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (provisionWatcher_EntityInfo) GeneratorVersion() int {
	return 2
}

// AddToModel is called by ObjectBox during model build
func (provisionWatcher_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("ProvisionWatcher", 12, 7116913470590795583)
	model.Property("Created", objectbox.PropertyType_Long, 1, 1038690664624318335)
	model.Property("Modified", objectbox.PropertyType_Long, 2, 580560849680174672)
	model.Property("Origin", objectbox.PropertyType_Long, 3, 8778504568219244319)
	model.Property("Id", objectbox.PropertyType_Long, 4, 5681197862269956959)
	model.PropertyFlags(objectbox.PropertyFlags_ID | objectbox.PropertyFlags_UNSIGNED)
	model.Property("Name", objectbox.PropertyType_String, 5, 4473016749449217532)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(12, 5394211294854382713)
	model.Property("Identifiers", objectbox.PropertyType_ByteVector, 6, 7746019713902694696)
	model.Property("Profile", objectbox.PropertyType_Relation, 7, 2681033652138232267)
	model.PropertyFlags(objectbox.PropertyFlags_UNSIGNED)
	model.PropertyRelation("DeviceProfile", 13, 5881067106611825786)
	model.Property("Service", objectbox.PropertyType_Relation, 8, 788498092666620726)
	model.PropertyFlags(objectbox.PropertyFlags_UNSIGNED)
	model.PropertyRelation("DeviceService", 14, 3456748412572729670)
	model.Property("OperatingState", objectbox.PropertyType_String, 9, 4988194061826642670)
	model.EntityLastPropertyId(9, 4988194061826642670)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (provisionWatcher_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*ProvisionWatcher); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(ProvisionWatcher).Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (provisionWatcher_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*ProvisionWatcher); ok {
		obj.Id = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(ProvisionWatcher).Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (provisionWatcher_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	if rel := &object.(*ProvisionWatcher).Profile; rel != nil {
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
	if rel := &object.(*ProvisionWatcher).Service; rel != nil {
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
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (provisionWatcher_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *ProvisionWatcher
	if objPtr, ok := object.(*ProvisionWatcher); ok {
		obj = objPtr
	} else {
		objVal := object.(ProvisionWatcher)
		obj = &objVal
	}

	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetIdentifiers = fbutils.CreateByteVectorOffset(fbb, mapStringStringJsonToDatabaseValue(obj.Identifiers))
	var offsetOperatingState = fbutils.CreateStringOffset(fbb, string(obj.OperatingState))

	var rIdProfile uint64
	if rel := &obj.Profile; rel != nil {
		if rId, err := DeviceProfileBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdProfile = rId
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

	// build the FlatBuffers object
	fbb.StartObject(9)
	fbutils.SetInt64Slot(fbb, 0, obj.BaseObject.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.BaseObject.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.BaseObject.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetIdentifiers)
	fbutils.SetUint64Slot(fbb, 6, rIdProfile)
	fbutils.SetUint64Slot(fbb, 7, rIdService)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetOperatingState)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (provisionWatcher_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(10, 0)

	var relProfile *DeviceProfile
	if rId := table.GetUint64Slot(16, 0); rId > 0 {
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

	var relService *DeviceService
	if rId := table.GetUint64Slot(18, 0); rId > 0 {
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

	return &ProvisionWatcher{
		BaseObject: models.BaseObject{
			Created:  table.GetInt64Slot(4, 0),
			Modified: table.GetInt64Slot(6, 0),
			Origin:   table.GetInt64Slot(8, 0),
		},
		Id:             objectbox.StringIdConvertToEntityProperty(id),
		Name:           fbutils.GetStringSlot(table, 12),
		Identifiers:    mapStringStringJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 14)),
		Profile:        *relProfile,
		Service:        *relService,
		OperatingState: models.OperatingState(fbutils.GetStringSlot(table, 20)),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (provisionWatcher_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]ProvisionWatcher, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (provisionWatcher_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]ProvisionWatcher), *object.(*ProvisionWatcher))
}

// Box provides CRUD access to ProvisionWatcher objects
type ProvisionWatcherBox struct {
	*objectbox.Box
}

// BoxForProvisionWatcher opens a box of ProvisionWatcher objects
func BoxForProvisionWatcher(ob *objectbox.ObjectBox) *ProvisionWatcherBox {
	return &ProvisionWatcherBox{
		Box: ob.InternalBox(12),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the ProvisionWatcher.Id property on the passed object will be assigned the new ID as well.
func (box *ProvisionWatcherBox) Put(object *ProvisionWatcher) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the ProvisionWatcher.Id property on the passed object will be assigned the new ID as well.
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
func (box *ProvisionWatcherBox) PutAsync(object *ProvisionWatcher) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the ProvisionWatcher.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the ProvisionWatcher.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ProvisionWatcherBox) PutAll(objects []ProvisionWatcher) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ProvisionWatcherBox) Get(id uint64) (*ProvisionWatcher, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*ProvisionWatcher), nil
}

// Get reads all stored objects
func (box *ProvisionWatcherBox) GetAll() ([]ProvisionWatcher, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]ProvisionWatcher), nil
}

// Remove deletes a single object
func (box *ProvisionWatcherBox) Remove(object *ProvisionWatcher) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.Id))
}

// Creates a query with the given conditions. Use the fields of the ProvisionWatcher_ struct to create conditions.
// Keep the *ProvisionWatcherQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ProvisionWatcherBox) Query(conditions ...objectbox.Condition) *ProvisionWatcherQuery {
	return &ProvisionWatcherQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the ProvisionWatcher_ struct to create conditions.
// Keep the *ProvisionWatcherQuery if you intend to execute the query multiple times.
func (box *ProvisionWatcherBox) QueryOrError(conditions ...objectbox.Condition) (*ProvisionWatcherQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ProvisionWatcherQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all ProvisionWatcher which Id is either 42 or 47:
// 		box.Query(ProvisionWatcher_.Id.In(42, 47)).Find()
type ProvisionWatcherQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ProvisionWatcherQuery) Find() ([]ProvisionWatcher, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]ProvisionWatcher), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ProvisionWatcherQuery) Offset(offset uint64) *ProvisionWatcherQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ProvisionWatcherQuery) Limit(limit uint64) *ProvisionWatcherQuery {
	query.Query.Limit(limit)
	return query
}
