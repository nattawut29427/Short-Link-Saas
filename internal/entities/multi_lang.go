package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)


type MultiLang struct {
	Th string `json:"th"`
	En string `json:"en"`
}


func (m MultiLang) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MultiLang) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, m)
}
