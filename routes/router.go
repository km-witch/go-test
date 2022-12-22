package routes

import (
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
	docs.SwaggerInfo.Host = "localhost:8080"
	route_block := r.Group("/api/block")
	{
		route_block.GET("/:userid", controller.FindUserAndCreateBlock) // 유저 블록 확인 &&생성
	}
	route_obj := r.Group("/api/obj")
	{
		route_obj.POST("/airdrop", controller.Airdrop)
		route_obj.GET("/:blockid", controller.GetObj_by_blockid) // 블록 조회 => Obj
		route_obj.GET("/msg/paging/:page/:limit", controller.GetObjMsgs)
		route_obj.GET("/msg/:id", controller.GetObjMsg)
	}
	route_item := r.Group("/api/item")
	{
		route_item.GET("/nft/:nftid", controller.GetNftById)
		route_item.GET("/group/:groupid", controller.GetProductGroupById)
		route_item.GET("/collection/:collectionid", controller.GetCollectionById)
		route_item.POST("/nft", controller.CreateNftByGroupId)
	}
	configs.DB.AutoMigrate(&model.Obj{}, &model.Block{}, &model.NftTx{})
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
