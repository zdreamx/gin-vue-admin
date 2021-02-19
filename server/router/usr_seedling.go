package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSeedlingRouter(Router *gin.RouterGroup) {
	SeedlingRouter := Router.Group("seed").Use(middleware.OperationRecord())
	{
		SeedlingRouter.POST("createSeedling", v1.CreateSeedling)   // 新建Seedling
		SeedlingRouter.DELETE("deleteSeedling", v1.DeleteSeedling) // 删除Seedling
		SeedlingRouter.DELETE("deleteSeedlingByIds", v1.DeleteSeedlingByIds) // 批量删除Seedling
		SeedlingRouter.PUT("updateSeedling", v1.UpdateSeedling)    // 更新Seedling
		SeedlingRouter.GET("findSeedling", v1.FindSeedling)        // 根据ID获取Seedling
		SeedlingRouter.GET("getSeedlingList", v1.GetSeedlingList)  // 获取Seedling列表
	}
}
