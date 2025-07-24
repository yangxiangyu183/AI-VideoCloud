import service from '@/utils/request'
// @Tags CasbinRule
// @Summary 创建casbinRule表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "创建casbinRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /casbinRule/createCasbinRule [post]
export const createCasbinRule = (data) => {
  return service({
    url: '/casbinRule/createCasbinRule',
    method: 'post',
    data
  })
}

// @Tags CasbinRule
// @Summary 删除casbinRule表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "删除casbinRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /casbinRule/deleteCasbinRule [delete]
export const deleteCasbinRule = (params) => {
  return service({
    url: '/casbinRule/deleteCasbinRule',
    method: 'delete',
    params
  })
}

// @Tags CasbinRule
// @Summary 批量删除casbinRule表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除casbinRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /casbinRule/deleteCasbinRule [delete]
export const deleteCasbinRuleByIds = (params) => {
  return service({
    url: '/casbinRule/deleteCasbinRuleByIds',
    method: 'delete',
    params
  })
}

// @Tags CasbinRule
// @Summary 更新casbinRule表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "更新casbinRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbinRule/updateCasbinRule [put]
export const updateCasbinRule = (data) => {
  return service({
    url: '/casbinRule/updateCasbinRule',
    method: 'put',
    data
  })
}

// @Tags CasbinRule
// @Summary 用id查询casbinRule表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CasbinRule true "用id查询casbinRule表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /casbinRule/findCasbinRule [get]
export const findCasbinRule = (params) => {
  return service({
    url: '/casbinRule/findCasbinRule',
    method: 'get',
    params
  })
}

// @Tags CasbinRule
// @Summary 分页获取casbinRule表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取casbinRule表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbinRule/getCasbinRuleList [get]
export const getCasbinRuleList = (params) => {
  return service({
    url: '/casbinRule/getCasbinRuleList',
    method: 'get',
    params
  })
}

// @Tags CasbinRule
// @Summary 不需要鉴权的casbinRule表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.CasbinRuleSearch true "分页获取casbinRule表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /casbinRule/getCasbinRulePublic [get]
export const getCasbinRulePublic = () => {
  return service({
    url: '/casbinRule/getCasbinRulePublic',
    method: 'get',
  })
}
