package objectbox

// implements core-metadata service contract

import (
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox/obx"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type coreMetaDataClient struct {
	objectBox *objectbox.ObjectBox
	cmdClient *coreCommandClient // Temporary until core-command is completely separated from core-metadata

	addressableBox      *obx.AddressableBox      // no async - a config
	deviceBox           *obx.DeviceBox           // no async - a config/state
	deviceProfileBox    *obx.DeviceProfileBox    // no async - a config/state
	deviceReportBox     *obx.DeviceReportBox     // no async - has unique and requires insert/update to fail
	deviceServiceBox    *obx.DeviceServiceBox    // no async - a config/state, not for data ingress; has relations
	provisionWatcherBox *obx.ProvisionWatcherBox // no async - has relations

	queries coreMetaDataQueries
}

//region Queries
type coreMetaDataQueries struct {
	addressable struct {
		address   addressableQuery
		name      addressableQuery
		port      addressableQuery
		publisher addressableQuery
		topic     addressableQuery
	}
	deviceProfile struct {
		labels               deviceProfileQuery
		manufacturer         deviceProfileQuery
		manufacturerAndModel deviceProfileQuery
		model                deviceProfileQuery
		name                 deviceProfileQuery
	}
	device struct {
		labels  deviceQuery
		name    deviceQuery
		profile deviceQuery
		service deviceQuery
	}
	deviceReport struct {
		device deviceReportQuery
		event  deviceReportQuery
		name   deviceReportQuery
	}
	deviceService struct {
		addressable deviceServiceQuery
		labels      deviceServiceQuery
		name        deviceServiceQuery
	}
	provisionWatcher struct {
		//identifiers provisionWatcherQuery byte-vector .contains() query is currently not available
		name    provisionWatcherQuery
		profile provisionWatcherQuery
		service provisionWatcherQuery
	}
}

type addressableQuery struct {
	*obx.AddressableQuery
	sync.Mutex
}

type deviceQuery struct {
	*obx.DeviceQuery
	sync.Mutex
}

type deviceProfileQuery struct {
	*obx.DeviceProfileQuery
	sync.Mutex
}

type deviceReportQuery struct {
	*obx.DeviceReportQuery
	sync.Mutex
}

type deviceServiceQuery struct {
	*obx.DeviceServiceQuery
	sync.Mutex
}

type provisionWatcherQuery struct {
	*obx.ProvisionWatcherQuery
	sync.Mutex
}

//endregion

