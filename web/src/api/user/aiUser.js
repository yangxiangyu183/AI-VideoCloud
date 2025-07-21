import service from '@/utils/request'
// @Tags AiUser
// @Summary 创建aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AiUser true "创建aiUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aiUser/createAiUser [post]
export const createAiUser = (data) => {
  return service({
    url: '/aiUser/createAiUser',
    method: 'post',
    data
  })
}

// @Tags AiUser
// @Summary 删除aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AiUser true "删除aiUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aiUser/deleteAiUser [delete]
export const deleteAiUser = (params) => {
  return service({
    url: '/aiUser/deleteAiUser',
    method: 'delete',
    params
  })
}

// @Tags AiUser
// @Summary 批量删除aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除aiUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aiUser/deleteAiUser [delete]
export const deleteAiUserByIds = (params) => {
  return service({
    url: '/aiUser/deleteAiUserByIds',
    method: 'delete',
    params
  })
}

// @Tags AiUser
// @Summary 更新aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AiUser true "更新aiUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aiUser/updateAiUser [put]
export const updateAiUser = (data) => {
  return service({
    url: '/aiUser/updateAiUser',
    method: 'put',
    data
  })
}

// @Tags AiUser
// @Summary 用id查询aiUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AiUser true "用id查询aiUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aiUser/findAiUser [get]
export const findAiUser = (params) => {
  return service({
    url: '/aiUser/findAiUser',
    method: 'get',
    params
  })
}

// @Tags AiUser
// @Summary 分页获取aiUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取aiUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aiUser/getAiUserList [get]
export const getAiUserList = (params) => {
  return service({
    url: '/aiUser/getAiUserList',
    method: 'get',
    params
  })
}

// @Tags AiUser
// @Summary 不需要鉴权的aiUser表接口
// @Accept application/json
// @Produce application/json
// @Param data query userReq.AiUserSearch true "分页获取aiUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /aiUser/getAiUserPublic [get]
export const getAiUserPublic = () => {
  return service({
    url: '/aiUser/getAiUserPublic',
    method: 'get',
  })
}
