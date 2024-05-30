package models

import (
	"time"
)

type Dalam struct {
	Id        string  `gorm:"type:varchar(300); primaryKey" json:"id"`
	Nama       string  `gorm:"type:varchar(250)" json:"nama"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}