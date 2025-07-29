<template>
  <div class="unified-monitor-container">
    <!-- 左侧设备树 -->
    <div class="device-tree-panel">
      <div class="tree-header">
        <h3>设备列表</h3>
        <div class="tree-controls">
          <el-button size="small" type="primary" @click="refreshDeviceTree">
            <el-icon><Refresh /></el-icon>
          </el-button>
          <el-button size="small" type="success" @click="showDeviceManager = true">
            <el-icon><Setting /></el-icon>
          </el-button>
        </div>
      </div>
      
      <!-- 分组时间选择 -->
      <div class="time-selector">
        <el-select v-model="selectedTimeRange" size="small" @change="handleTimeRangeChange">
          <el-option label="1分钟" value="1" />
          <el-option label="5分钟" value="5" />
          <el-option label="4分钟" value="4" />
        </el-select>
      </div>

      <!-- 设备树 -->
      <div class="device-tree">
        <el-tree
          ref="deviceTreeRef"
          :data="deviceTreeData"
          :props="treeProps"
          node-key="id"
          :expand-on-click-node="false"
          :check-on-click-node="true"
          show-checkbox
          @check="handleDeviceCheck"
        >
          <template #default="{ node, data }">
            <div class="tree-node">
              <el-icon v-if="data.type === 'group'">
                <Folder />
              </el-icon>
              <el-icon v-else :color="getDeviceStatusColor(data.status)">
                <VideoCamera />
              </el-icon>
              <span class="node-label">{{ data.label }}</span>
              <span v-if="data.type === 'device'" class="device-status" :class="getDeviceStatusClass(data.status)">
                {{ getDeviceStatusText(data.status) }}
              </span>
            </div>
          </template>
        </el-tree>
      </div>
    </div>

    <!-- 右侧视频监控区域 -->
    <div class="video-monitor-panel">
      <!-- 顶部控制栏 -->
      <div class="monitor-header">
        <div class="left-controls">
          <span class="title">视频监控</span>
          <el-select 
            v-model="currentSplitMode" 
            class="split-mode-selector"
            @change="handleSplitModeChange"
          >
            <el-option label="4分屏" value="4" />
            <el-option label="9分屏" value="9" />
            <el-option label="16分屏" value="16" />
          </el-select>
        </div>
        <div class="right-controls">
          <el-button size="small" type="success" @click="startAllStreams">全部播放</el-button>
          <el-button size="small" type="warning" @click="stopAllStreams">全部停止</el-button>
          <el-button size="small" type="info" @click="clearAllSlots">清空</el-button>
        </div>
      </div>

      <!-- 视频网格 -->
      <div class="video-grid" :class="`grid-${currentSplitMode}`">
        <div 
          v-for="(slot, index) in videoSlots" 
          :key="index" 
          class="video-slot"
          :class="{ 'has-stream': slot.device }"
        >
          <!-- 空槽位 -->
          <div v-if="!slot.device" class="empty-slot">
            <el-icon :size="48" color="#409EFF">
              <Plus />
            </el-icon>
            <span>拖拽设备到此处</span>
          </div>
          
          <!-- 视频流 -->
          <div v-else class="video-container">
            <div class="video-header">
              <span class="device-name">{{ slot.device.deviceName }}</span>
              <div class="video-controls">
                <el-button 
                  :type="slot.isPlaying ? 'warning' : 'success'" 
                  size="small" 
                  @click="togglePlayPause(index)"
                >
                  <el-icon>
                    <VideoPlay v-if="!slot.isPlaying" />
                    <VideoPause v-else />
                  </el-icon>
                </el-button>
                <el-button type="info" size="small" @click="toggleFullscreen(index)">
                  <el-icon><FullScreen /></el-icon>
                </el-button>
                <el-button type="danger" size="small" @click="removeDeviceFromSlot(index)">
                  <el-icon><Close /></el-icon>
                </el-button>
              </div>
            </div>
            <div class="video-content">
              <video
                :id="`video-${index}`"
                :src="slot.device.streamUrl"
                :muted="true"
                :autoplay="slot.isPlaying"
                :controls="false"
                class="video-player"
                @error="handleVideoError(index)"
                @loadstart="handleVideoLoadStart(index)"
                @canplay="handleVideoCanPlay(index)"
              >
                您的浏览器不支持视频播放
              </video>
              <div v-if="slot.loading" class="loading-overlay">
                <el-icon class="is-loading" :size="32">
                  <Loading />
                </el-icon>
                <span>加载中...</span>
              </div>
              <div v-if="slot.error" class="error-overlay">
                <el-icon :size="32" color="#F56C6C">
                  <Warning />
                </el-icon>
                <span>视频加载失败</span>
              </div>
            </div>
            <div class="video-footer">
              <span class="device-status-badge" :class="getDeviceStatusClass(slot.device.status)">
                {{ getDeviceStatusText(slot.device.status) }}
              </span>
              <span class="device-resolution" v-if="slot.device.resolution">
                {{ slot.device.resolution }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 设备管理弹窗 -->
    <el-dialog 
      v-model="showDeviceManager" 
      title="设备管理" 
      width="90%"
      :before-close="closeDeviceManager"
    >
      <div class="device-manager-content">
        <el-tabs v-model="activeTab" type="border-card">
          <!-- 设备列表 -->
          <el-tab-pane label="设备列表" name="devices">
            <div class="manager-section">
              <div class="search-bar">
                <el-input 
                  v-model="deviceSearchText" 
                  placeholder="搜索设备名称" 
                  style="width: 300px; margin-right: 16px;"
                  clearable
                />
                <el-button type="primary" @click="searchDevices">搜索</el-button>
                <el-button @click="resetDeviceSearch">重置</el-button>
              </div>
              
              <el-table :data="filteredDevices" style="width: 100%; margin-top: 16px;">
                <el-table-column prop="deviceName" label="设备名称" />
                <el-table-column prop="streamUrl" label="视频流地址" show-overflow-tooltip />
                <el-table-column prop="status" label="状态">
                  <template #default="scope">
                    <el-tag :type="getStatusTagType(scope.row.status)">
                      {{ getDeviceStatusText(scope.row.status) }}
                    </el-tag>
                  </template>
                </el-table-column>
                <el-table-column prop="resolution" label="分辨率" />
                <el-table-column label="操作" width="120">
                  <template #default="scope">
                    <el-button 
                      type="primary" 
                      size="small" 
                      @click="addDeviceToMonitor(scope.row)"
                    >
                      添加监控
                    </el-button>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </el-tab-pane>
          
          <!-- 分组管理 -->
          <el-tab-pane label="分组管理" name="groups">
            <div class="manager-section">
              <el-table :data="allGroups" style="width: 100%;">
                <el-table-column prop="groupName" label="分组名称" />
                <el-table-column prop="deviceCount" label="设备数量" />
                <el-table-column prop="CreatedAt" label="创建时间">
                  <template #default="scope">
                    {{ formatDate(scope.row.CreatedAt) }}
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { 
  Plus, 
  VideoPlay, 
  VideoPause, 
  Close, 
  Loading, 
  Warning,
  FullScreen,
  Folder,
  VideoCamera,
  Refresh,
  Setting
} from '@element-plus/icons-vue'
import { getMonitorDeviceList } from '@/api/device/monitorDevice'
import { getDeviceGroupList } from '@/api/device/deviceGroup'

defineOptions({
  name: 'UnifiedMonitor'
})

// 响应式数据
const currentSplitMode = ref('9')
const selectedTimeRange = ref('5')
const showDeviceManager = ref(false)
const deviceTreeRef = ref()

// 设备树相关
const deviceTreeData = ref([])
const treeProps = {
  children: 'children',
  label: 'label'
}

// 视频槽位数据
const videoSlots = ref([])

// 设备管理相关
const activeTab = ref('devices')
const deviceSearchText = ref('')
const allDevices = ref([])
const allGroups = ref([])
const filteredDevices = ref([])

// 初始化视频槽位
const initVideoSlots = () => {
  const slotCount = parseInt(currentSplitMode.value)
  videoSlots.value = Array(slotCount).fill(null).map(() => ({
    device: null,
    isPlaying: false,
    loading: false,
    error: false
  }))
}

// 构建设备树数据
const buildDeviceTree = (groups, devices) => {
  const tree = []
  const groupMap = new Map()
  
  // 创建分组节点
  groups.forEach(group => {
    const node = {
      id: `group-${group.ID}`,
      label: group.groupName,
      type: 'group',
      children: [],
      groupId: group.ID,
      pid: group.pid
    }
    groupMap.set(group.ID, node)
  })
  
  // 构建分组层级关系
  groups.forEach(group => {
    const node = groupMap.get(group.ID)
    if (group.pid && group.pid !== 0) {
      const parent = groupMap.get(group.pid)
      if (parent) {
        parent.children.push(node)
      } else {
        tree.push(node)
      }
    } else {
      tree.push(node)
    }
  })
  
  // 添加设备到对应分组
  devices.forEach(device => {
    const deviceNode = {
      id: `device-${device.ID}`,
      label: device.deviceName,
      type: 'device',
      device: device,
      status: device.status
    }
    
    if (device.groupId) {
      const group = groupMap.get(device.groupId)
      if (group) {
        group.children.push(deviceNode)
      } else {
        tree.push(deviceNode)
      }
    } else {
      tree.push(deviceNode)
    }
  })
  
  return tree
}

// 加载设备树数据
const loadDeviceTree = async () => {
  try {
    const [groupsRes, devicesRes] = await Promise.all([
      getDeviceGroupList({ page: 1, pageSize: 1000 }),
      getMonitorDeviceList({ page: 1, pageSize: 1000 })
    ])
    
    if (groupsRes.code === 0 && devicesRes.code === 0) {
      const groups = groupsRes.data.list || []
      const devices = devicesRes.data.list || []
      deviceTreeData.value = buildDeviceTree(groups, devices)
    }
  } catch (error) {
    ElMessage.error('加载设备树失败')
    console.error('加载设备树失败:', error)
  }
}

// 刷新设备树
const refreshDeviceTree = () => {
  loadDeviceTree()
  ElMessage.success('设备树已刷新')
}

// 处理设备选中
const handleDeviceCheck = (data, { checkedNodes }) => {
  const deviceNodes = checkedNodes.filter(node => node.type === 'device')
  
  deviceNodes.forEach(node => {
    const device = node.device
    if (device && device.streamUrl) {
      addDeviceToEmptySlot(device)
    }
  })
}

// 添加设备到空槽位
const addDeviceToEmptySlot = (device) => {
  const emptySlotIndex = videoSlots.value.findIndex(slot => !slot.device)
  if (emptySlotIndex === -1) {
    ElMessage.warning('没有可用的视频槽位')
    return
  }
  
  // 检查设备是否已经在其他槽位中
  const existingSlot = videoSlots.value.find(slot => slot.device && slot.device.ID === device.ID)
  if (existingSlot) {
    ElMessage.warning('该设备已在监控中')
    return
  }
  
  videoSlots.value[emptySlotIndex] = {
    device: device,
    isPlaying: false,
    loading: true,
    error: false
  }
  
  ElMessage.success(`设备 ${device.deviceName} 已添加`)
  
  // 自动开始播放
  setTimeout(() => {
    togglePlayPause(emptySlotIndex)
  }, 500)
}

// 分屏模式切换
const handleSplitModeChange = () => {
  const oldSlots = [...videoSlots.value]
  initVideoSlots()
  
  // 保留现有设备到新的槽位中
  const newSlotCount = parseInt(currentSplitMode.value)
  oldSlots.slice(0, newSlotCount).forEach((slot, index) => {
    if (slot.device) {
      videoSlots.value[index] = slot
    }
  })
}

// 时间范围变化
const handleTimeRangeChange = () => {
  ElMessage.info(`已切换到 ${selectedTimeRange.value} 分钟视图`)
}

// 播放/暂停切换
const togglePlayPause = (slotIndex) => {
  const slot = videoSlots.value[slotIndex]
  if (!slot.device) return

  slot.isPlaying = !slot.isPlaying
  
  nextTick(() => {
    const videoElement = document.querySelector(`#video-${slotIndex}`)
    if (videoElement) {
      if (slot.isPlaying) {
        slot.loading = true
        slot.error = false
        videoElement.play().catch(error => {
          console.error('视频播放失败:', error)
          slot.error = true
          slot.loading = false
          slot.isPlaying = false
          ElMessage.error(`设备 ${slot.device.deviceName} 视频播放失败`)
        })
      } else {
        videoElement.pause()
        slot.loading = false
      }
    }
  })
}

// 全部播放
const startAllStreams = () => {
  videoSlots.value.forEach((slot, index) => {
    if (slot.device && !slot.isPlaying) {
      togglePlayPause(index)
    }
  })
}

// 全部停止
const stopAllStreams = () => {
  videoSlots.value.forEach((slot, index) => {
    if (slot.device && slot.isPlaying) {
      togglePlayPause(index)
    }
  })
}

// 清空所有槽位
const clearAllSlots = () => {
  stopAllStreams()
  initVideoSlots()
  // 清除树的选中状态
  if (deviceTreeRef.value) {
    deviceTreeRef.value.setCheckedKeys([])
  }
  ElMessage.success('已清空所有视频槽位')
}

// 从槽位移除设备
const removeDeviceFromSlot = (slotIndex) => {
  const slot = videoSlots.value[slotIndex]
  if (slot.device) {
    // 停止播放
    if (slot.isPlaying) {
      togglePlayPause(slotIndex)
    }
    
    // 清除树中对应设备的选中状态
    if (deviceTreeRef.value) {
      const deviceKey = `device-${slot.device.ID}`
      deviceTreeRef.value.setChecked(deviceKey, false)
    }
    
    videoSlots.value[slotIndex] = {
      device: null,
      isPlaying: false,
      loading: false,
      error: false
    }
    
    ElMessage.success('设备已移除')
  }
}

// 全屏切换
const toggleFullscreen = (slotIndex) => {
  const slot = videoSlots.value[slotIndex]
  if (!slot.device) return

  const videoElement = document.querySelector(`#video-${slotIndex}`)
  if (videoElement) {
    if (document.fullscreenElement) {
      document.exitFullscreen()
    } else {
      videoElement.requestFullscreen().catch(error => {
        console.error('全屏失败:', error)
        ElMessage.error('全屏功能不可用')
      })
    }
  }
}

// 视频事件处理
const handleVideoError = (slotIndex) => {
  videoSlots.value[slotIndex].error = true
  videoSlots.value[slotIndex].loading = false
  videoSlots.value[slotIndex].isPlaying = false
}

const handleVideoLoadStart = (slotIndex) => {
  videoSlots.value[slotIndex].loading = true
  videoSlots.value[slotIndex].error = false
}

const handleVideoCanPlay = (slotIndex) => {
  videoSlots.value[slotIndex].loading = false
  videoSlots.value[slotIndex].error = false
}

// 关闭设备管理器
const closeDeviceManager = () => {
  showDeviceManager.value = false
}

// 状态相关方法
const getDeviceStatusText = (status) => {
  const statusMap = {
    '1': '在线',
    '2': '离线',
    '3': '故障'
  }
  return statusMap[status] || '未知'
}

const getDeviceStatusClass = (status) => {
  const classMap = {
    '1': 'status-online',
    '2': 'status-offline',
    '3': 'status-error'
  }
  return classMap[status] || 'status-unknown'
}

const getDeviceStatusColor = (status) => {
  const colorMap = {
    '1': '#67C23A',
    '2': '#909399',
    '3': '#F56C6C'
  }
  return colorMap[status] || '#E6A23C'
}

// 生命周期
onMounted(() => {
  initVideoSlots()
  loadDeviceTree()
  loadDeviceManagerData()
  
  // 定时刷新设备状态
  const statusInterval = setInterval(() => {
    refreshDeviceStatus()
  }, 30000)
  
  onUnmounted(() => {
    clearInterval(statusInterval)
    stopAllStreams()
  })
})

// 刷新设备状态
const refreshDeviceStatus = async () => {
  try {
    const response = await getMonitorDeviceList({ page: 1, pageSize: 1000 })
    if (response.code === 0) {
      const updatedDevices = response.data.list
      
      // 更新槽位中的设备状态
      videoSlots.value.forEach((slot, index) => {
        if (slot.device) {
          const updatedDevice = updatedDevices.find(d => d.ID === slot.device.ID)
          if (updatedDevice) {
            slot.device = updatedDevice
          }
        }
      })
      
      // 更新设备树
      loadDeviceTree()
      
      // 更新设备管理数据
      loadDeviceManagerData()
    }
  } catch (error) {
    console.error('刷新设备状态失败:', error)
  }
}

// 加载设备管理数据
const loadDeviceManagerData = async () => {
  try {
    const [devicesRes, groupsRes] = await Promise.all([
      getMonitorDeviceList({ page: 1, pageSize: 1000 }),
      getDeviceGroupList({ page: 1, pageSize: 1000 })
    ])
    
    if (devicesRes.code === 0) {
      allDevices.value = devicesRes.data.list || []
      filteredDevices.value = allDevices.value
    }
    
    if (groupsRes.code === 0) {
      allGroups.value = (groupsRes.data.list || []).map(group => ({
        ...group,
        deviceCount: allDevices.value.filter(device => device.groupId === group.ID).length
      }))
    }
  } catch (error) {
    ElMessage.error('加载设备管理数据失败')
    console.error('加载设备管理数据失败:', error)
  }
}

// 搜索设备
const searchDevices = () => {
  if (!deviceSearchText.value.trim()) {
    filteredDevices.value = allDevices.value
  } else {
    filteredDevices.value = allDevices.value.filter(device => 
      device.deviceName.toLowerCase().includes(deviceSearchText.value.toLowerCase())
    )
  }
}

// 重置设备搜索
const resetDeviceSearch = () => {
  deviceSearchText.value = ''
  filteredDevices.value = allDevices.value
}

// 从设备管理添加设备到监控
const addDeviceToMonitor = (device) => {
  addDeviceToEmptySlot(device)
  closeDeviceManager()
}

// 获取状态标签类型
const getStatusTagType = (status) => {
  const typeMap = {
    '1': 'success',
    '2': 'info', 
    '3': 'danger'
  }
  return typeMap[status] || 'info'
}

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<style scoped>
.unified-monitor-container {
  display: flex;
  height: 100vh;
  background-color: #f5f5f5;
}

/* 左侧设备树面板 */
.device-tree-panel {
  width: 300px;
  background-color: white;
  border-right: 1px solid #e4e7ed;
  display: flex;
  flex-direction: column;
}

.tree-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #e4e7ed;
}

