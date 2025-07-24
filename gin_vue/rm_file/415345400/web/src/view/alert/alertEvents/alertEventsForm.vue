
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联任务id:" prop="taskId">
    <el-input v-model.number="formData.taskId" :clearable="true" placeholder="请输入关联任务id" />
</el-form-item>
        <el-form-item label="关联摄像头id:" prop="cameraId">
    <el-input v-model.number="formData.cameraId" :clearable="true" placeholder="请输入关联摄像头id" />
</el-form-item>
        <el-form-item label="告警类型:" prop="alertType">
    <el-select v-model="formData.alertType" placeholder="请选择告警类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in alert_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="警告级别:" prop="servenrity">
    <el-select v-model="formData.servenrity" placeholder="请选择警告级别" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in servenrityOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="截图路径:" prop="imagePath">
    <el-input v-model="formData.imagePath" :clearable="true" placeholder="请输入截图路径" />
</el-form-item>
        <el-form-item label="视频片段路径:" prop="videoPath">
    <el-input v-model="formData.videoPath" :clearable="true" placeholder="请输入视频片段路径" />
</el-form-item>
        <el-form-item label="是否已读:" prop="isRead">
    <el-select v-model="formData.isRead" placeholder="请选择是否已读" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in is_readOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="是否已经处理:" prop="isProcessed">
    <el-select v-model="formData.isProcessed" placeholder="请选择是否已经处理" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in is_processedOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="处理用户id:" prop="processedBy">
    <el-input v-model.number="formData.processedBy" :clearable="true" placeholder="请输入处理用户id" />
</el-form-item>
        <el-form-item label="处理时间:" prop="processedTime">
    <el-date-picker v-model="formData.processedTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="处理结果:" prop="processedResult">
    <el-input v-model="formData.processedResult" :clearable="true" placeholder="请输入处理结果" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAlertEvents,
  updateAlertEvents,
  findAlertEvents
} from '@/api/alert/alertEvents'

defineOptions({
    name: 'AlertEventsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const alert_typeOptions = ref([])
const servenrityOptions = ref([])
const is_readOptions = ref([])
const is_processedOptions = ref([])
const formData = ref({
            taskId: undefined,
            cameraId: undefined,
            alertType: '',
            servenrity: '',
            imagePath: '',
            videoPath: '',
            isRead: '',
            isProcessed: '',
            processedBy: undefined,
            processedTime: new Date(),
            processedResult: '',
        })
// 验证规则
const rule = reactive({
               taskId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cameraId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               alertType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               servenrity : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               imagePath : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               videoPath : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isRead : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isProcessed : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               processedBy : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               processedTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               processedResult : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAlertEvents({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    alert_typeOptions.value = await getDictFunc('alert_type')
    servenrityOptions.value = await getDictFunc('servenrity')
    is_readOptions.value = await getDictFunc('is_read')
    is_processedOptions.value = await getDictFunc('is_processed')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createAlertEvents(formData.value)
               break
             case 'update':
               res = await updateAlertEvents(formData.value)
               break
             default:
               res = await createAlertEvents(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
