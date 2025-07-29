<template>
  <div class="alert-page">
    <!-- 顶部导航栏 -->
    <div class="top-nav">
      <div class="nav-menu">
        <span class="nav-item active">事件告警</span>
        <el-select
            v-model="topNavCameraAddress"
            placeholder="摄像头点位"
            size="small"
            clearable
            style="width: 120px; margin-right: 16px;"
            @change="onTopNavCameraAddressChange"
        >
          <el-option
              v-for="item in cameraAddressOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
          />
        </el-select>
        <el-select
            v-model="topNavTimeRange"
            placeholder="时间范围"
            size="small"
            clearable
            style="width: 120px;"
            @change="onTopNavTimeRangeChange"
        >
          <el-option
              v-for="item in timeRangeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
          />
        </el-select>
      </div>
      <div class="nav-user">
        <el-button type="text">退出</el-button>
      </div>
    </div>
    <div class="main-content">
      <!-- 左侧筛选栏 -->
      <el-card class="filter-bar">
        <div class="filter-title">筛选条件</div>

        <!-- 设备状态筛选 -->
        <div class="filter-group-title">设备状态</div>
        <el-form :model="searchInfo" label-width="0px" class="filter-form">
          <el-form-item>
            <div class="filter-checkbox-list">
              <el-checkbox v-model="searchInfo.deviceStatusAll" @change="onDeviceStatusAllChange">全选</el-checkbox>
              <el-checkbox v-for="statusItem in deviceStatusOptions" :key="statusItem.value" v-model="statusItem.checked" @change="onDeviceStatusChange">
                <span class="status-dot" :style="{ backgroundColor: getDeviceStatusColor(statusItem.value) }"></span>
                {{ statusItem.label }}
              </el-checkbox>
            </div>
          </el-form-item>

          <!-- 预警类型筛选 -->
          <el-form-item>
            <div class="filter-checkbox-list">
              <div class="filter-group-title">预警类型</div>
              <el-checkbox v-model="searchInfo.alertTypeAll" @change="onAlertTypeAllChange">全选</el-checkbox>
              <el-checkbox v-for="item in alert_typeOptions" :key="item.value" v-model="item.checked" @change="onAlertTypeChange">{{ item.label }}</el-checkbox>
            </div>
          </el-form-item>

          <!-- 摄像头点位筛选 -->
          <el-form-item>
            <div class="filter-checkbox-list">
              <div class="filter-group-title">摄像头点位</div>
              <el-checkbox v-model="searchInfo.cameraAddressAll" @change="onCameraAddressAllChange">全选</el-checkbox>
              <el-checkbox v-for="cameraItem in cameraAddressOptions" :key="cameraItem.value" v-model="cameraItem.checked" @change="onCameraAddressChange">{{ cameraItem.label }}</el-checkbox>
            </div>
          </el-form-item>

          <!-- 时间范围筛选 -->
          <el-form-item>
            <div class="filter-checkbox-list">
              <div class="filter-group-title">时间范围</div>
              <el-checkbox v-model="searchInfo.timeRangeAll" @change="onTimeRangeAllChange">全选</el-checkbox>
              <el-checkbox v-for="timeItem in timeRangeOptions" :key="timeItem.value" v-model="timeItem.checked" @change="onTimeRangeChange">{{ timeItem.label }}</el-checkbox>
            </div>
          </el-form-item>

          <el-form-item>
            <div class="filter-group-title">时间范围</div>
            <el-date-picker
                v-model="searchInfo.createdAtRange"
                type="datetimerange"
                range-separator="至"
                start-placeholder="开始时间"
                end-placeholder="结束时间"
                size="small"
                style="width: 100%"
                @change="onCustomTimeRangeChange"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit" size="small">搜索</el-button>
            <el-button @click="onReset" size="small">重置</el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <!-- 右侧卡片内容区 -->
      <div class="card-list-area">
        <div class="card-list-header">
          <div class="statistics-section">
            <span class="total-count">共有 {{ total }} 条警告事件</span>
            <div v-if="hasActiveFilters" class="filter-summary">
              <span class="filter-label">当前筛选:</span>
              <el-tag 
                v-for="tag in activeFilterTags" 
                :key="tag.key"
                size="small"
                type="info"
                closable
                @close="removeFilterTag(tag)"
                style="margin-right: 8px;"
              >
                {{ tag.label }}
              </el-tag>
              <el-button 
                type="text" 
                size="small" 
                @click="clearAllFilters"
                style="color: #409eff;"
              >
                清空筛选
              </el-button>
            </div>
          </div>
          <el-input
              v-model="searchInfo.keyword"
              placeholder="摄像头点位/报警类型/设备名称"
              size="small"
              class="search-input"
              @input="onKeywordChange"
              @keyup.enter="onSubmit"
              style="width: 260px; margin-left: 20px;"
          >
            <template #append>
              <el-button icon="el-icon-search" @click="onSubmit" />
            </template>
          </el-input>
        </div>
        <div class="card-scroll-container">
          <!-- 空状态提示 -->
          <div v-if="tableData.length === 0" class="empty-state">
            <div class="empty-icon">
              <el-icon size="64" color="#666">
                <Search />
              </el-icon>
            </div>
            <div class="empty-text">
              <h3>暂无数据</h3>
              <p v-if="hasActiveFilters">当前筛选条件下没有找到相关告警事件，请尝试调整筛选条件</p>
              <p v-else>暂时没有告警事件数据</p>
            </div>
            <div v-if="hasActiveFilters" class="empty-actions">
              <el-button type="primary" @click="clearAllFilters">清空筛选条件</el-button>
            </div>
          </div>
          
          <!-- 卡片网格 -->
          <div v-else class="card-grid">
            <div v-for="item in tableData" :key="item.ID" class="card-item">
              <el-card shadow="hover" class="event-card">
                <div class="card-img-box" @click="getDetails(item)">
                  <img :src="item.video || '/src/assets/default-event.jpg'" class="card-img" />
                  <div class="img-red-rect"></div>
                </div>
                <div class="card-info">
                  <div class="card-info-row">
                    <span class="device-status" :style="{ color: getDeviceStatusColor(item.deviceStatus) }">
                      <span class="status-dot" :style="{ backgroundColor: getDeviceStatusColor(item.deviceStatus) }"></span>
                      {{ getDeviceStatusText(item.deviceStatus) }}
                    </span>
                  </div>
                  <div class="card-info-row"><b>摄像头点位：</b>{{ item.cameraAddress || '未知点位' }}</div>
                  <div class="card-info-row"><b>预警类型：</b>{{ getAlertTypeLabel(item.alertType) }}</div>
                  <div class="card-info-row"><b>预警时间：</b>{{ formatDate(item.alertTime) }}</div>
                  <div class="card-info-row"><b>分辨率：</b>{{ item.resolution || '未知' }}</div>
                  <div class="card-info-row"><b>视频流地址：</b>{{ item.streamUrl || '未配置' }}</div>
                </div>
                <div class="card-actions">
                  <el-button type="primary" link @click="getDetails(item)">查看</el-button>
                </div>
              </el-card>
            </div>
          </div>
        </div>
        <div class="gva-pagination">
          <el-pagination
              layout="total, sizes, prev, pager, next, jumper"
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[6, 12, 24, 48]"
              :total="total"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
          />
        </div>
      </div>
    </div>
    <!-- 详情弹窗 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="告警详情">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="设备状态">
          <span :style="{ color: getDeviceStatusColor(detailForm.deviceStatus) }">
            <span class="status-dot" :style="{ backgroundColor: getDeviceStatusColor(detailForm.deviceStatus) }"></span>
            {{ getDeviceStatusText(detailForm.deviceStatus) }}
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="摄像头点位">{{ detailForm.cameraAddress || '未知点位' }}</el-descriptions-item>
        <el-descriptions-item label="预警类型">{{ getAlertTypeLabel(detailForm.alertType) }}</el-descriptions-item>
        <el-descriptions-item label="预警时间">{{ formatDate(detailForm.alertTime) }}</el-descriptions-item>
        <el-descriptions-item label="分辨率">{{ detailForm.resolution || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="视频流地址">{{ detailForm.streamUrl || '未配置' }}</el-descriptions-item>
        <el-descriptions-item label="监控视频">{{ detailForm.video || '无' }}</el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  findAlert,
  getAlertList
} from '@/api/alert_video/alert'
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { getDeviceStatusText, getDeviceStatusColor, getDeviceStatusOptions } from '@/utils/deviceStatus'
import { ref, watch, nextTick, computed } from 'vue'

// 防抖函数
const debounce = (func, wait) => {
  let timeout
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}
import { useAppStore } from "@/pinia"
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'

const appStore = useAppStore()

// 字典
const alert_typeOptions = ref([])
const deviceStatusOptions = ref([])
const cameraAddressOptions = ref([])
const eventScopeOptions = ref([])

// 选择模式
const alertTypeMultipleMode = ref(false)
const deviceStatusSelectMode = ref('multiple')

const setOptions = async () =>{
  try {
    const alertTypes = await getDictFunc('alert_type')
    alert_typeOptions.value = alertTypes.map(item => ({
      ...item,
      checked: false
    }))
  } catch (error) {
    console.error('获取预警类型数据失败:', error)
    // 如果API获取失败，使用模拟数据
    alert_typeOptions.value = [
      { value: 'smoking_detection', label: '吸烟检测', checked: false },
      { value: 'ground_garbage', label: '地面垃圾', checked: false },
      { value: 'yacht_boat_recognition', label: '游艇小艇识别', checked: false },
      { value: 'vehicle_type_recognition', label: '车辆类型识别', checked: false },
      { value: 'non_motor_vehicle_recognition', label: '非机动车识别', checked: false },
      { value: 'drone_recognition', label: '无人机识别', checked: false },
      { value: 'small_boat_raft_recognition', label: '小船皮筏识别', checked: false },
      { value: 'person_intrusion', label: '人员入侵', checked: false }
    ]
  }

  deviceStatusOptions.value = getDeviceStatusOptions().map(item => ({
    ...item,
    checked: false
  }))
  // 获取摄像头点位数据
  await getCameraAddressOptions()
  // 获取事件范围数据
  await getEventScopeOptions()

  // 在选项初始化完成后，加载数据
  // getTableData() // 暂时注释掉，使用直接数据初始化

  // 直接初始化数据
  initializeData()
}
setOptions()

// 直接初始化数据
const initializeData = () => {
  console.log('=== 初始化数据 ===')

  // 设置完整的筛选结果为所有数据
  filteredResults.value = [...allMockData.value]
  total.value = allMockData.value.length
  page.value = 1
  
  // 应用分页显示
  applyPagination()

  console.log('✅ 初始化完成，总数据数量:', filteredResults.value.length)
  console.log('✅ 当前页显示数量:', tableData.value.length)
  console.log('✅ 初始化数据:', tableData.value.map(item => `ID${item.ID}:${item.alertType}`))
}

// 获取摄像头点位选项
const getCameraAddressOptions = async () => {
  try {
    // 根据卡片数据中的摄像头点位生成选项
    cameraAddressOptions.value = [
      { value: '杭州市区', label: '杭州市区', checked: false },
      { value: '马路路段', label: '马路路段', checked: false },
      { value: '十字路', label: '十字路', checked: false },
      { value: '村村桥中', label: '村村桥中', checked: false },
      { value: '怀柔密云', label: '怀柔密云', checked: false },
      { value: '新济南广场', label: '新济南广场', checked: false }
    ]
  } catch (error) {
    console.error('获取摄像头点位数据失败:', error)
  }
}

// 获取事件范围选项
const getEventScopeOptions = async () => {
  try {
    // 这里可以调用API获取事件范围数据，暂时使用模拟数据
    eventScopeOptions.value = [
      { value: '全部区域', label: '全部区域', checked: false },
      { value: '重点区域', label: '重点区域', checked: false },
      { value: '一般区域', label: '一般区域', checked: false },
      { value: '敏感区域', label: '敏感区域', checked: false },
      { value: '公共区域', label: '公共区域', checked: false }
    ]
  } catch (error) {
    console.error('获取事件范围数据失败:', error)
  }
}

// 顶部导航栏选择
const topNavCameraAddress = ref('')
const topNavTimeRange = ref('')

// 时间范围选项
const timeRangeOptions = ref([
  { value: 'today', label: '今天', checked: false },
  { value: 'yesterday', label: '昨天', checked: false },
  { value: 'last3days', label: '最近3天', checked: false },
  { value: 'last7days', label: '最近7天', checked: false },
  { value: 'last30days', label: '最近30天', checked: false },
  { value: 'thisMonth', label: '本月', checked: false },
  { value: 'lastMonth', label: '上月', checked: false }
])

// 统一的筛选状态管理
const filterState = ref({
  deviceStatus: {
    all: false,
    selected: []
  },
  alertType: {
    all: false,
    selected: []
  },
  cameraAddress: {
    all: false,
    selected: []
  },
  timeRange: {
    all: false,
    selected: [],
    customRange: []
  },
  keyword: ''
})

// 保持向后兼容的查询条件
const searchInfo = ref({
  alertTypeAll: false,
  alertTypeList: [],
  alertTypeSingle: '',
  deviceStatusAll: false,
  deviceStatusList: [],
  deviceStatusSingle: '',
  cameraAddressAll: false,
  timeRangeAll: false,
  cameraAddress: '',
  deviceName: '',
  createdAtRange: [],
  keyword: ''
})

// 完整的模拟数据集
const allMockData = ref([
  {
    ID: 1,
    video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
    cameraAddress: '杭州市区',
    alertType: 'small_boat_raft_recognition',
    alertTime: '2025-07-17 11:40:00',
    deviceName: '杭州市区小船监控设备HK-001',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.100:1935/live/stream1001'
  },
  {
    ID: 2,
    video: 'https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=400&h=300&fit=crop',
    cameraAddress: '杭州市区',
    alertType: 'small_boat_raft_recognition',
    alertTime: '2025-07-17 10:28:00',
    deviceName: '杭州市区小船监控设备DH-002',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.101:1935/live/stream1002'
  },
  {
    ID: 3,
    video: 'https://images.unsplash.com/photo-1449824913935-59a10b8d2000?w=400&h=300&fit=crop',
    cameraAddress: '马路路段',
    alertType: 'vehicle_type_recognition',
    alertTime: '2025-07-17 16:27:00',
    deviceName: '马路路段车辆监控设备UV-003',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.102:1935/live/stream1003'
  },
  {
    ID: 4,
    video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
    cameraAddress: '十字路',
    alertType: 'drone_recognition',
    alertTime: '2025-07-17 16:20:00',
    deviceName: '十字路口无人机监控设备HK-004',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.103:1935/live/stream1004'
  },
  {
    ID: 5,
    video: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
    cameraAddress: '村村桥中',
    alertType: 'non_motor_vehicle_recognition',
    alertTime: '2025-07-17 17:32:00',
    deviceName: '村村桥非机动车监控设备DH-005',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.104:1935/live/stream1005'
  },
  {
    ID: 6,
    video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
    cameraAddress: '怀柔密云',
    alertType: 'yacht_boat_recognition',
    alertTime: '2025-07-17 09:23:00',
    deviceName: '怀柔密云水域监控设备UV-006',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.105:1935/live/stream1006'
  },
  {
    ID: 7,
    video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
    cameraAddress: '新济南广场',
    alertType: 'person_intrusion',
    alertTime: '2025-07-17 09:05:00',
    deviceName: '新济南广场人员监控设备DH-007',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.106:1935/live/stream1007'
  },
  {
    ID: 8,
    video: 'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
    cameraAddress: '村村桥中',
    alertType: 'ground_garbage',
    alertTime: '2025-07-17 11:26:00',
    deviceName: '村村桥垃圾监控设备HK-008',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.107:1935/live/stream1008'
  },
  {
    ID: 9,
    video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
    cameraAddress: '杭州市区',
    alertType: 'smoking_detection',
    alertTime: '2025-07-17 08:15:00',
    deviceName: '杭州市区吸烟监控设备HK-009',
    deviceStatus: '1',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.108:1935/live/stream1009'
  },
  {
    ID: 10,
    video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
    cameraAddress: '马路路段',
    alertType: 'smoking_detection',
    alertTime: '2025-07-17 14:22:00',
    deviceName: '马路路段吸烟监控设备UV-010',
    deviceStatus: '2',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.109:1935/live/stream1010'
  },
  {
    ID: 11,
    video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
    cameraAddress: '十字路',
    alertType: 'ground_garbage',
    alertTime: '2025-07-17 13:45:00',
    deviceName: '十字路垃圾监控设备HK-011',
    deviceStatus: '3',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.110:1935/live/stream1011'
  },
  {
    ID: 12,
    video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
    cameraAddress: '怀柔密云',
    alertType: 'person_intrusion',
    alertTime: '2025-07-17 15:30:00',
    deviceName: '怀柔密云人员监控设备UV-012',
    deviceStatus: '2',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.111:1935/live/stream1012'
  },
  {
    ID: 13,
    video: 'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
    cameraAddress: '新济南广场',
    alertType: 'yacht_boat_recognition',
    alertTime: '2025-07-17 12:15:00',
    deviceName: '新济南广场水域监控设备HK-013',
    deviceStatus: '3',
    resolution: '1920x1080',
    streamUrl: 'rtmp://192.168.1.112:1935/live/stream1013'
  }
])

// 核心筛选引擎
const applyAllFilters = (data, filters) => {
  try {
    if (!data || !Array.isArray(data)) {
      console.warn('筛选数据无效:', data)
      return []
    }
    
    let result = [...data]
    let filterCount = 0
    
    // 按设备状态筛选
    if (!filters.deviceStatus.all && filters.deviceStatus.selected.length > 0) {
      result = result.filter(item => 
        filters.deviceStatus.selected.includes(item.deviceStatus)
      )
      filterCount++
      console.log(`设备状态筛选后数量: ${result.length}`)
    }
    
    // 按预警类型筛选
    if (!filters.alertType.all && filters.alertType.selected.length > 0) {
      result = result.filter(item => 
        filters.alertType.selected.includes(item.alertType)
      )
      filterCount++
      console.log(`预警类型筛选后数量: ${result.length}`)
    }
    
    // 按摄像头点位筛选
    if (!filters.cameraAddress.all && filters.cameraAddress.selected.length > 0) {
      result = result.filter(item => 
        filters.cameraAddress.selected.includes(item.cameraAddress)
      )
      filterCount++
      console.log(`摄像头点位筛选后数量: ${result.length}`)
    }
    
    // 按时间范围筛选
    if (filters.timeRange.customRange && filters.timeRange.customRange.length === 2) {
      const [startTime, endTime] = filters.timeRange.customRange
      result = result.filter(item => {
        try {
          const itemTime = new Date(item.alertTime)
          return itemTime >= startTime && itemTime <= endTime
        } catch (error) {
          console.warn('时间解析错误:', item.alertTime, error)
          return false
        }
      })
      filterCount++
      console.log(`时间范围筛选后数量: ${result.length}`)
    }
    
    // 按关键词筛选
    if (filters.keyword && filters.keyword.trim()) {
      const keyword = filters.keyword.trim().toLowerCase()
      result = result.filter(item => {
        try {
          return item.cameraAddress.toLowerCase().includes(keyword) ||
                 getAlertTypeLabel(item.alertType).toLowerCase().includes(keyword) ||
                 item.deviceName.toLowerCase().includes(keyword)
        } catch (error) {
          console.warn('关键词筛选错误:', item, error)
          return false
        }
      })
      filterCount++
      console.log(`关键词筛选后数量: ${result.length}`)
    }
    
    console.log(`应用了 ${filterCount} 个筛选条件，最终结果数量: ${result.length}`)
    return result
    
  } catch (error) {
    console.error('筛选过程中发生错误:', error)
    return data || []
  }
}

// 单独的筛选函数，用于性能优化
const filterByDeviceStatus = (data, statusList) => {
  if (!statusList || statusList.length === 0) return data
  return data.filter(item => statusList.includes(item.deviceStatus))
}

const filterByAlertType = (data, typeList) => {
  if (!typeList || typeList.length === 0) return data
  return data.filter(item => typeList.includes(item.alertType))
}

const filterByCameraAddress = (data, addressList) => {
  if (!addressList || addressList.length === 0) return data
  return data.filter(item => addressList.includes(item.cameraAddress))
}

const filterByTimeRange = (data, timeRange) => {
  if (!timeRange || timeRange.length !== 2) return data
  const [startTime, endTime] = timeRange
  return data.filter(item => {
    try {
      const itemTime = new Date(item.alertTime)
      return itemTime >= startTime && itemTime <= endTime
    } catch (error) {
      console.warn('时间解析错误:', item.alertTime, error)
      return false
    }
  })
}

const filterByKeyword = (data, keyword) => {
  if (!keyword || !keyword.trim()) return data
  const searchTerm = keyword.trim().toLowerCase()
  return data.filter(item => {
    try {
      return item.cameraAddress.toLowerCase().includes(searchTerm) ||
             getAlertTypeLabel(item.alertType).toLowerCase().includes(searchTerm) ||
             item.deviceName.toLowerCase().includes(searchTerm)
    } catch (error) {
      console.warn('关键词筛选错误:', item, error)
      return false
    }
  })
}

// 筛选结果缓存
const filterCache = ref(new Map())
const lastFilterState = ref(null)

// 生成筛选条件的缓存键
const generateCacheKey = (filters) => {
  return JSON.stringify({
    deviceStatus: filters.deviceStatus.selected.sort(),
    alertType: filters.alertType.selected.sort(),
    cameraAddress: filters.cameraAddress.selected.sort(),
    timeRange: filters.timeRange.customRange,
    keyword: filters.keyword
  })
}

// 检查筛选条件是否发生变化
const hasFilterChanged = (newFilters, oldFilters) => {
  if (!oldFilters) return true
  return generateCacheKey(newFilters) !== generateCacheKey(oldFilters)
}

// 存储完整的筛选结果，用于分页
const filteredResults = ref([])

// 优化的筛选函数，支持缓存和增量筛选
const optimizedFilter = () => {
  const startTime = performance.now()
  
  try {
    // 检查是否有缓存
    const cacheKey = generateCacheKey(filterState.value)
    if (filterCache.value.has(cacheKey)) {
      const cachedResult = filterCache.value.get(cacheKey)
      filteredResults.value = cachedResult
      total.value = cachedResult.length
      page.value = 1
      
      // 应用分页显示
      applyPagination()
      
      const endTime = performance.now()
      console.log(`✅ 使用缓存筛选完成，结果数量: ${cachedResult.length}，耗时: ${(endTime - startTime).toFixed(2)}ms`)
      return
    }
    
    // 执行筛选
    const filteredData = applyAllFilters(allMockData.value, filterState.value)
    
    // 缓存结果（限制缓存大小）
    if (filterCache.value.size > 50) {
      // 清除最旧的缓存项
      const firstKey = filterCache.value.keys().next().value
      filterCache.value.delete(firstKey)
    }
    filterCache.value.set(cacheKey, filteredData)
    
    // 更新筛选结果
    filteredResults.value = filteredData
    total.value = filteredData.length
    page.value = 1
    
    // 应用分页显示
    applyPagination()
    
    // 记录最后的筛选状态
    lastFilterState.value = JSON.parse(JSON.stringify(filterState.value))
    
    const endTime = performance.now()
    console.log(`✅ 筛选完成，结果数量: ${filteredData.length}，耗时: ${(endTime - startTime).toFixed(2)}ms`)
    
  } catch (error) {
    console.error('筛选过程中发生错误:', error)
    // 发生错误时显示所有数据
    filteredResults.value = [...allMockData.value]
    total.value = allMockData.value.length
    page.value = 1
    applyPagination()
  }
}

// 应用分页逻辑
const applyPagination = () => {
  const startIndex = (page.value - 1) * pageSize.value
  const endIndex = startIndex + pageSize.value
  tableData.value = filteredResults.value.slice(startIndex, endIndex)
  
  console.log(`📄 分页显示: 第${page.value}页，每页${pageSize.value}条，显示${startIndex + 1}-${Math.min(endIndex, filteredResults.value.length)}条，共${filteredResults.value.length}条`)
  console.log(`📄 当前页数据:`, tableData.value.map(item => `ID${item.ID}`))
}

// 防抖处理的筛选函数
const debouncedFilter = debounce(optimizedFilter, 300)

// 智能触发筛选（只在条件真正变化时触发）
const triggerFilter = () => {
  if (hasFilterChanged(filterState.value, lastFilterState.value)) {
    console.log('筛选条件发生变化，触发筛选')
    debouncedFilter()
  } else {
    console.log('筛选条件未变化，跳过筛选')
  }
}

// 更新筛选条件并触发筛选
const updateFilterCondition = (category, selectedValues, isAll = false) => {
  filterState.value[category].selected = selectedValues
  filterState.value[category].all = isAll
  
  // 同步到旧的searchInfo结构以保持兼容性
  if (category === 'deviceStatus') {
    searchInfo.value.deviceStatusAll = isAll
    searchInfo.value.deviceStatusList = selectedValues
  } else if (category === 'alertType') {
    searchInfo.value.alertTypeAll = isAll
    searchInfo.value.alertTypeList = selectedValues
  } else if (category === 'cameraAddress') {
    searchInfo.value.cameraAddressAll = isAll
  } else if (category === 'timeRange') {
    searchInfo.value.timeRangeAll = isAll
  }
  
  triggerFilter()
}

// 全选逻辑处理
const onAlertTypeAllChange = (val) => {
  alert_typeOptions.value.forEach(item => {
    item.checked = val
  })
  
  const selectedValues = val ? alert_typeOptions.value.map(item => item.value) : []
  updateFilterCondition('alertType', selectedValues, val)
}

const onAlertTypeChange = () => {
  console.log('=== 预警类型筛选触发 ===')

  // 获取选中的预警类型
  const selectedAlertTypes = alert_typeOptions.value.filter(item => item.checked).map(item => item.value)
  console.log('选中的预警类型:', selectedAlertTypes)

  // 检查是否需要更新全选状态
  const checkedCount = selectedAlertTypes.length
  const isAllSelected = checkedCount === alert_typeOptions.value.length
  searchInfo.value.alertTypeAll = isAllSelected

  // 更新筛选条件并触发筛选
  updateFilterCondition('alertType', selectedAlertTypes, isAllSelected)
}

// 测试筛选功能
const testFilter = () => {
  console.log('=== 测试筛选功能 ===')

  // 获取选中的预警类型
  const selectedAlertTypes = alert_typeOptions.value.filter(item => item.checked).map(item => item.value)
  console.log('选中的预警类型:', selectedAlertTypes)

  // 使用完整的模拟数据
  const allMockData = [
    {
      ID: 1,
      video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
      cameraAddress: '杭州市区',
      alertType: 'small_boat_raft_recognition',
      alertTime: '2025-07-17 11:40:00',
      deviceName: '杭州市区小船监控设备HK-001',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.100:1935/live/stream1001'
    },
    {
      ID: 2,
      video: 'https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=400&h=300&fit=crop',
      cameraAddress: '杭州市区',
      alertType: 'small_boat_raft_recognition',
      alertTime: '2025-07-17 10:28:00',
      deviceName: '杭州市区小船监控设备DH-002',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.101:1935/live/stream1002'
    },
    {
      ID: 3,
      video: 'https://images.unsplash.com/photo-1449824913935-59a10b8d2000?w=400&h=300&fit=crop',
      cameraAddress: '马路路段',
      alertType: 'vehicle_type_recognition',
      alertTime: '2025-07-17 16:27:00',
      deviceName: '马路路段车辆监控设备UV-003',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.102:1935/live/stream1003'
    },
    {
      ID: 4,
      video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
      cameraAddress: '十字路',
      alertType: 'drone_recognition',
      alertTime: '2025-07-17 16:20:00',
      deviceName: '十字路口无人机监控设备HK-004',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.103:1935/live/stream1004'
    },
    {
      ID: 5,
      video: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
      cameraAddress: '村村桥中',
      alertType: 'non_motor_vehicle_recognition',
      alertTime: '2025-07-17 17:32:00',
      deviceName: '村村桥非机动车监控设备DH-005',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.104:1935/live/stream1005'
    },
    {
      ID: 6,
      video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
      cameraAddress: '怀柔密云',
      alertType: 'yacht_boat_recognition',
      alertTime: '2025-07-17 09:23:00',
      deviceName: '怀柔密云水域监控设备UV-006',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.105:1935/live/stream1006'
    },
    {
      ID: 7,
      video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
      cameraAddress: '新济南广场',
      alertType: 'person_intrusion',
      alertTime: '2025-07-17 09:05:00',
      deviceName: '新济南广场人员监控设备DH-007',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.106:1935/live/stream1007'
    },
    {
      ID: 8,
      video: 'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
      cameraAddress: '村村桥中',
      alertType: 'ground_garbage',
      alertTime: '2025-07-17 11:26:00',
      deviceName: '村村桥垃圾监控设备HK-008',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.107:1935/live/stream1008'
    }
  ]

  if (selectedAlertTypes.length > 0) {
    const filteredData = allMockData.filter(item => selectedAlertTypes.includes(item.alertType))
    console.log('筛选后的数据:', filteredData)
    console.log('筛选后的数据数量:', filteredData.length)

    // 直接设置到tableData进行显示
    tableData.value = filteredData
    total.value = filteredData.length
    console.log('✅ 成功设置tableData，数量:', tableData.value.length)
  } else {
    console.log('没有选中预警类型，显示所有数据')
    tableData.value = allMockData
    total.value = allMockData.length
    console.log('✅ 显示所有数据，数量:', tableData.value.length)
  }
}

const onCameraAddressAllChange = (val) => {
  cameraAddressOptions.value.forEach(item => {
    item.checked = val
  })
  
  const selectedValues = val ? cameraAddressOptions.value.map(item => item.value) : []
  updateFilterCondition('cameraAddress', selectedValues, val)
}

const onCameraAddressChange = () => {
  console.log('=== 摄像头点位筛选触发 ===')
  
  // 获取选中的摄像头点位
  const selectedCameraAddress = cameraAddressOptions.value.filter(item => item.checked).map(item => item.value)
  console.log('选中的摄像头点位:', selectedCameraAddress)
  
  // 检查是否需要更新全选状态
  const checkedCount = selectedCameraAddress.length
  const isAllSelected = checkedCount === cameraAddressOptions.value.length
  searchInfo.value.cameraAddressAll = isAllSelected
  
  // 更新筛选条件并触发筛选
  updateFilterCondition('cameraAddress', selectedCameraAddress, isAllSelected)
}

const onTimeRangeAllChange = (val) => {
  timeRangeOptions.value.forEach(item => {
    item.checked = val
  })
  
  // 如果全选，则清空具体时间范围
  if (val) {
    filterState.value.timeRange.customRange = []
    searchInfo.value.createdAtRange = []
  }
  
  const selectedValues = val ? timeRangeOptions.value.map(item => item.value) : []
  updateFilterCondition('timeRange', selectedValues, val)
}

const onTimeRangeChange = () => {
  console.log('=== 时间范围筛选触发 ===')
  
  // 获取选中的时间范围
  const selectedTimeRanges = timeRangeOptions.value.filter(item => item.checked).map(item => item.value)
  console.log('选中的时间范围:', selectedTimeRanges)
  
  // 检查是否需要更新全选状态
  const checkedCount = selectedTimeRanges.length
  const isAllSelected = checkedCount === timeRangeOptions.value.length
  searchInfo.value.timeRangeAll = isAllSelected

  // 根据选中的时间范围设置时间筛选条件
  if (selectedTimeRanges.length === 1) {
    // 如果只选择了一个时间范围，设置对应的时间区间
    const selectedRange = selectedTimeRanges[0]
    setTimeRangeFilter(selectedRange)
  } else if (selectedTimeRanges.length === 0) {
    // 如果没有选择，清空时间范围
    filterState.value.timeRange.customRange = []
    searchInfo.value.createdAtRange = []
  }
  
  // 更新筛选条件并触发筛选
  updateFilterCondition('timeRange', selectedTimeRanges, isAllSelected)
}

// 关键词搜索处理
const onKeywordChange = (value) => {
  console.log('=== 关键词搜索触发 ===', value)
  
  filterState.value.keyword = value
  triggerFilter()
}

// 自定义时间范围变化处理
const onCustomTimeRangeChange = (value) => {
  console.log('=== 自定义时间范围变化 ===', value)
  
  if (value && value.length === 2) {
    filterState.value.timeRange.customRange = value
    
    // 清空预设时间范围的选择
    timeRangeOptions.value.forEach(item => {
      item.checked = false
    })
    searchInfo.value.timeRangeAll = false
    topNavTimeRange.value = ''
  } else {
    filterState.value.timeRange.customRange = []
  }
  
  triggerFilter()
}

// 计算是否有活跃的筛选条件
const hasActiveFilters = computed(() => {
  return filterState.value.deviceStatus.selected.length > 0 ||
         filterState.value.alertType.selected.length > 0 ||
         filterState.value.cameraAddress.selected.length > 0 ||
         filterState.value.timeRange.selected.length > 0 ||
         (filterState.value.timeRange.customRange && filterState.value.timeRange.customRange.length > 0) ||
         (filterState.value.keyword && filterState.value.keyword.trim())
})

// 计算活跃筛选条件的标签
const activeFilterTags = computed(() => {
  const tags = []
  
  // 设备状态标签
  if (filterState.value.deviceStatus.selected.length > 0) {
    const statusLabels = filterState.value.deviceStatus.selected.map(status => 
      getDeviceStatusText(status)
    ).join('、')
    tags.push({
      key: 'deviceStatus',
      label: `设备状态: ${statusLabels}`,
      category: 'deviceStatus'
    })
  }
  
  // 预警类型标签
  if (filterState.value.alertType.selected.length > 0) {
    const typeLabels = filterState.value.alertType.selected.map(type => 
      getAlertTypeLabel(type)
    ).join('、')
    tags.push({
      key: 'alertType',
      label: `预警类型: ${typeLabels}`,
      category: 'alertType'
    })
  }
  
  // 摄像头点位标签
  if (filterState.value.cameraAddress.selected.length > 0) {
    const addressLabels = filterState.value.cameraAddress.selected.join('、')
    tags.push({
      key: 'cameraAddress',
      label: `摄像头点位: ${addressLabels}`,
      category: 'cameraAddress'
    })
  }
  
  // 时间范围标签
  if (filterState.value.timeRange.selected.length > 0) {
    const rangeLabels = filterState.value.timeRange.selected.map(range => {
      const option = timeRangeOptions.value.find(opt => opt.value === range)
      return option ? option.label : range
    }).join('、')
    tags.push({
      key: 'timeRange',
      label: `时间范围: ${rangeLabels}`,
      category: 'timeRange'
    })
  }
  
  // 自定义时间范围标签
  if (filterState.value.timeRange.customRange && filterState.value.timeRange.customRange.length === 2) {
    const [start, end] = filterState.value.timeRange.customRange
    const startStr = formatDate(start).split(' ')[0]
    const endStr = formatDate(end).split(' ')[0]
    tags.push({
      key: 'customTimeRange',
      label: `自定义时间: ${startStr} 至 ${endStr}`,
      category: 'customTimeRange'
    })
  }
  
  // 关键词标签
  if (filterState.value.keyword && filterState.value.keyword.trim()) {
    tags.push({
      key: 'keyword',
      label: `关键词: ${filterState.value.keyword.trim()}`,
      category: 'keyword'
    })
  }
  
  return tags
})

// 移除单个筛选标签
const removeFilterTag = (tag) => {
  console.log('=== 移除筛选标签 ===', tag)
  
  switch (tag.category) {
    case 'deviceStatus':
      deviceStatusOptions.value.forEach(item => {
        item.checked = false
      })
      updateFilterCondition('deviceStatus', [], false)
      break
    case 'alertType':
      alert_typeOptions.value.forEach(item => {
        item.checked = false
      })
      updateFilterCondition('alertType', [], false)
      break
    case 'cameraAddress':
      cameraAddressOptions.value.forEach(item => {
        item.checked = false
      })
      updateFilterCondition('cameraAddress', [], false)
      topNavCameraAddress.value = ''
      break
    case 'timeRange':
      timeRangeOptions.value.forEach(item => {
        item.checked = false
      })
      updateFilterCondition('timeRange', [], false)
      topNavTimeRange.value = ''
      break
    case 'customTimeRange':
      filterState.value.timeRange.customRange = []
      searchInfo.value.createdAtRange = []
      triggerFilter()
      break
    case 'keyword':
      filterState.value.keyword = ''
      searchInfo.value.keyword = ''
      triggerFilter()
      break
  }
}

// 清空所有筛选条件
const clearAllFilters = () => {
  console.log('=== 清空所有筛选条件 ===')
  onReset()
}

// 设置时间范围筛选
const setTimeRangeFilter = (rangeValue) => {
  const now = new Date()
  let startTime, endTime

  switch (rangeValue) {
    case 'today':
      startTime = new Date(now.getFullYear(), now.getMonth(), now.getDate())
      endTime = new Date(now.getFullYear(), now.getMonth(), now.getDate(), 23, 59, 59)
      break
    case 'yesterday':
      const yesterday = new Date(now.getTime() - 24 * 60 * 60 * 1000)
      startTime = new Date(yesterday.getFullYear(), yesterday.getMonth(), yesterday.getDate())
      endTime = new Date(yesterday.getFullYear(), yesterday.getMonth(), yesterday.getDate(), 23, 59, 59)
      break
    case 'last3days':
      startTime = new Date(now.getTime() - 3 * 24 * 60 * 60 * 1000)
      endTime = now
      break
    case 'last7days':
      startTime = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000)
      endTime = now
      break
    case 'last30days':
      startTime = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000)
      endTime = now
      break
    case 'thisMonth':
      startTime = new Date(now.getFullYear(), now.getMonth(), 1)
      endTime = new Date(now.getFullYear(), now.getMonth() + 1, 0, 23, 59, 59)
      break
    case 'lastMonth':
      const lastMonth = new Date(now.getFullYear(), now.getMonth() - 1, 1)
      startTime = lastMonth
      endTime = new Date(now.getFullYear(), now.getMonth(), 0, 23, 59, 59)
      break
    default:
      startTime = null
      endTime = null
  }

  if (startTime && endTime) {
    filterState.value.timeRange.customRange = [startTime, endTime]
    searchInfo.value.createdAtRange = [startTime, endTime]
  }
}

