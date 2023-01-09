package produkcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/model"
)

func TambahProduk(c *gin.Context) {
	var produk model.Product

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.DB.Create(&produk)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil ditambahkan"})
}

func LihatProduk(c *gin.Context) {
	var produk []model.Product

	if err := model.DB.Where("produk_tampil = ?", 1).Find(&produk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data Kosong", "data": produk})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diambil", "data": produk})
	}
}

func UpdateProduk(c *gin.Context) {
	var produk model.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.DB.Model(&produk).Where("id = ?", id).Updates(&produk)

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate", "data": produk})
}

// Soft Delete - Menambah Value Delete At
func HapusProduk(c *gin.Context) {
	var produk model.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := model.DB.Model(&produk).Where("id = ?", id).Delete(&produk).Error; err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"message": "Data gagal dihapus", "data": produk})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus", "data": produk})
	}
}

// Sangat Soft Delete - Hanya Menonaktifkan Produk Agar Tidak Tampil
func ArchiveProduk(c *gin.Context) {
	var produk model.Product

	id := c.Param("id")

	if err := model.DB.Model(&produk).Where("id = ?", id).Update("produk_tampil", 0).Error; err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"message": "Data gagal di Archieve", "data": produk})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di Archieve", "data": produk})
	}
}

// Sangat Soft Delete - Mengaktifkan Produk Agar Tampil Kembali
func RestoreProduk(c *gin.Context) {
	var produk model.Product

	id := c.Param("id")

	if err := model.DB.Model(&produk).Where("id = ?", id).Update("produk_tampil", 1).Error; err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"message": "Data gagal di Restore", "data": produk})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di Restore", "data": produk})
	}
}
