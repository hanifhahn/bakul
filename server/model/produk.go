package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint   `gorm:"primaryKey"`
	Nama         string `gorm:"varchar" json:"nama"`
	Kategori     string `gorm:"varchar" json:"kategori"`
	Detail       string `gorm:"varchar" json:"detail"`
	Harga        int    `gorm:"int" json:"harga"`
	Stok         int    `gorm:"int" json:"stok"`
	Foto         string `gorm:"varchar" json:"foto"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ProdukTampil int            `gorm:"default:1" json:"produkTampil" `
}