const onDeviceStatusAllChange = (val) => {
  deviceStatusOptions.value.forEach(item => {
    item.checked = val
  })
  
  const selectedValues = val ? deviceStatusOptions.value.map(item => item.value) : []
  updateFilterCondition('deviceStatus', selectedValues, val)
}

const onDeviceStatusChange = () => {
  console.log('=== 设备状态筛选触发 ===')
  
  // 获取选中的设备状态
  const selectedDeviceStatus = deviceStatusOptions.value.filter(item => item.checked).map(item => item.value)
  console.log('选中的设备状态:', selectedDeviceStatus)
  console.log('设备状态选项:', deviceStatusOptions.value.map(item => `${item.label}(${item.value}): ${item.checked}`))
  
  // 检查是否需要更新全选状态
  const checkedCount = selectedDeviceStatus.length
  const isAllSelected = checkedCount === deviceStatusOptions.value.length
  searchInfo.value.deviceStatusAll = isAllSelected
  
  console.log(`设备状态全选状态: ${isAllSelected} (${checkedCount}/${deviceStatusOptions.value.length})`)
  
  // 更新筛选条件并触发筛选
  updateFilterCondition('deviceStatus', selectedDeviceStatus, isAllSelected)
}

// 顶部导航栏下拉列表处理
const onTopNavCameraAddressChange = (value) => {
  console.log('=== 顶部导航栏摄像头点位变化 ===', value)
  
  // 根据选择的摄像头点位进行筛选
  searchInfo.value.cameraAddress = value

  // 同步到左侧筛选栏：如果选择了特定点位，则只选中该点位
  if (value) {
    // 先清空所有选择
    cameraAddressOptions.value.forEach(item => {
      item.checked = false
    })
    // 选中对应的点位
    const selectedItem = cameraAddressOptions.value.find(item => item.value === value)
    if (selectedItem) {
      selectedItem.checked = true
    }
    
    // 更新筛选条件
    updateFilterCondition('cameraAddress', [value], false)
  } else {
    // 如果清空选择，则重置左侧筛选栏
    cameraAddressOptions.value.forEach(item => {
      item.checked = false
    })
    
    // 更新筛选条件
    updateFilterCondition('cameraAddress', [], false)
  }
}

