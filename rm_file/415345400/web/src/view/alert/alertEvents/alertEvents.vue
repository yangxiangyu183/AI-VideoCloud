
<template>
  <div class="alert-page">
    <!-- 应用标题栏 -->
    <div class="app-header">
      <div class="app-title">思通数科AI视频卫士</div>
      <div class="app-nav">
        <a href="#" class="nav-link">开源社区</a>
        <a href="#" class="nav-link">新手入门</a>
        <a href="#" class="nav-link">系统设置</a>
        <a href="#" class="nav-link">退出</a>
      </div>
    </div>

    <!-- 顶部导航栏 -->
    <div class="top-nav">
      <div class="nav-item" :class="{ active: currentNav === 'dashboard' }" @click="currentNav = 'dashboard'">
        数据看板
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'monitor' }" @click="currentNav = 'monitor'">
        监测任务
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'alert' }" @click="currentNav = 'alert'">
        事件告警
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'video' }" @click="currentNav = 'video'">
        视频接入
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'algorithm' }" @click="currentNav = 'algorithm'">
        算法管理
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'orchestration' }" @click="currentNav = 'orchestration'">
        算法编排
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'training' }" @click="currentNav = 'training'">
        模型训练
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'screen' }" @click="currentNav = 'screen'">
        监控大屏
      </div>
      <div class="nav-item" :class="{ active: currentNav === 'config' }" @click="currentNav = 'config'">
        系统配置
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="main-content">
      <!-- 左侧筛选条件 -->
      <div class="filter-sidebar">
        <div class="filter-header">
          <span class="filter-icon">≡</span>
          <span class="filter-title">筛选条件</span>
        </div>
        <div class="filter-section">
          <div class="section-title">预警类型</div>
          <div class="checkbox-group">
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.selectAll" @change="handleSelectAll">
              <span class="checkmark"></span>
              <span class="checkbox-text">■ 全选</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.smoking" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">吸烟监测</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.litter" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">地面垃圾</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.vendor" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">游摊小贩识别</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.vehicle" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">车辆类型识别</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.nonMotorized" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">非机动车识别</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.drone" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">无人机识别</span>
            </label>
            <label class="checkbox-item">
              <input type="checkbox" v-model="filterOptions.boat" @change="handleFilterChange">
              <span class="checkmark"></span>
              <span class="checkbox-text">小船皮筏艇识别</span>
            </label>
          </div>
        </div>
      </div>

      <!-- 右侧内容区域 -->
      <div class="content-area">
        <!-- 搜索栏 -->
        <div class="search-bar">
          <div class="search-input-group">
            <input 
              type="text" 
              v-model="searchKeywords" 
              placeholder="请输入多个关键词,用空格隔开"
              class="search-input"
            >
            <button class="search-btn" @click="handleSearch">
              <i class="search-icon">🔍</i>
              搜索
            </button>
          </div>
        </div>

        <!-- 告警统计和筛选 -->
        <div class="alert-stats">
          <div class="stats-info">共有{{ totalAlerts }}条告警事件</div>
          <div class="filter-controls">
            <select v-model="selectedCamera" class="filter-select">
              <option value="">摄像头点位</option>
              <option value="hangzhou">杭州西湖</option>
              <option value="guilin">桂林阳桥</option>
              <option value="qiandao">千岛湖</option>
              <option value="xinyi">新沂新悦广场</option>
            </select>
            <select v-model="selectedTimeRange" class="filter-select">
              <option value="">时间范围</option>
              <option value="today">今天</option>
              <option value="week">本周</option>
              <option value="month">本月</option>
            </select>
          </div>
        </div>

        <!-- 告警网格 -->
        <div class="alert-grid">
          <div 
            v-for="alert in filteredAlerts" 
            :key="alert.id" 
            class="alert-card"
            @click="viewAlertDetail(alert)"
          >
            <div class="alert-image">
              <img :src="alert.imageUrl" :alt="alert.type" class="alert-img">
              <div v-if="alert.confidence" class="confidence-overlay">
                {{ alert.confidence }}%
              </div>
              <div class="bounding-box"></div>
            </div>
            <div class="alert-info">
              <div class="camera-location">{{ alert.cameraLocation }}</div>
              <div class="alert-type">{{ alert.type }}</div>
              <div class="alert-time">{{ alert.time }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 联系我们按钮 -->
    <div class="contact-btn" @click="contactUs">
      <i class="contact-icon">💬</i>
      联系我们
    </div>

    <!-- 告警详情弹窗 -->
    <div v-if="showDetailModal" class="modal-overlay" @click="closeDetailModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>告警详情</h3>
          <button class="close-btn" @click="closeDetailModal">×</button>
        </div>
        <div class="modal-body" v-if="selectedAlert">
          <div class="detail-image">
            <img :src="selectedAlert.imageUrl" :alt="selectedAlert.type" class="detail-img">
            <div v-if="selectedAlert.confidence" class="detail-confidence">
              {{ selectedAlert.confidence }}%
            </div>
          </div>
          <div class="detail-info">
            <div class="detail-item">
              <span class="detail-label">摄像头点位:</span>
              <span class="detail-value">{{ selectedAlert.cameraLocation }}</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">预警类型:</span>
              <span class="detail-value">{{ selectedAlert.type }}</span>
            </div>
            <div class="detail-item">
              <span class="detail-label">预警时间:</span>
              <span class="detail-value">{{ selectedAlert.time }}</span>
            </div>
            <div class="detail-item" v-if="selectedAlert.confidence">
              <span class="detail-label">置信度:</span>
              <span class="detail-value">{{ selectedAlert.confidence }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'AlertEvents'
})

