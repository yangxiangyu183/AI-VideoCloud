<template>
  <div class="dashboard-container">
    <!-- Top Information Cards (Header Section) -->
    <div class="top-cards-section">
      <div class="info-card">
        <div class="card-content">
          <div class="card-title">摄像头布控</div>
          <div class="card-number">12</div>
          <div class="card-status success">100% 设备状态全部正常</div>
        </div>
      </div>
      <div class="info-card">
        <div class="card-content">
          <div class="card-title">今日预警数量</div>
          <div class="card-number">192</div>
          <div class="card-status decrease">▼10% 比昨天预警数量降低</div>
        </div>
      </div>
      <div class="info-card">
        <div class="card-content">
          <div class="card-title">昨日预警数量</div>
          <div class="card-number">220</div>
          <div class="card-status increase">▲12% 比前天预警数量升高</div>
        </div>
      </div>
    </div>

    <!-- Main Content Area -->
    <div class="main-content-section">
      <!-- Left Panel: Trend Analysis -->
      <div class="trend-analysis-panel">
        <gva-card title="预警排名走势分析" custom-class="trend-card">
          <div class="chart-container">
            <div class="chart-tabs">
              <el-button 
                :type="activeTimeRange === '24h' ? 'primary' : ''" 
                size="small"
                @click="handleTimeRangeChange('24h')"
              >
                24小时
              </el-button>
              <el-button 
                :type="activeTimeRange === '7d' ? 'primary' : ''" 
                size="small"
                @click="handleTimeRangeChange('7d')"
              >
                7天
              </el-button>
              <el-button 
                :type="activeTimeRange === '30d' ? 'primary' : ''" 
                size="small"
                @click="handleTimeRangeChange('30d')"
              >
                30天
              </el-button>
            </div>
            <gva-chart :type="4" :time-range="activeTimeRange" />
          </div>
        </gva-card>
      </div>

      <!-- Right Panel: Video and Alerts -->
      <div class="right-panel">
        <!-- Video Monitoring -->
        <div class="video-panel">
          <gva-quick-link />
        </div>
        
        <!-- Alert List -->
        <div class="alert-panel">
          <gva-alert-list />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import {
    GvaChart,
    GvaQuickLink,
    GvaCard,
    GvaAlertList
  } from './components'
  
  defineOptions({
    name: 'Dashboard'
  })

  // 时间范围状态
  const activeTimeRange = ref('24h')

  // 处理时间范围切换
  const handleTimeRangeChange = (range) => {
    activeTimeRange.value = range
  }
</script>

<style lang="scss" scoped>
.dashboard-container {
  padding: 20px;
  background: #f5f7fa;
  min-height: calc(100vh - 120px);
}

.top-cards-section {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  margin-bottom: 20px;
}

.info-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 24px;
  text-align: center;
  transition: transform 0.2s ease;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }

  .card-content {
    .card-title {
      font-size: 14px;
      color: #606266;
      margin-bottom: 12px;
    }

    .card-number {
      font-size: 36px;
      font-weight: bold;
      color: #303133;
      margin-bottom: 8px;
    }

    .card-status {
      font-size: 12px;
      font-weight: 500;

      &.success {
        color: #67c23a;
      }

      &.decrease {
        color: #e6a23c;
      }

      &.increase {
        color: #f56c6c;
      }
    }
  }
}

.main-content-section {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 20px;
  height: calc(100vh - 300px);
}

.trend-analysis-panel {
  .trend-card {
    :deep(.el-card__body) {
      padding: 16px 20px 20px 20px;
    }
  }

  .chart-container {
    display: flex;
    flex-direction: column;

    .chart-tabs {
      display: flex;
      gap: 8px;
      margin-bottom: 16px;
      justify-content: flex-end;
      flex-shrink: 0;

      .el-button {
        font-size: 12px;
        padding: 6px 12px;
        transition: all 0.3s ease;
        
        &:not(.el-button--primary) {
          background: #f5f7fa;
          border-color: #dcdfe6;
          color: #606266;
          
          &:hover {
            background: #ecf5ff;
            border-color: #409eff;
            color: #409eff;
          }
        }
      }
    }
  }
}

.right-panel {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.video-panel {
  flex: 1;
  min-height: 0;
}

.alert-panel {
  flex: 1;
  min-height: 0;
}



// 响应式设计
@media (max-width: 1200px) {
  .main-content-section {
    grid-template-columns: 1fr;
    grid-template-rows: auto auto;
    height: auto;
    gap: 16px;
  }

  .trend-analysis-panel {
    order: 1;
  }

  .right-panel {
    order: 2;
    flex-direction: row;
    gap: 16px;
  }

  .video-panel,
  .alert-panel {
    flex: 1;
    min-height: 300px;
  }
}

@media (max-width: 768px) {
  .dashboard-container {
    padding: 12px;
  }

  .top-cards-section {
    grid-template-columns: 1fr;
    gap: 12px;
  }

  .info-card {
    padding: 16px;

    .card-content {
      .card-number {
        font-size: 28px;
      }
    }
  }

  .right-panel {
    flex-direction: column;
  }
}
</style>
