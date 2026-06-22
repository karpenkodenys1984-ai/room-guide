package models

import "time"

type NodeRecord struct {
	NodeId           int64
	Label            string
	X                float32
	Y                float32
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}