// 当前导航
const currentNav = ref('alert')

// 搜索关键词
const searchKeywords = ref('')

// 筛选选项
const filterOptions = reactive({
  selectAll: true,
  smoking: false,
  litter: false,
  vendor: false,
  vehicle: false,
  nonMotorized: false,
  drone: false,
  boat: false
})

// 筛选控制
const selectedCamera = ref('')
const selectedTimeRange = ref('')

// 弹窗控制
const showDetailModal = ref(false)
const selectedAlert = ref(null)

// 告警数据
const alerts = ref([
  {
    id: 1,
    imageUrl: 'https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=300&h=200&fit=crop',
    confidence: 87,
    cameraLocation: '杭州西湖',
    type: '小船皮筏艇识别',
    time: '2025-07-17 13:40',
    category: 'boat'
  },
  {
    id: 2,
    imageUrl: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '杭州西湖',
    type: '小船皮筏艇识别',
    time: '2025-07-17 10:28',
    category: 'boat'
  },
  {
    id: 3,
    imageUrl: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '桂林阳桥',
    type: '非机动车识别',
    time: '2025-07-17 09:27',
    category: 'nonMotorized'
  },
  {
    id: 4,
    imageUrl: 'https://images.unsplash.com/photo-1579829366248-204fe8413f31?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '千岛湖',
    type: '无人机识别',
    time: '2025-07-17 06:10',
    category: 'drone'
  },
  {
    id: 5,
    imageUrl: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '桂林阳桥',
    type: '非机动车识别',
    time: '2025-07-17 17:32',
    category: 'nonMotorized'
  },
  {
    id: 6,
    imageUrl: 'https://images.unsplash.com/photo-1558618666-fcd25c85cd64?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '桂林阳桥',
    type: '非机动车识别',
    time: '2025-07-17 11:20',
    category: 'nonMotorized'
  },
  {
    id: 7,
    imageUrl: 'https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '杭州西湖',
    type: '小船皮筏艇识别',
    time: '2025-07-17 09:53',
    category: 'boat'
  },
  {
    id: 8,
    imageUrl: 'https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=300&h=200&fit=crop',
    confidence: null,
    cameraLocation: '新沂新悦广场',
    type: '游摊小贩识别',
    time: '2025-07-17 09:05',
    category: 'vendor'
  }
])

// 计算属性
const totalAlerts = computed(() => filteredAlerts.value.length)

