import service from '@/utils/request'

// @Tags Seedling
// @Summary 创建Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "创建Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /seed/createSeedling [post]
export const createSeedling = (data) => {
     return service({
         url: "/seed/createSeedling",
         method: 'post',
         data
     })
 }


// @Tags Seedling
// @Summary 删除Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "删除Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /seed/deleteSeedling [delete]
 export const deleteSeedling = (data) => {
     return service({
         url: "/seed/deleteSeedling",
         method: 'delete',
         data
     })
 }

// @Tags Seedling
// @Summary 删除Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /seed/deleteSeedling [delete]
 export const deleteSeedlingByIds = (data) => {
     return service({
         url: "/seed/deleteSeedlingByIds",
         method: 'delete',
         data
     })
 }

// @Tags Seedling
// @Summary 更新Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "更新Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /seed/updateSeedling [put]
 export const updateSeedling = (data) => {
     return service({
         url: "/seed/updateSeedling",
         method: 'put',
         data
     })
 }


// @Tags Seedling
// @Summary 用id查询Seedling
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seedling true "用id查询Seedling"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /seed/findSeedling [get]
 export const findSeedling = (params) => {
     return service({
         url: "/seed/findSeedling",
         method: 'get',
         params
     })
 }


// @Tags Seedling
// @Summary 分页获取Seedling列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Seedling列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /seed/getSeedlingList [get]
 export const getSeedlingList = (params) => {
     return service({
         url: "/seed/getSeedlingList",
         method: 'get',
         params
     })
 }