<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->
<template>
  <div class="monitoring-interface">
    <!-- 头部区域 -->
    <div class="header-section">
              <div class="header-left">
          <el-icon class="camera-icon"><VideoCamera /></el-icon>
          <span class="header-title">摄像机实时监控画面</span>
        </div>
      <div class="header-right">
        <el-select v-model="selectedEnvironment" placeholder="选择监控环境" class="environment-selector">
          <el-option
            v-for="item in environments"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </div>
    </div>

    <!-- 视频监控区域 -->
    <div class="video-monitor-area">
      <div class="video-container">
        <!-- 模拟视频画面 -->
        <div class="video-feed">
          <div class="video-overlay">
            <el-icon class="play-icon"><VideoPlay /></el-icon>
            <span>实时监控画面</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 右侧联系按钮 -->
    <div class="contact-button">
      <el-button type="primary" circle class="contact-btn">
        <el-icon><ChatDotRound /></el-icon>
      </el-button>
      <span class="contact-text">联系我们</span>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { VideoCamera, VideoPlay, ChatDotRound } from '@element-plus/icons-vue'

  // 环境选择数据
  const selectedEnvironment = ref('env3')
  const environments = [
    { value: 'env1', label: '模拟环境1 / 北京天安门' },
    { value: 'env2', label: '模拟环境2 / 上海外滩' },
    { value: 'env3', label: '模拟环境3 / 山东泰安泰山' },
    { value: 'env4', label: '模拟环境4 / 深圳湾' }
  ]
</script>

<style scoped lang="scss">
.monitoring-interface {
  position: relative;
  width: 100%;
  height: 100%;
  background: #f5f5f5;
  border-radius: 8px;
  overflow: hidden;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: white;
  border-bottom: 1px solid #e4e7ed;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);

  .header-left {
    display: flex;
    align-items: center;
    gap: 6px;

    .camera-icon {
      font-size: 16px;
      color: #409eff;
    }

    .header-title {
      font-size: 14px;
      font-weight: 600;
      color: #303133;
    }
  }

  .header-right {
    .environment-selector {
      width: 200px;
      font-size: 12px;
    }
  }
}

.video-monitor-area {
  position: relative;
  width: 100%;
  height: calc(100% - 60px);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px;

  .video-container {
    width: 100%;
    height: 100%;
    border-radius: 6px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);

    .video-feed {
      width: 100%;
      height: 100%;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      display: flex;
      align-items: center;
      justify-content: center;
      position: relative;

      .video-overlay {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 12px;
        color: white;
        text-align: center;

        .play-icon {
          font-size: 32px;
          opacity: 0.8;
        }

        span {
          font-size: 12px;
          opacity: 0.9;
        }
      }
    }
  }
}

.contact-button {
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  z-index: 10;

  .contact-btn {
    width: 40px;
    height: 40px;
    font-size: 16px;
    box-shadow: 0 2px 8px rgba(64, 158, 255, 0.3);
  }

  .contact-text {
    font-size: 10px;
    color: #606266;
    white-space: nowrap;
  }
}

// 响应式设计
@media (max-width: 768px) {
  .header-section {
    padding: 8px 12px;

    .header-left .header-title {
      font-size: 12px;
    }

    .header-right .environment-selector {
      width: 150px;
    }
  }

  .contact-button {
    right: 12px;

    .contact-btn {
      width: 36px;
      height: 36px;
      font-size: 14px;
    }

    .contact-text {
      font-size: 9px;
    }
  }
}
</style>
