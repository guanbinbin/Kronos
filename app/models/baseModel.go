package models

import "time"

var Models = []interface{}{
	&User{}, &Article{},
}

type BaseModel struct {
	ID        uint64     `gorm:"primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" structs:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt" structs:"deletedAt"`
}
