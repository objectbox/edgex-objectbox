package objectbox

// implements export-client service contract

import (
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox/obx"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type exportClient struct {
	objectBox *objectbox.ObjectBox

	registrationBox *obx.RegistrationBox // no async - a config

	queries exportQueries
}

//region Queries
type exportQueries struct {
	registration struct {
		name registrationQuery
	}
}

type registrationQuery struct {
	*obx.RegistrationQuery
	sync.Mutex
}

//endregion

func newExportClient(objectBox *objectbox.ObjectBox) (*exportClient, error) {
	var client = &exportClient{objectBox: objectBox}
	var err error

	client.registrationBox = obx.BoxForRegistration(objectBox)

	//region Registration
	if err == nil {
		client.queries.registration.name.RegistrationQuery, err =
			client.registrationBox.QueryOrError(obx.Registration_.Name.Equals("", true))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, mapError(err)
	}
}

func (client *exportClient) Registrations() ([]contract.Registration, error) {
	result, err := client.registrationBox.GetAll()
	return result, mapError(err)
}

func (client *exportClient) AddRegistration(reg contract.Registration) (string, error) {
	// NOTE this is done instead of onCreate because there is no reg.Timestamps
	if reg.Created == 0 {
		reg.Created = db.MakeTimestamp()
	}

	id, err := client.registrationBox.Put(&reg)
	return obx.IdToString(id), mapError(err)
}

func (client *exportClient) UpdateRegistration(reg contract.Registration) error {
	// NOTE this is done instead of onUpdate because there is no reg.Timestamps
	reg.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(reg.ID); err != nil {
		return mapError(err)
	} else if exists, err := client.registrationBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.registrationBox.Put(&reg)
	return mapError(err)
}

func (client *exportClient) RegistrationById(id string) (contract.Registration, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Registration{}, mapError(err)
	} else if object, err := client.registrationBox.Get(id); err != nil {
		return contract.Registration{}, mapError(err)
	} else if object == nil {
		return contract.Registration{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *exportClient) DeleteRegistrationById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return mapError(err)
	} else {
		return mapError(client.registrationBox.RemoveId(id))
	}
}

func (client *exportClient) RegistrationByName(name string) (contract.Registration, error) {
	var query = &client.queries.registration.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Registration_.Name, name); err != nil {
		return contract.Registration{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.Registration{}, mapError(err)
	} else if len(list) == 0 {
		return contract.Registration{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *exportClient) DeleteRegistrationByName(name string) error {
	if obj, err := client.RegistrationByName(name); err != nil {
		return mapError(err)
	} else {
		return mapError(client.registrationBox.Remove(&obj))
	}
}

func (client *exportClient) ScrubAllRegistrations() error {
	return mapError(client.registrationBox.RemoveAll())
}
