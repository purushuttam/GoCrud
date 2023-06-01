package model

type Note struct {
	Id      int    `gorm:"type:int;primary_key"`
	Conetnt string `gorm:"not null" json:"content,omitempty"`
}
