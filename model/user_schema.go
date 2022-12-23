package model

import (
	"errors"
	"net/http"
	"pkg/configs"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Id             int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Uid            int       `gorm:"column:uid" json:"uid"`
	Email          string    `gorm:"column:email" json:"email"`
	IsActive       bool      `gorm:"column:is_active" json:"is_active"`
	EmailConfirmed bool      `gorm:"column:email_confirmed" json:"email_confirmed"`
	Role           int       `gorm:"column:role" json:"role"`
	Create_time    time.Time `gorm:"column:cteate_time;autoCreateTime" json:"create_time"`
	Password       int       `gorm:"column:password" json:"password"`
	AccessId       int       `gorm:"column:access_id" json:"access_id"`
	BlockId        int       `gorm:"column:block_id" json:"block_id"`
}

type AccessLog struct {
	Id         int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`
	BlockId    int       `gorm:"column:block_id" json:"block_id"`
	AccessTime time.Time `gorm:"autoCreateTime"`
}

type Profile struct {
	Id         int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`
	NickName   bool      `gorm:"column:nickname" json:"nickname"` //string
	Language   string    `gorm:"column:language" json:"language"`
	PfpNftId   int       `gorm:"column:pfp_nft_id" json:"pfp_nft_id"`
	Avatar     int       `gorm:"column:avatar" json:"avatar"` //string
	DefaultImg int       `gorm:"column:default_img_url" json:"default_img_url"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime:milli"`
}
type Wallet struct {
	Id         int       `gorm:"id;primarykey;autoIncrement" json:"id"`
	User_id    int       `json:"user_id" binding:"required"`
	Nft        []Nft     `json:"nft_id" gorm:"foreignKey:Id"`
	Wit        int       `json:"wit" gorm:"column:wit;default:300"`
	Created_at time.Time `json:"create_at" gorm:"column:created_at;autoCreateTime"`
	Updated_at time.Time `json:"update_at" gorm:"column:updated_at;autoUpdateTime"`
}

var UserSchema *User
var AccessLogSchema *AccessLog
var ProfileSchema *Profile
var WalletSchema *Wallet

// 유저 조회
func (u *User) FindUserByUid(db *gorm.DB, uid string) (User, error) {
	var result User
	db.Model(&result).Where("uid=?", uid).Find(&result)
	if result.Id == 0 {
		return result, errors.New("No User")
	}
	return result, nil
}
func (a *AccessLog) BlockAccess(db *gorm.DB, uid, bid int) (AccessLog, error) {
	var userInput AccessLog

	// access 기록 있는지 조회
	err := db.Model(&userInput).Where("user_id=?", uid).Find(&userInput).Error
	if err != nil {
		return userInput, err
	}

	// 있으면 update
	if userInput.UserId != 0 {
		err = db.Model(&userInput).Where("user_id=?", uid).Updates(userInput).Error
		return userInput, err
	}

	// 없으면 새로 생성
	userInput.UserId = uid
	userInput.BlockId = bid
	err = db.Model(&userInput).Create(&userInput).Error
	return userInput, err
}

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

// uid로 유저 생성
func (u *User) GetUserByUid(db *gorm.DB, uid int) (User, error) {
	var userInput User

	userInput.Uid = uid

	db.Model(&userInput).Where("uid=?", uid).Find(&userInput)

	if userInput.Id == 0 {
		return userInput, errors.New("NO Content")
	}
	return userInput, nil
}

// uid로 유저 생성
func (u *User) CreateUserByUid(db *gorm.DB, uid int) (User, error) {
	var userInput User

	userInput.Uid = uid

	err := db.Model(&userInput).Create(&userInput).Error

	return userInput, err
}

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

// 프로필 조회
func (p *Profile) GetProfileByUserId(db *gorm.DB, userId int) (int, error) {
	var profileForm Profile

	db.Model(&profileForm).Where("user_id=?", userId).Find(&profileForm)

	if profileForm.Id == 0 {
		return 0, errors.New("No wallet")
	}

	return profileForm.Id, nil
}

// 프로필 생성
func (p *Profile) CreateProfile(db *gorm.DB, userId int) (Profile, error) {
	var profileForm Profile
	profileForm.UserId = userId

	err := db.Model(&profileForm).Create(&profileForm).Error
	return profileForm, err
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
