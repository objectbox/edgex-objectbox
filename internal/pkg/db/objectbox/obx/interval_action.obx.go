// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type intervalAction_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var IntervalActionBinding = intervalAction_EntityInfo{
	Entity: objectbox.Entity{
		Id: 9,
	},
	Uid: 2756744796714279591,
}

// IntervalAction_ contains type-based Property helpers to facilitate some common operations such as Queries.
var IntervalAction_ = struct {
	ID         *objectbox.PropertyUint64
	Created    *objectbox.PropertyInt64
	Modified   *objectbox.PropertyInt64
	Origin     *objectbox.PropertyInt64
	Name       *objectbox.PropertyString
	Interval   *objectbox.PropertyString
	Parameters *objectbox.PropertyString
	Target     *objectbox.PropertyString
	Protocol   *objectbox.PropertyString
	HTTPMethod *objectbox.PropertyString
	Address    *objectbox.PropertyString
	Port       *objectbox.PropertyInt
	Path       *objectbox.PropertyString
	Publisher  *objectbox.PropertyString
	User       *objectbox.PropertyString
	Password   *objectbox.PropertyString
	Topic      *objectbox.PropertyString
}{
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Interval: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Parameters: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Target: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Protocol: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	HTTPMethod: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Address: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Port: &objectbox.PropertyInt{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Path: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Publisher: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     14,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	User: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     15,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Password: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     16,
			Entity: &IntervalActionBinding.Entity,
		},
	},
	Topic: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     17,
			Entity: &IntervalActionBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (intervalAction_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by ObjectBox during model build
func (intervalAction_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("IntervalAction", 9, 2756744796714279591)
	model.Property("ID", objectbox.PropertyType_Long, 1, 7420223465596233035)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Created", objectbox.PropertyType_Long, 2, 326969378931318024)
	model.Property("Modified", objectbox.PropertyType_Long, 3, 8212849254297936236)
	model.Property("Origin", objectbox.PropertyType_Long, 4, 738028464208107115)
	model.Property("Name", objectbox.PropertyType_String, 5, 4933633578353384552)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(23, 4596954898919109231)
	model.Property("Interval", objectbox.PropertyType_String, 6, 5283703314264311266)
	model.Property("Parameters", objectbox.PropertyType_String, 7, 3543707062021049254)
	model.Property("Target", objectbox.PropertyType_String, 8, 3745173396511053028)
	model.Property("Protocol", objectbox.PropertyType_String, 9, 133680632503980738)
	model.Property("HTTPMethod", objectbox.PropertyType_String, 10, 7777723197401595499)
	model.Property("Address", objectbox.PropertyType_String, 11, 7874742624379905478)
	model.Property("Port", objectbox.PropertyType_Long, 12, 4919605734646660559)
	model.Property("Path", objectbox.PropertyType_String, 13, 8245929477933611564)
	model.Property("Publisher", objectbox.PropertyType_String, 14, 4267442333403608480)
	model.Property("User", objectbox.PropertyType_String, 15, 1111650812930557467)
	model.Property("Password", objectbox.PropertyType_String, 16, 72978472924665700)
	model.Property("Topic", objectbox.PropertyType_String, 17, 8496152850176280729)
	model.EntityLastPropertyId(17, 8496152850176280729)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (intervalAction_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*IntervalAction); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.ID), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(IntervalAction).ID), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (intervalAction_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*IntervalAction); ok {
		obj.ID = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(IntervalAction).ID
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (intervalAction_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (intervalAction_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*IntervalAction)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetInterval = fbutils.CreateStringOffset(fbb, obj.Interval)
	var offsetParameters = fbutils.CreateStringOffset(fbb, obj.Parameters)
	var offsetTarget = fbutils.CreateStringOffset(fbb, obj.Target)
	var offsetProtocol = fbutils.CreateStringOffset(fbb, obj.Protocol)
	var offsetHTTPMethod = fbutils.CreateStringOffset(fbb, obj.HTTPMethod)
	var offsetAddress = fbutils.CreateStringOffset(fbb, obj.Address)
	var offsetPath = fbutils.CreateStringOffset(fbb, obj.Path)
	var offsetPublisher = fbutils.CreateStringOffset(fbb, obj.Publisher)
	var offsetUser = fbutils.CreateStringOffset(fbb, obj.User)
	var offsetPassword = fbutils.CreateStringOffset(fbb, obj.Password)
	var offsetTopic = fbutils.CreateStringOffset(fbb, obj.Topic)

	// build the FlatBuffers object
	fbb.StartObject(17)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, obj.Created)
	fbutils.SetInt64Slot(fbb, 2, obj.Modified)
	fbutils.SetInt64Slot(fbb, 3, obj.Origin)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetInterval)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetParameters)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetTarget)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetProtocol)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetHTTPMethod)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetAddress)
	fbutils.SetInt64Slot(fbb, 11, int64(obj.Port))
	fbutils.SetUOffsetTSlot(fbb, 12, offsetPath)
	fbutils.SetUOffsetTSlot(fbb, 13, offsetPublisher)
	fbutils.SetUOffsetTSlot(fbb, 14, offsetUser)
	fbutils.SetUOffsetTSlot(fbb, 15, offsetPassword)
	fbutils.SetUOffsetTSlot(fbb, 16, offsetTopic)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (intervalAction_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(4, 0)

	return &IntervalAction{
		ID:         objectbox.StringIdConvertToEntityProperty(id),
		Created:    table.GetInt64Slot(6, 0),
		Modified:   table.GetInt64Slot(8, 0),
		Origin:     table.GetInt64Slot(10, 0),
		Name:       fbutils.GetStringSlot(table, 12),
		Interval:   fbutils.GetStringSlot(table, 14),
		Parameters: fbutils.GetStringSlot(table, 16),
		Target:     fbutils.GetStringSlot(table, 18),
		Protocol:   fbutils.GetStringSlot(table, 20),
		HTTPMethod: fbutils.GetStringSlot(table, 22),
		Address:    fbutils.GetStringSlot(table, 24),
		Port:       int(table.GetUint64Slot(26, 0)),
		Path:       fbutils.GetStringSlot(table, 28),
		Publisher:  fbutils.GetStringSlot(table, 30),
		User:       fbutils.GetStringSlot(table, 32),
		Password:   fbutils.GetStringSlot(table, 34),
		Topic:      fbutils.GetStringSlot(table, 36),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (intervalAction_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]IntervalAction, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (intervalAction_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]IntervalAction), *object.(*IntervalAction))
}

