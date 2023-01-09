package model

import (
	"gorm.io/gorm"
)

type Product struct {
	Nama     string `gorm:"varchar" json:"nama"`
	Kategori string `gorm:"varchar" json:"kategori"`
	Detail   string `gorm:"varchar" json:"detail"`
	Harga    int    `gorm:"int" json:"harga"`
	Stok     int    `gorm:"int" json:"stok"`
	Foto     string `gorm:"varchar" json:"foto"`
	gorm.Model
	ProdukTampil int `gorm:"default:1" json:"produkTampil" `
}
