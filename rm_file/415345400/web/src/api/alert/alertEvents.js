import service from '@/utils/request'
// @Tags AlertEvents
// @Summary 创建alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AlertEvents true "创建alertEvents表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /alertEvents/createAlertEvents [post]
export const createAlertEvents = (data) => {
  return service({
    url: '/alertEvents/createAlertEvents',
    method: 'post',
    data
  })
}

// @Tags AlertEvents
// @Summary 删除alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AlertEvents true "删除alertEvents表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alertEvents/deleteAlertEvents [delete]
export const deleteAlertEvents = (params) => {
  return service({
    url: '/alertEvents/deleteAlertEvents',
    method: 'delete',
    params
  })
}

// @Tags AlertEvents
// @Summary 批量删除alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除alertEvents表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alertEvents/deleteAlertEvents [delete]
export const deleteAlertEventsByIds = (params) => {
  return service({
    url: '/alertEvents/deleteAlertEventsByIds',
    method: 'delete',
    params
  })
}

// @Tags AlertEvents
// @Summary 更新alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AlertEvents true "更新alertEvents表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /alertEvents/updateAlertEvents [put]
export const updateAlertEvents = (data) => {
  return service({
    url: '/alertEvents/updateAlertEvents',
    method: 'put',
    data
  })
}

// @Tags AlertEvents
// @Summary 用id查询alertEvents表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AlertEvents true "用id查询alertEvents表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alertEvents/findAlertEvents [get]
export const findAlertEvents = (params) => {
  return service({
    url: '/alertEvents/findAlertEvents',
    method: 'get',
    params
  })
}

// @Tags AlertEvents
// @Summary 分页获取alertEvents表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取alertEvents表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /alertEvents/getAlertEventsList [get]
export const getAlertEventsList = (params) => {
  return service({
    url: '/alertEvents/getAlertEventsList',
    method: 'get',
    params
  })
}

// @Tags AlertEvents
// @Summary 不需要鉴权的alertEvents表接口
// @Accept application/json
// @Produce application/json
// @Param data query alertReq.AlertEventsSearch true "分页获取alertEvents表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alertEvents/getAlertEventsPublic [get]
export const getAlertEventsPublic = () => {
  return service({
    url: '/alertEvents/getAlertEventsPublic',
    method: 'get',
  })
}
