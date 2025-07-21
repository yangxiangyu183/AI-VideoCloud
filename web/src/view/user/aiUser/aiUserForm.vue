
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="userName字段:" prop="userName">
    <el-input v-model="formData.userName" :clearable="true" placeholder="请输入userName字段" />
</el-form-item>
        <el-form-item label="password字段:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入password字段" />
</el-form-item>
        <el-form-item label="userEmail字段:" prop="userEmail">
    <el-input v-model="formData.userEmail" :clearable="true" placeholder="请输入userEmail字段" />
</el-form-item>
        <el-form-item label="用户权限:" prop="permissions">
    <el-select v-model="formData.permissions" placeholder="请选择用户权限" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in permissionsOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="用户状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择用户状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in StatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
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
  createAiUser,
  updateAiUser,
  findAiUser
} from '@/api/user/aiUser'

defineOptions({
    name: 'AiUserForm'
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
const permissionsOptions = ref([])
const StatusOptions = ref([])
const formData = ref({
            userName: '',
            password: '',
            userEmail: '',
            permissions: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               userName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               permissions : [{
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
      const res = await findAiUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    permissionsOptions.value = await getDictFunc('permissions')
    StatusOptions.value = await getDictFunc('Status')
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
               res = await createAiUser(formData.value)
               break
             case 'update':
               res = await updateAiUser(formData.value)
               break
             default:
               res = await createAiUser(formData.value)
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
