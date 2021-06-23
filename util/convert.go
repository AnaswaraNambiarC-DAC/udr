// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package util

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udr/logger"
)

func MapToByte(data map[string]interface{}) []byte {
	ret, err := json.Marshal(data)
	if err != nil {
		logger.UtilLog.Error(err)
	}
	return ret
}

func MapArrayToByte(data []map[string]interface{}) []byte {
	ret, err := json.Marshal(data)
	if err != nil {
		logger.UtilLog.Error(err)
	}
	return ret
}

func PrimitiveAToByte(data []interface{}) []byte {
	ret, err := json.Marshal(data)
	if err != nil {
		logger.UtilLog.Error(err)
	}
	return ret
}

func ToBsonM(data interface{}) bson.M {
	tmp, err := json.Marshal(data)
	if err != nil {
		logger.UtilLog.Error(err)
	}
	putData := bson.M{}
	err = json.Unmarshal(tmp, &putData)
	if err != nil {
		logger.UtilLog.Error(err)
	}
	return putData
}

func SnssaiHexToModels(hexString string) (*models.Snssai, error) {
	sst, err := strconv.ParseInt(hexString[:2], 16, 32)
	if err != nil {
		return nil, err
	}
	sNssai := &models.Snssai{
		Sst: int32(sst),
		Sd:  hexString[2:],
	}
	return sNssai, nil
}

func SnssaiModelsToHex(snssai models.Snssai) string {
	sst := fmt.Sprintf("%02x", snssai.Sst)
	return sst + snssai.Sd
}
