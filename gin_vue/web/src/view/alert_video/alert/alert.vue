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
          <div class="statistics-info">
            <span class="total-count">共有 {{ total }} 条警告事件</span>
            <div class="status-statistics">
              <span class="status-item">
                <span class="status-dot" style="background-color: #52c41a;"></span>
                在线设备: {{ onlineDeviceCount }}
              </span>
              <span class="status-item">
                <span class="status-dot" style="background-color: #ff4d4f;"></span>
                离线设备: {{ offlineDeviceCount }}
              </span>
              <span class="status-item">
                <span class="status-dot" style="background-color: #fa8c16;"></span>
                故障设备: {{ faultDeviceCount }}
              </span>
            </div>
          </div>
          <el-input
              v-model="searchInfo.keyword"
              placeholder="摄像头点位/报警类型/设备名称"
              size="small"
              class="search-input"
              @keyup.enter="onSubmit"
              style="width: 260px; margin-left: 20px;"
          >
            <template #append>
              <el-button icon="el-icon-search" @click="onSubmit" />
            </template>
          </el-input>
        </div>
        <div class="card-scroll-container">
          <div class="card-grid">
            <div v-for="item in tableData" :key="item.ID" class="card-item">
              <el-card shadow="hover" class="event-card">
                <div class="card-img-box" @click="getDetails(item)">
                  <img :src="item.video || '/src/assets/default-event.jpg'" class="card-img" />
                  <div class="img-red-rect"></div>
                </div>
                <div class="card-info">
                  <!-- 设备状态行 -->
                  <div class="card-info-row device-status-row">
                    <span class="device-status" :style="{ color: getDeviceStatusColor(item.deviceStatus) }">
                      <span class="status-dot" :style="{ backgroundColor: getDeviceStatusColor(item.deviceStatus) }"></span>
                      {{ getDeviceStatusText(item.deviceStatus) }}
                    </span>
                    <span class="device-group">分组{{ item.groupId || 'N/A' }}</span>
                  </div>

                  <!-- 设备基本信息 -->
                  <div class="card-info-section">
                    <div class="section-title">设备信息</div>
                    <div class="card-info-row"><b>设备名称：</b>{{ item.deviceName || '未知设备' }}</div>
                    <div class="card-info-row"><b>设备厂商：</b>{{ item.manufacturer || '未知' }}</div>
                    <div class="card-info-row"><b>设备型号：</b>{{ item.deviceModel || '未知' }}</div>
                  </div>

                  <!-- 网络配置信息 -->
                  <div class="card-info-section">
                    <div class="section-title">网络配置</div>
                    <div class="card-info-row"><b>IP地址：</b>{{ item.ipAddress || '未配置' }}</div>
                    <div class="card-info-row"><b>分辨率：</b>{{ item.resolution || '未知' }}</div>
                    <div class="card-info-row"><b>视频流：</b>{{ item.streamUrl ? '已配置' : '未配置' }}</div>
                  </div>

                  <!-- 告警信息 -->
                  <div class="card-info-section">
                    <div class="section-title">告警信息</div>
                    <div class="card-info-row"><b>摄像头点位：</b>{{ item.cameraAddress || '未知点位' }}</div>
                    <div class="card-info-row"><b>预警类型：</b>{{ getAlertTypeLabel(item.alertType) }}</div>
                    <div class="card-info-row"><b>预警时间：</b>{{ formatDate(item.alertTime) }}</div>
                  </div>

                  <!-- 维护信息 -->
                  <div class="card-info-section">
                    <div class="section-title">维护信息</div>
                    <div class="card-info-row"><b>安装日期：</b>{{ item.installDate || '未知' }}</div>
                    <div class="card-info-row"><b>维护日期：</b>{{ item.maintenanceDate || '未知' }}</div>
                  </div>
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
        <el-descriptions-item label="告警ID">{{ detailForm.ID }}</el-descriptions-item>
        <el-descriptions-item label="设备状态">
          <span :style="{ color: getDeviceStatusColor(detailForm.deviceStatus) }">
            <span class="status-dot" :style="{ backgroundColor: getDeviceStatusColor(detailForm.deviceStatus) }"></span>
            {{ getDeviceStatusText(detailForm.deviceStatus) }}
          </span>
        </el-descriptions-item>
        <el-descriptions-item label="设备名称">{{ detailForm.deviceName || '未知设备' }}</el-descriptions-item>
        <el-descriptions-item label="摄像头点位">{{ detailForm.cameraAddress || '未知点位' }}</el-descriptions-item>
        <el-descriptions-item label="预警类型">{{ getAlertTypeLabel(detailForm.alertType) }}</el-descriptions-item>
        <el-descriptions-item label="预警时间">{{ formatDate(detailForm.alertTime) }}</el-descriptions-item>
        <el-descriptions-item label="设备厂商">{{ detailForm.manufacturer || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="设备型号">{{ detailForm.deviceModel || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="分辨率">{{ detailForm.resolution || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="IP地址">{{ detailForm.ipAddress || '未配置' }}</el-descriptions-item>
        <el-descriptions-item label="视频流地址">{{ detailForm.streamUrl || '未配置' }}</el-descriptions-item>
        <el-descriptions-item label="设备分组ID">{{ detailForm.groupId || '未分组' }}</el-descriptions-item>
        <el-descriptions-item label="安装日期">{{ detailForm.installDate || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="维护日期">{{ detailForm.maintenanceDate || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="摄像头ID">{{ detailForm.cameraId || '未知' }}</el-descriptions-item>
        <el-descriptions-item label="监控视频">
          <img :src="detailForm.video" style="max-width: 100%; height: auto;" />
        </el-descriptions-item>
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
import { ref, watch, computed } from 'vue'
import { useAppStore } from "@/pinia"
import { ElMessage } from 'element-plus'

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
      { value: 'vendor_recognition', label: '游摊小贩识别', checked: false },
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

  const initialData = [
    {
      // alert表字段
      ID: 1,
      video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
      cameraAddress: '1228106',
      alertType: 'small_boat_raft_recognition',
      alertTime: '2025-07-17 11:40:00',
      cameraId: 1001,
      // monitor_device表字段（基于真实数据）
      deviceName: '1228106',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.100:1935/live/stream1001',
      groupId: 1,
      manufacturer: '海康威视',
      installDate: '2024-03-15',
      maintenanceDate: '2025-01-10',
      deviceModel: 'DS-2CD2T47G1-L',
      ipAddress: '192.168.1.100'
    },
    {
      // alert表字段
      ID: 2,
      video: 'https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=400&h=300&fit=crop',
      cameraAddress: '1',
      alertType: 'small_boat_raft_recognition',
      alertTime: '2025-07-17 10:28:00',
      cameraId: 1002,
      // monitor_device表字段（基于真实数据）
      deviceName: '1',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.101:1935/live/stream1002',
      groupId: 1,
      manufacturer: '大华技术',
      installDate: '2024-03-15',
      maintenanceDate: '2025-01-10',
      deviceModel: 'DH-IPC-HFW4433M-I2',
      ipAddress: '192.168.1.101'
    },
    {
      // alert表字段
      ID: 3,
      video: 'https://images.unsplash.com/photo-1449824913935-59a10b8d2000?w=400&h=300&fit=crop',
      cameraAddress: '2',
      alertType: 'vehicle_type_recognition',
      alertTime: '2025-07-17 16:27:00',
      cameraId: 1003,
      // monitor_device表字段（基于真实数据）
      deviceName: '2',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.102:1935/live/stream1003',
      groupId: 2,
      manufacturer: '宇视科技',
      installDate: '2024-04-20',
      maintenanceDate: '2025-02-15',
      deviceModel: 'IPC6322LR-X',
      ipAddress: '192.168.1.102'
    },
    {
      // alert表字段
      ID: 4,
      video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
      cameraAddress: 'seventeen',
      alertType: 'drone_recognition',
      alertTime: '2025-07-17 16:20:00',
      cameraId: 1004,
      // monitor_device表字段（基于真实数据）
      deviceName: 'seventeen',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.103:1935/live/stream1004',
      groupId: 2,
      manufacturer: '海康威视',
      installDate: '2024-05-10',
      maintenanceDate: '2025-03-05',
      deviceModel: 'DS-2CD2T85G1-I8',
      ipAddress: '192.168.1.103'
    },
    {
      // alert表字段
      ID: 5,
      video: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=400&h=300&fit=crop',
      cameraAddress: 's',
      alertType: 'non_motor_vehicle_recognition',
      alertTime: '2025-07-17 17:32:00',
      cameraId: 1005,
      // monitor_device表字段（基于真实数据）
      deviceName: 's',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.104:1935/live/stream1005',
      groupId: 3,
      manufacturer: '大华技术',
      installDate: '2024-06-01',
      maintenanceDate: '2025-04-01',
      deviceModel: 'DH-IPC-HFW4831E-SE',
      ipAddress: '192.168.1.104'
    },
    {
      // alert表字段
      ID: 6,
      video: 'https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop',
      cameraAddress: '怀柔密云',
      alertType: 'yacht_boat_recognition',
      alertTime: '2025-07-17 09:23:00',
      cameraId: 1006,
      // monitor_device表字段（基于真实数据）
      deviceName: '怀柔密云水域监控设备UV-006',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.105:1935/live/stream1006',
      groupId: 4,
      manufacturer: '宇视科技',
      installDate: '2024-07-15',
      maintenanceDate: '2025-05-10',
      deviceModel: 'IPC6618SR-X',
      ipAddress: '192.168.1.105'
    },
    {
      // alert表字段
      ID: 7,
      video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
      cameraAddress: '新济南广场',
      alertType: 'person_intrusion',
      alertTime: '2025-07-17 09:05:00',
      cameraId: 1007,
      // monitor_device表字段
      deviceName: '新济南广场人员监控设备DH-007',
      deviceStatus: '2',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.106:1935/live/stream1007',
      groupId: 4,
      manufacturer: '大华技术',
      installDate: '2024-08-20',
      maintenanceDate: '2025-06-15',
      deviceModel: 'DH-IPC-HFW5831E-ZE',
      ipAddress: '192.168.1.106'
    },
    {
      // alert表字段
      ID: 8,
      video: 'https://images.unsplash.com/photo-1573164713714-d95e436ab8d6?w=400&h=300&fit=crop',
      cameraAddress: '村村桥中',
      alertType: 'ground_garbage',
      alertTime: '2025-07-17 11:26:00',
      cameraId: 1008,
      // monitor_device表字段
      deviceName: '村村桥垃圾监控设备HK-008',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.107:1935/live/stream1008',
      groupId: 3,
      manufacturer: '海康威视',
      installDate: '2024-06-01',
      maintenanceDate: '2025-04-01',
      deviceModel: 'DS-2CD2T47G2-L',
      ipAddress: '192.168.1.107'
    },
    {
      // alert表字段
      ID: 9,
      video: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop',
      cameraAddress: '杭州市区',
      alertType: 'smoking_detection',
      alertTime: '2025-07-17 08:15:00',
      cameraId: 1009,
      // monitor_device表字段
      deviceName: '杭州市区吸烟监控设备HK-009',
      deviceStatus: '3',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.108:1935/live/stream1009',
      groupId: 1,
      manufacturer: '海康威视',
      installDate: '2024-03-15',
      maintenanceDate: '2025-01-10',
      deviceModel: 'DS-2CD2T47G1-L',
      ipAddress: '192.168.1.108'
    },
    {
      // alert表字段
      ID: 10,
      video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
      cameraAddress: '十字路',
      alertType: 'vendor_recognition',
      alertTime: '2025-07-17 14:30:00',
      cameraId: 1010,
      // monitor_device表字段
      deviceName: '十字路口小贩监控设备UV-010',
      deviceStatus: '2',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.109:1935/live/stream1010',
      groupId: 2,
      manufacturer: '宇视科技',
      installDate: '2024-05-10',
      maintenanceDate: '2025-03-05',
      deviceModel: 'IPC6322LR-X',
      ipAddress: '192.168.1.109'
    }
  ]

  // 直接设置数据
  tableData.value = initialData
  allTableData.value = initialData // 保存完整数据
  total.value = initialData.length
  page.value = 1

  console.log('✅ 初始化完成，数据数量:', tableData.value.length)
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

// 查询条件
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

// 全选逻辑处理
const onAlertTypeAllChange = (val) => {
  alert_typeOptions.value.forEach(item => {
    item.checked = val
  })
  // 自动触发查询
  onSubmit()
}

const onAlertTypeChange = () => {
  console.log('=== 预警类型筛选触发 ===')
  const selectedAlertTypes = alert_typeOptions.value.filter(item => item.checked).map(item => item.value)
  console.log('选中的预警类型:', selectedAlertTypes)
  console.log('全部数据数量:', allTableData.value.length)

  if (selectedAlertTypes.length > 0) {
    tableData.value = allTableData.value.filter(item => selectedAlertTypes.includes(item.alertType))
    console.log('筛选后数据数量:', tableData.value.length)
    console.log('筛选后的数据:', tableData.value.map(item => `${item.ID}:${item.alertType}`))
  } else {
    tableData.value = allTableData.value
    console.log('显示所有数据，数量:', tableData.value.length)
  }
  total.value = tableData.value.length
  page.value = 1
  console.log('✅ 筛选完成，当前显示数量:', total.value)
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
      video: 'https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=400&h=300&fit=crop',
      cameraAddress: '十字路',
      alertType: 'vendor_recognition',
      alertTime: '2025-07-17 14:30:00',
      deviceName: '十字路口小贩监控设备',
      deviceStatus: '1',
      resolution: '1920x1080',
      streamUrl: 'rtmp://192.168.1.109:1935/live/stream1010'
    }
  ]

  // 根据选中的预警类型筛选数据
  let filteredData = allMockData
  if (selectedAlertTypes.length > 0) {
    filteredData = allMockData.filter(item => selectedAlertTypes.includes(item.alertType))
    console.log('筛选后的数据:', filteredData)
    console.log('筛选后的数据数量:', filteredData.length)
  } else {
    console.log('没有选中预警类型，显示所有数据')
  }

  // 直接设置数据
  tableData.value = filteredData
  total.value = filteredData.length
  page.value = 1

  console.log('✅ 设置完成 - tableData数量:', tableData.value.length)
  console.log('✅ 设置完成 - total:', total.value)
  console.log('✅ 当前显示的数据:', tableData.value.map(item => `ID${item.ID}:${item.alertType}`))
}

const onCameraAddressAllChange = (val) => {
  cameraAddressOptions.value.forEach(item => {
    item.checked = val
  })
  // 自动触发查询
  onSubmit()
}

const onCameraAddressChange = () => {
  const checkedCount = cameraAddressOptions.value.filter(item => item.checked).length
  searchInfo.value.cameraAddressAll = checkedCount === cameraAddressOptions.value.length
  // 自动触发查询
  onSubmit()
}

const onTimeRangeAllChange = (val) => {
  timeRangeOptions.value.forEach(item => {
    item.checked = val
  })
  // 如果全选，则清空具体时间范围
  if (val) {
    searchInfo.value.createdAtRange = []
  }
  // 自动触发查询
  onSubmit()
}

const onTimeRangeChange = () => {
  const checkedCount = timeRangeOptions.value.filter(item => item.checked).length
  searchInfo.value.timeRangeAll = checkedCount === timeRangeOptions.value.length

  // 根据选中的时间范围设置时间筛选条件
  const checkedRanges = timeRangeOptions.value.filter(item => item.checked)
  if (checkedRanges.length === 1) {
    // 如果只选择了一个时间范围，设置对应的时间区间
    const selectedRange = checkedRanges[0].value
    onTopNavTimeRangeChange(selectedRange)
  } else {
    // 如果选择了多个或没有选择，清空时间范围
    searchInfo.value.createdAtRange = []
    onSubmit()
  }
}

const onDeviceStatusAllChange = (val) => {
  deviceStatusOptions.value.forEach(item => {
    item.checked = val
  })
  // 自动触发查询
  onSubmit()
}

const onDeviceStatusChange = () => {
  const checkedCount = deviceStatusOptions.value.filter(item => item.checked).length
  searchInfo.value.deviceStatusAll = checkedCount === deviceStatusOptions.value.length
  // 自动触发查询
  onSubmit()
}

// 顶部导航栏下拉列表处理
const onTopNavCameraAddressChange = (value) => {
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
    searchInfo.value.cameraAddressAll = false
  } else {
    // 如果清空选择，则重置左侧筛选栏
    cameraAddressOptions.value.forEach(item => {
      item.checked = false
    })
    searchInfo.value.cameraAddressAll = false
  }

  onSubmit()
}

const onTopNavTimeRangeChange = (value) => {
  // 根据选择的时间范围设置时间筛选条件
  const now = new Date()
  let startTime, endTime

  switch (value) {
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
    searchInfo.value.createdAtRange = [startTime, endTime]
  } else {
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
    searchInfo.value.timeRangeAll = false
  } else {
    // 如果清空选择，则重置左侧筛选栏
    timeRangeOptions.value.forEach(item => {
      item.checked = false
    })
    searchInfo.value.timeRangeAll = false
  }

  onSubmit()
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

// 实时统计计算
const onlineDeviceCount = computed(() => {
  return tableData.value.filter(item => item.deviceStatus === '1').length
})

const offlineDeviceCount = computed(() => {
  return tableData.value.filter(item => item.deviceStatus === '2').length
})

const faultDeviceCount = computed(() => {
  return tableData.value.filter(item => item.deviceStatus === '3').length
})
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
  page.value = 1
  getTableData()
}
const onReset = () => {
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
  page.value = 1
  getTableData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
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
  align-items: center;
  margin-bottom: 24px;
  padding: 20px 24px;
  background: #2d2d2d;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  flex-shrink: 0;
}

.card-list-header span {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.search-input {
  max-width: 300px;
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

.alert-page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #1a1a1a;
  color: #ffffff;
}

.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  background-color: #2a2a2a;
  border-bottom: 1px solid #3a3a3a;
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 16px;
}

.nav-item {
  font-size: 16px;
  font-weight: 500;
  color: #ffffff;
}

.nav-item.active {
  color: #409eff;
}

.main-content {
  flex: 1;
  display: flex;
  gap: 20px;
  padding: 20px;
  overflow: hidden;
}

.filter-bar {
  width: 280px;
  background-color: #2a2a2a !important;
  border: 1px solid #3a3a3a !important;
}

.filter-title {
  font-size: 16px;
  font-weight: 600;
  color: #ffffff;
  margin-bottom: 16px;
  padding-bottom: 8px;
  border-bottom: 1px solid #3a3a3a;
}

.filter-group-title {
  font-size: 14px;
  font-weight: 500;
  color: #e1e5e9;
  margin: 16px 0 8px 0;
}

.filter-checkbox-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.card-list-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.card-list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding: 16px 20px;
  background-color: #2a2a2a;
  border-radius: 8px;
  border: 1px solid #3a3a3a;
}

.statistics-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.total-count {
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
}

.status-statistics {
  display: flex;
  gap: 24px;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: #b0b0b0;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  display: inline-block;
}

.card-scroll-container {
  flex: 1;
  overflow-y: auto;
  padding-right: 8px;
}

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
  padding: 4px;
}

.card-item {
  transition: transform 0.2s ease;
}

.card-item:hover {
  transform: translateY(-2px);
}

.event-card {
  background-color: #2a2a2a !important;
  border: 1px solid #3a3a3a !important;
  border-radius: 8px;
  overflow: hidden;
}

.card-img-box {
  position: relative;
  cursor: pointer;
  overflow: hidden;
}

.card-img {
  width: 100%;
  height: 180px;
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
  height: 40px;
  border: 2px solid #ff4d4f;
  border-radius: 4px;
  background-color: rgba(255, 77, 79, 0.1);
}

.card-info {
  padding: 16px;
}

.card-info-row {
  margin-bottom: 8px;
  font-size: 12px;
  line-height: 1.4;
  color: #b0b0b0;
}

.card-info-row:last-child {
  margin-bottom: 0;
}

.card-info-row b {
  color: #e1e5e9;
  font-weight: 600;
  margin-right: 4px;
}

/* 设备状态行样式 */
.device-status-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #3a3a3a;
}

.device-group {
  font-size: 11px;
  color: #888;
  background-color: #3a3a3a;
  padding: 2px 6px;
  border-radius: 3px;
}

/* 信息分区样式 */
.card-info-section {
  margin-bottom: 12px;
  padding: 8px;
  background-color: rgba(58, 58, 58, 0.3);
  border-radius: 4px;
  border-left: 3px solid #409eff;
}

.card-info-section:last-child {
  margin-bottom: 0;
}

.section-title {
  font-size: 11px;
  font-weight: 600;
  color: #409eff;
  margin-bottom: 6px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.card-info-section .card-info-row {
  margin-bottom: 4px;
  font-size: 11px;
}

.card-info-section .card-info-row:last-child {
  margin-bottom: 0;
}

.device-status {
  display: flex;
  align-items: center;
  gap: 6px;
  font-weight: 600;
  font-size: 13px;
}

.card-actions {
  padding: 0 16px 16px;
  display: flex;
  justify-content: flex-end;
}

.gva-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}

/* Element Plus 组件样式覆盖 */
:deep(.el-card__body) {
  padding: 0 !important;
}

:deep(.el-checkbox) {
  color: #b0b0b0 !important;
}

:deep(.el-checkbox__label) {
  color: #b0b0b0 !important;
}

:deep(.el-checkbox.is-checked .el-checkbox__label) {
  color: #409eff !important;
}

:deep(.el-form-item) {
  margin-bottom: 16px;
}

:deep(.el-input__inner) {
  background-color: #3a3a3a !important;
  border-color: #4a4a4a !important;
  color: #ffffff !important;
}

:deep(.el-input__inner::placeholder) {
  color: #888888 !important;
}

:deep(.el-button--primary) {
  background-color: #409eff !important;
  border-color: #409eff !important;
}

:deep(.el-button) {
  background-color: #3a3a3a !important;
  border-color: #4a4a4a !important;
  color: #b0b0b0 !important;
}

:deep(.el-select .el-input__inner) {
  background-color: #3a3a3a !important;
  border-color: #4a4a4a !important;
  color: #ffffff !important;
}

:deep(.el-date-editor .el-input__inner) {
  background-color: #3a3a3a !important;
  border-color: #4a4a4a !important;
  color: #ffffff !important;
}

:deep(.el-pagination) {
  color: #b0b0b0 !important;
}

:deep(.el-pagination .el-pager li) {
  background-color: #3a3a3a !important;
  color: #b0b0b0 !important;
  border: 1px solid #4a4a4a !important;
}

:deep(.el-pagination .el-pager li.active) {
  background-color: #409eff !important;
  color: #ffffff !important;
}

:deep(.el-pagination .btn-prev),
:deep(.el-pagination .btn-next) {
  background-color: #3a3a3a !important;
  color: #b0b0b0 !important;
  border: 1px solid #4a4a4a !important;
}

:deep(.el-descriptions) {
  background-color: #2a2a2a !important;
}

:deep(.el-descriptions__label) {
  color: #e1e5e9 !important;
  background-color: #3a3a3a !important;
}

:deep(.el-descriptions__content) {
  color: #b0b0b0 !important;
  background-color: #2a2a2a !important;
}

:deep(.el-drawer) {
  background-color: #2a2a2a !important;
}

:deep(.el-drawer__header) {
  color: #ffffff !important;
  border-bottom: 1px solid #3a3a3a !important;
}

/* 滚动条样式 */
.card-scroll-container::-webkit-scrollbar {
  width: 6px;
}

.card-scroll-container::-webkit-scrollbar-track {
  background: #3a3a3a;
  border-radius: 3px;
}

.card-scroll-container::-webkit-scrollbar-thumb {
  background: #5a5a5a;
  border-radius: 3px;
}

.card-scroll-container::-webkit-scrollbar-thumb:hover {
  background: #6a6a6a;
}
</style>