// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	"github.com/edgexfoundry/edgex-go/pkg/models"
	. "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type scheduleEvent_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var ScheduleEventBinding = scheduleEvent_EntityInfo{
	Id:  6,
	Uid: 6239202972921638608,
}

// ScheduleEvent_ contains type-based Property helpers to facilitate some common operations such as Queries.
var ScheduleEvent_ = struct {
	Created     *objectbox.PropertyInt64
	Modified    *objectbox.PropertyInt64
	Origin      *objectbox.PropertyInt64
	Id          *objectbox.PropertyUint64
	Name        *objectbox.PropertyString
	Schedule    *objectbox.PropertyString
	Addressable *objectbox.PropertyUint64
	Parameters  *objectbox.PropertyString
	Service     *objectbox.PropertyString
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id: 1,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id: 2,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id: 3,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id: 4,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id: 5,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Schedule: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id: 6,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Addressable: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id: 9,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Parameters: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id: 7,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
	Service: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id: 8,
			Entity: &objectbox.Entity{
				Id: 6,
			},
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (scheduleEvent_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by ObjectBox during model build
func (scheduleEvent_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("ScheduleEvent", 6, 6239202972921638608)
	model.Property("Created", objectbox.PropertyType_Long, 1, 8889399643704751207)
	model.Property("Modified", objectbox.PropertyType_Long, 2, 2715634671260664229)
	model.Property("Origin", objectbox.PropertyType_Long, 3, 2309830902551992394)
	model.Property("Id", objectbox.PropertyType_Long, 4, 1449553450431761107)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Name", objectbox.PropertyType_String, 5, 2643597168165916624)
	model.Property("Schedule", objectbox.PropertyType_String, 6, 5692020356904915195)
	model.Property("Addressable", objectbox.PropertyType_Relation, 9, 2699934787885094203)
	model.PropertyRelation("Addressable", 5, 118391878970996196)
	model.Property("Parameters", objectbox.PropertyType_String, 7, 7056282312459005473)
	model.Property("Service", objectbox.PropertyType_String, 8, 1885883891643404243)
	model.EntityLastPropertyId(9, 2699934787885094203)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (scheduleEvent_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*ScheduleEvent); ok {
		return bsonIdToDatabaseValue(obj.Id), nil
	} else {
		return bsonIdToDatabaseValue(object.(ScheduleEvent).Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (scheduleEvent_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*ScheduleEvent); ok {
		obj.Id = bsonIdToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(ScheduleEvent).Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (scheduleEvent_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {

	if rel := &object.(*ScheduleEvent).Addressable; rel != nil {
		rId, err := AddressableBinding.GetId(rel)
		if err != nil {
			return err
		} else if rId == 0 {
			if rCursor, err := txn.CursorForName("Addressable"); err != nil {
				return err
			} else if rId, err = rCursor.Put(rel); err != nil {
				return err
			}
		}
		// NOTE Put/PutAsync() has a side-effect of setting the rel.ID, so at this point, it is already set
	}

	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (scheduleEvent_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*ScheduleEvent)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetSchedule = fbutils.CreateStringOffset(fbb, obj.Schedule)
	var offsetParameters = fbutils.CreateStringOffset(fbb, obj.Parameters)
	var offsetService = fbutils.CreateStringOffset(fbb, obj.Service)

	var rIdAddressable uint64
	if rel := &obj.Addressable; rel != nil {
		if rId, err := AddressableBinding.GetId(rel); err != nil {
			panic(err) // this must never happen but let's keep the check just to be sure
		} else {
			rIdAddressable = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(9)
	fbutils.SetInt64Slot(fbb, 0, obj.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetSchedule)
	fbutils.SetUint64Slot(fbb, 8, rIdAddressable)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetParameters)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetService)
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (scheduleEvent_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) interface{} {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(10, 0)

	var relAddressable *Addressable
	if rId := table.GetUint64Slot(20, 0); rId > 0 {
		if cursor, err := txn.CursorForName("Addressable"); err != nil {
			panic(err)
		} else if relObject, err := cursor.Get(rId); err != nil {
			panic(err)
		} else if relObj, ok := relObject.(*Addressable); ok {
			relAddressable = relObj
		} else {
			var relObj = relObject.(Addressable)
			relAddressable = &relObj
		}
	} else {
		relAddressable = &Addressable{}
	}

	return &ScheduleEvent{
		BaseObject: models.BaseObject{
			Created:  table.GetInt64Slot(4, 0),
			Modified: table.GetInt64Slot(6, 0),
			Origin:   table.GetInt64Slot(8, 0),
		},
		Id:          bsonIdToEntityProperty(id),
		Name:        fbutils.GetStringSlot(table, 12),
		Schedule:    fbutils.GetStringSlot(table, 14),
		Addressable: *relAddressable,
		Parameters:  fbutils.GetStringSlot(table, 16),
		Service:     fbutils.GetStringSlot(table, 18),
	}
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (scheduleEvent_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]ScheduleEvent, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (scheduleEvent_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]ScheduleEvent), *object.(*ScheduleEvent))
}

// Box provides CRUD access to ScheduleEvent objects
type ScheduleEventBox struct {
	*objectbox.Box
}

// BoxForScheduleEvent opens a box of ScheduleEvent objects
func BoxForScheduleEvent(ob *objectbox.ObjectBox) *ScheduleEventBox {
	return &ScheduleEventBox{
		Box: ob.InternalBox(6),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the ScheduleEvent.Id property on the passed object will be assigned the new ID as well.
func (box *ScheduleEventBox) Put(object *ScheduleEvent) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the ScheduleEvent.Id property on the passed object will be assigned the new ID as well.
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
func (box *ScheduleEventBox) PutAsync(object *ScheduleEvent) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the ScheduleEvent.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the ScheduleEvent.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ScheduleEventBox) PutAll(objects []ScheduleEvent) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ScheduleEventBox) Get(id uint64) (*ScheduleEvent, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*ScheduleEvent), nil
}

// Get reads all stored objects
func (box *ScheduleEventBox) GetAll() ([]ScheduleEvent, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]ScheduleEvent), nil
}

// Remove deletes a single object
func (box *ScheduleEventBox) Remove(object *ScheduleEvent) (err error) {
	return box.Box.Remove(bsonIdToDatabaseValue(object.Id))
}

// Creates a query with the given conditions. Use the fields of the ScheduleEvent_ struct to create conditions.
// Keep the *ScheduleEventQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ScheduleEventBox) Query(conditions ...objectbox.Condition) *ScheduleEventQuery {
	return &ScheduleEventQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the ScheduleEvent_ struct to create conditions.
// Keep the *ScheduleEventQuery if you intend to execute the query multiple times.
func (box *ScheduleEventBox) QueryOrError(conditions ...objectbox.Condition) (*ScheduleEventQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ScheduleEventQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all ScheduleEvent which Id is either 42 or 47:
// 		box.Query(ScheduleEvent_.Id.In(42, 47)).Find()
type ScheduleEventQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ScheduleEventQuery) Find() ([]ScheduleEvent, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]ScheduleEvent), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ScheduleEventQuery) Offset(offset uint64) *ScheduleEventQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ScheduleEventQuery) Limit(limit uint64) *ScheduleEventQuery {
	query.Query.Limit(limit)
	return query
}