const onTopNavTimeRangeChange = (value) => {
  console.log('=== 顶部导航栏时间范围变化 ===', value)
  
  // 设置时间范围筛选
  if (value) {
    setTimeRangeFilter(value)
  } else {
    filterState.value.timeRange.customRange = []
    searchInfo.value.createdAtRange = []
  }

  // 同步到左侧筛选栏：如果选择了特定时间范围，则只选中该时间范围
  if (value) {
    // 先清空所有选择
    timeRangeOptions.value.forEach(item => {
      item.checked = false
    })
    // 选中对应的时间范围
    const selectedItem = timeRangeOptions.value.find(item => item.value === value)
    if (selectedItem) {
      selectedItem.checked = true
    }
    
    // 更新筛选条件
    updateFilterCondition('timeRange', [value], false)
  } else {
    // 如果清空选择，则重置左侧筛选栏
    timeRangeOptions.value.forEach(item => {
      item.checked = false
    })
    
    // 更新筛选条件
    updateFilterCondition('timeRange', [], false)
  }
}

// 选择模式切换处理
const onAlertTypeSelectModeChange = () => {
  searchInfo.value.alertTypeList = []
  searchInfo.value.alertTypeSingle = ''
}

const onDeviceStatusSelectModeChange = () => {
  searchInfo.value.deviceStatusList = []
  searchInfo.value.deviceStatusSingle = ''
}

