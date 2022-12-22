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

type ReqBody_Token struct {
	NickName string
}

type Resp_FindUserBlockData struct {
	Block model.Block
	Objs  []model.Obj
}

// FindUserAndCreateBlock           godoc
// @Summary      					유저의 블록 보유 확인 및 생성
// @Description  					유저의 블록 보유 확인 후 (없으면 생성 후)리턴
// @Tags        					Main
// @Param        					ReqBody_Token body ReqBody_Token true "Write User Token"
// @Produce      					json
// @Success      					200  {object}  Resp_FindUserBlockData
// @Router       					/api/user/ [post]
func UserBlockAccess(ctx *gin.Context) {
	// body에 담아서 토큰 담아오기
	var reqBody ReqBody_Token
	user_uid := ctx.MustGet("user_uid").(string)

	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		log.Fatal(err)
		return
	}
	claim, err := ValidateJWT(user_uid)
	if err != nil {
		ctx.JSON(http.StatusForbidden, nil)
		return
	}

	uid, _ := strconv.Atoi(claim.UID)

	// 유저 table에 존재하나 확인 없으면 생성
	user, err := model.UserSchema.GetUserByUid(configs.DB, uid)
	if err != nil {
		user, err = model.UserSchema.CreateUserByUid(configs.DB, uid)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
	}
	uid = user.Uid

	// 지갑 존재 하나 확인 없으면 생성
	_, werr := model.WalletSchema.GetWalletByUserId(configs.DB, claim.UID)
	if werr != nil {
		_, wcrr := model.WalletSchema.CreateWallet(configs.DB, uid)

		if wcrr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": wcrr,
			})
			return
		}
	}

	// profile table 존재하나 확인 없으면 생성
	_, perr := model.ProfileSchema.GetProfileByUserId(configs.DB, uid)
	if perr != nil {
		_, perr := model.WalletSchema.CreateWallet(configs.DB, uid)

		if perr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": perr,
			})
			return
		}
	}

	// block 존재하나 확인 없으면 생성 -> user 정보에 block id추가
	block, berr := model.BlockSchema.GetBlock_ByUserId(configs.DB, claim.UID)
	fmt.Println(claim.UID, block.Id)
	if berr != nil {
		block.User_id = uid
		block.Thema = "Empty"
		block.Name = reqBody.NickName
		block, berr = model.BlockSchema.CreateBlock(configs.DB, block)
		if berr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": berr,
			})
			return
		}
	}

	// block의 obj 정보 return
	objs, _ := model.ObjSchema.GetObjsByUserId(configs.DB, claim.UID)

	// block access log 생성(최신화) 후 값 저장 -> access id
	fmt.Println("7")
	_, aerr := model.AccessLogSchema.BlockAccess(configs.DB, uid, block.Id)
	if aerr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": aerr,
		})
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
	ObjMessage string
	ObjId      int
}

// WriteObjMessage                  godoc
// @Summary      					obj message 작성
// @Description  					유저의 블록 보유 확인 후 (없으면 생성 후)리턴
// @Tags        					Main
// @Param        					ReqBody_ObjMessage body ReqBody_ObjMessage true "Write User obj message data"
// @Produce      					json
// @Success      					200  {object}  model.Obj_msg
// @Router       					/api/obj/msg [post]
func WriteObjMessage(ctx *gin.Context) {

	fmt.Println("writeObjMessage")

	// body에 담아서 토큰 담아오기
	var reqBody ReqBody_ObjMessage
	user_uid := ctx.MustGet("user_uid").(string)
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		//log.Fatal(err)
		fmt.Println(err)
		return
	}

	uid, _ := strconv.Atoi(user_uid)
	oid := strconv.Itoa(reqBody.ObjId)
	message := reqBody.ObjMessage

	fmt.Println("uid: ", uid)
	fmt.Println("objectID: ", oid)
	fmt.Println("objectMsg: ", message)

	// obj 주인과 obj 작성 타입 확인
	obj, err := model.ObjSchema.GetObjByObjId(configs.DB, oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	// 작성을 owner만 가능한 경우
	if obj.MsgRole == 3 {
		if obj.User_id != uid {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		fmt.Println("2")
		// 이미 작성했나 확인 *UserAll 이 아니라, User 가 맞지 않나?
		amount, err := model.Obj_msgSchema.GetObjMsgCountByUserAll(configs.DB, uid, oid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		if amount > 0 {
			fmt.Println("3")
			// update
			msg, err := model.Obj_msgSchema.UpdateObjMsg(configs.DB, message, oid, uid)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"payload": msg,
			})
			return
		}

		// 성공
		fmt.Println("4")
		msg, err := model.Obj_msgSchema.CreateObjMsg(configs.DB, message, oid, uid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
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
		fmt.Println("5")

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

		fmt.Println("6")
		// 성공
		msg, err := model.Obj_msgSchema.CreateObjMsg(configs.DB, message, oid, uid)
		ctx.JSON(http.StatusOK, gin.H{
			"payload": msg,
		})
		return
	}

	fmt.Println("7")
	ctx.JSON(http.StatusInternalServerError, nil)
	return
}

type ReqBody_ObjDel struct {
	ObjMsgId int
	ObjId    int
}

// DeleteObjMsg                     godoc
// @Summary      					Obj msg 삭제
// @Description  					Obj msg의 is_active 값을 변경해 삭제처리 한다
// @Tags        					Main
// @Param        					ReqBody_ObjDel body ReqBody_ObjDel true "Write User obj message data"
// @Produce      					json
// @Success      					200  {object}  model.Obj_msg
// @Router       					/api/obj/msg/del [post]
func DeleteObjMsg(ctx *gin.Context) {
	// body에 담아서 토큰 담아오기
	var reqBody ReqBody_ObjDel
	user_uid := ctx.MustGet("user_uid").(string)
	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		log.Fatal(err)
		return
	}

	uid, _ := strconv.Atoi(user_uid)
	oid := strconv.Itoa(reqBody.ObjMsgId)
	omid := strconv.Itoa(reqBody.ObjMsgId)
	fmt.Println("1")

	// obj 주인과 obj 작성 타입 확인
	obj, err := model.ObjSchema.GetObjByObjId(configs.DB, oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	fmt.Println("2")
	obj_msg, err := model.Obj_msgSchema.GetObjMsgByObjId(configs.DB, oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
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

	fmt.Println("3")
	if obj.MsgRole == 3 { // owner only
		if obj.User_id != uid {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		fmt.Println("4")
		// del
		obj_msg, err := model.Obj_msgSchema.UpdateObjMsgIsActive(configs.DB, omid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"payload": obj_msg,
		})
		return
	}

	fmt.Println("5")
	if obj.MsgRole == 6 { // owner + writer
		if obj.User_id != uid || obj_msg.Created_user != uid {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		fmt.Println("6")
		obj_msg, err = model.Obj_msgSchema.UpdateObjMsgIsActive(configs.DB, omid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"payload": obj_msg,
		})
		return
	}

	fmt.Println("7")
	ctx.JSON(http.StatusInternalServerError, nil)
	return
}
