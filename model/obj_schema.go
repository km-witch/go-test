package model

import (
	"errors"
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
	Updated_user int       `gorm:"column:updated_user" json:"updated_user" binding:"required"`
	IsActive     bool      `gorm:"column:is_active" json:"is_active"`
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoUpdateTime:milli"`
}

var ObjSchema *Obj
var Obj_msgSchema *Obj_msg

// Obj Msg 수정
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
	err := db.Model(&result).Where("is_active=true AND id=?", objid).Find(&result).Error

	return len(result), err
}

func (om *Obj_msg) GetObjMsgCountByUser(db *gorm.DB, uid int, objid string) (int, error) {
	var result []Obj_msg
	err := db.Model(&result).Where("is_active=true AND id=? AND created_user=?", objid, uid).Find(&result).Error

	return len(result), err
}
func (om *Obj_msg) GetObjMsgCountByUserAll(db *gorm.DB, uid int, objid string) (int, error) {
	var result []Obj_msg
	err := db.Model(&result).Where("id=? AND created_user=?", objid, uid).Find(&result).Error

	return len(result), err
}

// Obj Msg 조회
func (om *Obj_msg) CreateObjMsg(db *gorm.DB, message, oid string, uid int) (Obj_msg, error) {
	var result Obj_msg
	var obj Obj

	result.Message = message
	result.ObjId, _ = strconv.Atoi(oid)
	result.Created_user = uid
	result.Updated_user = uid
	result.IsActive = true

	db.Model(&result).Create(&result)
	db.Model(&obj).Find(&obj, result.ObjId)
	db.Model(Obj{}).Where("id=?", result.ObjId).UpdateColumn("amount", obj.Amount+1)

	return result, nil
}

// Block 조회 API ✅ [Block Key ⇒ 배치된 Obj.model]
func (o *Obj) GetObjByObjId(db *gorm.DB, obj_id string) (Obj, error) {
	var result Obj
	db.Model(&result).Find(&result, obj_id)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Obj id로 Obj Msg 조회
func (om *Obj_msg) GetObjMsgByObjId(db *gorm.DB, oid string) (Obj_msg, error) {
	var result Obj_msg

	db.Model(&result).Where("obj_id=?", oid).Find(&result)

	if result.Id == 0 {
		return result, errors.New("No Message")
	}
	return result, nil
}

// Obj Msg active 변경
func (om *Obj_msg) UpdateObjMsgIsActive(db *gorm.DB, obj_msg_id string) (Obj_msg, error) {
	var result Obj_msg
	var obj Obj

	db.Model(&result).Where("id=?", obj_msg_id).UpdateColumn("is_active", false)
	db.Model(&result).Where("id=?", obj_msg_id).Find(&result)

	db.Model(&obj).Find(&obj, result.ObjId)
	db.Model(Obj{}).Where("id=?", result.ObjId).UpdateColumn("amount", obj.Amount-1)

	return result, nil
}

// Block 조회 API ✅ [Block Key ⇒ 배치된 Obj.model]
func (o *Obj) GetObjByBlockId(db *gorm.DB, block_id string) (Obj, error) {
	var result Obj
	db.Where("block_id=?", block_id).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Block 조회단체로 API ✅ [Block Key ⇒ 배치된 Obj.model]
func (o *Obj) GetObjsByBlockId(db *gorm.DB, block_id string) ([]Obj, error) {
	var result []Obj
	db.Model(&result).Where("block_id=?", block_id).Find(&result)
	return result, nil
}

// Obj 조회 (Block 조회단체로 API ✅ [Block Key ⇒ 배치된 Obj.model])
func (o *Obj) GetObjsByUserId(db *gorm.DB, userId string) ([]Obj, error) {
	var result []Obj
	db.Model(&result).Where("user_id=?", userId).Find(&result)
	return result, nil
}

// Obj 생성
func (o *Obj) CreateObj(db *gorm.DB, objForm Obj) (Obj, error) {
	var userInput Obj
	userInput = objForm
	// NFT 생성 후 ID 값 받아서, 아래 Obj 생성 로직 추가 예정
	db.Create(&userInput)
	return userInput, nil
}

// Obj Msg 조회 Paging
func (om *Obj_msg) GetObjMsgs(db *gorm.DB, page, limit string) ([]Obj_msg, error) {
	var result []Obj_msg
	page_num, _ := strconv.Atoi(page)
	limit_num, _ := strconv.Atoi(limit)
	offset := (page_num - 1) * limit_num
	db.Limit(limit_num).Offset(offset).Order("created_at DESC").Where("is_active=1").Find(&result)

	if result == nil {
		return nil, errors.New("No ObjMsgs")
	}

	return result, nil
}

// Obj Msg 조회
func (om *Obj_msg) GetObjMsg(db *gorm.DB, ctx *gin.Context) {
	var result Obj_msg
	id := ctx.Param("id")

	db.Find(&result, id)

	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}

// Obj Msg 조회
func (om *Obj_msg) DeleteObjMsg(db *gorm.DB, ctx *gin.Context) {
	var result Obj_msg
	id := ctx.Param("id")

	db.Delete(&result, id)

	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}

// // Obj Msg 수정
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

// // 트리 Airdrop API [Airdrop NFT → OBJ 생성 → Obj Return;]
// // Block 조회 API ✅ [Block Key ⇒ 배치된 Obj.model]
// func (o *Obj) ObjAirdrop(db *gorm.DB, block_id string) (Obj, error) {
// 	var result Obj
// 	db.Where("block_id=?", block_id).Find(&result)

// 	if result.Id == 0 {
// 		return result, errors.New("NoContent")
// 	} else {
// 		return result, nil
// 	}
// }

// // Obj Msg 생성
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
