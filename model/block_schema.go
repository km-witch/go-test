package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Block struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Thema        string    `gorm:"column:thema" json:"thema"`
	User_id      int       `gorm:"column:user_id" json:"user_id"`
	Name         string    `gorm:"column:name" json:"name"`
	Created_time time.Time `gorm:"autoCreateTime" json:"created_time"`
	Updated_time time.Time `gorm:"autoUpdateTime" json:"updated_time"`
}

var BlockSchema *Block

// # 블록조회 By UserID
func (b *Block) GetBlock_ByUserId(db *gorm.DB, userid string) (Block, error) {
	var result Block
	db.Where("user_id=?", userid).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// 블록 DB 컬럼 생성
func (b *Block) CreateBlock(db *gorm.DB, blockData Block) (Block, error) {
	var result Block
	result = blockData
	db.Create(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// // 블록 DB 컬럼 삭제 By ID
// func (b *Block) RemoveBlock(db *gorm.DB, ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	db.Delete(&Block{}, id)
// 	ctx.JSON(http.StatusOK, nil)
// }