.tree-header h3 {
  margin: 0;
  font-size: 16px;
  color: #303133;
}

.tree-controls {
  display: flex;
  gap: 8px;
}

.time-selector {
  padding: 12px 16px;
  border-bottom: 1px solid #e4e7ed;
}

.device-tree {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
}

.tree-node {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.node-label {
  flex: 1;
  font-size: 14px;
}

.device-status {
  font-size: 12px;
  padding: 2px 6px;
  border-radius: 4px;
  color: white;
}

/* 右侧视频监控面板 */
.video-monitor-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
}

.monitor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  background-color: #f8f9fa;
  border-bottom: 1px solid #e4e7ed;
}

.left-controls {
  display: flex;
  gap: 16px;
  align-items: center;
}

.title {
  color: #303133;
  font-size: 18px;
  font-weight: 500;
}

.right-controls {
  display: flex;
  gap: 12px;
  align-items: center;
}

/* 分屏选择器样式 */
.split-mode-selector {
  width: 160px;
  font-size: 16px;
}

.split-mode-selector :deep(.el-input__wrapper) {
  height: 40px;
  font-size: 16px;
  font-weight: 500;
  background-color: #ffffff;
  border: 2px solid #409EFF;
  border-radius: 8px;
}

.split-mode-selector :deep(.el-input__inner) {
  color: #303133;
  font-size: 16px;
  font-weight: 500;
  text-align: center;
}

