package controller

import (
	"log"
	"net/http"
	"pkg/configs"
	"pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqBody_Token struct {
	NickName string
}

type Resp_FindUserBlockData struct {
	Block model.Block
	Objs  []model.Obj_with_productid
}

// UserBlockAccess          		godoc
// @Summary      					#유저 접속시 호출 초기화 및 유저 조회
// @Description  					유저의 블록 보유 확인 후 (없으면 생성 후)리턴
// @Tags        					Main
// @Param                           Authorization header string true "Bearer"
// @Param        					ReqBody_Token body ReqBody_Token true "Write User Token"
// @Produce      					json
// @Security 					    Authorization
// @Success      					200  {object}  Resp_FindUserBlockData
// @Router       					/api/user/ [post]
func UserBlockAccess(ctx *gin.Context) {
	// body에 담아서 토큰 담아오기
	log.Println("UserBlockAccess")
	var reqBody ReqBody_Token
	user_uid := ctx.MustGet("user_uid").(string)
	log.Println("1")
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println("Bind ERR:", err)
		// log.Fatal(err)
		log.Println(err)
		return
	}

	log.Println("2")
	uid, _ := strconv.Atoi(user_uid)

	// 유저 table에 존재하나 확인 없으면 생성
	user, err := model.UserSchema.GetUserByUid(configs.DB, uid)
	if err != nil {
		user, err = model.UserSchema.CreateUserByUid(configs.DB, uid)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	uid = user.Uid
	user_id_string := strconv.Itoa(user.Id)
	log.Println("3")

	// 지갑 존재 하나 확인 없으면 생성
	_, werr := model.WalletSchema.GetWalletByUserId(configs.DB, user_id_string)
	if werr != nil {
		// 월렛이 없다면 월렛 생성
		_, wcrr := model.WalletSchema.CreateWallet(configs.DB, user.Id)

		if wcrr != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	log.Println("4")
	// profile table 존재하나 확인 없으면 생성
	_, perr := model.ProfileSchema.GetProfileByUserId(configs.DB, user.Id)
	if perr != nil {
		_, perr := model.ProfileSchema.CreateProfile(configs.DB, user.Id)

		if perr != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	log.Println("5")
	// block 존재하나 확인 없으면 생성 -> user 정보에 block id추가
	block, berr := model.BlockSchema.GetBlock_ByUserId(configs.DB, user_id_string)
	log.Println(user_uid, block.Id)
	if berr != nil {
		block.User_id = user.Id
		block.Thema = "Empty"
		block.Name = reqBody.NickName
		block, berr = model.BlockSchema.CreateBlock(configs.DB, block)
		if berr != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	log.Println("6")
	// block의 obj 정보 return
	objs, _ := model.ObjSchema.GetObjsByUserIdWithProductId(configs.DB, user_id_string)

	// block access log 생성(최신화) 후 값 저장 -> access id
	log.Println("7")
	_, aerr := model.AccessLogSchema.BlockAccess(configs.DB, user.Id, block.Id)
	if aerr != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// return
	ctx.JSON(http.StatusOK, gin.H{
		"blockdata": block,
		"objs":      objs,
	})
	return
}

type ReqBody_ObjMessage struct {
	ObjMessage   string
	ObjId        int
	UserNickname string
}

// WriteObjMessage                  godoc
// @Summary      					obj message 작성
// @Description  					유저의 블록 보유 확인 후 (없으면 생성 후)리턴
// @Tags        					Main
// @Param                           Authorization header string true "Bearer"
// @Param        					ReqBody_ObjMessage body ReqBody_ObjMessage true "Write User obj id and message"
// @Produce      					json
// @Security 					    Authorization
// @Success      					200  {object}  model.Obj_msg
// @Router       					/api/obj/msg [post]
func WriteObjMessage(ctx *gin.Context) {
	// body에 담아서 토큰 담아오기
	var reqBody ReqBody_ObjMessage
	user_uid := ctx.MustGet("user_uid").(string)
	log.Println(1)
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// log.Fatal(err)
		log.Println(err)
		return
	}

	log.Println("리퀘바디: ", reqBody)
	uid, _ := strconv.Atoi(user_uid)
	result_user, err := model.UserSchema.FindUserByUid(configs.DB, user_uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	uid = result_user.Id
	oid := strconv.Itoa(reqBody.ObjId)
	message := reqBody.ObjMessage
	log.Println("object ID :", oid)
	// obj 주인과 obj 작성 타입 확인
	obj, err := model.ObjSchema.GetObjByObjId(configs.DB, oid)
	if err != nil {
		log.Println("GetObjByObjId Failed: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "GetObjByObjId Failed",
		})
		return
	}

	log.Println("4")
	// 작성을 owner만 가능한 경우
	if obj.MsgRole == 3 {
		if obj.User_id != uid {
			log.Println("메세지 롤이 다릅니다: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "메세지 롤이 다릅니다",
			})
			return
		}

		log.Println("2")
		// 이미 작성했나 확인
		// GetAllObjMsgCountByUser -> is_active false 된 메시지까지 체크라 all obj msg
		amount, err := model.Obj_msgSchema.GetAllObjMsgCountByUser(configs.DB, uid, oid)
		if err != nil {
			log.Println("이미 작성됨: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "이미 작성됨",
			})
			return
		}
		if amount > 0 {
			log.Println("3")
			// update
			msg, err := model.Obj_msgSchema.UpdateObjMsg(configs.DB, message, oid, uid)
			if err != nil {
				log.Println("메세지 업데이트 실패: ", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "메세지 업데이트 실패",
				})
				return
			}

			log.Println("업데이트 성공")
			ctx.JSON(http.StatusOK, gin.H{
				"payload": msg,
			})
			return
		}

		// 성공
		log.Println("4")
		msg, err := model.Obj_msgSchema.CreateObjMsg(configs.DB, message, oid, reqBody.UserNickname, uid)
		if err != nil {
			log.Println("메세지 생성 실패: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "메세지 생성 실패",
			})
			return
		}

		log.Println("신규 작성 성공")
		ctx.JSON(http.StatusOK, gin.H{
			"payload": msg,
		})
		return
	}

	// 작성을 guest만 가능한 경우
	if obj.MsgRole == 6 {
		if obj.User_id == uid {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		log.Println("5")

		// 작성 개수 확인 3개 미만일것
		amount, err := model.Obj_msgSchema.GetObjMsgCountByUser(configs.DB, uid, oid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		if amount >= 3 {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		log.Println("6")
		// 성공
		msg, err := model.Obj_msgSchema.CreateObjMsg(configs.DB, message, oid, reqBody.UserNickname, uid)
		ctx.JSON(http.StatusOK, gin.H{
			"payload": msg,
		})
		return
	}

	log.Println("7")
	ctx.JSON(http.StatusInternalServerError, nil)
	return
}

type Response_ReadObjMessages struct {
	Payload []model.Obj_msg
}

// Obj messages 페이징 조회            godoc
// @Summary      					Obj messages 페이징 조회
// @Description  					Obj messages 페이징 조회
// @Tags        					Main
// @Param        					page  	path    string  true  "페이지입력"
// @Param        					limit  	path    string  true  "조회갯수제한"
// @Produce      					json
// @Success      					200  {object}  Response_ReadObjMessages
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

type ReqBody_ObjDel struct {
	ObjMsgId int
	ObjId    int
}

// DeleteObjMsg                     godoc
// @Summary      					Obj msg 삭제
// @Description  					Obj msg의 is_active 값을 변경해 삭제처리 한다
// @Tags        					Main
// @Param                           Authorization header string true "Bearer"
// @Param        					ReqBody_ObjDel body ReqBody_ObjDel true "Write User obj id and obj_msg id"
// @Produce      					json
// @Security 					    Authorization
// @Success      					200  {object}  model.Obj_msg
// @Router       					/api/obj/msg/del [post]
func DeleteObjMsg(ctx *gin.Context) {
	// body에 담아서 토큰 담아오기
	var reqBody ReqBody_ObjDel
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// log.Fatal(err)
		log.Println(err)
		return
	}

	user_uid := ctx.MustGet("user_uid").(string)
	user_uid_toInt, _ := strconv.Atoi(user_uid)
	result_user, err := model.UserSchema.GetUserByUid(configs.DB, user_uid_toInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	uid := result_user.Id
	oid := strconv.Itoa(reqBody.ObjId)
	omid := strconv.Itoa(reqBody.ObjMsgId)
	log.Println("1")

	// obj 주인과 obj 작성 타입 확인
	obj, err := model.ObjSchema.GetObjByObjId(configs.DB, oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}

	log.Println("2")
	obj_msg, err := model.Obj_msgSchema.GetObjMsgByObjId(configs.DB, oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	if !obj_msg.IsActive {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	if obj_msg.ObjId != obj.Id {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	log.Println("3")
	if obj.MsgRole == 3 { // owner only
		if obj.User_id != uid {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		log.Println("4")
		// del
		obj_msg, err := model.Obj_msgSchema.UpdateObjMsgIsActive(configs.DB, omid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"payload": obj_msg,
		})
		return
	}

	log.Println("5")
	if obj.MsgRole == 6 { // owner + writer
		if obj.User_id != uid || obj_msg.Created_user != uid {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		log.Println("6")
		obj_msg, err = model.Obj_msgSchema.UpdateObjMsgIsActive(configs.DB, omid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"payload": obj_msg,
		})
		return
	}

	log.Println("7")
	ctx.JSON(http.StatusInternalServerError, nil)
}
