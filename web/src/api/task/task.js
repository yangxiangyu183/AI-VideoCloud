import service from '@/utils/request'
// @Tags Task
// @Summary 创建监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Task true "创建监测任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /task表/createTask [post]
export const createTask = (data) => {
  return service({
    url: '/task表/createTask',
    method: 'post',
    data
  })
}

// @Tags Task
// @Summary 删除监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Task true "删除监测任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task表/deleteTask [delete]
export const deleteTask = (params) => {
  return service({
    url: '/task表/deleteTask',
    method: 'delete',
    params
  })
}

// @Tags Task
// @Summary 批量删除监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除监测任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task表/deleteTask [delete]
export const deleteTaskByIds = (params) => {
  return service({
    url: '/task表/deleteTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags Task
// @Summary 更新监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Task true "更新监测任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task表/updateTask [put]
export const updateTask = (data) => {
  return service({
    url: '/task表/updateTask',
    method: 'put',
    data
  })
}

// @Tags Task
// @Summary 用id查询监测任务
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Task true "用id查询监测任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task表/findTask [get]
export const findTask = (params) => {
  return service({
    url: '/task表/findTask',
    method: 'get',
    params
  })
}

// @Tags Task
// @Summary 分页获取监测任务列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取监测任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task表/getTaskList [get]
export const getTaskList = (params) => {
  return service({
    url: '/task表/getTaskList',
    method: 'get',
    params
  })
}

// @Tags Task
// @Summary 不需要鉴权的监测任务接口
// @Accept application/json
// @Produce application/json
// @Param data query task_borReq.TaskSearch true "分页获取监测任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /task表/getTaskPublic [get]
export const getTaskPublic = () => {
  return service({
    url: '/task表/getTaskPublic',
    method: 'get',
  })
}