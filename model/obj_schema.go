package model

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Obj struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement"` // PK
	User_id      int       `gorm:"column:user_id" json:"user_id" binding:"required"`
	Nft_id       int       `gorm:"column:nft_id" json:"nft_id"`
	Building_id  int       `gorm:"column:building_id" json:"building_id"`
	Block_id     int       `gorm:"column:block_id" json:"block_id" binding:"required"`
	Pos          string    `gorm:"column:pos" json:"pos" binding:"required"`
	Rot          string    `gorm:"column:rot" json:"rot" binding:"required"`
	Updated_user int       `gorm:"column:updated_user" json:"updated_user"`
	Amount       int       `gorm:"column:amount" json:"amount"`
	MsgRole      int       `gorm:"column:msg_role" json:"msg_role"`
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoUpdateTime:milli"`
}

type Obj_msg struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement"` // PK
	ObjId        int       `gorm:"column:obj_id" json:"obj_id" binding:"required"`
	Message      string    `gorm:"column:message" json:"message" binding:"required"`
	Created_user int       `gorm:"column:created_user" json:"created_user" binding:"required"`
	UserNickname string    `gorm:"column:user_nickname" json:"user_nickname" binding:"required"`
	Updated_user int       `gorm:"column:updated_user" json:"updated_user" binding:"required"`
	IsActive     bool      `gorm:"column:is_active" json:"is_active"`
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoUpdateTime:milli"`
}

type Obj_with_productid struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement"` // PK
	User_id      int       `gorm:"column:user_id" json:"user_id" binding:"required"`
	Nft_id       int       `gorm:"column:nft_id" json:"nft_id"`
	Product_id   int       `gorm:"column:product_id" json:"product_id"`
	Building_id  int       `gorm:"column:building_id" json:"building_id"`
	Block_id     int       `gorm:"column:block_id" json:"block_id" binding:"required"`
	Pos          string    `gorm:"column:pos" json:"pos" binding:"required"`
	Rot          string    `gorm:"column:rot" json:"rot" binding:"required"`
	Updated_user int       `gorm:"column:updated_user" json:"updated_user"`
	Amount       int       `gorm:"column:amount" json:"amount"`
	MsgRole      int       `gorm:"column:msg_role" json:"msg_role"`
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoUpdateTime:milli"`
}

var ObjSchema *Obj
var Obj_msgSchema *Obj_msg

// Obj Msg ??????
func (om *Obj_msg) UpdateObjMsg(db *gorm.DB, message, oid string, uid int) (Obj_msg, error) {
	var userInput Obj_msg
	userInput.Message = message
	userInput.ObjId, _ = strconv.Atoi(oid)
	userInput.Updated_user = uid
	userInput.IsActive = true

	db.Model(&userInput).Where("obj_id=?", oid).Updates(&userInput)
	db.Model(&userInput).Where("obj_id=?", oid).Find(&userInput)

	return userInput, nil
}

func (om *Obj_msg) GetObjMsgCount(db *gorm.DB, objid string) (int, error) {
	var result Obj
	err := db.Model(&result).Where("id=?", objid).Find(&result).Error

	return result.Amount, err
}

func (om *Obj_msg) GetObjActiveMsgCount(db *gorm.DB, objid string) (int, error) {
	var result []Obj_msg
	err := db.Model(&result).Where("is_active=true AND obj_id=?", objid).Find(&result).Error

	return len(result), err
}

func (om *Obj_msg) GetObjMsgCountByUser(db *gorm.DB, uid int, objid string) (int, error) {
	var result []Obj_msg
	err := db.Model(&result).Where("is_active=true AND id=? AND created_user=?", objid, uid).Find(&result).Error

	return len(result), err
}
func (om *Obj_msg) GetAllObjMsgCountByUser(db *gorm.DB, uid int, objid string) (int, error) {
	var result []Obj_msg
	err := db.Model(&result).Where("id=? AND created_user=?", objid, uid).Find(&result).Error

	return len(result), err
}

// Obj Msg ??????
func (om *Obj_msg) CreateObjMsg(db *gorm.DB, message, oid, nickname string, uid int) (Obj_msg, error) {
	var result Obj_msg
	var obj Obj

	result.Message = message
	result.ObjId, _ = strconv.Atoi(oid)
	result.Created_user = uid
	result.Updated_user = uid
	result.IsActive = true
	result.UserNickname = nickname

	err := db.Model(&result).Create(&result).Error
	if err != nil {
		return result, errors.New("Create Failed")
	}

	db.Model(&obj).Find(&obj, result.ObjId)
	db.Model(Obj{}).Where("id=?", result.ObjId).UpdateColumn("amount", obj.Amount+1)

	return result, nil
}

