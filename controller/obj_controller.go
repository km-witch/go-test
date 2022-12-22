package controller

import (
	"net/http"
	"pkg/configs"
	"pkg/model"

	"github.com/gin-gonic/gin"
)

type Resp_GetMessageCount struct {
	Payload int `json:"payload" example:"24"`
}

// GetObjMessageCount     	godoc
// @Summary         		Obj Message 갯수 조회 기능 (트리 크기 조정용)
// @Description     		Obj Message 갯수 조회 기능 (트리 크기 조정용)
// @Tags            		Obj
// @Param           		obj_id path int true "Object id"
// @Produce         		json
// @Success         		200 {object} Resp_GetMessageCount
// @Failure         		400
// @Router          		/api/obj/msg/count/{obj_id} [get]
func GetObjMessageCount(ctx *gin.Context) {
	obj_id := ctx.Param("obj_id")
	objAmount, err := model.Obj_msgSchema.GetObjMsgCount(configs.DB, obj_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": objAmount,
	})
}

// Obj messages 페이징 조회            godoc
// @Summary      					Obj messages 페이징 조회
// @Description  					Obj messages 페이징 조회
// @Tags        					Main
// @Param        					page  	path    string  true  "페이지입력"
// @Param        					limit  	path    string  true  "조회갯수제한"
// @Produce      					json
// @Success      					200  {object}  model.Obj_msg
// @Router       					/api/obj/msg/paging/{page}/{limit} [get]
func ReadObjMessages(ctx *gin.Context) {
	page := ctx.Param("page")
	limit := ctx.Param("limit")

	resultObjs, err := model.Obj_msgSchema.GetObjMsgs(configs.DB, page, limit)
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"payload": resultObjs,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": resultObjs,
	})
}

// // GetObj_by_blockid           	godoc
// // @Summary      				블록아이디로 오브제를 가져오는 API
// // @Description  				블록아이디를 넣으면 오브제를 리턴함.
// // @Tags        				Obj
// // @Param        				blockid  	path    string  true  "Write Block ID"
// // @Produce      				json
// // @Success      				200  {object}  model.Obj
// // @Router       				/api/obj/{blockid} [get]
// func GetObj_by_blockid(ctx *gin.Context) {
// 	block_id := ctx.Param("blockid")
// 	result, err := model.ObjSchema.GetObjByBlockId(configs.ConnectDB(), block_id)
// 	if err != nil {
// 		ctx.JSON(http.StatusNoContent, nil)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"payload": result,
// 	})
// }

// type ReqBody_Airdrop struct {
// 	Nft_id   int    `json:"nft_id"`
// 	Block_id int    `json:"block_id"`
// 	Pos      string `json:"pos"`
// 	Rot      string `json:"rot"`
// 	Amount   int    `json:"amount"`
// 	MsgRole  int    `json:"msg_role"`
// 	Sales_id string `json:"sale_id"`
// 	User_id  string `json:"user_id"`
// }

// // 에어드랍		            		godoc
// // @Summary      				에어드랍진행 !
// // @Description  				Wallet생성도 해주고, NFT 생성 해주고, Obj 생성!
// // @Tags        				Obj
// // @Param        				ReqBody_Airdrop  	body    ReqBody_Airdrop  true  "Write Block ID"
// // @Produce      				json
// // @Success      				200  {object}  model.Obj
// // @Router       				/api/obj/airdrop [post]
// func Airdrop(ctx *gin.Context) {
// 	var reqBody ReqBody_Airdrop
// 	var sale model.Sale
// 	if err := ctx.ShouldBind(&reqBody); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		fmt.Println(err)
// 		return
// 	}
// 	// 세일로그확인 ✅
// 	logLen := model.SalesLogSchema.GetSalesLog(configs.ConnectDB(), reqBody.Sales_id, reqBody.User_id)
// 	if logLen > 0 {
// 		fmt.Println("Len Log", logLen)
// 		ctx.JSON(http.StatusBadRequest, nil)
// 		return
// 	}

// 	// NFT 생성 & 로그 ✅ -> Sale에서 가져오는 그룹아이디가 없음. -> 운영자가  박으면됨.
// 	configs.DB.Model(&sale).Where("id=?", reqBody.Sales_id).Find(&sale)
// 	productid_to_string := strconv.Itoa(sale.Product_id)
// 	userid_to_numb, _ := strconv.Atoi(reqBody.User_id)
// 	saleid_to_numb, _ := strconv.Atoi(reqBody.Sales_id)
// 	result, err := model.NftSchema.CreateNftByGroupId(configs.ConnectDB(), productid_to_string)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, nil)
// 	}

// 	var wallet_id int
// 	wallet_id, err = model.WalletSchema.GetWalletByUserId(configs.ConnectDB(), reqBody.User_id)
// 	if err != nil {
// 		// 월렛이 없는 유저라면 월렛생성
// 		w, err := model.WalletSchema.CreateWallet(configs.ConnectDB(), userid_to_numb)
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, nil)
// 		}
// 		wallet_id = w.Id
// 	}

// 	//트랜잭션
// 	//Create Transaction
// 	var TxForm model.NftTx
// 	TxForm.Method = 0
// 	TxForm.From = 0
// 	TxForm.Nftid = result.Id
// 	TxForm.To = wallet_id
// 	fmt.Println(4)
// 	_, err = model.NftTxSchema.CreateTx(configs.ConnectDB(), TxForm)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// 세일 로그 남기기
// 	var salelogForm model.Saleslog
// 	salelogForm.Sale_id = saleid_to_numb
// 	salelogForm.User_id = userid_to_numb
// 	model.SalesLogSchema.CreateSalesLog(configs.ConnectDB(), salelogForm)

// 	// 오브제 생성
// 	var objForm model.Obj
// 	objForm.Nft_id = result.Id
// 	objForm.User_id = userid_to_numb
// 	objForm.Building_id = 0
// 	objForm.Block_id = reqBody.Block_id
// 	objForm.Pos = reqBody.Pos
// 	objForm.Rot = reqBody.Rot
// 	objForm.Amount = 1
// 	objForm.MsgRole = 1
// 	realObj, err := model.ObjSchema.CreateObj(configs.ConnectDB(), objForm)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"payload": realObj,
// 	})
// }

// Obj 메세지 조회           	   godoc
// @Summary      				Obj Msg 단일 조회
// @Description  				Obj Msg 단일 조회 기능 (단일, ID값 기준)
// @Tags        				Obj
// @Param        				id  	path    string  true  "오브제 메세지 ID"
// @Produce      				json
// @Success      				200  {object}  model.Obj_msg
// @Router       				/api/obj/msg/{id} [get]
func GetObjMsg(ctx *gin.Context) {
	model.Obj_msgSchema.GetObjMsg(configs.ConnectDB(), ctx)
}

// func GetObjMsgs(ctx *gin.Context) {
// 	model.Obj_msgSchema.GetObjMsgs(configs.ConnectDB(), ctx)
// }
