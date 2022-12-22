package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement" json:"id"` // PK
	Name_ko      string    `gorm:"column:name_ko" json:"name_ko" binding:"required"`
	Name_en      string    `gorm:"column:name_en" json:"name_en" binding:"required"`
	Publisher_ko string    `gorm:"column:publisher_ko" json:"publisher_ko" binding:"required"`
	Publisher_en string    `gorm:"column:publisher_en" json:"publisher_en" binding:"required"`
	Opensea_url  string    `gorm:"column:opensea_url" json:"opensea_url"`
	Twitter_url  string    `gorm:"column:twitter_url" json:"twitter_url"`
	Discord_url  string    `gorm:"column:discord_url" json:"discord_url"`
	Ww_url       string    `gorm:"column:ww_url" json:"ww_url"`
	Created_time time.Time `gorm:"autoCreateTime" json:"created_time"`
}

type ProductGroup struct {
	Id             int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Collection_id  int       `gorm:"column:collection_id" json:"collection_id" binding:"required"`
	Thumbnail_url  string    `gorm:"column:thumbnail_url" json:"thumbnail_url"`
	Contract       string    `gorm:"column:contract" json:"contract" binding:"required"`
	Name_ko        string    `gorm:"column:name_ko" json:"name_ko"`
	Name_en        string    `gorm:"column:name_en" json:"name_en"`
	Description    string    `gorm:"column:description" json:"description"`
	Amount         int       `gorm:"column:amount" json:"amount"`
	Properties     string    `gorm:"column:properties" json:"properties"`
	Image_Url      string    `gorm:"column:image_url" json:"image_url"`
	Snap           string    `gorm:"column:snap" json:"snap"`
	Metadata       string    `gorm:"column:metadata" json:"metadata"`
	Message_role   string    `gorm:"column:message_role" json:"message_role"` // MSG Role을 여기서 주어야 하는가..에 대한 의문.
	Message_amount int       `gorm:"column:message_amount" json:"message_amount"`
	Created_time   time.Time `gorm:"autoCreateTime" json:"created_time"`
	Updated_time   time.Time `gorm:"autoUpdateTime" json:"updated_time"`
}

// FK 적용방식 구성하시오.
type Nft struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	TokenId      int       `gorm:"column:token_id" json:"token_id" binding:"required"`
	Product_id   int       `gorm:"column:product_id" json:"product_id" binding:"required"`
	Name_ko      string    `gorm:"column:name_ko" json:"name_ko" binding:"required"`
	Name_en      string    `gorm:"column:name_en" json:"name_en" binding:"required"`
	Description  string    `gorm:"column:description" json:"description" binding:"required"`
	Properties   string    `gorm:"column:properties" json:"properties"`
	Created_time time.Time `gorm:"autoCreateTime" json:"created_time"`
	Updated_time time.Time `gorm:"autoUpdateTime" json:"updated_time"`
}

type NftTx struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Method       int       `gorm:"column:method" json:"method" binding:"required"`
	From         int       `gorm:"column:from" json:"from" binding:"required"`
	To           int       `gorm:"column:to" json:"to" binding:"required"`
	Nftid        int       `gorm:"column:nft_id" json:"nft_id" binding:"required"`
	Created_time time.Time `gorm:"autoCreateTime" json:"created_time"`
}

var CollectionSchema *Collection
var ProductGroupSchema *ProductGroup
var NftSchema *Nft
var NftTxSchema *NftTx