func newCoreMetaDataClient(objectBox *objectbox.ObjectBox, commandClient *coreCommandClient) (*coreMetaDataClient, error) {
	var client = &coreMetaDataClient{objectBox: objectBox, cmdClient: commandClient}
	var err error

	client.addressableBox = obx.BoxForAddressable(objectBox)
	client.deviceBox = obx.BoxForDevice(objectBox)
	client.deviceProfileBox = obx.BoxForDeviceProfile(objectBox)
	client.deviceReportBox = obx.BoxForDeviceReport(objectBox)
	client.deviceServiceBox = obx.BoxForDeviceService(objectBox)
	client.provisionWatcherBox = obx.BoxForProvisionWatcher(objectBox)

	//region Addressable
	if err == nil {
		client.queries.addressable.address.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Address.Equals("", true))
	}

	if err == nil {
		client.queries.addressable.name.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.addressable.port.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Port.Equals(0))
	}

	if err == nil {
		client.queries.addressable.publisher.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Publisher.Equals("", true))
	}

	if err == nil {
		client.queries.addressable.topic.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Topic.Equals("", true))
	}
	//endregion

	//region Device
	if err == nil {
		client.queries.device.labels.DeviceQuery, err =
			client.deviceBox.QueryOrError(obx.Device_.Labels.Contains("", true))
	}
	if err == nil {
		client.queries.device.name.DeviceQuery, err =
			client.deviceBox.QueryOrError(obx.Device_.Name.Equals("", true))
	}
	if err == nil {
		client.queries.device.profile.DeviceQuery, err =
			client.deviceBox.QueryOrError(obx.Device_.Profile.Equals(0))
	}
	if err == nil {
		client.queries.device.service.DeviceQuery, err =
			client.deviceBox.QueryOrError(obx.Device_.Service.Equals(0))
	}
	//endregion

	//region DeviceProfile
	if err == nil {
		client.queries.deviceProfile.labels.DeviceProfileQuery, err =
			client.deviceProfileBox.QueryOrError(obx.DeviceProfile_.Labels.Contains("", true))
	}
	if err == nil {
		client.queries.deviceProfile.manufacturer.DeviceProfileQuery, err =
			client.deviceProfileBox.QueryOrError(obx.DeviceProfile_.Manufacturer.Equals("", true))
	}
	if err == nil {
		client.queries.deviceProfile.manufacturerAndModel.DeviceProfileQuery, err =
			client.deviceProfileBox.QueryOrError(
				obx.DeviceProfile_.Manufacturer.Equals("", true),
				obx.DeviceProfile_.Model.Equals("", true))
	}
	if err == nil {
		client.queries.deviceProfile.model.DeviceProfileQuery, err =
			client.deviceProfileBox.QueryOrError(obx.DeviceProfile_.Model.Equals("", true))
	}
	if err == nil {
		client.queries.deviceProfile.name.DeviceProfileQuery, err =
			client.deviceProfileBox.QueryOrError(obx.DeviceProfile_.Name.Equals("", true))
	}
	//endregion

	//region DeviceReport
	if err == nil {
		client.queries.deviceReport.device.DeviceReportQuery, err =
			client.deviceReportBox.QueryOrError(obx.DeviceReport_.Device.Equals("", true))
	}
	if err == nil {
		client.queries.deviceReport.event.DeviceReportQuery, err =
			client.deviceReportBox.QueryOrError(obx.DeviceReport_.Action.Equals("", true))
	}
	if err == nil {
		client.queries.deviceReport.name.DeviceReportQuery, err =
			client.deviceReportBox.QueryOrError(obx.DeviceReport_.Name.Equals("", true))
	}
	//endregion

	//region DeviceService
	if err == nil {
		client.queries.deviceService.addressable.DeviceServiceQuery, err =
			client.deviceServiceBox.QueryOrError(obx.DeviceService_.Addressable.Equals(0))
	}

	if err == nil {
		client.queries.deviceService.labels.DeviceServiceQuery, err =
			client.deviceServiceBox.QueryOrError(obx.DeviceService_.Labels.Contains("", true))
	}

	if err == nil {
		client.queries.deviceService.name.DeviceServiceQuery, err =
			client.deviceServiceBox.QueryOrError(obx.DeviceService_.Name.Equals("", true))
	}
	//endregion

	//region ProvisionWatcher
	//if err == nil {
	//	client.queries.provisionWatcher.identifiers.ProvisionWatcherQuery, err =
	//		client.provisionWatcherBox.QueryOrError(obx.ProvisionWatcher_.Identifiers.Contains(nil))
	//}
	if err == nil {
		client.queries.provisionWatcher.name.ProvisionWatcherQuery, err =
			client.provisionWatcherBox.QueryOrError(obx.ProvisionWatcher_.Name.Equals("", true))
	}
	if err == nil {
		client.queries.provisionWatcher.profile.ProvisionWatcherQuery, err =
			client.provisionWatcherBox.QueryOrError(obx.ProvisionWatcher_.Profile.Equals(0))
	}
	if err == nil {
		client.queries.provisionWatcher.service.ProvisionWatcherQuery, err =
			client.provisionWatcherBox.QueryOrError(obx.ProvisionWatcher_.Service.Equals(0))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, mapError(err)
	}
}

