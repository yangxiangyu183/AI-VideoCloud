// 设备状态工具函数

// 设备状态常量
export const DEVICE_STATUS = {
  ONLINE: '1',  // 在线
  OFFLINE: '2', // 离线
  FAULT: '3'    // 故障
}

// 设备状态中文映射
export const DEVICE_STATUS_TEXT = {
  [DEVICE_STATUS.ONLINE]: '在线',
  [DEVICE_STATUS.OFFLINE]: '离线',
  [DEVICE_STATUS.FAULT]: '故障'
}

// 设备状态颜色映射
export const DEVICE_STATUS_COLOR = {
  [DEVICE_STATUS.ONLINE]: '#52c41a',   // 绿色
  [DEVICE_STATUS.OFFLINE]: '#ff4d4f',  // 红色
  [DEVICE_STATUS.FAULT]: '#fa8c16'     // 橙色
}

// 获取设备状态文本
export function getDeviceStatusText(status) {
  return DEVICE_STATUS_TEXT[status] || '未知'
}

// 获取设备状态颜色
export function getDeviceStatusColor(status) {
  return DEVICE_STATUS_COLOR[status] || '#d9d9d9'
}

// 获取设备状态选项（用于筛选）
export function getDeviceStatusOptions() {
  return [
    { value: DEVICE_STATUS.ONLINE, label: '在线' },
    { value: DEVICE_STATUS.OFFLINE, label: '离线' },
    { value: DEVICE_STATUS.FAULT, label: '故障' }
  ]
}

// 判断设备是否在线
export function isDeviceOnline(status) {
  return status === DEVICE_STATUS.ONLINE
}

// 判断设备是否离线
export function isDeviceOffline(status) {
  return status === DEVICE_STATUS.OFFLINE
}

// 判断设备是否故障
export function isDeviceFault(status) {
  return status === DEVICE_STATUS.FAULT
}