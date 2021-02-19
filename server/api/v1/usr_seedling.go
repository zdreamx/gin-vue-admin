package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Seedling
// @Summary 创建Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "创建Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /seed/createSeedling [post]
func CreateSeedling(c *gin.Context) {
	var seed model.Seedling
	_ = c.ShouldBindJSON(&seed)
	if err := service.CreateSeedling(seed); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Seedling
// @Summary 删除Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "删除Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /seed/deleteSeedling [delete]
func DeleteSeedling(c *gin.Context) {
	var seed model.Seedling
	_ = c.ShouldBindJSON(&seed)
	if err := service.DeleteSeedling(seed); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Seedling
// @Summary 批量删除Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /seed/deleteSeedlingByIds [delete]
func DeleteSeedlingByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSeedlingByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Seedling
// @Summary 更新Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "更新Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /seed/updateSeedling [put]
func UpdateSeedling(c *gin.Context) {
	var seed model.Seedling
	_ = c.ShouldBindJSON(&seed)
	if err := service.UpdateSeedling(seed); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Seedling
// @Summary 用id查询Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "用id查询Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /seed/findSeedling [get]
func FindSeedling(c *gin.Context) {
	var seed model.Seedling
	_ = c.ShouldBindQuery(&seed)
	if err, reseed := service.GetSeedling(seed.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reseed": reseed}, c)
	}
}

// @Tags Seedling
// @Summary 分页获取Seedling列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SeedlingSearch true "分页获取Seedling列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /seed/getSeedlingList [get]
func GetSeedlingList(c *gin.Context) {
	var pageInfo request.SeedlingSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSeedlingInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
