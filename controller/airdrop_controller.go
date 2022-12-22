package controller

import (
	"fmt"
	"log"
	"net/http"
	"pkg/configs"
	"pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqBody_Airdrop struct {
	Pos      string `json:"pos"`
	Rot      string `json:"rot"`
	Amount   int    `json:"amount"`
	MsgRole  int    `json:"msg_role"`
	Sales_id string `json:"sale_id"` // 얘로 Tree면 2번을, 카드면 3번을 넣어주세요.
	User_id  string `json:"user_id"`
}

// #에어드랍		              godoc
// @Summary      				#에어드랍진행
// @Description  				#에어드랍진행 (SaleID를 기준으로 트리이거나 또는 카드가 될 수 있음, 1인당 1개씩 수령가능)
// @Tags        				Main
// @Security 					Authorization
// @Param        				ReqBody_Airdrop  	body    ReqBody_Airdrop  true  "Plz Write"
// @Produce      				json
// @Success      				200  {object}  model.Obj
// @Router       				/api/item/airdrop [post]
func Airdrop_Item(ctx *gin.Context) {
	var reqBody ReqBody_Airdrop
	var sale model.Sale

	// ## UID로 유저를 찾아 User ID를 반환해줘야함
	user_uid := ctx.MustGet("user_uid").(string)
	user_result, err := model.UserSchema.FindUserByUid(configs.DB, user_uid)
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"error": "NO USER EXIST",
		})
	}
	userId_int := user_result.Id
	userId_string := strconv.Itoa(userId_int)

	// ## 바디 파싱
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		log.Fatal(err)
		return
	}
	fmt.Println("🦾 Request Body Parsing Successed")

	// ## 세일 로그 확인 (판매 로그를 확인해 이미 받았는지 확인함.) ✅
	// 트리 1개, 방명록 1개 에어드랍가능. ✅
	logLen := model.SalesLogSchema.GetSalesLog(configs.DB, reqBody.Sales_id, userId_string)
	if logLen >= 1 {
		fmt.Println("Len Log", logLen)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Already Received User",
		})
		return
	}

	// ## 세일 확인
	configs.DB.Model(&sale).Where("id=?", reqBody.Sales_id).Find(&sale)
	productid_to_string := strconv.Itoa(sale.Product_id)
	saleid_to_numb, _ := strconv.Atoi(reqBody.Sales_id)
	result, err := model.NftSchema.CreateNftByGroupId(configs.DB, productid_to_string)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Create NFT By Group ID Failed",
		})
		return
	}

	// ## 월렛조회
	var wallet_id int
	wallet_id, err = model.WalletSchema.GetWalletByUserId(configs.DB, userId_string)
	if err != nil {
		// 월렛이 없는 유저라면 월렛생성
		w, err := model.WalletSchema.CreateWallet(configs.DB, userId_int)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Create Wallet Failed",
			})
		}
		wallet_id = w.Id
	}

	// ## NFT 트랜잭션 생성
	// Create Transaction
	var TxForm model.NftTx
	TxForm.Method = 0
	TxForm.From = 0
	TxForm.To = wallet_id
	TxForm.Nftid = result.Id
	_, err = model.NftTxSchema.CreateTx(configs.DB, TxForm)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "TX Creation Failed",
		})
		return
	}

	// ## 세일데이터 업데이트

	// 세일 ID를 통해 세일 조회
	salesResult, err := model.SalesSchema.GetSalesById(configs.DB, reqBody.Sales_id)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Sales Failed",
		})
		return
	}

	// 세일 로그 남기기
	var salelogForm model.Saleslog
	salelogForm.Sale_id = saleid_to_numb
	salelogForm.User_id = userId_int
	salelogForm.Type = salesResult.Sale_type
	salelogForm.Amount = 1
	salelogForm.Nft_id = result.Id
	salelogForm.Won_price = salesResult.Won_price
	salelogForm.Doller_price = salesResult.Doller_price
	model.SalesLogSchema.CreateSalesLog(configs.DB, salelogForm)

	// 블록조회
	result_block, err := model.BlockSchema.GetBlock_ByUserId(configs.DB, userId_string)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Block ID Failed",
		})
	}

	// 오브제 생성
	var objForm model.Obj
	objForm.Nft_id = result.Id
	objForm.User_id = userId_int
	objForm.Building_id = 0
	objForm.Block_id = result_block.Id
	objForm.Pos = reqBody.Pos
	objForm.Rot = reqBody.Rot
	objForm.Amount = reqBody.Amount
	objForm.MsgRole = reqBody.MsgRole
	realObj, err := model.ObjSchema.CreateObj(configs.DB, objForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Create Obj Failed",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": realObj,
	})
}

// #Obj 조회 By UserID		     godoc
// @Summary      				#JWT Token을 헤더에 포함하면 Obj를 조회함
// @Description  				#JWT Token을 헤더에 포함하면 Obj를 조회함
// @Tags        				Main
// @Security 					Authorization
// @Produce      				json
// @Success      				200  {array}  []model.Obj
// @Router       				/api/obj/userid [get]
func GetObjsByUserId(ctx *gin.Context) {
	user_uid := ctx.MustGet("user_uid").(string)
	// Token받아서 아래에 넣어줄 것.
	objs_result, err := model.ObjSchema.GetObjsByUserId(configs.DB, user_uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": objs_result,
	})
}

// #Obj 조회 By BlockID		     godoc
// @Summary      				#Obj 조회 By BlockID
// @Description  				#Obj 조회 By BlockID
// @Tags        				Main
// @Param        				blockid  	path    string   true  "Write Block ID"
// @Produce      				json
// @Success      				200  {object}  []model.Obj
// @Router       				/api/obj/block/{blockid} [get]
func GetObjsByBlockId(ctx *gin.Context) {
	blockId := ctx.Param("blockid")
	objs_result, err := model.ObjSchema.GetObjsByBlockId(configs.DB, blockId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": objs_result,
	})
}