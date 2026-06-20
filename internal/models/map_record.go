package models

import "time"

type MapRecord struct {
	MapId            int64
	MapPath          string
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}
