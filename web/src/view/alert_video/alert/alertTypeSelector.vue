<template>
  <div class="alert-type-container">
    <!-- 预警类型选择区域 -->
    <div class="alert-type-selector">
      <h3 class="selector-title">预警类型</h3>
      <div class="checkbox-group">
        <div 
          v-for="type in alertTypes" 
          :key="type.id" 
          class="checkbox-item"
        >
          <input
            type="checkbox"
            :id="'alert-type-' + type.id"
            :value="type.id"
            v-model="selectedTypes"
            @change="handleTypeChange"
            class="checkbox-input"
          />
          <label :for="'alert-type-' + type.id" class="checkbox-label">
            {{ type.name }}
          </label>
        </div>
      </div>
    </div>

    <!-- 数据展示区域 -->
    <div class="data-display">
      <h3 class="display-title">预警数据</h3>
      <div v-if="filteredData.length === 0" class="no-data">
        请选择预警类型查看数据
      </div>
      <div v-else class="data-list">
        <div 
          v-for="item in filteredData" 
          :key="item.id" 
          class="data-item"
        >
          <div class="item-header">
            <span class="item-type" :class="'type-' + item.typeId">
              {{ getTypeName(item.typeId) }}
            </span>
            <span class="item-time">{{ formatTime(item.createTime) }}</span>
          </div>
          <div class="item-content">
            <div class="item-title">{{ item.title }}</div>
            <div class="item-description">{{ item.description }}</div>
          </div>
          <div class="item-footer">
            <span class="item-status" :class="'status-' + item.status">
              {{ getStatusText(item.status) }}
            </span>
            <span class="item-device">设备: {{ item.deviceName }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

// 预警类型数据
const alertTypes = ref([
  { id: 1, name: '设备离线', color: '#ff4757' },
  { id: 2, name: '温度异常', color: '#ffa502' },
  { id: 3, name: '湿度异常', color: '#2ed573' },
  { id: 4, name: '电量不足', color: '#ff6348' },
  { id: 5, name: '网络异常', color: '#5352ed' },
  { id: 6, name: '存储空间不足', color: '#ff6b6b' }
])

// 选中的预警类型
const selectedTypes = ref([])

// 模拟的预警数据
const alertData = ref([
  {
    id: 1,
    typeId: 1,
    title: '设备离线预警',
    description: '设备ID: DEV001 在 2024-01-15 10:30:00 检测到离线',
    status: 'active',
    createTime: '2024-01-15T10:30:00',
    deviceName: '监控设备001'
  },
  {
    id: 2,
    typeId: 2,
    title: '温度异常预警',
    description: '设备ID: DEV002 温度达到 45°C，超过正常范围',
    status: 'resolved',
    createTime: '2024-01-15T09:15:00',
    deviceName: '温控设备002'
  },
  {
    id: 3,
    typeId: 3,
    title: '湿度异常预警',
    description: '设备ID: DEV003 湿度低于 30%，需要加湿',
    status: 'active',
    createTime: '2024-01-15T08:45:00',
    deviceName: '湿度设备003'
  },
  {
    id: 4,
    typeId: 4,
    title: '电量不足预警',
    description: '设备ID: DEV004 电量剩余 15%，请及时充电',
    status: 'warning',
    createTime: '2024-01-15T07:20:00',
    deviceName: '移动设备004'
  },
  {
    id: 5,
    typeId: 5,
    title: '网络异常预警',
    description: '设备ID: DEV005 网络连接不稳定，丢包率 15%',
    status: 'active',
    createTime: '2024-01-15T06:30:00',
    deviceName: '网络设备005'
  },
  {
    id: 6,
    typeId: 6,
    title: '存储空间不足预警',
    description: '设备ID: DEV006 存储空间剩余 5GB，建议清理',
    status: 'warning',
    createTime: '2024-01-15T05:45:00',
    deviceName: '存储设备006'
  }
])

// 根据选中类型过滤数据
const filteredData = computed(() => {
  if (selectedTypes.value.length === 0) {
    return []
  }
  return alertData.value.filter(item => selectedTypes.value.includes(item.typeId))
})

// 处理类型选择变化
const handleTypeChange = () => {
  console.log('选中的预警类型:', selectedTypes.value)
  console.log('过滤后的数据:', filteredData.value)
}

// 获取类型名称
const getTypeName = (typeId) => {
  const type = alertTypes.value.find(t => t.id === typeId)
  return type ? type.name : '未知类型'
}

// 获取状态文本
const getStatusText = (status) => {
  const statusMap = {
    'active': '活跃',
    'resolved': '已解决',
    'warning': '警告'
  }
  return statusMap[status] || '未知状态'
}

// 格式化时间
const formatTime = (timeStr) => {
  const date = new Date(timeStr)
  return date.toLocaleString('zh-CN')
}

// 组件挂载时初始化
onMounted(() => {
  console.log('预警类型选择器组件已挂载')
})
</script>

<style scoped>
.alert-type-container {
  display: flex;
  gap: 20px;
  padding: 20px;
  background: #f8f9fa;
  min-height: 600px;
}

.alert-type-selector {
  flex: 0 0 250px;
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.selector-title {
  margin: 0 0 20px 0;
  font-size: 18px;
  font-weight: 600;
  color: #333;
  border-bottom: 2px solid #007bff;
  padding-bottom: 10px;
}

.checkbox-group {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.checkbox-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  border-radius: 6px;
  transition: background-color 0.2s;
}

.checkbox-item:hover {
  background-color: #f8f9fa;
}

.checkbox-input {
  width: 16px;
  height: 16px;
  cursor: pointer;
}

.checkbox-label {
  cursor: pointer;
  font-size: 14px;
  color: #555;
  user-select: none;
}

.data-display {
  flex: 1;
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.display-title {
  margin: 0 0 20px 0;
  font-size: 18px;
  font-weight: 600;
  color: #333;
  border-bottom: 2px solid #28a745;
  padding-bottom: 10px;
}

.no-data {
  text-align: center;
  color: #999;
  font-size: 16px;
  padding: 40px 0;
}

.data-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.data-item {
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 16px;
  transition: box-shadow 0.2s;
}

.data-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.item-type {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  color: white;
}

.type-1 { background-color: #ff4757; }
.type-2 { background-color: #ffa502; }
.type-3 { background-color: #2ed573; }
.type-4 { background-color: #ff6348; }
.type-5 { background-color: #5352ed; }
.type-6 { background-color: #ff6b6b; }

.item-time {
  font-size: 12px;
  color: #666;
}

.item-content {
  margin-bottom: 12px;
}

.item-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.item-description {
  font-size: 14px;
  color: #666;
  line-height: 1.5;
}

.item-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.item-status {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.status-active {
  background-color: #dc3545;
  color: white;
}

.status-resolved {
  background-color: #28a745;
  color: white;
}

.status-warning {
  background-color: #ffc107;
  color: #212529;
}

.item-device {
  font-size: 12px;
  color: #666;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .alert-type-container {
    flex-direction: column;
  }
  
  .alert-type-selector {
    flex: none;
  }
}
</style> 