// 全选逻辑
watch(() => searchInfo.value.alertTypeList, (val) => {
  if (val.includes('all')) {
    searchInfo.value.alertTypeList = ['all', ...alert_typeOptions.value.map(i => i.value)]
  } else if (val.length === 0) {
    // 保持空
  } else {
    searchInfo.value.alertTypeList = val.filter(v => v !== 'all')
  }
})

// 分页
const page = ref(1)
const total = ref(0)
const pageSize = ref(6)
const tableData = ref([])
const allTableData = ref([]) // 存储所有数据用于滚动显示

// 查询
const getTableData = async() => {
  try {
    // 组装查询参数
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value,
      alertType: alert_typeOptions.value.filter(item => item.checked).map(item => item.value).join(','),
      cameraAddress: cameraAddressOptions.value.filter(item => item.checked).map(item => item.value).join(','),
      eventScope: eventScopeOptions.value.filter(item => item.checked).map(item => item.value).join(','),
      deviceStatus: deviceStatusOptions.value.filter(item => item.checked).map(item => item.value).join(',')
    }

    // 清理空参数
    Object.keys(params).forEach(key => {
      if (params[key] === '' || params[key] === null || params[key] === undefined) {
        delete params[key]
      }
    })

    // 强制使用模拟数据进行演示
    const table = { code: -1 } // 强制进入模拟数据分支
    // const table = await getAlertList(params)
    if (table.code === 0) {
      // 确保数据结构正确，为每个告警项添加默认的设备信息
      const mockImages = [
        'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
        'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
        'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
        'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
        'https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=400&h=300&fit=crop',
        'https://images.unsplash.com/photo-1563013544-824ae1b704d3?w=400&h=300&fit=crop'
      ]

      tableData.value = (table.data.list || []).map((item, index) => ({
        ...item,
        deviceName: item.deviceName || item.cameraAddress || '未知设备',
        deviceStatus: item.deviceStatus || '2', // 默认离线状态
        resolution: item.resolution || '未知',
        streamUrl: item.streamUrl || '',
        // 如果没有图片数据，使用模拟图片
        video: item.video || mockImages[index % mockImages.length],
        // 如果没有预警类型数据，添加模拟数据用于显示测试
        alertType: item.alertType || ['smoking_detection', 'ground_garbage', 'yacht_boat_recognition', 'vehicle_type_recognition', 'non_motor_vehicle_recognition', 'drone_recognition', 'small_boat_raft_recognition', 'person_intrusion'][index % 8]
      })) ;
      total.value = table.data.total || 0
      page.value = table.data.page || 1
      pageSize.value = table.data.pageSize || 6
    } else {
      // 如果API调用失败，使用模拟数据（基于图片上显示的卡片数据）
      let mockData = [
        {
          // alert表字段
          ID: 1,
          video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
          cameraAddress: '杭州市区',
          alertType: 'small_boat_raft_recognition',
          alertTime: '2025-07-17 11:40:00',
          cameraId: 1001,
          createdAt: '2025-07-17 11:40:00',
          updatedAt: '2025-07-17 11:40:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '杭州市区小船监控设备HK-001',
          deviceStatus: '1', // status字段
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.100:1935/live/stream1001',
          groupId: 1,
          manufacturer: '海康威视',
          installDate: '2024-03-15',
          maintenanceDate: '2025-01-10'
        },
        {
          // alert表字段
          ID: 2,
          video: 'https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=400&h=300&fit=crop',
          cameraAddress: '杭州市区',
          alertType: 'small_boat_raft_recognition',
          alertTime: '2025-07-17 10:28:00',
          cameraId: 1002,
          createdAt: '2025-07-17 10:28:00',
          updatedAt: '2025-07-17 10:28:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '杭州市区小船监控设备DH-002',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.101:1935/live/stream1002',
          groupId: 1,
          manufacturer: '大华技术',
          installDate: '2024-03-15',
          maintenanceDate: '2025-01-10'
        },
        {
          // alert表字段
          ID: 3,
          video: 'https://images.unsplash.com/photo-1449824913935-59a10b8d2000?w=400&h=300&fit=crop',
          cameraAddress: '马路路段',
          alertType: 'vehicle_type_recognition',
          alertTime: '2025-07-17 16:27:00',
          cameraId: 1003,
          createdAt: '2025-07-17 16:27:00',
          updatedAt: '2025-07-17 16:27:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '马路路段车辆监控设备UV-003',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.102:1935/live/stream1003',
          groupId: 2,
          manufacturer: '宇视科技',
          installDate: '2024-04-20',
          maintenanceDate: '2025-02-15'
        },
        {
          // alert表字段
          ID: 4,
          video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
          cameraAddress: '十字路',
          alertType: 'drone_recognition',
          alertTime: '2025-07-17 16:20:00',
          cameraId: 1004,
          createdAt: '2025-07-17 16:20:00',
          updatedAt: '2025-07-17 16:20:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '十字路口无人机监控设备HK-004',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.103:1935/live/stream1004',
          groupId: 2,
          manufacturer: '海康威视',
          installDate: '2024-05-10',
          maintenanceDate: '2025-03-05'
        },
        {
          // alert表字段
          ID: 5,
          video: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
          cameraAddress: '村村桥中',
          alertType: 'non_motor_vehicle_recognition',
          alertTime: '2025-07-17 17:32:00',
          cameraId: 1005,
          createdAt: '2025-07-17 17:32:00',
          updatedAt: '2025-07-17 17:32:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '村村桥非机动车监控设备DH-005',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.104:1935/live/stream1005',
          groupId: 3,
          manufacturer: '大华技术',
          installDate: '2024-06-01',
          maintenanceDate: '2025-04-01'
        },
        {
          // alert表字段
          ID: 6,
          video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
          cameraAddress: '怀柔密云',
          alertType: 'yacht_boat_recognition',
          alertTime: '2025-07-17 09:23:00',
          cameraId: 1006,
          createdAt: '2025-07-17 09:23:00',
          updatedAt: '2025-07-17 09:23:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '怀柔密云水域监控设备UV-006',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.105:1935/live/stream1006',
          groupId: 3,
          manufacturer: '宇视科技',
          installDate: '2024-07-15',
          maintenanceDate: '2025-05-10'
        },
        {
          // alert表字段
          ID: 7,
          video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
          cameraAddress: '新济南广场',
          alertType: 'person_intrusion',
          alertTime: '2025-07-17 09:05:00',
          cameraId: 1007,
          createdAt: '2025-07-17 09:05:00',
          updatedAt: '2025-07-17 09:05:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '新济南广场人员监控设备DH-007',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.106:1935/live/stream1007',
          groupId: 4,
          manufacturer: '大华技术',
          installDate: '2024-08-20',
          maintenanceDate: '2025-06-15'
        },
        {
          // alert表字段
          ID: 8,
          video: 'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
          cameraAddress: '村村桥中',
          alertType: 'ground_garbage',
          alertTime: '2025-07-17 11:26:00',
          cameraId: 1008,
          createdAt: '2025-07-17 11:26:00',
          updatedAt: '2025-07-17 11:26:00',
          // monitor_device表字段（通过LEFT JOIN获取）
          deviceName: '村村桥垃圾监控设备HK-008',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.107:1935/live/stream1008',
          groupId: 3,
          manufacturer: '海康威视',
          installDate: '2024-06-01',
          maintenanceDate: '2025-04-01'
        },
        // 添加更多数据以测试分页功能
        {
          ID: 9,
          video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
          cameraAddress: '杭州市区',
          alertType: 'smoking_detection',
          alertTime: '2025-07-17 08:15:00',
          cameraId: 1009,
          createdAt: '2025-07-17 08:15:00',
          updatedAt: '2025-07-17 08:15:00',
          deviceName: '杭州市区吸烟监控设备HK-009',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.108:1935/live/stream1009',
          groupId: 1,
          manufacturer: '海康威视',
          installDate: '2024-03-15',
          maintenanceDate: '2025-01-10'
        },
        {
          ID: 10,
          video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
          cameraAddress: '马路路段',
          alertType: 'non_motor_vehicle_recognition',
          alertTime: '2025-07-17 07:30:00',
          cameraId: 1010,
          createdAt: '2025-07-17 07:30:00',
          updatedAt: '2025-07-17 07:30:00',
          deviceName: '马路路段非机动车监控设备UV-010',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.109:1935/live/stream1010',
          groupId: 2,
          manufacturer: '宇视科技',
          installDate: '2024-04-20',
          maintenanceDate: '2025-02-15'
        },
        {
          ID: 11,
          video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
          cameraAddress: '十字路',
          alertType: 'person_intrusion',
          alertTime: '2025-07-17 06:45:00',
          cameraId: 1011,
          createdAt: '2025-07-17 06:45:00',
          updatedAt: '2025-07-17 06:45:00',
          deviceName: '十字路口人员监控设备HK-011',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.110:1935/live/stream1011',
          groupId: 2,
          manufacturer: '海康威视',
          installDate: '2024-05-10',
          maintenanceDate: '2025-03-05'
        },
        {
          ID: 12,
          video: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
          cameraAddress: '怀柔密云',
          alertType: 'small_boat_raft_recognition',
          alertTime: '2025-07-17 05:20:00',
          cameraId: 1012,
          createdAt: '2025-07-17 05:20:00',
          updatedAt: '2025-07-17 05:20:00',
          deviceName: '怀柔密云小船监控设备UV-012',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.111:1935/live/stream1012',
          groupId: 4,
          manufacturer: '宇视科技',
          installDate: '2024-07-15',
          maintenanceDate: '2025-05-10'
        },
        {
          ID: 13,
          video: 'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
          cameraAddress: '新济南广场',
          alertType: 'ground_garbage',
          alertTime: '2025-07-17 04:10:00',
          cameraId: 1013,
          createdAt: '2025-07-17 04:10:00',
          updatedAt: '2025-07-17 04:10:00',
          deviceName: '新济南广场垃圾监控设备DH-013',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.112:1935/live/stream1013',
          groupId: 4,
          manufacturer: '大华技术',
          installDate: '2024-08-20',
          maintenanceDate: '2025-06-15'
        },
        {
          ID: 14,
          video: 'https://images.unsplash.com/photo-1449824913935-59a10b8d2000?w=400&h=300&fit=crop',
          cameraAddress: '村村桥中',
          alertType: 'vehicle_type_recognition',
          alertTime: '2025-07-17 03:25:00',
          cameraId: 1014,
          createdAt: '2025-07-17 03:25:00',
          updatedAt: '2025-07-17 03:25:00',
          deviceName: '村村桥车辆监控设备HK-014',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.113:1935/live/stream1014',
          groupId: 3,
          manufacturer: '海康威视',
          installDate: '2024-06-01',
          maintenanceDate: '2025-04-01'
        },
        // 添加更多数据确保分页功能可见
        {
          ID: 15,
          video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
          cameraAddress: '杭州市区',
          alertType: 'yacht_boat_recognition',
          alertTime: '2025-07-17 02:15:00',
          cameraId: 1015,
          createdAt: '2025-07-17 02:15:00',
          updatedAt: '2025-07-17 02:15:00',
          deviceName: '杭州市区游艇监控设备HK-015',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.114:1935/live/stream1015',
          groupId: 1,
          manufacturer: '海康威视',
          installDate: '2024-03-15',
          maintenanceDate: '2025-01-10'
        },
        {
          ID: 16,
          video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
          cameraAddress: '马路路段',
          alertType: 'smoking_detection',
          alertTime: '2025-07-17 01:30:00',
          cameraId: 1016,
          createdAt: '2025-07-17 01:30:00',
          updatedAt: '2025-07-17 01:30:00',
          deviceName: '马路路段吸烟监控设备UV-016',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.115:1935/live/stream1016',
          groupId: 2,
          manufacturer: '宇视科技',
          installDate: '2024-04-20',
          maintenanceDate: '2025-02-15'
        },
        {
          ID: 17,
          video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
          cameraAddress: '十字路',
          alertType: 'ground_garbage',
          alertTime: '2025-07-17 00:45:00',
          cameraId: 1017,
          createdAt: '2025-07-17 00:45:00',
          updatedAt: '2025-07-17 00:45:00',
          deviceName: '十字路口垃圾监控设备HK-017',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.116:1935/live/stream1017',
          groupId: 2,
          manufacturer: '海康威视',
          installDate: '2024-05-10',
          maintenanceDate: '2025-03-05'
        },
        {
          ID: 18,
          video: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
          cameraAddress: '怀柔密云',
          alertType: 'drone_recognition',
          alertTime: '2025-07-16 23:20:00',
          cameraId: 1018,
          createdAt: '2025-07-16 23:20:00',
          updatedAt: '2025-07-16 23:20:00',
          deviceName: '怀柔密云无人机监控设备UV-018',
          deviceStatus: '1',
          resolution: '1920x1080',
          streamUrl: 'rtmp://192.168.1.117:1935/live/stream1018',
          groupId: 4,
          manufacturer: '宇视科技',
          installDate: '2024-07-15',
          maintenanceDate: '2025-05-10'
        }
      ]

      // 根据选中的预警类型进行筛选
      const selectedAlertTypes = alert_typeOptions.value.filter(item => item.checked).map(item => item.value)
      console.log('=== 筛选逻辑开始 ===')
      console.log('选中的预警类型值:', selectedAlertTypes)
      console.log('筛选前的数据数量:', mockData.length)
      console.log('所有数据的预警类型:', mockData.map(item => `ID${item.ID}:${item.alertType}`))

      if (selectedAlertTypes.length > 0) {
        console.log('开始筛选数据...')
        const originalData = [...mockData] // 保存原始数据
        mockData = mockData.filter(item => {
          const isMatch = selectedAlertTypes.includes(item.alertType)
          console.log(`ID ${item.ID}: ${item.alertType} -> ${isMatch ? '✓匹配' : '✗不匹配'}`)
          return isMatch
        })
        console.log('筛选后的数据数量:', mockData.length)
        console.log('筛选后的数据:', mockData.map(item => `ID${item.ID}:${item.alertType}:${item.cameraAddress}`))

        // 如果筛选后没有数据，显示提示信息
        if (mockData.length === 0) {
          console.log('⚠️ 筛选后没有匹配的数据')
        }
      } else {
        console.log('没有选中任何预警类型，显示所有数据')
        console.log('显示所有数据，数量:', mockData.length)
      }

      // 根据选中的设备状态进行筛选（只有选中时才筛选）
      const selectedDeviceStatus = deviceStatusOptions.value.filter(item => item.checked).map(item => item.value)
      if (selectedDeviceStatus.length > 0) {
        mockData = mockData.filter(item => selectedDeviceStatus.includes(item.deviceStatus))
      }

      // 根据选中的摄像头点位进行筛选（只有选中时才筛选）
      const selectedCameraAddress = cameraAddressOptions.value.filter(item => item.checked).map(item => item.value)
      if (selectedCameraAddress.length > 0) {
        mockData = mockData.filter(item => selectedCameraAddress.includes(item.cameraAddress))
      }

      // 根据设备名称搜索进行筛选
      if (searchInfo.value.deviceName && searchInfo.value.deviceName.trim() !== '') {
        const keyword = searchInfo.value.deviceName.toLowerCase()
        mockData = mockData.filter(item =>
            item.deviceName.toLowerCase().includes(keyword) ||
            item.cameraAddress.toLowerCase().includes(keyword)
        )
      }

      // 根据关键词搜索进行筛选
      if (searchInfo.value.keyword && searchInfo.value.keyword.trim() !== '') {
        const keyword = searchInfo.value.keyword.toLowerCase()
        mockData = mockData.filter(item =>
            item.cameraAddress.toLowerCase().includes(keyword) ||
            item.deviceName.toLowerCase().includes(keyword) ||
            getAlertTypeLabel(item.alertType).toLowerCase().includes(keyword)
        )
      }

      // 设置总数
      total.value = mockData.length

      // 实现分页逻辑，每页显示6张卡片
      const startIndex = (page.value - 1) * pageSize.value
      const endIndex = startIndex + pageSize.value
      tableData.value = mockData.slice(startIndex, endIndex)
      allTableData.value = mockData // 保持兼容性

      // 调试信息
      console.log('分页信息:', {
        total: total.value,
        page: page.value,
        pageSize: pageSize.value,
        startIndex,
        endIndex,
        displayedData: tableData.value.length,
        totalPages: Math.ceil(total.value / pageSize.value)
      })
    }
  } catch (error) {
    console.error('获取告警数据失败:', error)
    // 显示错误提示
    ElMessage.error('获取告警数据失败，请稍后重试')
  }
}

