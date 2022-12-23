package controller

import (
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
}

type Resp_Airdrop_Item struct {
	Payload model.Obj
}

// #에어드랍		              godoc
// @Summary      				#에어드랍진행
// @Description  				#에어드랍진행 (SaleID를 기준으로 트리이거나 또는 카드가 될 수 있음, 1인당 1개씩 수령가능)
// @Tags        				Main
// @Security 					Authorization
// @Param                       Authorization header string true "Bearer"
// @Param        				ReqBody_Airdrop  	body    ReqBody_Airdrop  true  "Plz Write"
// @Produce      				json
// @Success      				200  {object}  Resp_Airdrop_Item
// @Router       				/api/obj/airdrop [post]
func Airdrop_Item(ctx *gin.Context) {
	var reqBody ReqBody_Airdrop
	var sale model.Sale

	// ## UID로 유저를 찾아 User ID를 반환해줘야함
	user_uid := ctx.MustGet("user_uid").(string)
	// UID로 UserID 찾기
	// user_uid_int, _ := strconv.Atoi(user_uid)
	// user_result, err := model.UserSchema.GetUserByUid(configs.DB, user_uid_int)
	// userId_string := strconv.Itoa(user_result.Id)

	user_result, err := model.UserSchema.FindUserByUid(configs.DB, user_uid)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		log.Println("UID 확인 실패")
		return
	}
	userId_int := user_result.Id
	userId_string := strconv.Itoa(userId_int)

	// //Wallet 조회
	// walletResult, err := model.WalletSchema.GetWalletByUserId(configs.DB, userId_string)
	// if err != nil {
	// 	ctx.JSON(http.StatusNoContent, nil)
	// 	log.Println("Wallet 조회 실패")
	// 	return
	// }

	// ## 바디 파싱
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "바디 파싱 실패"})
		log.Println("바디 파싱 실패")
		return
	}
	log.Println("🦾 Request Body Parsing Successed")

	// ## 세일 로그 확인 (판매 로그를 확인해 이미 받았는지 확인함.) ✅
	// 트리 1개, 방명록 1개 에어드랍가능. ✅
	logLen := model.SalesLogSchema.GetSalesLog(configs.DB, reqBody.Sales_id, userId_string)
	if logLen >= 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "이미 받음"})
		log.Println("이미받음")
		return
	}

	// ## 월렛조회
	var wallet_id int
	wallet_id, err = model.WalletSchema.GetWalletByUserId(configs.DB, userId_string)
	if err != nil {
		// 월렛이 없는 유저라면 월렛생성
		w, err := model.WalletSchema.CreateWallet(configs.DB, userId_int)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "월렛 조회 실패"})
			log.Println("월렛 조회 실패")
			return
		}
		wallet_id = w.Id
	}

	// ## 세일 확인
	configs.DB.Model(&sale).Where("id=?", reqBody.Sales_id).Find(&sale)
	productid_to_string := strconv.Itoa(sale.Product_id)
	saleid_to_numb, _ := strconv.Atoi(reqBody.Sales_id)
	result, err := model.NftSchema.CreateNftByGroupId(configs.DB, productid_to_string, wallet_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println("세일 확인 및 NFT 생성 실패")
		return
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
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println("NFT 트랜잭션 생성 실패")
		return
	}

	// ## 세일데이터 업데이트

	// 세일 ID를 통해 세일 조회
	salesResult, err := model.SalesSchema.GetSalesById(configs.DB, reqBody.Sales_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println("세일조회실패")
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
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println("블록 조회 실패")
		return
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
	objForm.MsgRole = reqBody.MsgRole // 3=OWNER || 6=Guest || 9=ALL // 트리는 =6 || 카드=3
	realObj, err := model.ObjSchema.CreateObj(configs.DB, objForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println("오브제 생성 실패")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": realObj,
	})
}

type Resp_GetObjsByUserId struct {
	Payload []model.Obj
}

// #Obj 조회 By UserID		     godoc
// @Summary      				#JWT Token을 헤더에 포함하면 Obj를 조회함
// @Description  				#JWT Token을 헤더에 포함하면 Obj를 조회함
// @Tags        				Main
// @Security 					Authorization
// @Param                       Authorization header string true "Bearer"
// @Produce      				json
// @Success      				200  {object}  Resp_GetObjsByUserId
// @Router       				/api/obj/userid [get]
func GetObjsByUserId(ctx *gin.Context) {
	user_uid := ctx.MustGet("user_uid").(string)
	user_uid_int, _ := strconv.Atoi(user_uid)

	// Token받아서 아래에 넣어줄 것.
	user_result, _ := model.UserSchema.GetUserByUid(configs.DB, user_uid_int)
	userId_int := strconv.Itoa(user_result.Id)

	objs_result, err := model.ObjSchema.GetObjsByUserId(configs.DB, userId_int)
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
