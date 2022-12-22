package controller

import (
	"net/http"
	"pkg/configs"
	"pkg/model"

	"github.com/gin-gonic/gin"
)

// KPI : JWT 회원가입 , 로그인, 로그아웃 구현.
func RegisterUser(ctx *gin.Context) {
	var userInput model.User // 받을 데이터 객체를 정의해야함. model 내용과 같으므로 model Folder의 Collection Struct.
	// request JSON BODY(JSON)로 받은 userInput을 받고, 형식에 맞지 않으면 에러처리.
	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	// Create : 실제로 DB에 삽입함. 에러가 발생할 시 자동으로 Panic ERROR 로그를 남기기에 별도로 에러처리가 없음.
	configs.DB2.AutoMigrate(&userInput) // AutoMirgate는 Model에 적용된 스키마를 기준으로 새 데이터베이스 적용시, 테이블이 없으면 생성하도록 함.
	configs.DB2.Create(&userInput)

	// Return To FE : 첫번째 인자로는 상태코드 두번째 인자로는 JSON데이터를 리턴.
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Whatthe",
	})
}

// Create user
func CreateUser(ctx *gin.Context) {
	model.UserSchema.CreateUser(configs.DB2, ctx)
}

func DeleteUser(ctx *gin.Context) {
	var userInput model.User
	uid := ctx.Param("uid")

	configs.DB2.Table("user").Where("uid=?", uid).Delete(&userInput)
	// configs.DB2.AutoMigrate(&userInput)
	// configs.DB2.Create(&userInput)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}

// /*
// 	Wallet
// 	-> user의 wallet data (wit나 nft등의 보유를 알 수 있다)
// 	{wallet_id, user_id, nft_id, wit, created_time, updated_time}
// */
// // Create Wallet
// func CreateWallet(ctx *gin.Context) {
// 	var userInput model.Wallet

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	configs.DB2.Table("wallet").Create(&userInput)

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": userInput,
// 	})
// }

// /*
// 	ItemOwned
// 	-> 유저의 NFT 보유 정보 -> wallet의 하위 테이블 개념
// */
// // create ItemOwned
// func CreateItemOwned(ctx *gin.Context) {
// 	var userInput model.ItemOwned

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		log.Fatal(err)
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	configs.DB2.Table("ItemOwned").Create(&userInput)

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": userInput,
// 	})
// }

// // read ItemOwned
// func CheckItemOwned(ctx *gin.Context) {
// 	var userInput model.ItemOwned

// 	owner_id := ctx.Param("owner")
// 	collection_id := ctx.Param("collection")
// 	fmt.Println(owner_id, collection_id)
// 	configs.DB2.Table("ItemOwned").Where("owner_id=? AND collection_id=?", owner_id, collection_id).Find(&userInput)

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": userInput,
// 	})
// }

// /*
// 	access
// 	-> 유저의 block 방문 기록 로그용
// 	{user_id(FK), access_time, block_id}
// */
// // create access
// func CreateAccess(ctx *gin.Context) {
// 	var userInput model.Access

// 	if err := ctx.ShouldBind(&userInput); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	configs.DB2.Table("Access").Create(&userInput)

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": userInput,
// 	})
// }

// update access
func UpdateAccess(ctx *gin.Context) {
	model.AccessSchema.UpdateAccess(configs.DB2, ctx)
}

// read access
func ReadAccess(ctx *gin.Context) {
	model.AccessSchema.ReadAccess(configs.DB2, ctx)
}
