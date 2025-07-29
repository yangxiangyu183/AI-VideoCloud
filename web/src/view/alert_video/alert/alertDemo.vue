<template>
  <div class="alert-demo-container">
    <div class="demo-header">
      <h1>预警管理系统</h1>
      <p>选择预警类型查看对应的数据</p>
    </div>
    
    <!-- 使用预警类型选择器组件 -->
    <AlertTypeSelector />
    
    <!-- 额外的功能区域 -->
    <div class="additional-features">
      <div class="feature-card">
        <h3>统计信息</h3>
        <div class="stats-grid">
          <div class="stat-item">
            <div class="stat-number">{{ totalAlerts }}</div>
            <div class="stat-label">总预警数</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ activeAlerts }}</div>
            <div class="stat-label">活跃预警</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ resolvedAlerts }}</div>
            <div class="stat-label">已解决</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import AlertTypeSelector from './alertTypeSelector.vue'

// 模拟统计数据
const totalAlerts = ref(156)
const activeAlerts = ref(23)
const resolvedAlerts = ref(133)
</script>

<style scoped>
.alert-demo-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.demo-header {
  text-align: center;
  color: white;
  margin-bottom: 30px;
}

.demo-header h1 {
  font-size: 2.5rem;
  margin: 0 0 10px 0;
  font-weight: 700;
}

.demo-header p {
  font-size: 1.1rem;
  margin: 0;
  opacity: 0.9;
}

.additional-features {
  margin-top: 30px;
  display: flex;
  justify-content: center;
}

.feature-card {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  min-width: 300px;
}

.feature-card h3 {
  margin: 0 0 20px 0;
  font-size: 18px;
  font-weight: 600;
  color: #333;
  text-align: center;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.stat-item {
  text-align: center;
}

.stat-number {
  font-size: 2rem;
  font-weight: 700;
  color: #007bff;
  margin-bottom: 5px;
}

.stat-label {
  font-size: 14px;
  color: #666;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .demo-header h1 {
    font-size: 2rem;
  }
  
  .feature-card {
    min-width: auto;
    width: 100%;
  }
}
</style> 