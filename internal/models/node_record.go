package models

import "time"

type NodeRecord struct {
	NodeId           int64
	Label            string
	X                int16
	Y                int16
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}
