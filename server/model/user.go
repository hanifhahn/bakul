package model

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/hanifhahn/bakul/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	ID           uint   `gorm:"primaryKey"`
	Nama         string `gorm:"varchar" json:"nama"`
	NomorTelepon string `gorm:"varchar unique" json:"nomorTelepon"`
	Email        string `gorm:"varchar unique" json:"email"`
	Password     string `gorm:"varchar" json:"password"`
	Foto         string `gorm:"varchar" json:"foto"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserTampil   int            `gorm:"default:1" json:"userTampil" `
}

func GetUserByID(uid uint) (Users, error) {

	var u Users

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *Users) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := Users{}

	err = DB.Model(Users{}).Where("nama = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *Users) SaveUser() (*Users, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}

func (u *Users) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Nama = html.EscapeString(strings.TrimSpace(u.Nama))

	return nil
}