const onSubmit = () => {
  console.log('=== 手动搜索触发 ===')
  page.value = 1
  triggerFilter()
}

const onReset = () => {
  console.log('=== 重置筛选条件 ===')
  
  // 重置筛选状态
  filterState.value = {
    deviceStatus: {
      all: false,
      selected: []
    },
    alertType: {
      all: false,
      selected: []
    },
    cameraAddress: {
      all: false,
      selected: []
    },
    timeRange: {
      all: false,
      selected: [],
      customRange: []
    },
    keyword: ''
  }
  
  // 重置旧的搜索信息结构
  searchInfo.value = {
    alertTypeAll: false,
    alertTypeList: [],
    alertTypeSingle: '',
    deviceStatusAll: false,
    deviceStatusList: [],
    deviceStatusSingle: '',
    cameraAddressAll: false,
    timeRangeAll: false,
    cameraAddress: '',
    deviceName: '',
    createdAtRange: [],
    keyword: ''
  }
  
  // 重置所有复选框状态
  alert_typeOptions.value.forEach(item => {
    item.checked = false
  })
  deviceStatusOptions.value.forEach(item => {
    item.checked = false
  })
  cameraAddressOptions.value.forEach(item => {
    item.checked = false
  })
  timeRangeOptions.value.forEach(item => {
    item.checked = false
  })
  
  // 重置顶部导航栏选择
  topNavCameraAddress.value = ''
  topNavTimeRange.value = ''
  
  // 重置分页并显示所有数据
  page.value = 1
  filteredResults.value = [...allMockData.value]
  total.value = allMockData.value.length
  
  // 应用分页显示
  applyPagination()
  
  console.log('✅ 重置完成，显示所有数据，数量:', total.value)
}
const handleSizeChange = (val) => {
  pageSize.value = val
  page.value = 1 // 重置到第一页
  applyPagination() // 重新应用分页
}

