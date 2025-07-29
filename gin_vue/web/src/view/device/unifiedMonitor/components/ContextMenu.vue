<template>
  <teleport to="body">
    <div
      v-if="visible"
      ref="menuRef"
      class="context-menu"
      :style="menuStyle"
      @click.stop
      @contextmenu.prevent
    >
      <div class="menu-item-group">
        <!-- 分组节点菜单项 -->
        <template v-if="nodeType === 'group'">
          <div class="menu-item" @click="handleEdit">
            <el-icon><Edit /></el-icon>
            <span>编辑分组</span>
          </div>
          <div class="menu-item" @click="handleDelete">
            <el-icon><Delete /></el-icon>
            <span>删除分组</span>
          </div>
          <div class="menu-divider"></div>
          <div class="menu-item" @click="handleAddSubGroup">
            <el-icon><FolderAdd /></el-icon>
            <span>新增子分组</span>
          </div>
          <div class="menu-item" @click="handleAddDevice">
            <el-icon><Plus /></el-icon>
            <span>添加设备</span>
          </div>
        </template>
        
        <!-- 设备节点菜单项 -->
        <template v-else-if="nodeType === 'device'">
          <div class="menu-item" @click="handleEdit">
            <el-icon><Edit /></el-icon>
            <span>编辑设备</span>
          </div>
          <div class="menu-item" @click="handleDelete">
            <el-icon><Delete /></el-icon>
            <span>删除设备</span>
          </div>
          <div class="menu-divider"></div>
          <div class="menu-item" @click="handleAddToMonitor">
            <el-icon><VideoCamera /></el-icon>
            <span>添加到监控</span>
          </div>
          <div class="menu-item" @click="handleTestConnection">
            <el-icon><Link /></el-icon>
            <span>测试连接</span>
          </div>
        </template>
      </div>
    </div>
  </teleport>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onUnmounted } from 'vue'
import { 
  Edit, 
  Delete, 
  FolderAdd, 
  Plus, 
  VideoCamera, 
  Link 
} from '@element-plus/icons-vue'

// Props
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  position: {
    type: Object,
    default: () => ({ x: 0, y: 0 })
  },
  nodeType: {
    type: String,
    validator: (value) => ['group', 'device'].includes(value)
  },
  nodeData: {
    type: Object,
    default: null
  }
})

// Emits
const emit = defineEmits([
  'edit',
  'delete', 
  'add-sub-group',
  'add-device',
  'add-to-monitor',
  'test-connection',
  'close'
])

// 响应式数据
const menuRef = ref()

// 计算菜单样式
const menuStyle = computed(() => {
  const { x, y } = props.position
  
  // 获取视口尺寸
  const viewportWidth = window.innerWidth
  const viewportHeight = window.innerHeight
  
  // 菜单预估尺寸
  const menuWidth = 160
  const menuHeight = props.nodeType === 'group' ? 120 : 120
  
  // 计算最终位置，避免超出视口
  let finalX = x
  let finalY = y
  
  if (x + menuWidth > viewportWidth) {
    finalX = x - menuWidth
  }
  
  if (y + menuHeight > viewportHeight) {
    finalY = y - menuHeight
  }
  
  return {
    position: 'fixed',
    left: `${Math.max(0, finalX)}px`,
    top: `${Math.max(0, finalY)}px`,
    zIndex: 9999
  }
})

// 菜单项点击处理
const handleEdit = () => {
  emit('edit', props.nodeData)
  emit('close')
}

const handleDelete = () => {
  emit('delete', props.nodeData)
  emit('close')
}

const handleAddSubGroup = () => {
  emit('add-sub-group', props.nodeData)
  emit('close')
}

const handleAddDevice = () => {
  emit('add-device', props.nodeData)
  emit('close')
}

const handleAddToMonitor = () => {
  emit('add-to-monitor', props.nodeData)
  emit('close')
}

const handleTestConnection = () => {
  emit('test-connection', props.nodeData)
  emit('close')
}

// 点击外部关闭菜单
const handleClickOutside = (event) => {
  if (props.visible && menuRef.value && !menuRef.value.contains(event.target)) {
    emit('close')
  }
}

// 按ESC键关闭菜单
const handleKeydown = (event) => {
  if (event.key === 'Escape' && props.visible) {
    emit('close')
  }
}

// 生命周期
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<style scoped>
.context-menu {
  background: white;
  border: 1px solid #e4e7ed;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  padding: 4px 0;
  min-width: 160px;
  user-select: none;
}

.menu-item-group {
  display: flex;
  flex-direction: column;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 16px;
  font-size: 14px;
  color: #606266;
  cursor: pointer;
  transition: all 0.2s;
}

.menu-item:hover {
  background-color: #f5f7fa;
  color: #409eff;
}

.menu-item:active {
  background-color: #ecf5ff;
}

.menu-divider {
  height: 1px;
  background-color: #e4e7ed;
  margin: 4px 0;
}

.menu-item .el-icon {
  font-size: 16px;
  color: inherit;
}

/* 动画效果 */
.context-menu {
  animation: menuFadeIn 0.15s ease-out;
}

@keyframes menuFadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
</style>