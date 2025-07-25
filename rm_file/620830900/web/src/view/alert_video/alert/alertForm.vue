
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="监控视频:" prop="video">
    <el-input v-model="formData.video" :clearable="true" placeholder="请输入监控视频" />
</el-form-item>
        <el-form-item label="摄像头点位:" prop="cameraAddress">
    <el-input v-model="formData.cameraAddress" :clearable="true" placeholder="请输入摄像头点位" />
</el-form-item>
        <el-form-item label="预警类型:" prop="alertType">
    <el-select v-model="formData.alertType" placeholder="请选择预警类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in alert_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="预警时间:" prop="alertTime">
    <el-date-picker v-model="formData.alertTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createAlert,
  updateAlert,
  findAlert
} from '@/api/alert_video/alert'

defineOptions({
    name: 'AlertForm'
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
const formData = ref({
            video: '',
            cameraAddress: '',
            alertType: '',
            alertTime: new Date(),
        })
// 验证规则
const rule = reactive({
               video : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cameraAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               alertType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               alertTime : [{
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
      const res = await findAlert({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    alert_typeOptions.value = await getDictFunc('alert_type')
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
               res = await createAlert(formData.value)
               break
             case 'update':
               res = await updateAlert(formData.value)
               break
             default:
               res = await createAlert(formData.value)
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