const handleCurrentChange = (val) => {
  page.value = val
  applyPagination() // 重新应用分页
}

// 详情弹窗
const detailForm = ref({})
const detailShow = ref(false)
const getDetails = async (row) => {
  try {
    const res = await findAlert({ ID: row.ID })
    if (res.code === 0) {
      // 确保API返回的数据包含完整信息，如果没有则使用卡片数据补充
      detailForm.value = {
        ...res.data,
        // 确保所有字段数据一致，优先使用卡片中的数据保持一致性
        deviceName: row.deviceName || res.data.deviceName,
        deviceStatus: row.deviceStatus || res.data.deviceStatus,
        cameraAddress: row.cameraAddress || res.data.cameraAddress,
        alertType: row.alertType || res.data.alertType,
        alertTime: row.alertTime || res.data.alertTime,
        resolution: row.resolution || res.data.resolution,
        streamUrl: row.streamUrl || res.data.streamUrl,
        video: row.video || res.data.video,
        groupId: row.groupId || res.data.groupId
      }
    } else {
      // 如果API调用失败，直接使用卡片数据
      detailForm.value = { ...row }
    }
  } catch (error) {
    console.error('获取详情失败:', error)
    // 如果API调用失败，直接使用卡片数据
    detailForm.value = { ...row }
  }
  detailShow.value = true
}

