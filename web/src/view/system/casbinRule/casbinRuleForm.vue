
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ptype字段:" prop="ptype">
    <el-input v-model="formData.ptype" :clearable="true" placeholder="请输入ptype字段" />
</el-form-item>
        <el-form-item label="v0字段:" prop="v0">
    <el-input v-model="formData.v0" :clearable="true" placeholder="请输入v0字段" />
</el-form-item>
        <el-form-item label="v1字段:" prop="v1">
    <el-input v-model="formData.v1" :clearable="true" placeholder="请输入v1字段" />
</el-form-item>
        <el-form-item label="v2字段:" prop="v2">
    <el-input v-model="formData.v2" :clearable="true" placeholder="请输入v2字段" />
</el-form-item>
        <el-form-item label="v3字段:" prop="v3">
    <el-input v-model="formData.v3" :clearable="true" placeholder="请输入v3字段" />
</el-form-item>
        <el-form-item label="v4字段:" prop="v4">
    <el-input v-model="formData.v4" :clearable="true" placeholder="请输入v4字段" />
</el-form-item>
        <el-form-item label="v5字段:" prop="v5">
    <el-input v-model="formData.v5" :clearable="true" placeholder="请输入v5字段" />
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
  createCasbinRule,
  updateCasbinRule,
  findCasbinRule
} from '@/api/system/casbinRule'

defineOptions({
    name: 'CasbinRuleForm'
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
            ptype: '',
            v0: '',
            v1: '',
            v2: '',
            v3: '',
            v4: '',
            v5: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCasbinRule({ ID: route.query.id })
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
               res = await createCasbinRule(formData.value)
               break
             case 'update':
               res = await updateCasbinRule(formData.value)
               break
             default:
               res = await createCasbinRule(formData.value)
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
