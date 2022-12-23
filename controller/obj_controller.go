package controller

import (
	"net/http"
	"pkg/configs"
	"pkg/model"

	"github.com/gin-gonic/gin"
)

type Resp_GetMessageCount struct {
	Payload int `json:"payload" example:"2"`
}

// GetObjMessageCount     	godoc
// @Summary         		Obj Message 갯수 조회 기능 (트리 크기 조정용)
// @Description     		Obj Message 갯수 조회 기능 (트리 크기 조정용)
// @Tags            		Main
// @Param           		obj_id path int true "Object id"
// @Produce         		json
// @Success         		200 {object} Resp_GetMessageCount
// @Failure         		400
// @Router          		/api/obj/msg/count/{obj_id} [get]
func GetObjMessageCount(ctx *gin.Context) {
	obj_id := ctx.Param("obj_id")
	objAmount, err := model.Obj_msgSchema.GetObjActiveMsgCount(configs.DB, obj_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": objAmount,
	})
}

// // GetObj_by_blockid           	godoc
// // @Summary      				블록아이디로 오브제를 가져오는 API
// // @Description  				블록아이디를 넣으면 오브제를 리턴함.
// // @Tags        					Ready
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

// Obj 메세지 조회           	   godoc
// @Summary      				Obj Msg 단일 조회
// @Description  				Obj Msg 단일 조회 기능 (단일, ID값 기준)
// @Tags        				Ready
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
