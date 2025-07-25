<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑设备' : '新增设备'"
    width="600px"
    :before-close="handleClose"
    destroy-on-close
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="120px"
      @submit.prevent
    >
      <el-form-item label="设备名称" prop="deviceName">
        <el-input 
          v-model="form.deviceName" 
          placeholder="请输入设备名称"
          maxlength="50"
          show-word-limit
        />
      </el-form-item>
      
      <el-form-item label="所属分组" prop="groupId">
        <el-select 
          v-model="form.groupId" 
          placeholder="请选择分组" 
          style="width: 100%"
          filterable
        >
          <el-option 
            v-for="group in groups" 
            :key="group.ID" 
            :label="group.groupName" 
            :value="group.ID" 
          />
        </el-select>
      </el-form-item>
      
      <el-form-item label="视频流地址" prop="streamUrl">
        <el-input 
          v-model="form.streamUrl" 
          placeholder="请输入视频流地址，如：rtsp://192.168.1.100:554/stream"
          type="textarea"
          :rows="2"
        />
        <div class="form-tip">
          支持 RTSP、RTMP、HTTP、HTTPS 协议的视频流地址
        </div>
      </el-form-item>
      
      <el-form-item label="设备状态" prop="status">
        <el-select 
          v-model="form.status" 
          placeholder="请选择状态" 
          style="width: 100%"
        >
          <el-option label="在线" value="1" />
          <el-option label="离线" value="2" />
          <el-option label="故障" value="3" />
        </el-select>
      </el-form-item>
      
      <el-form-item label="分辨率" prop="resolution">
        <el-select 
          v-model="form.resolution" 
          placeholder="请选择或输入分辨率" 
          style="width: 100%"
          filterable
          allow-create
        >
          <el-option label="1920x1080" value="1920x1080" />
          <el-option label="1280x720" value="1280x720" />
          <el-option label="640x480" value="640x480" />
          <el-option label="1024x768" value="1024x768" />
        </el-select>
      </el-form-item>
      
      <el-form-item label="设备描述" prop="description">
        <el-input 
          v-model="form.description" 
          type="textarea" 
          :rows="3"
          placeholder="请输入设备描述（可选）"
          maxlength="200"
          show-word-limit
        />
      </el-form-item>
    </el-form>
    
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button 
          type="primary" 
          @click="handleSubmit" 
          :loading="loading"
        >
          {{ isEdit ? '更新' : '创建' }}
        </el-button>
        <el-button 
          v-if="!isEdit" 
          type="success" 
          @click="handleSubmitAndAdd" 
          :loading="loading"
        >
          创建并添加到监控
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'

// Props
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  deviceData: {
    type: Object,
    default: null
  },
  groups: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  }
})

// Emits
const emit = defineEmits([
  'update:modelValue',
  'submit',
  'submit-and-add'
])

// 响应式数据
const formRef = ref()
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const isEdit = computed(() => {
  return props.deviceData && props.deviceData.ID
})

// 表单数据
const form = reactive({
  ID: null,
  deviceName: '',
  groupId: '',
  streamUrl: '',
  status: '1',
  resolution: '',
  description: ''
})

// URL验证函数
const validateUrl = (rule, value, callback) => {
  if (!value) {
    callback(new Error('请输入视频流地址'))
    return
  }
  
  const urlPattern = /^(rtsp|rtmp|http|https):\/\/.+/i
  if (!urlPattern.test(value)) {
    callback(new Error('请输入有效的视频流地址（支持RTSP、RTMP、HTTP、HTTPS协议）'))
    return
  }
  
  callback()
}

// 设备名称验证函数
const validateDeviceName = (rule, value, callback) => {
  if (!value || !value.trim()) {
    callback(new Error('请输入设备名称'))
    return
  }
  
  if (value.length < 2) {
    callback(new Error('设备名称至少2个字符'))
    return
  }
  
  if (value.length > 50) {
    callback(new Error('设备名称不能超过50个字符'))
    return
  }
  
  callback()
}

// 表单验证规则
const rules = {
  deviceName: [
    { validator: validateDeviceName, trigger: 'blur' }
  ],
  groupId: [
    { required: true, message: '请选择设备分组', trigger: 'change' }
  ],
  streamUrl: [
    { validator: validateUrl, trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择设备状态', trigger: 'change' }
  ]
}

// 监听设备数据变化，初始化表单
watch(() => props.deviceData, (newData) => {
  if (newData) {
    // 编辑模式，填充数据
    Object.assign(form, {
      ID: newData.ID || null,
      deviceName: newData.deviceName || '',
      groupId: newData.groupId || '',
      streamUrl: newData.streamUrl || '',
      status: newData.status || '1',
      resolution: newData.resolution || '',
      description: newData.description || ''
    })
  } else {
    // 新增模式，重置表单
    resetForm()
  }
}, { immediate: true })

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    ID: null,
    deviceName: '',
    groupId: '',
    streamUrl: '',
    status: '1',
    resolution: '',
    description: ''
  })
  
  nextTick(() => {
    formRef.value?.clearValidate()
  })
}

// 关闭弹窗
const handleClose = () => {
  visible.value = false
  resetForm()
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    
    const submitData = {
      ...form,
      deviceName: form.deviceName.trim(),
      streamUrl: form.streamUrl.trim(),
      description: form.description.trim()
    }
    
    emit('submit', submitData)
  } catch (error) {
    ElMessage.error('请检查表单填写是否正确')
  }
}

// 提交并添加到监控
const handleSubmitAndAdd = async () => {
  try {
    await formRef.value.validate()
    
    const submitData = {
      ...form,
      deviceName: form.deviceName.trim(),
      streamUrl: form.streamUrl.trim(),
      description: form.description.trim()
    }
    
    emit('submit-and-add', submitData)
  } catch (error) {
    ElMessage.error('请检查表单填写是否正确')
  }
}

// 暴露方法给父组件
defineExpose({
  resetForm
})
</script>

<style scoped>
.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
  line-height: 1.4;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

:deep(.el-form-item__label) {
  font-weight: 500;
}

:deep(.el-textarea__inner) {
  resize: vertical;
}

:deep(.el-select) {
  width: 100%;
}
</style>