.split-mode-selector :deep(.el-input__suffix) {
  color: #303133;
}

.split-mode-selector :deep(.el-select__caret) {
  color: #303133;
  font-size: 16px;
}

.video-grid {
  flex: 1;
  display: grid;
  gap: 8px;
  padding: 16px;
  overflow: hidden;
}

.grid-4 {
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 1fr);
}

.grid-9 {
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(3, 1fr);
}

.grid-16 {
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: repeat(4, 1fr);
}

.video-slot {
  background-color: #ffffff;
  border: 2px solid #e4e7ed;
  border-radius: 8px;
  overflow: hidden;
  position: relative;
  min-height: 200px;
}

.video-slot.has-stream {
  border-color: #409EFF;
}

.empty-slot {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #909399;
  transition: all 0.3s ease;
}

.empty-slot:hover {
  background-color: #f5f7fa;
  border-color: #409EFF;
}

.empty-slot span {
  margin-top: 8px;
  font-size: 14px;
}

.video-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.video-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 12px;
  background-color: rgba(0, 0, 0, 0.7);
  border-bottom: 1px solid #3a3a3a;
}

.device-name {
  font-size: 12px;
  font-weight: 500;
  color: #ffffff;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 60%;
}

.video-controls {
  display: flex;
  gap: 4px;
}

