package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/middlewares"
	"github.com/hanifhahn/bakul/model"
	"github.com/hanifhahn/bakul/routes/produkcontroller"
	"github.com/hanifhahn/bakul/routes/usercontroller"
)

func main() {
	r := gin.Default()
	r.POST("/produk", produkcontroller.TambahProduk)
	r.GET("/produk", produkcontroller.LihatProduk)
	r.PATCH("/produk/:id", produkcontroller.UpdateProduk)
	r.DELETE("/produk/:id", produkcontroller.HapusProduk)

	r.PATCH("/archiveProduk/:id", produkcontroller.ArchiveProduk)
	r.PATCH("/restoreProduk/:id", produkcontroller.RestoreProduk)

	public := r.Group("/api")
	public.POST("/register", usercontroller.Register)
	public.POST("/login", usercontroller.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", usercontroller.CurrentUser)

	model.ConnectDB()

	fmt.Println("Hello world")

	r.Run()
}
