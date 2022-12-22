package controller

import (
	"net/http"
	"pkg/configs"
	"pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBlock_ByUserId(ctx *gin.Context) {
	id := ctx.Param("userid")
	result, err := model.BlockSchema.GetBlock_ByUserId(configs.ConnectDB(), id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}

// FindUserAndCreateBlock            godoc
// @Summary      					유저의 블록 보유 확인 및 생성
// @Description  					유저의 블록 보유 확인 후 (없으면 생성 후)리턴
// @Tags        					Block
// @Param        					userid  	path    string  true  "Write Block ID"
// @Produce      					json
// @Success      					200  {object}  model.Block
// @Router       					/api/block/{userid} [get]
func FindUserAndCreateBlock(ctx *gin.Context) {
	// 1. 유저 조회
	// 2. 유저는 블록을 가지고 있는가?
	id := ctx.Param("userid")
	id_toNum, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	getBlock, err := model.BlockSchema.GetBlock_ByUserId(configs.ConnectDB(), id)

	var result model.Block
	var blockData model.Block

	if err != nil {
		// 3. 블록이 없는것임. -> 블록을 생성하라.
		blockData.User_id = id_toNum
		blockData.Name = "christmas"
		blockData.Thema = "christmas"
		createdBlock, err := model.BlockSchema.CreateBlock(configs.ConnectDB(), blockData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		} else {
			result = createdBlock
		}
	} else {
		// 블록이 이미 존재함 -> Result에 데이터 삽입
		result = getBlock
	}

	// Return Result To FE
	// 4. 가지고 있다면 블록을 리턴한다.
	ctx.JSON(http.StatusOK, gin.H{
		"payload": result,
	})
}
