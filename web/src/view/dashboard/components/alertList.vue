<template>
  <div class="alert-list-container">
    <!-- Header -->
    <div class="alert-header">
      <div class="header-left">
        <el-icon class="bell-icon"><Bell /></el-icon>
        <span class="title">最近预警列表清单</span>
      </div>
      <div class="header-right">
        <span class="more-link">更多</span>
      </div>
    </div>

    <!-- Alert List -->
    <div class="alert-list">
      <div
        v-for="(alert, index) in alertList"
        :key="index"
        class="alert-item"
      >
        <div class="alert-content">
          <div class="alert-left">
            <div class="alert-icon">
              <el-icon><VideoCamera /></el-icon>
            </div>
            <div class="alert-info">
              <div class="alert-time">{{ alert.time }}</div>
              <div class="alert-location">
                <span class="label">监控区域:</span>
                <span class="value">{{ alert.location }}</span>
              </div>
              <div class="alert-type">
                <span class="label">预警类型:</span>
                <span class="value">{{ alert.type }}</span>
              </div>
            </div>
          </div>
          <div class="alert-right">
            <div class="alert-image">
              <img :src="alert.image" :alt="alert.type" />
            </div>
          </div>
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
import { Bell, VideoCamera } from '@element-plus/icons-vue'

// 预警数据 - 根据图片中的实际数据
const alertList = ref([
  {
    time: '2025-07-17 13:40',
    location: '杭州西湖',
    type: '小船皮筏艇识别',
    image: 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIwIiBoZWlnaHQ9IjgwIiB2aWV3Qm94PSIwIDAgMTIwIDgwIiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgo8cmVjdCB3aWR0aD0iMTIwIiBoZWlnaHQ9IjgwIiBmaWxsPSIjRjVGNUY1Ii8+CjxwYXRoIGQ9Ik00MCAzMEg4MFY1MEg0MFYzMFoiIGZpbGw9IiNEREREREQiLz4KPHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4PSI0OCIgeT0iMjgiPgo8cGF0aCBkPSJNOSAyTDcuMTcgNEg0QzIuOSA0IDIgNC45IDIgNlYxOEMyIDEwLjEgMi45IDIwIDQgMjBIMjBDMjEuMSAyMCAyMiAxOS4xIDIyIDE4VjZDMjIgNC45IDIxLjEgNCAyMCA0SDE2LjgzTDE1IDJIOVpNMTIgMTdDOS4yNCAyNyA3IDE0Ljc2IDcgMTJDNyA5LjI0IDkuMjQgNyAxMiA3QzE0Ljc2IDcgMTcgOS4yNCAxNyAxMkMxNyAxNC43NiAxNC43NiAxNyAxMiAxN1oiIGZpbGw9IiM5OTk5OTkiLz4KPC9zdmc+Cjwvc3ZnPgo='
  },
  {
    time: '2025-07-17 10:28',
    location: '杭州西湖',
    type: '小船皮筏艇识别',
    image: 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIwIiBoZWlnaHQ9IjgwIiB2aWV3Qm94PSIwIDAgMTIwIDgwIiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgo8cmVjdCB3aWR0aD0iMTIwIiBoZWlnaHQ9IjgwIiBmaWxsPSIjRjVGNUY1Ii8+CjxwYXRoIGQ9Ik00MCAzMEg4MFY1MEg0MFYzMFoiIGZpbGw9IiNEREREREQiLz4KPHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4PSI0OCIgeT0iMjgiPgo8cGF0aCBkPSJNOSAyTDcuMTcgNEg0QzIuOSA0IDIgNC45IDIgNlYxOEMyIDEwLjEgMi45IDIwIDQgMjBIMjBDMjEuMSAyMCAyMiAxOS4xIDIyIDE4VjZDMjIgNC45IDIxLjEgNCAyMCA0SDE2LjgzTDE1IDJIOVpNMTIgMTdDOS4yNCAyNyA3IDE0Ljc2IDcgMTJDNyA5LjI0IDkuMjQgNyAxMiA3QzE0Ljc2IDcgMTcgOS4yNCAxNyAxMkMxNyAxNC43NiAxNC43NiAxNyAxMiAxN1oiIGZpbGw9IiM5OTk5OTkiLz4KPC9zdmc+Cjwvc3ZnPgo='
  }
])
</script>

<style lang="scss" scoped>
.alert-list-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  background: #fff;
}

.alert-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;

  .header-left {
    display: flex;
    align-items: center;
    gap: 8px;

    .bell-icon {
      color: #666;
      font-size: 16px;
    }

    .title {
      font-size: 16px;
      font-weight: 500;
      color: #333;
    }
  }

  .header-right {
    .more-link {
      color: #409eff;
      font-size: 14px;
      cursor: pointer;
      
      &:hover {
        color: #66b1ff;
      }
    }
  }
}

.alert-list {
  flex: 1;
  overflow-y: auto;
  padding: 0 20px;
}

.alert-item {
  padding: 16px 0;
  border-bottom: 1px solid #f5f5f5;

  &:last-child {
    border-bottom: none;
  }

  .alert-content {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 16px;
  }

  .alert-left {
    display: flex;
    gap: 12px;
    flex: 1;

    .alert-icon {
      width: 40px;
      height: 40px;
      background: #e8f4ff;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;

      .el-icon {
        color: #409eff;
        font-size: 18px;
      }
    }

    .alert-info {
      flex: 1;

      .alert-time {
        font-size: 16px;
        font-weight: 500;
        color: #333;
        margin-bottom: 8px;
      }

      .alert-location,
      .alert-type {
        display: flex;
        align-items: center;
        margin-bottom: 4px;
        font-size: 14px;

        .label {
          color: #666;
          margin-right: 4px;
        }

        .value {
          color: #333;
        }
      }
    }
  }

  .alert-right {
    .alert-image {
      width: 120px;
      height: 80px;
      border-radius: 6px;
      overflow: hidden;
      background: #f5f5f5;

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }
}

.empty-state {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #999;

  .empty-icon {
    font-size: 48px;
    margin-bottom: 16px;
    opacity: 0.5;
  }

  .empty-text {
    font-size: 14px;
  }
}

// 响应式设计
@media (max-width: 768px) {
  .alert-item {
    .alert-content {
      flex-direction: column;
      gap: 12px;
    }

    .alert-right {
      align-self: center;

      .alert-image {
        width: 100px;
        height: 66px;
      }
    }
  }
}
</style>