// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	"github.com/edgexfoundry/edgex-go/pkg/models"
	. "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type notification_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var NotificationBinding = notification_EntityInfo{
	Entity: objectbox.Entity{
		Id: 11,
	},
	Uid: 473217340861383887,
}

// Notification_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Notification_ = struct {
	Created     *objectbox.PropertyInt64
	Modified    *objectbox.PropertyInt64
	Origin      *objectbox.PropertyInt64
	ID          *objectbox.PropertyUint64
	Slug        *objectbox.PropertyString
	Sender      *objectbox.PropertyString
	Category    *objectbox.PropertyString
	Severity    *objectbox.PropertyString
	Content     *objectbox.PropertyString
	Description *objectbox.PropertyString
	Status      *objectbox.PropertyString
	Labels      *objectbox.PropertyStringVector
	ContentType *objectbox.PropertyString
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &NotificationBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &NotificationBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &NotificationBinding.Entity,
		},
	},
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &NotificationBinding.Entity,
		},
	},
	Slug: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &NotificationBinding.Entity,
		},
	},
	Sender: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &NotificationBinding.Entity,
		},
	},
	Category: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &NotificationBinding.Entity,
		},
	},
	Severity: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &NotificationBinding.Entity,
		},
	},
	Content: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &NotificationBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &NotificationBinding.Entity,
		},
	},
	Status: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &NotificationBinding.Entity,
		},
	},
	Labels: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &NotificationBinding.Entity,
		},
	},
	ContentType: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &NotificationBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (notification_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by ObjectBox during model build
func (notification_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Notification", 11, 473217340861383887)
	model.Property("Created", objectbox.PropertyType_Long, 1, 8097807178986063059)
	model.Property("Modified", objectbox.PropertyType_Long, 2, 8972589080235606015)
	model.Property("Origin", objectbox.PropertyType_Long, 3, 7004422487931328765)
	model.Property("ID", objectbox.PropertyType_Long, 4, 8366513294708486502)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Slug", objectbox.PropertyType_String, 5, 9191193423331162482)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(11, 5929069448268439040)
	model.Property("Sender", objectbox.PropertyType_String, 6, 3293128717989344156)
	model.Property("Category", objectbox.PropertyType_String, 7, 6717911931500064874)
	model.Property("Severity", objectbox.PropertyType_String, 8, 5601160557432732794)
	model.Property("Content", objectbox.PropertyType_String, 9, 2552815875989069884)
	model.Property("Description", objectbox.PropertyType_String, 10, 1885790405613350261)
	model.Property("Status", objectbox.PropertyType_String, 11, 4717815493415963654)
	model.Property("Labels", objectbox.PropertyType_StringVector, 12, 8144753170650843613)
	model.Property("ContentType", objectbox.PropertyType_String, 13, 1190301343113188450)
	model.EntityLastPropertyId(13, 1190301343113188450)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (notification_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Notification); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.ID), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Notification).ID), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (notification_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*Notification); ok {
		obj.ID = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Notification).ID
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (notification_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (notification_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Notification)
	var offsetSlug = fbutils.CreateStringOffset(fbb, obj.Slug)
	var offsetSender = fbutils.CreateStringOffset(fbb, obj.Sender)
	var offsetCategory = fbutils.CreateStringOffset(fbb, string(obj.Category))
	var offsetSeverity = fbutils.CreateStringOffset(fbb, string(obj.Severity))
	var offsetContent = fbutils.CreateStringOffset(fbb, obj.Content)
	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.Description)
	var offsetStatus = fbutils.CreateStringOffset(fbb, string(obj.Status))
	var offsetLabels = fbutils.CreateStringVectorOffset(fbb, obj.Labels)
	var offsetContentType = fbutils.CreateStringOffset(fbb, obj.ContentType)

	// build the FlatBuffers object
	fbb.StartObject(13)
	fbutils.SetInt64Slot(fbb, 0, obj.BaseObject.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.BaseObject.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.BaseObject.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetSlug)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetSender)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetCategory)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetSeverity)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetContent)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetDescription)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetStatus)
	fbutils.SetUOffsetTSlot(fbb, 11, offsetLabels)
	fbutils.SetUOffsetTSlot(fbb, 12, offsetContentType)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (notification_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(10, 0)

	return &Notification{
		BaseObject: models.BaseObject{
			Created:  table.GetInt64Slot(4, 0),
			Modified: table.GetInt64Slot(6, 0),
			Origin:   table.GetInt64Slot(8, 0),
		},
		ID:          objectbox.StringIdConvertToEntityProperty(id),
		Slug:        fbutils.GetStringSlot(table, 12),
		Sender:      fbutils.GetStringSlot(table, 14),
		Category:    models.NotificationsCategory(fbutils.GetStringSlot(table, 16)),
		Severity:    models.NotificationsSeverity(fbutils.GetStringSlot(table, 18)),
		Content:     fbutils.GetStringSlot(table, 20),
		Description: fbutils.GetStringSlot(table, 22),
		Status:      models.NotificationsStatus(fbutils.GetStringSlot(table, 24)),
		Labels:      fbutils.GetStringVectorSlot(table, 26),
		ContentType: fbutils.GetStringSlot(table, 28),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (notification_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Notification, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (notification_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]Notification), *object.(*Notification))
}

// Box provides CRUD access to Notification objects
type NotificationBox struct {
	*objectbox.Box
}

// BoxForNotification opens a box of Notification objects
func BoxForNotification(ob *objectbox.ObjectBox) *NotificationBox {
	return &NotificationBox{
		Box: ob.InternalBox(11),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Notification.ID property on the passed object will be assigned the new ID as well.
func (box *NotificationBox) Put(object *Notification) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Notification.ID property on the passed object will be assigned the new ID as well.
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
func (box *NotificationBox) PutAsync(object *Notification) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Notification.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Notification.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *NotificationBox) PutAll(objects []Notification) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *NotificationBox) Get(id uint64) (*Notification, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Notification), nil
}

// Get reads all stored objects
func (box *NotificationBox) GetAll() ([]Notification, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Notification), nil
}

// Remove deletes a single object
func (box *NotificationBox) Remove(object *Notification) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.ID))
}

// Creates a query with the given conditions. Use the fields of the Notification_ struct to create conditions.
// Keep the *NotificationQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *NotificationBox) Query(conditions ...objectbox.Condition) *NotificationQuery {
	return &NotificationQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Notification_ struct to create conditions.
// Keep the *NotificationQuery if you intend to execute the query multiple times.
func (box *NotificationBox) QueryOrError(conditions ...objectbox.Condition) (*NotificationQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &NotificationQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Notification which ID is either 42 or 47:
// 		box.Query(Notification_.ID.In(42, 47)).Find()
type NotificationQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *NotificationQuery) Find() ([]Notification, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Notification), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *NotificationQuery) Offset(offset uint64) *NotificationQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *NotificationQuery) Limit(limit uint64) *NotificationQuery {
	query.Query.Limit(limit)
	return query
}
