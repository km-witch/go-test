package model

import (
	"errors"
	"net/http"
	"pkg/configs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Wallet struct {
	Id      int `gorm:"id;primarykey"`
	User_id int `json:"user_id" binding:"required"`
	Wit     int `json:"wit" gorm:"default:300"`
	// Created_at time.Time `json:"create_at"`
	// Updated_at time.Time `json:"update_at"`
}

type ItemOwned struct {
	Id            int `gorm:"id;primarykey"`
	Owner_id      int `json:"owner_id" binding:"required"`
	Collection_id int `json:"collection_id" binding:"required"`
	Nft_id        int `json:"nft_id" binding:"required" gorm:"default:1"`
	Amount        int `json:"amount"`
	// Created_at time.Time `json:"create_at;not null"`
}

var WalletSchema *Wallet
var ItemOwnedSchema *ItemOwned

// 월렛 조회
func (w *Wallet) GetWalletByUserId(db *gorm.DB, userId string) (int, error) {
	var result Wallet
	configs.DB.Model(&result).Where("user_id=?", userId).Find(&result)
	if result.Id == 0 {
		return 0, errors.New("No wallet")
	}
	return result.Id, nil
}

// 월렛 생성
func (w *Wallet) CreateWallet(db *gorm.DB, userId int) (Wallet, error) {
	var walletForm Wallet
	walletForm.Wit = 300
	walletForm.User_id = userId

	db.Create(&walletForm)
	return walletForm, nil
}

// 월렛 수정
func (w *Wallet) UpdateWallet(db *gorm.DB, ctx *gin.Context) {
	var userInput Wallet
	id := ctx.Param("id")
	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	db.Table("wallet").Where("user_id=?", id).Updates(userInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}

// 아이템 보유생성
func (i *ItemOwned) CreateItemOwned(db *gorm.DB, ctx *gin.Context) {
	var userInput ItemOwned

	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	db.Table("ItemOwned").Create(&userInput)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": userInput,
	})
}

// 아이템 보유 업데이트
func (i *ItemOwned) UpdateItemOwned(db *gorm.DB, ctx *gin.Context) {
	var userInput ItemOwned
	id := ctx.Param("id")
	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	db.Table("ItemOwned").Where("user_id=?", id).Updates(userInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}

// 아이템 보유 확인
func (i *ItemOwned) CheckItemOwned(db *gorm.DB, ctx *gin.Context) {
	var userInput ItemOwned

	owner_id := ctx.Param("owner")
	collection_id := ctx.Param("collection")
	db.Table("ItemOwned").Where("owner_id=? AND collection_id=?", owner_id, collection_id).Find(&userInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}
