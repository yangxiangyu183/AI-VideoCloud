
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务名称:" prop="taskName">
    <el-input v-model="formData.taskName" :clearable="true" placeholder="请输入任务名称" />
</el-form-item>
        <el-form-item label="任务描述:" prop="taskDescription">
    <el-input v-model="formData.taskDescription" :clearable="true" placeholder="请输入任务描述" />
</el-form-item>
        <el-form-item label="任务优先级:" prop="taskPriority">
    <el-select v-model="formData.taskPriority" placeholder="请选择任务优先级" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in TaskpriorityOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="摄像头接口:" prop="cameraInterface">
    <el-input v-model="formData.cameraInterface" :clearable="true" placeholder="请输入摄像头接口" />
</el-form-item>
        <el-form-item label="监控点位:" prop="monitorPoints">
    <el-select v-model="formData.monitorPoints" placeholder="请选择监控点位" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in servenrityOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="应用场景:" prop="applicationScenario">
    <el-select v-model="formData.applicationScenario" placeholder="请选择应用场景" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in ApplicationscenarioOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="任务状态:" prop="taskStatus">
    <el-select v-model="formData.taskStatus" placeholder="请选择任务状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in TaskstatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="警告级别:" prop="warnLevel">
    <el-select v-model="formData.warnLevel" placeholder="请选择警告级别" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in servenrityOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="监测日期:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="监测轮询时长:" prop="deletedAt">
    <el-date-picker v-model="formData.deletedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createTask,
  updateTask,
  findTask
} from '@/api/task_bor/task'

defineOptions({
    name: 'TaskForm'
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
const TaskpriorityOptions = ref([])
const servenrityOptions = ref([])
const TaskstatusOptions = ref([])
const ApplicationscenarioOptions = ref([])
const formData = ref({
            taskName: '',
            taskDescription: '',
            taskPriority: '',
            cameraInterface: '',
            monitorPoints: '',
            applicationScenario: '',
            taskStatus: '',
            warnLevel: '',
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    TaskpriorityOptions.value = await getDictFunc('Taskpriority')
    servenrityOptions.value = await getDictFunc('servenrity')
    TaskstatusOptions.value = await getDictFunc('Taskstatus')
    ApplicationscenarioOptions.value = await getDictFunc('Applicationscenario')
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
               res = await createTask(formData.value)
               break
             case 'update':
               res = await updateTask(formData.value)
               break
             default:
               res = await createTask(formData.value)
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
