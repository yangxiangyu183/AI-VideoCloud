import service from '@/utils/request'
// @Tags Alert
// @Summary 创建alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Alert true "创建alert表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /alert/createAlert [post]
export const createAlert = (data) => {
  return service({
    url: '/alert/createAlert',
    method: 'post',
    data
  })
}

// @Tags Alert
// @Summary 删除alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Alert true "删除alert表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alert/deleteAlert [delete]
export const deleteAlert = (params) => {
  return service({
    url: '/alert/deleteAlert',
    method: 'delete',
    params
  })
}

// @Tags Alert
// @Summary 批量删除alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除alert表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alert/deleteAlert [delete]
export const deleteAlertByIds = (params) => {
  return service({
    url: '/alert/deleteAlertByIds',
    method: 'delete',
    params
  })
}

// @Tags Alert
// @Summary 更新alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Alert true "更新alert表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /alert/updateAlert [put]
export const updateAlert = (data) => {
  return service({
    url: '/alert/updateAlert',
    method: 'put',
    data
  })
}

// @Tags Alert
// @Summary 用id查询alert表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Alert true "用id查询alert表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alert/findAlert [get]
export const findAlert = (params) => {
  return service({
    url: '/alert/findAlert',
    method: 'get',
    params
  })
}

// @Tags Alert
// @Summary 不需要鉴权的alert表接口
// @Accept application/json
// @Produce application/json
// @Param data query alert_videoReq.AlertSearch true "分页获取alert表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alert/getAlertPublic [get]
export const getAlertPublic = () => {
  return service({
    url: '/alert/getAlertPublic',
    method: 'get',
  })
}