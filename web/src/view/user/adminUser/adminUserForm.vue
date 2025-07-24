
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户的ID:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户的ID" />
</el-form-item>
        <el-form-item label="管理员的密码:" prop="adminPassword">
    <el-input v-model="formData.adminPassword" :clearable="true" placeholder="请输入管理员的密码" />
</el-form-item>
        <el-form-item label="管理员账号:" prop="adminName">
    <el-input v-model="formData.adminName" :clearable="true" placeholder="请输入管理员账号" />
</el-form-item>
        <el-form-item label="用户的账号状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择用户的账号状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in StatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="用户的使用权限:" prop="permissions">
    <el-select v-model="formData.permissions" placeholder="请选择用户的使用权限" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in permissionsOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="管理员等级（1普通,2超级）:" prop="idAdmin">
    <el-select v-model="formData.idAdmin" placeholder="请选择管理员等级（1普通,2超级）" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in id_adminOptions" :key="key" :label="item.label" :value="item.value" />
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
  createAdminUser,
  updateAdminUser,
  findAdminUser
} from '@/api/user/adminUser'

defineOptions({
    name: 'AdminUserForm'
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
const id_adminOptions = ref([])
const StatusOptions = ref([])
const permissionsOptions = ref([])
const formData = ref({
            userId: undefined,
            adminPassword: '',
            adminName: '',
            status: '',
            permissions: '',
            idAdmin: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               adminPassword : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               adminName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               permissions : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               idAdmin : [{
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
      const res = await findAdminUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    id_adminOptions.value = await getDictFunc('id_admin')
    StatusOptions.value = await getDictFunc('Status')
    permissionsOptions.value = await getDictFunc('permissions')
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
               res = await createAdminUser(formData.value)
               break
             case 'update':
               res = await updateAdminUser(formData.value)
               break
             default:
               res = await createAdminUser(formData.value)
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