const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}

// 获取预警类型标签
const getAlertTypeLabel = (alertType) => {
  if (!alertType) return '未知类型'

  const alertTypeMapping = {
    'smoking_detection': '吸烟检测',
    'ground_garbage': '地面垃圾',
    'yacht_boat_recognition': '游艇小艇识别',
    'vehicle_type_recognition': '车辆类型识别',
    'non_motor_vehicle_recognition': '非机动车识别',
    'drone_recognition': '无人机识别',
    'small_boat_raft_recognition': '小船皮筏识别',
    'person_intrusion': '人员入侵'
  }

  return alertTypeMapping[alertType] || '未知类型'
}
</script>

<style scoped>
.alert-page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #000000;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  background: #000000;
  color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-item {
  font-size: 18px;
  font-weight: 600;
  color: white;
}

.nav-item.active {
  color: #ffd700;
}

.nav-user {
  color: white;
}

.main-content {
  flex: 1;
  display: flex;
  gap: 24px;
  padding: 24px;
  overflow: hidden;
  align-items: flex-start;
}

.filter-bar {
  width: 320px;
  height: calc(100vh - 140px);
  overflow-y: auto;
  flex-shrink: 0;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.filter-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 16px;
  color: #333;
  border-bottom: 2px solid #e1e5e9;
  padding-bottom: 8px;
}

