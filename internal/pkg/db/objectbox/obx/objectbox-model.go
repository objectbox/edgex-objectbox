// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

// ObjectBoxModel declares and builds the model from all the entities in the package.
// It is usually used when setting-up ObjectBox as an argument to the Builder.Model() function.
func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(1)

	model.RegisterBinding(AddressableBinding)
	model.RegisterBinding(CommandBinding)
	model.RegisterBinding(DeviceServiceBinding)
	model.RegisterBinding(EventBinding)
	model.RegisterBinding(ReadingBinding)
	model.RegisterBinding(ScheduleEventBinding)
	model.RegisterBinding(ValueDescriptorBinding)
	model.LastEntityId(7, 1150785711675845959)
	model.LastIndexId(5, 118391878970996196)
	model.LastRelationId(1, 7766060310030207431)

	return model
}
