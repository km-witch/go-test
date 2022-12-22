package controller

import (
	"log"
	"net/http"
	"pkg/configs"
	"pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetObjbyblockid            godoc
// @Summary      				블록아이디로 오브제를 가져오는 API
// @Description  				블록아이디를 넣으면 오브제를 리턴함.
// @Tags        				Obj
// @Param        				blockid  	path    string  true  "Write Block ID"
// @Produce      				json
// @Success      				200  {object}  model.Obj
// @Router       				/api/obj/{blockid} [get]
func GetObj_by_blockid(ctx *gin.Context) {
	block_id := ctx.Param("blockid")
	result, err := model.ObjSchema.GetObjByBlockId(configs.ConnectDB(), block_id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}

type ReqBody struct {
	Obj      model.Obj
	Sales_id string `json:"sale_id" binding:"required"`
	User_id  string `json:"user_id" binding:"required"`
}

// 대망의 에어드랍            		godoc
// @Summary      				에어드랍진행 !
// @Description  				Wallet생성도 해주고, NFT 생성 해주고, Obj 생성!
// @Tags        				Obj
// @Param        				ReqBody  	formData    string  true  "Write Block ID"
// @Produce      				json
// @Success      				200  {object}  model.Obj
// @Router       				/api/obj/airdrop [post]
func Airdrop(ctx *gin.Context) {
	var reqBody ReqBody
	var sale model.Sale
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	// 세일로그확인
	logLen := model.SalesLogSchema.GetSalesLog(configs.ConnectDB(), reqBody.Sales_id, reqBody.User_id)
	if logLen > 0 {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	// NFT 생성 & 로그
	configs.DB.Model(&sale).Where("id=?", reqBody.Sales_id).Find(&sale)
	productid_to_string := strconv.Itoa(sale.Product_id)
	userid_to_numb, _ := strconv.Atoi(reqBody.User_id)
	saleid_to_numb, _ := strconv.Atoi(reqBody.Sales_id)
	result, err := model.NftSchema.CreateNftByGroupId(configs.ConnectDB(), productid_to_string)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	var wallet_id int
	wallet_id, err = model.WalletSchema.GetWalletByUserId(configs.ConnectDB(), reqBody.User_id)
	if err != nil {
		// 월렛이 없는 유저라면
		w, err := model.WalletSchema.CreateWallet(configs.ConnectDB(), userid_to_numb)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		}
		wallet_id = w.Id
	}

	//트랜잭션
	//Create Transaction
	var TxForm model.NftTx
	TxForm.Method = 0
	TxForm.From = 0
	TxForm.Nftid = result.Id
	TxForm.To = wallet_id

	_, err = model.NftTxSchema.CreateTx(configs.ConnectDB(), TxForm)
	if err != nil {
		log.Fatal(err)
	}

	// 세일 로그 남기기
	var salelogForm model.Saleslog
	salelogForm.Sale_id = saleid_to_numb
	salelogForm.User_id = userid_to_numb
	model.SalesLogSchema.CreateSalesLog(configs.ConnectDB(), salelogForm)

	// 오브제 생성
	var objForm model.Obj
	objForm = reqBody.Obj
	objForm.Nft_id = result.Id
	realObj, err := model.ObjSchema.CreateObj(configs.ConnectDB(), objForm)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": realObj,
	})
}

// Obj 메세지 조회           	   godoc
// @Summary      				Obj Msg 단일 조회
// @Description  				Obj Msg 단일 조회 기능 (단일, ID값 기준)
// @Tags        				Obj
// @Param        				id  	path    string  true  "Write ObjMSG ID"
// @Produce      				json
// @Success      				200  {object}  model.Obj_msg
// @Router       				/api/obj/msg/{id} [get]
func GetObjMsg(ctx *gin.Context) {
	model.Obj_msgSchema.GetObjMsg(configs.ConnectDB(), ctx)
}

// Obj message 페이징 조회          godoc
// @Summary      				Obj message 페이징 조회
// @Description  				Obj message 페이징 조회
// @Tags        				Obj
// @Param        				page  	path    string  true  "페이지입력"
// @Param        				limit  	path    string  true  "조회갯수제한"
// @Produce      				json
// @Success      				200  {object}  model.Obj_msg
// @Router       				/api/obj/msg/paging/{page}/{limit} [get]
func GetObjMsgs(ctx *gin.Context) {
	model.Obj_msgSchema.GetObjMsgs(configs.ConnectDB(), ctx)
}