const filteredAlerts = computed(() => {
  let filtered = alerts.value

  // 按预警类型筛选
  if (!filterOptions.selectAll) {
    const selectedTypes = []
    if (filterOptions.smoking) selectedTypes.push('smoking')
    if (filterOptions.litter) selectedTypes.push('litter')
    if (filterOptions.vendor) selectedTypes.push('vendor')
    if (filterOptions.vehicle) selectedTypes.push('vehicle')
    if (filterOptions.nonMotorized) selectedTypes.push('nonMotorized')
    if (filterOptions.drone) selectedTypes.push('drone')
    if (filterOptions.boat) selectedTypes.push('boat')
    
    if (selectedTypes.length > 0) {
      filtered = filtered.filter(alert => selectedTypes.includes(alert.category))
    }
  }

  // 按摄像头点位筛选
  if (selectedCamera.value) {
    const cameraMap = {
      'hangzhou': '杭州西湖',
      'guilin': '桂林阳桥',
      'qiandao': '千岛湖',
      'xinyi': '新沂新悦广场'
    }
    filtered = filtered.filter(alert => alert.cameraLocation === cameraMap[selectedCamera.value])
  }

  // 按关键词搜索
  if (searchKeywords.value.trim()) {
    const keywords = searchKeywords.value.toLowerCase().split(' ')
    filtered = filtered.filter(alert => {
      return keywords.some(keyword => 
        alert.cameraLocation.toLowerCase().includes(keyword) ||
        alert.type.toLowerCase().includes(keyword) ||
        alert.time.includes(keyword)
      )
    })
  }

  return filtered
})

// 方法
const handleSelectAll = () => {
  if (filterOptions.selectAll) {
    filterOptions.smoking = false
    filterOptions.litter = false
    filterOptions.vendor = false
    filterOptions.vehicle = false
    filterOptions.nonMotorized = false
    filterOptions.drone = false
    filterOptions.boat = false
  }
}

const handleFilterChange = () => {
  const hasSelection = filterOptions.smoking || filterOptions.litter || 
                      filterOptions.vendor || filterOptions.vehicle || 
                      filterOptions.nonMotorized || filterOptions.drone || 
                      filterOptions.boat
  
  if (hasSelection) {
    filterOptions.selectAll = false
  } else {
    filterOptions.selectAll = true
  }
}

const handleSearch = () => {
  // 搜索逻辑已在计算属性中处理
  ElMessage.success('搜索完成')
}

const viewAlertDetail = (alert) => {
  selectedAlert.value = alert
  showDetailModal.value = true
}

const closeDetailModal = () => {
  showDetailModal.value = false
  selectedAlert.value = null
}

const contactUs = () => {
  ElMessage.info('联系我们功能')
}

// 生命周期
onMounted(() => {
  // 初始化数据
})
</script>

<style scoped>
.alert-page {
  min-height: 100vh;
  background-color: #f5f5f5;
  position: relative;
}

/* 应用标题栏 */
.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  padding: 10px 20px;
  border-bottom: 1px solid #e0e0e0;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.app-title {
  font-size: 18px;
  font-weight: bold;
  color: #333;
}

.app-nav {
  display: flex;
  gap: 20px;
}

.nav-link {
  color: #666;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.3s ease;
}

.nav-link:hover {
  color: #1890ff;
}

/* 顶部导航栏 */
.top-nav {
  display: flex;
  background-color: #fff;
  border-bottom: 1px solid #e0e0e0;
  padding: 0 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  height: 50px;
  align-items: center;
}

.nav-item {
  padding: 15px 20px;
  cursor: pointer;
  border-bottom: 3px solid transparent;
  transition: all 0.3s ease;
  font-size: 14px;
  color: #666;
  white-space: nowrap;
}

.nav-item:hover {
  color: #1890ff;
  background-color: #f0f8ff;
}

.nav-item.active {
  color: #1890ff;
  border-bottom-color: #1890ff;
  background-color: #f0f8ff;
  font-weight: 500;
}

/* 主要内容区域 */
.main-content {
  display: flex;
  height: calc(100vh - 120px);
}

/* 左侧筛选条件 */
.filter-sidebar {
  width: 280px;
  background-color: #fff;
  border-right: 1px solid #e0e0e0;
  padding: 20px;
  overflow-y: auto;
}

.filter-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 1px solid #e0e0e0;
}

.filter-icon {
  font-size: 18px;
  color: #666;
}

.filter-title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

.filter-section {
  margin-bottom: 20px;
}

.section-title {
  font-size: 14px;
  font-weight: bold;
  color: #333;
  margin-bottom: 15px;
}

