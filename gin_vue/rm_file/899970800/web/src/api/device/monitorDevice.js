import service from '@/utils/request'
// @Tags MonitorDevice
// @Summary 创建monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MonitorDevice true "创建monitorDevice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /monitorDevice/createMonitorDevice [post]
export const createMonitorDevice = (data) => {
  return service({
    url: '/monitorDevice/createMonitorDevice',
    method: 'post',
    data
  })
}

// @Tags MonitorDevice
// @Summary 删除monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MonitorDevice true "删除monitorDevice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monitorDevice/deleteMonitorDevice [delete]
export const deleteMonitorDevice = (params) => {
  return service({
    url: '/monitorDevice/deleteMonitorDevice',
    method: 'delete',
    params
  })
}

// @Tags MonitorDevice
// @Summary 批量删除monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除monitorDevice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monitorDevice/deleteMonitorDevice [delete]
export const deleteMonitorDeviceByIds = (params) => {
  return service({
    url: '/monitorDevice/deleteMonitorDeviceByIds',
    method: 'delete',
    params
  })
}

// @Tags MonitorDevice
// @Summary 更新monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MonitorDevice true "更新monitorDevice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /monitorDevice/updateMonitorDevice [put]
export const updateMonitorDevice = (data) => {
  return service({
    url: '/monitorDevice/updateMonitorDevice',
    method: 'put',
    data
  })
}

// @Tags MonitorDevice
// @Summary 用id查询monitorDevice表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MonitorDevice true "用id查询monitorDevice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monitorDevice/findMonitorDevice [get]
export const findMonitorDevice = (params) => {
  return service({
    url: '/monitorDevice/findMonitorDevice',
    method: 'get',
    params
  })
}

// @Tags MonitorDevice
// @Summary 分页获取monitorDevice表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取monitorDevice表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitorDevice/getMonitorDeviceList [get]
export const getMonitorDeviceList = (params) => {
  return service({
    url: '/monitorDevice/getMonitorDeviceList',
    method: 'get',
    params
  })
}

// @Tags MonitorDevice
// @Summary 不需要鉴权的monitorDevice表接口
// @Accept application/json
// @Produce application/json
// @Param data query deviceReq.MonitorDeviceSearch true "分页获取monitorDevice表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /monitorDevice/getMonitorDevicePublic [get]
export const getMonitorDevicePublic = () => {
  return service({
    url: '/monitorDevice/getMonitorDevicePublic',
    method: 'get',
  })
}
