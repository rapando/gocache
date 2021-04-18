package entities

import "time"

type DataStore struct {
	Data    string
	Created int64
	Expiry  int64
}

func (ds *DataStore) Save(data string, expiry int64) {
	ds.Data = data
	ds.Created = time.Now().Unix()
	ds.Expiry = expiry
}
