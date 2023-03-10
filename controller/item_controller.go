package controller

import (
	"net/http"
	"pkg/configs"
	"pkg/model"

	"github.com/gin-gonic/gin"
)

// GetCollectionById            godoc
// @Summary      				Collection By ID
// @Description  				Collection ID를 넣으면 Collection을 리턴함
// @Tags        				Ready
// @Param        				collectionid  	path    string  true  "Write Block ID"
// @Produce      				json
// @Success      				200  {object}  model.Collection
// @Router       				/api/item/collection/{collectionid} [get]
func GetCollectionById(ctx *gin.Context) {
	id := ctx.Param("collectionid")
	resultCollection, err := model.CollectionSchema.GetCollectionById(configs.ConnectDB(), id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"payload": resultCollection,
		})
	}
}

// GetProductGroupById          godoc
// @Summary      				GetProductGroup By ID
// @Description  				Product Group ID를 넣으면 -> Group을 반환함.
// @Tags        				Ready
// @Param        				groupid  	path    string  true  "Write Block ID"
// @Produce      				json
// @Success      				200  {object}  model.ProductGroup
// @Router       				/api/item/group/{groupid} [get]
func GetProductGroupById(ctx *gin.Context) {
	id := ctx.Param("groupid")
	result, err := model.ProductGroupSchema.GetGroupById(configs.ConnectDB(), id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"payload": result,
		})
	}
}

// GetNftById           		godoc
// @Summary      				GetNftById By ID
// @Description  				Nft ID를 넣으면 -> NFT를 반환함.
// @Tags        				Ready
// @Param        				nftid  	path    string  true  "Write Nft ID"
// @Produce      				json
// @Success      				200  {object}  model.Nft
// @Router       				/api/item/nft/{nftid} [get]
func GetNftById(ctx *gin.Context) {
	id := ctx.Param("nftid")
	result, err := model.NftSchema.GetNftById(configs.ConnectDB(), id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"payload": result,
		})
	}
}

type Resp_CreateNftByGroupId struct {
	Nft model.Nft   `json:"nft"`
	Tx  model.NftTx `json:"tx"`
}

type ReqForm_CreateNftByGroupId struct {
	Groupid  string `json:"group_id"`
	Walletid string `json:"wallet_id"`
}

// // CreateNftByGroupId           		godoc
// // @Summary      						Create Nft By GroupId
// // @Description  						Nft ID를 넣으면 -> NFT를 반환함.
// // @Tags        						Ready
// // @Accept  							json
// // @Produce      						json
// // @Param   							ReqForm_CreateNftByGroupId formData ReqForm_CreateNftByGroupId true "group_id, wallet_id"
// // @Success      						200  {object}  Resp_CreateNftByGroupId
// // @Router       						/api/item/nft/ [post]
// func CreateNftByGroupId(ctx *gin.Context) {
// 	var ReqBody ReqForm_CreateNftByGroupId
// 	// Group ID Bind
// 	if err := ctx.ShouldBind(&ReqBody); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	// Check Group Exist
// 	_, err := model.ProductGroupSchema.GetGroupById(configs.ConnectDB(), ReqBody.Groupid)
// 	if err != nil {
// 		ctx.JSON(http.StatusNoContent, gin.H{
// 			"error": "Group Not Exist",
// 		})
// 		return
// 	}

// 	// NFT 생성
// 	result, err := model.NftSchema.CreateNftByGroupId(configs.ConnectDB(), ReqBody.Groupid)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, nil)
// 	} else {
// 		//Create Transaction
// 		var TxForm model.NftTx
// 		TxForm.Method = 0
// 		TxForm.From = 0
// 		TxForm.Nftid = result.Id
// 		wallet_toInt, _ := strconv.Atoi(ReqBody.Walletid)
// 		TxForm.To = wallet_toInt

// 		tx, err := model.NftTxSchema.CreateTx(configs.ConnectDB(), TxForm)
// 		if err != nil {
// 			// log.Fatal(err)
// 			log.Println(err)
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"nft": result,
// 			"tx":  tx,
// 		})
// 	}
// }

// // Create Collection
// func CreateCollection(ctx *gin.Context) {
// 	model.CollectionSchema.CreateCollection(configs.DB, ctx)
// }

// // Read Collection (Paging)
// func ReadCollection(ctx *gin.Context) {
// 	model.CollectionSchema.ReadCollection(configs.DB, ctx)
// }

// // Create Group
// func CreateGroup(ctx *gin.Context) {
// 	model.ProductGroupSchema.CreateGroup(configs.DB, ctx)
// }

// // Read Group (Paging)
// func ReadGroup(ctx *gin.Context) {
// 	model.ProductGroupSchema.ReadGroup(configs.DB, ctx)
// }

// // Create NFT
// func CreateNft(ctx *gin.Context) {
// 	model.NftSchema.CreateNft(configs.DB, ctx)
// }

// // Read Nfts (Paging)
// func ReadNfts(ctx *gin.Context) {
// 	model.NftSchema.ReadNfts(configs.DB, ctx)
// }

// // Read One Nft By ID
// func ReadNftById(ctx *gin.Context) {
// 	model.NftSchema.ReadNftById(configs.DB, ctx)
// }
