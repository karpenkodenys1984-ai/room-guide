package models

import "time"

type EdgeRecord struct {
	EdgeId           int64
	Source           int64
	Target           int64
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}
