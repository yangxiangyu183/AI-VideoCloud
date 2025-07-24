
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="摄像头点位:" prop="deviceName">
    <el-input v-model="formData.deviceName" :clearable="true" placeholder="请输入摄像头点位" />
</el-form-item>
        <el-form-item label="关联分组id:" prop="groupId">
    <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入关联分组id" />
</el-form-item>
        <el-form-item label="视频流地址:" prop="streamUrl">
    <el-input v-model="formData.streamUrl" :clearable="true" placeholder="请输入视频流地址" />
</el-form-item>
        <el-form-item label="设备状态：1-在线 2-离线 3-故障:" prop="status">
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="分辨率:" prop="resolution">
    <el-input v-model="formData.resolution" :clearable="true" placeholder="请输入分辨率" />
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
  createMonitorDevice,
  updateMonitorDevice,
  findMonitorDevice
} from '@/api/device/monitorDevice'

defineOptions({
    name: 'MonitorDeviceForm'
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
const formData = ref({
            deviceName: '',
            groupId: undefined,
            streamUrl: '',
            status: false,
            resolution: '',
        })
// 验证规则
const rule = reactive({
               deviceName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               groupId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               streamUrl : [{
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
      const res = await findMonitorDevice({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createMonitorDevice(formData.value)
               break
             case 'update':
               res = await updateMonitorDevice(formData.value)
               break
             default:
               res = await createMonitorDevice(formData.value)
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
