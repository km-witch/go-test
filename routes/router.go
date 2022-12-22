package routes

import (
	"fmt"
	"net/http"
	"pkg/configs"
	"pkg/controller"
	"pkg/docs"
	"pkg/model"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/docs/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "dev-go.witchworld.io"

	route_block := r.Group("/api/block", AuthCheck())
	{
		route_block.GET("/get/:userid", controller.GetBlock_ByUserId) // 유저 블록 확인 &&생성
		route_block.GET("/data/:userid", controller.FindUserAndCreateBlock)
	}
	route_obj := r.Group("/api/obj")
	{
		route_obj.POST("/airdrop", AuthCheck(), controller.Airdrop_Item)        // 로그인 필수 (헤더에 토큰 넣어서 보내야함)
		route_obj.GET("/user/:userid", AuthCheck(), controller.GetObjsByUserId) // 로그인 필수. (헤더에 토큰 넣어서 보내야함)
		route_obj.GET("/block/:blockid", controller.GetObjsByBlockId)
		route_obj.GET("/msg/paging/:page/:limit", controller.ReadObjMessages)
		route_obj.GET("/msg/:id", controller.GetObjMsg)
		route_obj.GET("/msg/count/:obj_id", controller.GetObjMessageCount)
	}
	route_item := r.Group("/api/item")
	{
		route_item.GET("/nft/:nftid", controller.GetNftById)
		route_item.GET("/group/:groupid", controller.GetProductGroupById)
		route_item.GET("/collection/:collectionid", controller.GetCollectionById)
		route_item.POST("/nft", controller.CreateNftByGroupId)
	}
	route_main := r.Group("/api")
	{
		route_main.POST("/user", controller.UserBlockAccess)
		route_main.POST("/obj/msg", controller.WriteObjMessage)
		route_main.POST("/obj/msg/del", controller.DeleteObjMsg)
	}

	// route_block := r.Group("/api/block")
	// {
	// 	route_block.GET("/get/:userid", controller.GetBlock_ByUserId) // 유저 블록 확인 &&생성
	// 	route_block.GET("/data/:userid", controller.FindUserAndCreateBlock)
	// }
	// route_item := r.Group("/api/obj")
	// {
	// 	route_item.GET("/userid", AuthCheck(), controller.GetObjsByUserId) // 로그인 필수. (헤더에 토큰 넣어서 보내야함)
	// 	route_item.GET("/blockid/:blockid", controller.GetObjsByBlockId)
	// }
	configs.DB.AutoMigrate(&model.Obj_msg{}, &model.Sale{}, &model.Saleslog{}, &model.Obj{}, &model.Block{}, &model.NftTx{}, &model.Wallet{}, &model.User{}, &model.AccessLog{}, &model.Profile{})
	r.Use(CORS())
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// testToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1aWQiOiIxMyIsIlByZUxvZ2luIjoidHJ1ZSIsImFkbWluIjoiZmFsc2UiLCJleHAiOjE2NzY4NTc1ODUsInVzZXIiOiJ0cnVlIn0.oY70FvH1M0VhFTI2DI6z_RusvcxGPn-l-3zrIEUxn2g"
func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Plz Add Your Token",
			})
			return
		}

		claim, err := controller.ValidateJWT(h.IDToken)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "Token Verification Failed",
			})
			return
		}

		fmt.Println("User UID : ", claim.UID)

		// 만약 위 에러에 걸리지 않으면 토큰인증완료.
		// UID를 User 테이블에 등록한다. (Claim.UID)
		ctx.Set("user_uid", claim.UID)
		ctx.Next()
	}
}
