package models

type Products struct {
	Id    int64  `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Stok  int32  `json:"stock"`
	Harga int32  `json:"harga"`
}
