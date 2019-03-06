// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type event_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var EventBinding = event_EntityInfo{
	Entity: objectbox.Entity{
		Id: 8,
	},
	Uid: 3479337365075920111,
}

// Event_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Event_ = struct {
	ID       *objectbox.PropertyUint64
	Pushed   *objectbox.PropertyInt64
	Device   *objectbox.PropertyString
	Created  *objectbox.PropertyInt64
	Modified *objectbox.PropertyInt64
	Origin   *objectbox.PropertyInt64
	Event    *objectbox.PropertyString
	Readings *objectbox.RelationManyToMany
}{
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &EventBinding.Entity,
		},
	},
	Pushed: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &EventBinding.Entity,
		},
	},
	Device: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &EventBinding.Entity,
		},
	},
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &EventBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &EventBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &EventBinding.Entity,
		},
	},
	Event: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &EventBinding.Entity,
		},
	},
	Readings: &objectbox.RelationManyToMany{
		Id:     2,
		Source: &EventBinding.Entity,
		Target: &ReadingBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (event_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by ObjectBox during model build
func (event_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Event", 8, 3479337365075920111)
	model.Property("ID", objectbox.PropertyType_Long, 1, 7150544686233801611)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Pushed", objectbox.PropertyType_Long, 2, 2696612706629694030)
	model.Property("Device", objectbox.PropertyType_String, 3, 4778307009560157127)
	model.Property("Created", objectbox.PropertyType_Long, 4, 3583575667727096062)
	model.Property("Modified", objectbox.PropertyType_Long, 5, 3656471225173119384)
	model.Property("Origin", objectbox.PropertyType_Long, 6, 3524587915626770103)
	model.Property("Event", objectbox.PropertyType_String, 7, 9008105006658809123)
	model.EntityLastPropertyId(7, 9008105006658809123)
	model.Relation(2, 5533510483840274449, ReadingBinding.Id, ReadingBinding.Uid)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (event_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Event); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.ID), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Event).ID), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (event_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*Event); ok {
		obj.ID = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Event).ID
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (event_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	if err := txn.RunWithCursor(EventBinding.Id, func(cursor *objectbox.Cursor) error {
		return cursor.RelationReplace(2, ReadingBinding.Id, id, object, object.(*Event).Readings)
	}); err != nil {
		return err
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (event_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Event)
	var offsetDevice = fbutils.CreateStringOffset(fbb, obj.Device)
	var offsetEvent = fbutils.CreateStringOffset(fbb, obj.Event)

	// build the FlatBuffers object
	fbb.StartObject(7)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, obj.Pushed)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetDevice)
	fbutils.SetInt64Slot(fbb, 3, obj.Created)
	fbutils.SetInt64Slot(fbb, 4, obj.Modified)
	fbutils.SetInt64Slot(fbb, 5, obj.Origin)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetEvent)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (event_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(4, 0)

	var relReadings []Reading
	if err := txn.RunWithCursor(EventBinding.Id, func(cursor *objectbox.Cursor) error {
		if rSlice, err := cursor.RelationGetAll(2, ReadingBinding.Id, id); err != nil {
			return err
		} else {
			relReadings = rSlice.([]Reading)
			return nil
		}
	}); err != nil {
		return nil, err
	}

	return &Event{
		ID:       objectbox.StringIdConvertToEntityProperty(id),
		Pushed:   table.GetInt64Slot(6, 0),
		Device:   fbutils.GetStringSlot(table, 8),
		Created:  table.GetInt64Slot(10, 0),
		Modified: table.GetInt64Slot(12, 0),
		Origin:   table.GetInt64Slot(14, 0),
		Event:    fbutils.GetStringSlot(table, 16),
		Readings: relReadings,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (event_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Event, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (event_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]Event), *object.(*Event))
}

// Box provides CRUD access to Event objects
type EventBox struct {
	*objectbox.Box
}

// BoxForEvent opens a box of Event objects
func BoxForEvent(ob *objectbox.ObjectBox) *EventBox {
	return &EventBox{
		Box: ob.InternalBox(8),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Event.ID property on the passed object will be assigned the new ID as well.
func (box *EventBox) Put(object *Event) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Event.ID property on the passed object will be assigned the new ID as well.
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
func (box *EventBox) PutAsync(object *Event) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Event.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Event.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *EventBox) PutAll(objects []Event) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *EventBox) Get(id uint64) (*Event, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Event), nil
}

// Get reads all stored objects
func (box *EventBox) GetAll() ([]Event, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Event), nil
}

// Remove deletes a single object
func (box *EventBox) Remove(object *Event) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.ID))
}

// Creates a query with the given conditions. Use the fields of the Event_ struct to create conditions.
// Keep the *EventQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *EventBox) Query(conditions ...objectbox.Condition) *EventQuery {
	return &EventQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Event_ struct to create conditions.
// Keep the *EventQuery if you intend to execute the query multiple times.
func (box *EventBox) QueryOrError(conditions ...objectbox.Condition) (*EventQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &EventQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Event which ID is either 42 or 47:
// 		box.Query(Event_.ID.In(42, 47)).Find()
type EventQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *EventQuery) Find() ([]Event, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Event), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *EventQuery) Offset(offset uint64) *EventQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *EventQuery) Limit(limit uint64) *EventQuery {
	query.Query.Limit(limit)
	return query
}