// Box provides CRUD access to IntervalAction objects
type IntervalActionBox struct {
	*objectbox.Box
}

// BoxForIntervalAction opens a box of IntervalAction objects
func BoxForIntervalAction(ob *objectbox.ObjectBox) *IntervalActionBox {
	return &IntervalActionBox{
		Box: ob.InternalBox(9),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the IntervalAction.ID property on the passed object will be assigned the new ID as well.
func (box *IntervalActionBox) Put(object *IntervalAction) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the IntervalAction.ID property on the passed object will be assigned the new ID as well.
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
func (box *IntervalActionBox) PutAsync(object *IntervalAction) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the IntervalAction.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the IntervalAction.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *IntervalActionBox) PutAll(objects []IntervalAction) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *IntervalActionBox) Get(id uint64) (*IntervalAction, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*IntervalAction), nil
}

// Get reads all stored objects
func (box *IntervalActionBox) GetAll() ([]IntervalAction, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]IntervalAction), nil
}

// Remove deletes a single object
func (box *IntervalActionBox) Remove(object *IntervalAction) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.ID))
}

// Creates a query with the given conditions. Use the fields of the IntervalAction_ struct to create conditions.
// Keep the *IntervalActionQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *IntervalActionBox) Query(conditions ...objectbox.Condition) *IntervalActionQuery {
	return &IntervalActionQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the IntervalAction_ struct to create conditions.
// Keep the *IntervalActionQuery if you intend to execute the query multiple times.
func (box *IntervalActionBox) QueryOrError(conditions ...objectbox.Condition) (*IntervalActionQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &IntervalActionQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all IntervalAction which ID is either 42 or 47:
// 		box.Query(IntervalAction_.ID.In(42, 47)).Find()
type IntervalActionQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *IntervalActionQuery) Find() ([]IntervalAction, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]IntervalAction), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *IntervalActionQuery) Offset(offset uint64) *IntervalActionQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *IntervalActionQuery) Limit(limit uint64) *IntervalActionQuery {
	query.Query.Limit(limit)
	return query
}
