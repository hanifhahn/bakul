package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanifhahn/bakul/model"
	"github.com/hanifhahn/bakul/utils/token"
)

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := model.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Login(c *gin.Context) {

	var user model.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.Users{}

	u.Nama = user.Nama
	u.NomorTelepon = user.NomorTelepon
	u.Email = user.Email
	u.Password = user.Password
	u.Foto = user.Foto
	u.UserTampil = user.UserTampil

	token, err := model.LoginCheck(u.Nama, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var user model.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.Users{}

	u.Nama = user.Nama
	u.NomorTelepon = user.NomorTelepon
	u.Email = user.Email
	u.Password = user.Password
	u.Foto = user.Foto
	u.UserTampil = user.UserTampil

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
