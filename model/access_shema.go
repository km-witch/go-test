package model

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Access struct {
	User_id     int       `json:"user_id"`
	Access_time time.Time `json:"Access_time"`
	Block_id    int       `json:"block_id"`
}

var AccessSchema *Access

// 방문로그 생성
func (a *Access) CreateAccess(db *gorm.DB, ctx *gin.Context) {
	var userInput Access

	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	db.Table("Access").Create(&userInput)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": userInput,
	})
}

// 방문로그 수정
func (a *Access) UpdateAccess(db *gorm.DB, ctx *gin.Context) {
	var userInput Access
	id := ctx.Param("id")
	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	db.Table("Access").Where("user_id=?", id).Updates(userInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}

// 방문로그 조회 By ID
func (a *Access) ReadAccess(db *gorm.DB, ctx *gin.Context) {
	var userInput Access
	id := ctx.Param("id")

	db.Table("Access").Where("user_id=?", id).Find(&userInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}