// Block ?????? API ??? [Block Key ??? ????????? Obj.model]
func (o *Obj) GetObjByObjId(db *gorm.DB, obj_id string) (Obj, error) {
	var result Obj
	db.Model(&result).Find(&result, obj_id)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Obj id??? Obj Msg ??????
func (om *Obj_msg) GetObjMsgByObjId(db *gorm.DB, oid string) (Obj_msg, error) {
	var result Obj_msg

	db.Model(&result).Where("obj_id=?", oid).Find(&result)

	if result.Id == 0 {
		return result, errors.New("No Message")
	}
	return result, nil
}

// Obj Msg active ??????
func (om *Obj_msg) UpdateObjMsgIsActive(db *gorm.DB, obj_msg_id string) (Obj_msg, error) {
	var result Obj_msg
	var obj Obj

	err := db.Model(&result).Where("id=?", obj_msg_id).UpdateColumn("is_active", false).Error
	if err != nil {
		return result, errors.New("Updated Failed")
	}

	db.Model(&result).Where("id=?", obj_msg_id).Find(&result)

	db.Model(&obj).Find(&obj, result.ObjId)
	db.Model(Obj{}).Where("id=?", result.ObjId).UpdateColumn("amount", obj.Amount-1)

	return result, nil
}

// Block ?????? API ??? [Block Key ??? ????????? Obj.model]
func (o *Obj) GetObjByBlockId(db *gorm.DB, block_id string) (Obj, error) {
	var result Obj
	db.Where("block_id=?", block_id).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Block ??????????????? API ??? [Block Key ??? ????????? Obj.model]
func (o *Obj) GetObjsByBlockId(db *gorm.DB, block_id string) ([]Obj, error) {
	var result []Obj
	db.Model(&result).Where("block_id=?", block_id).Find(&result)
	return result, nil
}

// Obj ?????? (Block ??????????????? API ??? [Block Key ??? ????????? Obj.model])
func (o *Obj) GetObjsByUserId(db *gorm.DB, userId string) ([]Obj, error) {
	var result []Obj
	db.Model(&result).Where("user_id=?", userId).Find(&result)
	return result, nil
}

// Obj ?????? (Block ??????????????? API ??? [Block Key ??? ????????? Obj.model])
func (o *Obj) GetObjsByUserIdWithProductId(db *gorm.DB, userId string) ([]Obj_with_productid, error) {
	var result []Obj_with_productid
	// db.Model(Obj{}).Where("user_id=?", userId).Find(&result)
	db.Raw("SELECT objs.*, nfts.product_id FROM objs INNER JOIN nfts ON nfts.id = objs.nft_id WHERE objs.user_id=?", userId).Find(&result)
	log.Println(result)
	return result, nil
}

// Obj ??????
func (o *Obj) CreateObj(db *gorm.DB, objForm Obj) (Obj, error) {
	var userInput Obj
	userInput = objForm
	// NFT ?????? ??? ID ??? ?????????, ?????? Obj ?????? ?????? ?????? ??????
	db.Create(&userInput)
	return userInput, nil
}

// Obj Msg ?????? Paging
func (om *Obj_msg) GetObjMsgs(db *gorm.DB, page, limit, objid string) ([]Obj_msg, error) {
	var result []Obj_msg
	page_num, _ := strconv.Atoi(page)
	limit_num, _ := strconv.Atoi(limit)
	offset := (page_num - 1) * limit_num
	db.Limit(limit_num).Offset(offset).Order("id DESC").Where("is_active=1 AND obj_id=?", objid).Find(&result)

	if result == nil {
		return nil, errors.New("No ObjMsgs")
	}

	return result, nil
}

// Obj Msg ??????
func (om *Obj_msg) GetObjMsg(db *gorm.DB, ctx *gin.Context) {
	var result Obj_msg
	id := ctx.Param("id")

	db.Find(&result, id)

	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}

// Obj Msg ??????
func (om *Obj_msg) DeleteObjMsg(db *gorm.DB, ctx *gin.Context) {
	var result Obj_msg
	id := ctx.Param("id")

	db.Delete(&result, id)

	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}

// // Obj Msg ??????
// func (om *Obj_msg) UpdateObjMsg(db *gorm.DB, ctx *gin.Context) {
// 	var userInput Obj
// 	id := ctx.Param("id")
// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"err": err,
// 		})
// 		return
// 	}
// 	db.Model(&userInput).Where("id=?", id).Updates(userInput)

// 	ctx.JSON(http.StatusOK, nil)
// }

// // ?????? Airdrop API [Airdrop NFT ??? OBJ ?????? ??? Obj Return;]
// // Block ?????? API ??? [Block Key ??? ????????? Obj.model]
// func (o *Obj) ObjAirdrop(db *gorm.DB, block_id string) (Obj, error) {
// 	var result Obj
// 	db.Where("block_id=?", block_id).Find(&result)

// 	if result.Id == 0 {
// 		return result, errors.New("NoContent")
// 	} else {
// 		return result, nil
// 	}
// }

// // Obj Msg ??????
// func (om *Obj_msg) CreateObjMsg(db *gorm.DB, ctx *gin.Context) {
// 	var userInput Obj_msg

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	db.Create(&userInput)

// 	ctx.JSON(http.StatusCreated, nil)
// }
