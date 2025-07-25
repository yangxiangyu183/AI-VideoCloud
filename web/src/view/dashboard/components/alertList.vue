<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->
<template>
  <div class="alert-list-container">
    <!-- Header with More button -->
    <div class="alert-list-header">
      <div class="header-left">
        <span class="header-title">最近预警列表清单</span>
      </div>
      <div class="header-right">
        <el-button type="primary" size="small" text>更多</el-button>
      </div>
    </div>

    <!-- Alert List -->
    <div class="alert-list">
      <div
        v-for="(alert, index) in alertList"
        :key="index"
        class="alert-item"
        @click="handleAlertClick(alert)"
      >
        <div class="alert-thumbnail">
          <div class="thumbnail-placeholder">
            <el-icon><component :is="alert.icon" /></el-icon>
          </div>
        </div>
        <div class="alert-content">
          <div class="alert-time">{{ alert.time }}</div>
          <div class="alert-area">{{ alert.area }}</div>
          <div class="alert-type">{{ alert.type }}</div>
        </div>
      </div>
    </div>

    <!-- Empty state -->
    <div v-if="alertList.length === 0" class="empty-state">
      <el-icon class="empty-icon"><Bell /></el-icon>
      <span class="empty-text">暂无预警信息</span>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { Picture, Bell, Ship, Bicycle, Drone, TruckFilled, User } from '@element-plus/icons-vue'

  // 预警数据 - 根据图片中的实际数据
  const alertList = ref([
    {
      time: '2025-07-23 09:27',
      area: '杭州西湖',
      type: '小船皮筏艇识别',
      icon: Ship,
      thumbnail: '/path/to/boat-image.jpg'
    },
    {
      time: '2025-07-23 09:25',
      area: '桂林阳桥',
      type: '非机动车识别',
      icon: Bicycle,
      thumbnail: '/path/to/motorcycle-image.jpg'
    },
    {
      time: '2025-07-23 09:22',
      area: '千岛湖',
      type: '无人机识别',
      icon: Drone,
      thumbnail: '/path/to/drone-image.jpg'
    },
    {
      time: '2025-07-23 09:20',
      area: '北京天安门',
      type: '车辆违停识别',
      icon: TruckFilled,
      thumbnail: '/path/to/parking-image.jpg'
    },
    {
      time: '2025-07-23 09:18',
      area: '上海外滩',
      type: '人员聚集识别',
      icon: User,
      thumbnail: '/path/to/crowd-image.jpg'
    },
    {
      time: '2025-07-23 09:15',
      area: '深圳湾',
      type: '异常行为识别',
      icon: Picture,
      thumbnail: '/path/to/behavior-image.jpg'
    }
  ])

  // 处理预警点击
  const handleAlertClick = (alert) => {
    console.log('Alert clicked:', alert)
    // 这里可以添加跳转到预警详情页面的逻辑
  }
</script>

<style scoped lang="scss">
.alert-list-container {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.alert-list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #e4e7ed;
  background: #fafafa;

  .header-left {
    .header-title {
      font-size: 14px;
      font-weight: 600;
      color: #303133;
    }
  }

  .header-right {
    .el-button {
      font-size: 12px;
      color: #409eff;
    }
  }
}

.alert-list {
  flex: 1;
  overflow-y: auto;
  padding: 0;
}

.alert-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background-color: #f5f7fa;
  }

  &:last-child {
    border-bottom: none;
  }
}

.alert-thumbnail {
  width: 40px;
  height: 40px;
  margin-right: 12px;
  flex-shrink: 0;

  .thumbnail-placeholder {
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 6px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;

    .el-icon {
      font-size: 18px;
    }
  }
}

.alert-content {
  flex: 1;
  min-width: 0;

  .alert-time {
    font-size: 11px;
    color: #909399;
    margin-bottom: 3px;
  }

  .alert-area {
    font-size: 13px;
    color: #303133;
    font-weight: 500;
    margin-bottom: 2px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .alert-type {
    font-size: 11px;
    color: #606266;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
}

.empty-state {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #909399;

  .empty-icon {
    font-size: 32px;
    margin-bottom: 12px;
    opacity: 0.5;
  }

  .empty-text {
    font-size: 12px;
  }
}

// 滚动条样式
.alert-list::-webkit-scrollbar {
  width: 4px;
}

.alert-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 2px;
}

.alert-list::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 2px;
}

.alert-list::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style> 