func (client *coreMetaDataClient) GetAllDeviceReports() ([]contract.DeviceReport, error) {
	result, err := client.deviceReportBox.GetAll()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceReportByDeviceName(n string) ([]contract.DeviceReport, error) {
	var query = &client.queries.deviceReport.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceReport_.Device, n); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceReportByName(n string) (contract.DeviceReport, error) {
	var query = &client.queries.deviceReport.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceReport_.Name, n); err != nil {
		return contract.DeviceReport{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.DeviceReport{}, mapError(err)
	} else if len(list) == 0 {
		return contract.DeviceReport{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetDeviceReportById(id string) (contract.DeviceReport, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.DeviceReport{}, mapError(err)
	} else if object, err := client.deviceReportBox.Get(id); err != nil {
		return contract.DeviceReport{}, mapError(err)
	} else if object == nil {
		return contract.DeviceReport{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) AddDeviceReport(dr contract.DeviceReport) (string, error) {
	onCreate(&dr.Timestamps)

	id, err := client.deviceReportBox.Put(&dr)
	return obx.IdToString(id), mapError(err)
}

func (client *coreMetaDataClient) UpdateDeviceReport(dr contract.DeviceReport) error {
	onUpdate(&dr.Timestamps)

	if id, err := obx.IdFromString(dr.Id); err != nil {
		return mapError(err)
	} else if exists, err := client.deviceReportBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.deviceReportBox.Put(&dr)
	return mapError(err)
}

func (client *coreMetaDataClient) GetDeviceReportsByAction(n string) ([]contract.DeviceReport, error) {
	var query = &client.queries.deviceReport.event

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceReport_.Action, n); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) DeleteDeviceReportById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return mapError(err)
	} else {
		return mapError(client.deviceReportBox.RemoveId(id))
	}
}

func (client *coreMetaDataClient) UpdateDevice(d contract.Device) error {
	onUpdate(&d.Timestamps)

	if id, err := obx.IdFromString(d.Id); err != nil {
		return mapError(err)
	} else if exists, err := client.deviceBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.deviceBox.Put(&d)
	return mapError(err)
}

func (client *coreMetaDataClient) GetDeviceById(id string) (contract.Device, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Device{}, mapError(err)
	} else if object, err := client.deviceBox.Get(id); err != nil {
		return contract.Device{}, mapError(err)
	} else if object == nil {
		return contract.Device{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) GetDeviceByName(n string) (contract.Device, error) {
	var query = &client.queries.device.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Device_.Name, n); err != nil {
		return contract.Device{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.Device{}, mapError(err)
	} else if len(list) == 0 {
		return contract.Device{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetAllDevices() ([]contract.Device, error) {
	result, err := client.deviceBox.GetAll()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDevicesByProfileId(id string) ([]contract.Device, error) {
	var query = &client.queries.device.profile

	query.Lock()
	defer query.Unlock()

	if id, err := obx.IdFromString(id); err != nil {
		return nil, mapError(err)
	} else if err := query.SetInt64Params(obx.Device_.Profile, int64(id)); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDevicesByServiceId(id string) ([]contract.Device, error) {
	var query = &client.queries.device.service

	query.Lock()
	defer query.Unlock()

	if id, err := obx.IdFromString(id); err != nil {
		return nil, mapError(err)
	} else if err := query.SetInt64Params(obx.Device_.Service, int64(id)); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDevicesWithLabel(l string) ([]contract.Device, error) {
	var query = &client.queries.device.labels

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Device_.Labels, l); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) AddDevice(d contract.Device, commands []contract.Command) (string, error) {
	onCreate(&d.Timestamps)

	var id uint64
	var err = client.objectBox.RunInWriteTx(func() error {
		var err error
		id, err = client.deviceBox.Put(&d)
		if err != nil {
			return err
		}

		// In the future this should be a separate call, not part of AddDevice. See core-command.go
		for i := range commands {
			commands[i].Id, err = client.cmdClient.addCommand(commands[i], id)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return "", mapError(err)
	}

	return obx.IdToString(id), nil
}

func (client *coreMetaDataClient) DeleteDeviceById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return mapError(err)
	}

	return mapError(client.objectBox.RunInWriteTx(func() error {
		err = client.cmdClient.deleteCommandsByDeviceId(idString)
		if err != nil {
			return err
		}

		return client.deviceBox.RemoveId(id)
	}))
}

func (client *coreMetaDataClient) UpdateDeviceProfile(dp contract.DeviceProfile) error {
	onUpdate(&dp.Timestamps)

	// as we don't do lazy-loading externally, if the slice is nil, it's empty, not a "not-yet-loaded" lazy one
	if dp.DeviceCommands == nil {
		dp.DeviceCommands = []contract.ProfileResource{}
	}
	if dp.DeviceResources == nil {
		dp.DeviceResources = []contract.DeviceResource{}
	}
	if dp.CoreCommands == nil {
		dp.CoreCommands = []contract.Command{}
	}

	if id, err := obx.IdFromString(dp.Id); err != nil {
		return mapError(err)
	} else if exists, err := client.deviceProfileBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.deviceProfileBox.Put(&dp)
	return mapError(err)
}

func (client *coreMetaDataClient) AddDeviceProfile(d contract.DeviceProfile) (string, error) {
	onCreate(&d.Timestamps)

	id, err := client.deviceProfileBox.Put(&d)
	return obx.IdToString(id), mapError(err)
}

func (client *coreMetaDataClient) GetAllDeviceProfiles() ([]contract.DeviceProfile, error) {
	result, err := client.deviceProfileBox.GetAll()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceProfileById(id string) (contract.DeviceProfile, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.DeviceProfile{}, mapError(err)
	} else if object, err := client.deviceProfileBox.Get(id); err != nil {
		return contract.DeviceProfile{}, mapError(err)
	} else if object == nil {
		return contract.DeviceProfile{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) DeleteDeviceProfileById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return mapError(err)
	} else {
		return mapError(client.deviceProfileBox.RemoveId(id))
	}
}

func (client *coreMetaDataClient) GetDeviceProfilesByModel(m string) ([]contract.DeviceProfile, error) {
	var query = &client.queries.deviceProfile.model

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceProfile_.Model, m); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceProfilesWithLabel(l string) ([]contract.DeviceProfile, error) {
	var query = &client.queries.deviceProfile.labels

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceProfile_.Labels, l); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceProfilesByManufacturerModel(man string, mod string) ([]contract.DeviceProfile, error) {
	var query = &client.queries.deviceProfile.manufacturerAndModel

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceProfile_.Manufacturer, man); err != nil {
		return nil, mapError(err)
	}

	if err := query.SetStringParams(obx.DeviceProfile_.Model, mod); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceProfilesByManufacturer(man string) ([]contract.DeviceProfile, error) {
	var query = &client.queries.deviceProfile.manufacturer

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceProfile_.Manufacturer, man); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceProfileByName(n string) (contract.DeviceProfile, error) {
	var query = &client.queries.deviceProfile.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceProfile_.Name, n); err != nil {
		return contract.DeviceProfile{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.DeviceProfile{}, mapError(err)
	} else if len(list) == 0 {
		return contract.DeviceProfile{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) UpdateAddressable(a contract.Addressable) error {
	onUpdate(&a.Timestamps)

	if id, err := obx.IdFromString(a.Id); err != nil {
		return mapError(err)
	} else if exists, err := client.addressableBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.addressableBox.Put(&a)
	return mapError(err)
}

func (client *coreMetaDataClient) AddAddressable(a contract.Addressable) (string, error) {
	onCreate(&a.Timestamps)

	id, err := client.addressableBox.Put(&a)
	return obx.IdToString(id), mapError(err)
}

func (client *coreMetaDataClient) GetAddressableById(id string) (contract.Addressable, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Addressable{}, mapError(err)
	} else if object, err := client.addressableBox.Get(id); err != nil {
		return contract.Addressable{}, mapError(err)
	} else if object == nil {
		return contract.Addressable{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) GetAddressableByName(n string) (contract.Addressable, error) {
	var query = &client.queries.addressable.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Name, n); err != nil {
		return contract.Addressable{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.Addressable{}, mapError(err)
	} else if len(list) == 0 {
		return contract.Addressable{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetAddressablesByTopic(t string) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.topic

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Topic, t); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetAddressablesByPort(p int) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.port

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Addressable_.Port, int64(p)); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetAddressablesByPublisher(p string) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.publisher

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Publisher, p); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetAddressablesByAddress(add string) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.address

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Address, add); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetAddressables() ([]contract.Addressable, error) {
	result, err := client.addressableBox.GetAll()
	return result, mapError(err)
}

func (client *coreMetaDataClient) DeleteAddressableById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return mapError(err)
	} else {
		return mapError(client.addressableBox.RemoveId(id))
	}
}

func (client *coreMetaDataClient) UpdateDeviceService(ds contract.DeviceService) error {
	onUpdate(&ds.Timestamps)

	if id, err := obx.IdFromString(ds.Id); err != nil {
		return mapError(err)
	} else if exists, err := client.deviceServiceBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.deviceServiceBox.Put(&ds)
	return mapError(err)
}

func (client *coreMetaDataClient) GetDeviceServicesByAddressableId(id string) ([]contract.DeviceService, error) {
	var query = &client.queries.deviceService.addressable

	query.Lock()
	defer query.Unlock()

	if id, err := obx.IdFromString(id); err != nil {
		return nil, mapError(err)
	} else if err := query.SetInt64Params(obx.DeviceService_.Addressable, int64(id)); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceServicesWithLabel(l string) ([]contract.DeviceService, error) {
	var query = &client.queries.deviceService.labels

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceService_.Labels, l); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetDeviceServiceById(id string) (contract.DeviceService, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.DeviceService{}, mapError(err)
	} else if object, err := client.deviceServiceBox.Get(id); err != nil {
		return contract.DeviceService{}, mapError(err)
	} else if object == nil {
		return contract.DeviceService{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) GetDeviceServiceByName(n string) (contract.DeviceService, error) {
	var query = &client.queries.deviceService.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceService_.Name, n); err != nil {
		return contract.DeviceService{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.DeviceService{}, mapError(err)
	} else if len(list) == 0 {
		return contract.DeviceService{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetAllDeviceServices() ([]contract.DeviceService, error) {
	result, err := client.deviceServiceBox.GetAll()
	return result, mapError(err)
}

func (client *coreMetaDataClient) AddDeviceService(ds contract.DeviceService) (string, error) {
	onCreate(&ds.Timestamps)
	id, err := client.deviceServiceBox.Put(&ds)
	return obx.IdToString(id), mapError(err)
}

func (client *coreMetaDataClient) DeleteDeviceServiceById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return mapError(err)
	}

	return mapError(client.deviceServiceBox.RemoveId(id))
}

func (client *coreMetaDataClient) GetProvisionWatcherById(id string) (contract.ProvisionWatcher, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.ProvisionWatcher{}, mapError(err)
	} else if object, err := client.provisionWatcherBox.Get(id); err != nil {
		return contract.ProvisionWatcher{}, mapError(err)
	} else if object == nil {
		return contract.ProvisionWatcher{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) GetAllProvisionWatchers() ([]contract.ProvisionWatcher, error) {
	result, err := client.provisionWatcherBox.GetAll()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetProvisionWatcherByName(n string) (contract.ProvisionWatcher, error) {
	var query = &client.queries.provisionWatcher.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ProvisionWatcher_.Name, n); err != nil {
		return contract.ProvisionWatcher{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.ProvisionWatcher{}, mapError(err)
	} else if len(list) == 0 {
		return contract.ProvisionWatcher{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetProvisionWatchersByProfileId(id string) ([]contract.ProvisionWatcher, error) {
	var query = &client.queries.provisionWatcher.profile

	query.Lock()
	defer query.Unlock()

	if id, err := obx.IdFromString(id); err != nil {
		return nil, mapError(err)
	} else if err := query.SetInt64Params(obx.ProvisionWatcher_.Profile, int64(id)); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetProvisionWatchersByServiceId(id string) ([]contract.ProvisionWatcher, error) {
	var query = &client.queries.provisionWatcher.service

	query.Lock()
	defer query.Unlock()

	if id, err := obx.IdFromString(id); err != nil {
		return nil, mapError(err)
	} else if err := query.SetInt64Params(obx.ProvisionWatcher_.Service, int64(id)); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreMetaDataClient) GetProvisionWatchersByIdentifier(k string, v string) ([]contract.ProvisionWatcher, error) {
	// Identifiers is a map[string]string - and is stored as JSON []byte
	// We can create a query that contains "name":"value". While this is not 100 % correct (could have problems with
	// special characters), it's the way reference DB implementations do it so we assume there are no special characters
	// TODO currently not possible because ObjectBox ByteVector doesn't support contains query
	// var query = &client.queries.provisionWatcher.identifiers
	// ...

	// TODO can we make this more efficient?
	//  The biggest problem is that ProvisionWatcher contains relations which are loaded eagerly
	//  options are
	// 	 - query on identifiers.contains(`"name":"value"`)
	//   - current code with lazy-loading relations
	//   - "property" query (only returns a single property)
	if all, err := client.provisionWatcherBox.GetAll(); err != nil {
		return nil, mapError(err)
	} else {
		result := make([]contract.ProvisionWatcher, 0)
		for _, watcher := range all {
			if foundValue, found := watcher.Identifiers[k]; found {
				if foundValue == v {
					result = append(result, watcher)
				}
			}
		}
		return result, nil
	}
}

func (client *coreMetaDataClient) AddProvisionWatcher(pw contract.ProvisionWatcher) (string, error) {
	onCreate(&pw.Timestamps)

	id, err := client.provisionWatcherBox.Put(&pw)
	return obx.IdToString(id), mapError(err)
}

func (client *coreMetaDataClient) UpdateProvisionWatcher(pw contract.ProvisionWatcher) error {
	onUpdate(&pw.Timestamps)

	if id, err := obx.IdFromString(pw.Id); err != nil {
		return mapError(err)
	} else if exists, err := client.provisionWatcherBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.provisionWatcherBox.Put(&pw)
	return mapError(err)
}

func (client *coreMetaDataClient) DeleteProvisionWatcherById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return mapError(err)
	} else {
		return mapError(client.provisionWatcherBox.RemoveId(id))
	}
}

func (client *coreMetaDataClient) ScrubMetadata() error {
	var err error

	if err == nil {
		err = client.addressableBox.RemoveAll()
	}

	if err == nil {
		// TODO remove when core-command is completely detached from metadata
		err = obx.BoxForCommand(client.objectBox).RemoveAll()
	}

	if err == nil {
		err = client.deviceBox.RemoveAll()
	}

	if err == nil {
		err = client.deviceProfileBox.RemoveAll()
	}

	if err == nil {
		err = client.deviceReportBox.RemoveAll()
	}

	if err == nil {
		err = client.deviceServiceBox.RemoveAll()
	}

	if err == nil {
		err = client.provisionWatcherBox.RemoveAll()
	}

	return mapError(err)
}