.video-content {
  flex: 1;
  position: relative;
  background-color: #000;
}

.video-player {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.loading-overlay, .error-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.8);
  color: white;
}

.loading-overlay span, .error-overlay span {
  margin-top: 8px;
  font-size: 12px;
}

.video-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 4px 12px;
  background-color: rgba(0, 0, 0, 0.7);
  border-top: 1px solid #3a3a3a;
}

.device-status-badge {
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 4px;
}

.status-online {
  background-color: #67C23A;
  color: white;
}

.status-offline {
  background-color: #909399;
  color: white;
}

.status-error {
  background-color: #F56C6C;
  color: white;
}

.status-unknown {
  background-color: #E6A23C;
  color: white;
}

.device-resolution {
  font-size: 10px;
  color: #909399;
}

/* 设备管理弹窗样式 */
.device-manager-content {
  height: 600px;
}

.manager-section {
  padding: 20px;
}

.search-bar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  padding: 16px;
  background-color: #f5f5f5;
  border-radius: 4px;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .device-tree-panel {
    width: 250px;
  }
}

@media (max-width: 768px) {
  .unified-monitor-container {
    flex-direction: column;
  }
  
  .device-tree-panel {
    width: 100%;
    height: 200px;
  }
  
  .monitor-header {
    flex-direction: column;
    gap: 12px;
  }
  
  .grid-9 {
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(5, 1fr);
  }
  
  .grid-16 {
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(6, 1fr);
  }
  
  .search-bar {
    flex-direction: column;
    gap: 12px;
  }
  
  .search-bar .el-input {
    width: 100% !important;
    margin-right: 0 !important;
  }
}
</style>