.filter-group-title {
  font-size: 14px;
  font-weight: 600;
  margin: 16px 0 8px 0;
  color: #555;
}

.filter-form {
  padding: 0;
}

.filter-checkbox-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.filter-checkbox-list .el-checkbox {
  margin-right: 0;
  margin-bottom: 4px;
}

.status-dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 6px;
}

.card-list-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  height: calc(100vh - 140px);
}

.card-list-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
  padding: 20px 24px;
  background: #2d2d2d;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  flex-shrink: 0;
}

.statistics-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.total-count {
  font-size: 16px;
  font-weight: 600;
  color: #e1e5e9;
}

.filter-summary {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  margin-top: 4px;
}

.filter-label {
  font-size: 12px;
  color: #cccccc;
  white-space: nowrap;
}

.card-list-header span {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.search-input {
  max-width: 300px;
}

/* 筛选标签样式 */
.filter-summary :deep(.el-tag) {
  background-color: rgba(64, 158, 255, 0.1) !important;
  border-color: rgba(64, 158, 255, 0.3) !important;
  color: #409eff !important;
  font-size: 12px;
}

.filter-summary :deep(.el-tag .el-tag__close) {
  color: #409eff !important;
}

.filter-summary :deep(.el-tag .el-tag__close:hover) {
  background-color: rgba(64, 158, 255, 0.2) !important;
}

.filter-summary :deep(.el-button--text) {
  padding: 0 !important;
  font-size: 12px !important;
  height: auto !important;
  line-height: 1 !important;
  color: #409eff !important;
}

/* 空状态样式 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  text-align: center;
  color: #999;
  min-height: 400px;
}

.empty-icon {
  margin-bottom: 24px;
  opacity: 0.6;
}

.empty-text h3 {
  font-size: 18px;
  color: #e1e5e9;
  margin: 0 0 12px 0;
  font-weight: 500;
}

.empty-text p {
  font-size: 14px;
  color: #999;
  margin: 0;
  line-height: 1.5;
  max-width: 400px;
}

.empty-actions {
  margin-top: 24px;
}

/* 滚动容器样式 */
.card-scroll-container {
  flex: 1;
  overflow-y: auto;
  padding-right: 8px;
}

.card-scroll-container::-webkit-scrollbar {
  width: 8px;
}

.card-scroll-container::-webkit-scrollbar-track {
  background: #1a1a1a;
  border-radius: 4px;
}

.card-scroll-container::-webkit-scrollbar-thumb {
  background: #555;
  border-radius: 4px;
}

.card-scroll-container::-webkit-scrollbar-thumb:hover {
  background: #666;
}

/* 卡片网格布局 */
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
  padding: 4px;
}

.card-item {
  width: 100%;
}

.event-card {
  height: 320px;
  transition: all 0.3s ease;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #404040;
  background: #2d2d2d !important;
  display: flex;
  flex-direction: column;
}

.event-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.25);
  border-color: #555;
}

.card-img-box {
  position: relative;
  height: 140px;
  overflow: hidden;
  cursor: pointer;
  flex-shrink: 0;
}

.card-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.card-img-box:hover .card-img {
  transform: scale(1.05);
}

.img-red-rect {
  position: absolute;
  top: 10px;
  left: 10px;
  width: 60px;
  height: 30px;
  background: transparent;
  border: 2px solid #ff0000;
  border-radius: 4px;
}

.card-info {
  padding: 12px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.card-info-row {
  font-size: 12px;
  line-height: 1.3;
  color: #b0b0b0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.card-info-row b {
  color: #e1e5e9;
  font-weight: 600;
}

.device-status {
  display: inline-flex;
  align-items: center;
  font-weight: 600;
  font-size: 12px;
}

.card-actions {
  padding: 0 12px 12px;
  text-align: right;
  flex-shrink: 0;
}

.gva-pagination {
  margin-top: 24px;
  display: flex;
  justify-content: center;
  padding: 20px 24px;
  background: #2d2d2d;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  flex-shrink: 0;
}

/* 强制黑色主题 */
.filter-bar,
.card-list-header,
.gva-pagination {
  background: #2d2d2d !important;
  color: #e1e5e9 !important;
}

.event-card {
  background: #2d2d2d !important;
  border-color: #404040 !important;
}

.filter-title,
.filter-group-title,
.card-list-header span {
  color: #e1e5e9 !important;
}

.card-info-row {
  color: #b0b0b0 !important;
}

.card-info-row b {
  color: #e1e5e9 !important;
}

/* Element Plus 组件深色主题覆盖 */
.filter-bar :deep(.el-card__body) {
  background: #2d2d2d !important;
  color: #e1e5e9 !important;
}

.filter-bar :deep(.el-checkbox__label) {
  color: #e1e5e9 !important;
}

.filter-bar :deep(.el-input__inner) {
  background: #404040 !important;
  border-color: #666 !important;
  color: #e1e5e9 !important;
}

.filter-bar :deep(.el-button) {
  background: #404040 !important;
  border-color: #666 !important;
  color: #e1e5e9 !important;
}

.card-list-header :deep(.el-input__inner) {
  background: #404040 !important;
  border-color: #666 !important;
  color: #e1e5e9 !important;
}

/* 搜索框旁按钮的白色字体 */
.card-list-header :deep(.el-button) {
  background: #404040 !important;
  border-color: #666 !important;
  color: #ffffff !important;
}

.card-list-header :deep(.el-button:hover) {
  background: #555 !important;
  border-color: #777 !important;
  color: #ffffff !important;
}

.event-card :deep(.el-card__body) {
  background: #2d2d2d !important;
  color: #e1e5e9 !important;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .main-content {
    flex-direction: column;
  }

  .filter-bar {
    width: 100%;
    max-height: none;
  }

  .card-col {
    :span="12";
  }
}

@media (max-width: 768px) {
  .card-col {
    :span="24";
  }

  .nav-menu {
    flex-direction: column;
    gap: 8px;
  }

  .card-list-header {
    flex-direction: column;
    gap: 12px;
    align-items: stretch;
  }

  .search-input {
    max-width: none;
  }
}

/* Element Plus 组件样式覆盖 */
.gva-pagination :deep(.el-pagination) {
  justify-content: center;
}

.gva-pagination :deep(.el-pagination .el-select .el-input__inner) {
  background: #ffffff;
  border-color: #d1d5db;
  color: #000000;
}

.gva-pagination :deep(.el-pagination .el-input__inner) {
  background: #ffffff;
  border-color: #d1d5db;
  color: #000000;
}
</style>