.checkbox-group {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.checkbox-item {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: 14px;
  color: #666;
  position: relative;
  padding-left: 25px;
  transition: color 0.3s ease;
}

.checkbox-item:hover {
  color: #1890ff;
}

.checkbox-item input[type="checkbox"] {
  position: absolute;
  opacity: 0;
  cursor: pointer;
}

.checkmark {
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  height: 16px;
  width: 16px;
  background-color: #fff;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  transition: all 0.3s ease;
}

.checkbox-item input:checked ~ .checkmark {
  background-color: #1890ff;
  border-color: #1890ff;
}

.checkmark:after {
  content: "";
  position: absolute;
  display: none;
}

.checkbox-item input:checked ~ .checkmark:after {
  display: block;
}

.checkbox-item .checkmark:after {
  left: 5px;
  top: 2px;
  width: 4px;
  height: 8px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.checkbox-text {
  margin-left: 5px;
}

/* 右侧内容区域 */
.content-area {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background-color: #f5f5f5;
}

/* 搜索栏 */
.search-bar {
  margin-bottom: 20px;
}

.search-input-group {
  display: flex;
  gap: 10px;
  max-width: 600px;
}

.search-input {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  font-size: 14px;
  outline: none;
  transition: border-color 0.3s ease;
}

.search-input:focus {
  border-color: #1890ff;
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}

.search-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 12px 20px;
  background-color: #1890ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s ease;
  font-weight: 500;
}

.search-btn:hover {
  background-color: #40a9ff;
}

.search-icon {
  font-size: 14px;
}

/* 告警统计和筛选 */
.alert-stats {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding: 15px 0;
  border-bottom: 1px solid #e0e0e0;
}

.stats-info {
  font-size: 14px;
  color: #666;
  font-weight: 500;
}

.filter-controls {
  display: flex;
  gap: 15px;
}

.filter-select {
  padding: 8px 12px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  font-size: 14px;
  outline: none;
  background-color: white;
  transition: border-color 0.3s ease;
}

.filter-select:focus {
  border-color: #1890ff;
}

/* 告警网格 */
.alert-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.alert-card {
  background-color: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  border: 1px solid #f0f0f0;
}

.alert-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0,0,0,0.15);
}

.alert-image {
  position: relative;
  height: 200px;
  overflow: hidden;
}

.alert-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.confidence-overlay {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: bold;
}

.bounding-box {
  position: absolute;
  top: 20px;
  left: 20px;
  right: 20px;
  bottom: 20px;
  border: 2px solid #ff4d4f;
  border-radius: 4px;
}

.alert-info {
  padding: 15px;
}

.camera-location {
  font-size: 14px;
  color: #333;
  font-weight: bold;
  margin-bottom: 8px;
}

.alert-type {
  font-size: 13px;
  color: #666;
  margin-bottom: 6px;
}

.alert-time {
  font-size: 12px;
  color: #999;
}

/* 联系我们按钮 */
.contact-btn {
  position: fixed;
  right: 30px;
  bottom: 30px;
  width: 60px;
  height: 60px;
  background-color: #1890ff;
  color: white;
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.4);
  transition: all 0.3s ease;
  font-size: 12px;
  gap: 2px;
  z-index: 1000;
}

.contact-btn:hover {
  transform: scale(1.1);
  box-shadow: 0 6px 16px rgba(24, 144, 255, 0.6);
}

.contact-icon {
  font-size: 16px;
}

/* 告警详情弹窗 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1001;
}

.modal-content {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 20px;
  border-bottom: 1px solid #e0e0e0;
  background-color: #f5f5f5;
}

.modal-header h3 {
  margin: 0;
  font-size: 18px;
  color: #333;
}

.close-btn {
  background-color: #f0f0f0;
  border: none;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 20px;
  color: #666;
  transition: background-color 0.3s ease;
}

.close-btn:hover {
  background-color: #e0e0e0;
}

.modal-body {
  padding: 20px;
  overflow-y: auto;
  flex-grow: 1;
}

.detail-image {
  position: relative;
  height: 250px;
  margin-bottom: 15px;
  overflow: hidden;
}

.detail-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.detail-confidence {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: bold;
}

.detail-info {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.detail-item {
  display: flex;
  align-items: center;
}

.detail-label {
  font-size: 14px;
  color: #666;
  font-weight: bold;
  margin-right: 10px;
}

.detail-value {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    flex-direction: column;
  }
  
  .filter-sidebar {
    width: 100%;
    height: auto;
    border-right: none;
    border-bottom: 1px solid #e0e0e0;
  }
  
  .alert-grid {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  }
  
  .top-nav {
    overflow-x: auto;
    padding: 0 10px;
  }
  
  .nav-item {
    padding: 15px 15px;
    white-space: nowrap;
  }
  
  .app-header {
    flex-direction: column;
    gap: 10px;
  }
  
  .app-nav {
    gap: 15px;
  }
}
</style>