// Get Collection By ID
func (c *Collection) GetCollectionById(db *gorm.DB, id string) (Collection, error) {
	var result Collection
	db.Model(&Collection{}).Where("id=?", id).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Get Group By ID
func (pg *ProductGroup) GetGroupById(db *gorm.DB, id string) (ProductGroup, error) {
	var result ProductGroup
	db.Model(&ProductGroup{}).Where("id=?", id).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Get NFT By ID
func (n *Nft) GetNftById(db *gorm.DB, id string) (Nft, error) {
	var result Nft
	db.Model(&Nft{}).Where("id=?", id).Find(&result)
	if result.Id == 0 {
		return result, errors.New("NoContent")
	} else {
		return result, nil
	}
}

// Create NFT By GroupID
func (n *Nft) CreateNftByGroupId(db *gorm.DB, groupid string) (Nft, error) {
	// 그룹 데이터를 조회해 정보 가져오기
	var groupData ProductGroup
	db.Model(&ProductGroup{}).Where("id=?", groupid).Find(&groupData)
	// NFT 데이터 스키마 구성할 것.
	// NFT DB 최신 인덱스 조회
	var CurrentNft Nft
	var NftForm Nft
	db.Where("product_id=?", groupid).Last(&CurrentNft)
	fmt.Println(CurrentNft)

	// Req Form 생성
	groupid_numb, _ := strconv.Atoi(groupid)
	NftForm.TokenId = CurrentNft.TokenId + 1
	nftid_str := strconv.Itoa(NftForm.TokenId)
	NftForm.Product_id = groupid_numb
	NftForm.Name_ko = fmt.Sprintf("%s%s%s", groupData.Name_ko, "#", nftid_str)
	NftForm.Name_en = fmt.Sprintf("%s%s%s", groupData.Name_en, "#", nftid_str)
	NftForm.Properties = groupData.Properties
	NftForm.Description = groupData.Description

	// NFT Table에 Insert
	db.Create(&NftForm)

	// Return NFT Inserted
	return NftForm, nil
}

func (tx *NftTx) CreateTx(db *gorm.DB, txInfo NftTx) (NftTx, error) {
	var transaction NftTx
	transaction = txInfo
	fmt.Println(transaction)
	db.Create(&transaction)
	return transaction, nil
}

// // Create Collection
// func (c *Collection) CreateCollection(db *gorm.DB, ctx *gin.Context) {
// 	var userInput Collection

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	db.Create(&userInput)

// 	ctx.JSON(http.StatusCreated, nil)
// }

// // Read Collection (Paging)
// func (c *Collection) ReadCollection(db *gorm.DB, ctx *gin.Context) {
// 	var result []Collection
// 	page := ctx.Param("page")
// 	limit := ctx.Param("limit")
// 	page_num, err := strconv.Atoi(page)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	limit_num, err := strconv.Atoi(limit)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	offset := (page_num - 1) * limit_num

// 	db.Limit(limit_num).Offset(offset).Order("created_time DESC").Find(&result)

// 	if result == nil {
// 		ctx.JSON(http.StatusNoContent, nil)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"payload": result,
// 	})
// }

// // Create Group
// func (p *ProductGroup) CreateGroup(db *gorm.DB, ctx *gin.Context) {
// 	var userInput ProductGroup

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	db.Create(&userInput)

// 	ctx.JSON(http.StatusCreated, nil)
// }

// // Read Group (Paging)
// func (p *ProductGroup) ReadGroup(db *gorm.DB, ctx *gin.Context) {
// 	var result []ProductGroup
// 	page := ctx.Param("page")
// 	limit := ctx.Param("limit")
// 	page_num, err := strconv.Atoi(page)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	limit_num, err := strconv.Atoi(limit)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	offset := (page_num - 1) * limit_num

// 	db.Limit(limit_num).Offset(offset).Order("created_time DESC").Find(&result)

// 	if result == nil {
// 		ctx.JSON(http.StatusNoContent, nil)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"payload": result,
// 	})
// }

// // Create NFT
// func (n *Nft) CreateNft(db *gorm.DB, ctx *gin.Context) {
// 	var userInput Nft

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	db.Create(&userInput)

// 	ctx.JSON(http.StatusCreated, nil)
// }

// // Read Nfts (Paging)
// func (n *Nft) ReadNfts(db *gorm.DB, ctx *gin.Context) {
// 	var result []Nft
// 	page := ctx.Param("page")
// 	limit := ctx.Param("limit")
// 	page_num, err := strconv.Atoi(page)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	limit_num, err := strconv.Atoi(limit)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	offset := (page_num - 1) * limit_num

// 	db.Limit(limit_num).Offset(offset).Order("created_time DESC").Find(&result)

// 	if result == nil {
// 		ctx.JSON(http.StatusNoContent, nil)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"payload": result,
// 	})
// }

// // Read One Nft By ID
// func (n *Nft) ReadNftById(db *gorm.DB, ctx *gin.Context) {
// 	var result Nft
// 	id := ctx.Param("id")
// 	db.Model(&result).Where("id=?", id).Find(&result)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"payload": result,
// 	})
// }
