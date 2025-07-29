import service from '@/utils/request'
// @Tags DeviceGroup
// @Summary 创建deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DeviceGroup true "创建deviceGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceGroup/createDeviceGroup [post]
export const createDeviceGroup = (data) => {
  return service({
    url: '/deviceGroup/createDeviceGroup',
    method: 'post',
    data
  })
}

// @Tags DeviceGroup
// @Summary 删除deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DeviceGroup true "删除deviceGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceGroup/deleteDeviceGroup [delete]
export const deleteDeviceGroup = (params) => {
  return service({
    url: '/deviceGroup/deleteDeviceGroup',
    method: 'delete',
    params
  })
}

// @Tags DeviceGroup
// @Summary 批量删除deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除deviceGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceGroup/deleteDeviceGroup [delete]
export const deleteDeviceGroupByIds = (params) => {
  return service({
    url: '/deviceGroup/deleteDeviceGroupByIds',
    method: 'delete',
    params
  })
}

// @Tags DeviceGroup
// @Summary 更新deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DeviceGroup true "更新deviceGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceGroup/updateDeviceGroup [put]
export const updateDeviceGroup = (data) => {
  return service({
    url: '/deviceGroup/updateDeviceGroup',
    method: 'put',
    data
  })
}

// @Tags DeviceGroup
// @Summary 用id查询deviceGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DeviceGroup true "用id查询deviceGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceGroup/findDeviceGroup [get]
export const findDeviceGroup = (params) => {
  return service({
    url: '/deviceGroup/findDeviceGroup',
    method: 'get',
    params
  })
}

// @Tags DeviceGroup
// @Summary 分页获取deviceGroup表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取deviceGroup表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceGroup/getDeviceGroupList [get]
export const getDeviceGroupList = (params) => {
  return service({
    url: '/deviceGroup/getDeviceGroupList',
    method: 'get',
    params
  })
}

// @Tags DeviceGroup
// @Summary 不需要鉴权的deviceGroup表接口
// @Accept application/json
// @Produce application/json
// @Param data query deviceReq.DeviceGroupSearch true "分页获取deviceGroup表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /deviceGroup/getDeviceGroupPublic [get]
export const getDeviceGroupPublic = () => {
  return service({
    url: '/deviceGroup/getDeviceGroupPublic',
    method: 'get',
  })
}
