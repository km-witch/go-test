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
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoUpdateTime:milli"`
}

type Obj_msg struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement"` // PK
	ObjId        int       `gorm:"column:obj_id" json:"obj_id" binding:"required"`
	Type         string    `gorm:"column:type" json:"type" binding:"required"`
	Message      string    `gorm:"column:message" json:"message" binding:"required"`
	Created_user int       `gorm:"column:created_user" json:"created_user" binding:"required"`
	Updated_user int       `gorm:"column:updated_user" json:"updated_user" binding:"required"`
	IsActive     bool      `gorm:"column:is_active" json:"is_active"`
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoUpdateTime:milli"`
}

var ObjSchema *Obj
var Obj_msgSchema *Obj_msg

// Block 조회 API ✅ [Block Key ⇒ 배치된 Obj.model]
func (o *Obj) GetObjByBlockId(db *gorm.DB, block_id string) (Obj, error) {
	var result Obj
	db.Where("block_id=?", block_id).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
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
func (om *Obj_msg) GetObjMsgs(db *gorm.DB, ctx *gin.Context) {
	var result []Obj_msg
	page := ctx.Param("page")
	limit := ctx.Param("limit")
	page_num, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	limit_num, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	offset := (page_num - 1) * limit_num

	db.Limit(limit_num).Offset(offset).Order("created_time DESC").Where("is_active=?", true).Find(&result)

	if result == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
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
// // Block 조회 API ✅ [Block Key ⇒ 배치된 Obj.model]
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
