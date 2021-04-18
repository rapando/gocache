package models

import (
	"strings"

	"github.com/rapando/gocache/entities"
)

func Save(params []string) (dataLocation *entities.DataStore, key string) {
	var data entities.DataStore
	key = params[0]
	payload := strings.Join(params[1:], " ")
	data.Save(payload, -1)
	return &data, key

}
