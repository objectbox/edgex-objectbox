package obx

import (
	"encoding/json"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/globalsign/mgo/bson"
	"github.com/objectbox/objectbox-go/objectbox"
	"strconv"
)

func IdToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}
func IdFromString(id string) (uint64, error) {
	if id == "" {
		return 0, nil
	}
	return strconv.ParseUint(id, 10, 64)
}

func bsonIdToEntityProperty(dbValue uint64) bson.ObjectId {
	return bson.ObjectId(objectbox.StringIdConvertToEntityProperty(dbValue))
}

func bsonIdToDatabaseValue(goValue bson.ObjectId) uint64 {
	return objectbox.StringIdConvertToDatabaseValue(string(goValue))
}

func interfaceJsonToEntityProperty(dbValue []byte) interface{} {
	if dbValue == nil {
		return nil
	}

	var value interface{}
	if err := json.Unmarshal(dbValue, &value); err != nil {
		panic(err)
	} else {
		return value
	}
}

func interfaceJsonToDatabaseValue(goValue interface{}) []byte {
	if goValue == nil {
		return nil
	}

	if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func mapStringStringJsonToEntityProperty(dbValue []byte) map[string]string {
	if dbValue == nil {
		return nil
	}

	var value map[string]string
	if err := json.Unmarshal(dbValue, &value); err != nil {
		panic(err)
	} else {
		return value
	}
}

func mapStringStringJsonToDatabaseValue(goValue map[string]string) []byte {
	if goValue == nil {
		return nil
	}

	if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func notificationsCategoryToDatabaseValue(goValue []models.NotificationsCategory) []string {
	if goValue == nil {
		return nil
	}

	var result = make([]string, len(goValue))

	for k, v := range goValue {
		result[k] = string(v)
	}

	return result
}

func notificationsCategoryToEntityProperty(dbValue []string) []models.NotificationsCategory {
	if dbValue == nil {
		return nil
	}

	var result = make([]models.NotificationsCategory, len(dbValue))

	for k, v := range dbValue {
		result[k] = models.NotificationsCategory(v)
	}

	return result
}

func responsesJsonToEntityProperty(dbValue []byte) (result []models.Response) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func responsesJsonToDatabaseValue(goValue []models.Response) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func channelsJsonToEntityProperty(dbValue []byte) (result []models.Channel) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func channelsJsonToDatabaseValue(goValue []models.Channel) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func transmissionRecordsJsonToEntityProperty(dbValue []byte) (result []models.TransmissionRecord) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func transmissionRecordsJsonToDatabaseValue(goValue []models.TransmissionRecord) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func deviceResourcesJsonToEntityProperty(dbValue []byte) (result []models.DeviceResource) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func deviceResourcesJsonToDatabaseValue(goValue []models.DeviceResource) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func profileResourcesJsonToEntityProperty(dbValue []byte) (result []models.ProfileResource) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func profileResourcesJsonToDatabaseValue(goValue []models.ProfileResource) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}
