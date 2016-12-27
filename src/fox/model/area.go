package model

type Area struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name            string `json:"name" xorm:"default '' CHAR(50)"`
	NameEn          string `json:"name_en" xorm:"default '' VARCHAR(100)"`
	ParentId        int    `json:"parent_id" xorm:"default 0 index INT(11)"`
	Type            int    `json:"type" xorm:"default 0 TINYINT(4)"`
	NameTraditional string `json:"name_traditional" xorm:"default '' VARCHAR(50)"`
	Sort            int    `json:"sort" xorm:"default 0 INT(11)"`
}