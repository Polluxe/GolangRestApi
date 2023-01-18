package models

type item struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaItem 	string `gorm:"type:varchar(300)" json:"nama_item"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
