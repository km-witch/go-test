package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	UId       int  `gorm:"uid;primaryKey;autoIncrement"`
	Activated bool `gorm:"activated"`
	// Created_at       time.Time `gorm:"created_at;not null"`
	Created_by    string `gorm:"created_by"`
	Email         string `gorm:"column:email;not null;unique"`
	Email_confirm bool   `gorm:"column:email_confirm"`
	First_name    string `gorm:"column:first_name"`
	Gender        int    `gorm:"column:gender"`
	Id            string `gorm:"id;not null;unique"`
	Image_url     string `gorm:"image_url"`
	Lang_key      string `gorm:"lang_key"`
	// Last_modified    time.Time `gorm:"last_modified"`
	Last_name        string  `gorm:"last_name"`
	Nick_name        string  `gorm:"nick_name"`
	Nickname_changed bool    `gorm:"nickname_changed"`
	Nonce            float64 `gorm:"nonce"`
	Password         string  `gorm:"column:password"`
	Role             string  `gorm:"column:role"` // 0 UnSigned User || 1 Signed User || 9 Admin`
}

var UserSchema *User

// 유저 생성
func (u *User) CreateUser(db *gorm.DB, ctx *gin.Context) {
	var userInput User

	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	db.Table("user").Create(&userInput)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": userInput,
	})
}

// // 유저 조회
// func (u *User) UpdateUser(db *gorm.DB, ctx *gin.Context) {
// 	var userInput User
// 	id := ctx.Param("id")
// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// }
