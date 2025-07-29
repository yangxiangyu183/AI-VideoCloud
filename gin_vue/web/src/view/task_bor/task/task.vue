<template>
  <div>
    <!-- 顶部tab切换和筛选区 -->
    <div class="gva-search-box">
      <el-tabs v-model="activeTab" @tab-click="onTabChange">
        <el-tab-pane label="目标检测" name="target" />
        <el-tab-pane label="姿势检测" name="pose" />
        <el-tab-pane label="人脸识别" name="face" />
        <el-tab-pane label="目标统计" name="stat" />
        <el-tab-pane label="文字检测" name="text" />
      </el-tabs>
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item>
          <el-input v-model="searchInfo.taskName" placeholder="请输入任务名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-input v-model="searchInfo.searchContent" placeholder="请输入搜索内容" clearable style="width: 200px" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="searchInfo.warnLevel" placeholder="请选择告警级别" clearable style="width: 150px">
            <el-option label="高" value="高" />
            <el-option label="中" value="中" />
            <el-option label="低" value="低" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-date-picker v-model="searchInfo.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" clearable style="width: 250px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="plus" @click="openDialog()">+ 添加任务</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- 表格区 -->
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="taskName"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="任务名称" prop="taskName" width="120" />
        <el-table-column align="left" label="摄像头点位" prop="cameraInterface" width="120" />
        <el-table-column align="left" label="任务状态" prop="taskStatus" width="100">
          <template #default="scope">
            <span v-if="scope.row.taskStatus === '启动' || scope.row.taskStatus === '1'" class="status-badge status-started">启动</span>
            <span v-else class="status-badge status-stopped">{{ scope.row.taskStatus }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="告警级别" prop="warnLevel" width="100">
          <template #default="scope">
            <span v-if="scope.row.warnLevel === '高' || scope.row.warnLevel === '1'" class="warn-level-badge warn-level-高">高</span>
            <span v-else-if="scope.row.warnLevel === '中' || scope.row.warnLevel === '2'" class="warn-level-badge warn-level-中">中</span>
            <span v-else-if="scope.row.warnLevel === '低' || scope.row.warnLevel === '3'" class="warn-level-badge warn-level-低">低</span>
            <span v-else class="warn-level-badge warn-level-中">{{ scope.row.warnLevel }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" prop="createdAt" width="160">
          <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateTaskFunc(scope.row)">编辑</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <!-- 新增/编辑弹窗，两列布局 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'添加预警任务':'编辑预警任务'}}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">提交</el-button>
            <el-button @click="closeDialog">取消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="100px" class="form-two-col">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="任务名称" prop="taskName">
              <el-input v-model="formData.taskName" clearable placeholder="请输入任务名称" />
            </el-form-item>
            <el-form-item label="任务描述" prop="taskDescription">
              <el-input v-model="formData.taskDescription" clearable placeholder="请输入任务描述" />
            </el-form-item>
            <el-form-item label="任务优先级" prop="taskPriority">
              <el-select v-model="formData.taskPriority" placeholder="请选择任务优先级" filterable clearable>
                <el-option label="1" value="1" />
                <el-option label="2" value="2" />
                <el-option label="3" value="3" />
              </el-select>
            </el-form-item>
            <el-form-item label="监控点位" prop="monitorPoints">
              <el-select v-model="formData.monitorPoints" placeholder="请选择监控点位" filterable clearable>
                <el-option v-for="item in cameraOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="应用场景" prop="applicationScenario">
              <el-select v-model="formData.applicationScenario" placeholder="请选择应用场景" filterable clearable>
                <el-option v-for="item in ApplicationscenarioOptions" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="检测日期" prop="dateRange">
              <el-radio-group v-model="formData.dateType">
                <el-radio label="每天">每天</el-radio>
              </el-radio-group>
              <el-date-picker v-model="formData.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" clearable style="width: 100%; margin-top: 8px;" />
            </el-form-item>
            <el-form-item label="检测轮询时长" prop="pollingDuration">
              <el-radio-group v-model="formData.pollingType">
                <el-radio label="实时">实时</el-radio>
              </el-radio-group>
              <div style="margin-top: 8px;">
                <el-input v-model="formData.pollingDuration" clearable placeholder="请输入轮询时长" style="width: 60%" />
                <el-select v-model="formData.pollingUnit" placeholder="时长单位" style="width: 38%; margin-left: 2%" clearable>
                  <el-option label="秒" value="秒" />
                  <el-option label="分钟" value="分钟" />
                  <el-option label="小时" value="小时" />
                </el-select>
              </div>
            </el-form-item>
            <el-form-item label="任务状态" prop="taskStatus">
              <el-radio-group v-model="formData.taskStatus">
                <el-radio label="启动">启动</el-radio>
                <el-radio label="停用">停用</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="告警级别" prop="warnLevel">
              <el-select v-model="formData.warnLevel" placeholder="请选择告警级别" filterable clearable>
                <el-option label="高" value="高" />
                <el-option label="中" value="中" />
                <el-option label="低" value="低" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div style="text-align: right; margin-top: 20px;">
        <el-button @click="closeDialog">重置</el-button>
        <el-button type="primary" :loading="btnLoading" @click="enterDialog">提交</el-button>
      </div>
    </el-drawer>
    <!-- 详情弹窗，左字段右画面 -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="监测任务详情">
      <el-row>
        <el-col :span="14">
          <el-descriptions :column="1" border>
            <el-descriptions-item label="任务名称">{{ detailForm.taskName }}</el-descriptions-item>
            <el-descriptions-item label="任务描述">{{ detailForm.taskDescription || '无' }}</el-descriptions-item>
            <el-descriptions-item label="监控点位">{{ detailForm.cameraInterface }}</el-descriptions-item>
            <el-descriptions-item label="监控分组">{{ detailForm.monitorGroup || '模拟环境1' }}</el-descriptions-item>
            <el-descriptions-item label="任务状态">{{ detailForm.taskStatus }}</el-descriptions-item>
            <el-descriptions-item label="告警级别">{{ detailForm.warnLevel }}</el-descriptions-item>
            <el-descriptions-item label="任务启动时间">{{ detailForm.createdAt }}</el-descriptions-item>
            <el-descriptions-item label="任务停用时间">{{ detailForm.stoppedAt || '无' }}</el-descriptions-item>
          </el-descriptions>
        </el-col>
        <el-col :span="10">
          <div style="background: #000; height: 240px; display: flex; align-items: center; justify-content: center;">
            <el-icon style="font-size: 40px; color: #fff; animation: spin 1s linear infinite;">
              <Loading />
            </el-icon>
          </div>
          <div style="text-align: center; color: #888; margin-top: 10px;">当前监控摄像头画面</div>
        </el-col>
      </el-row>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createTask,
  deleteTask,
  deleteTaskByIds,
  updateTask,
  findTask,
  getTaskList
} from '@/api/task_bor/task'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

defineOptions({
    name: 'Task'
})
// 按钮权限实例化
const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const TaskpriorityOptions = ref([])
const servenrityOptions = ref([])
const ApplicationscenarioOptions = ref([])
const TaskstatusOptions = ref([])
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
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  try {
    const table = await getTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  } catch (error) {
    // 如果API调用失败，使用模拟数据
    console.log('API调用失败，使用模拟数据')
    tableData.value = [
      {
        taskName: '监测任务1',
        cameraInterface: '1',
        taskStatus: '1',
        warnLevel: '1',
        createdAt: '2025-07-22 09:07'
      },
      {
        taskName: '监测任务2',
        cameraInterface: '惠州市惠城区\n西湖',
        taskStatus: '启动',
        warnLevel: '低',
        createdAt: '2025-07-24 09:16'
      }
    ]
    total.value = 2
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    TaskpriorityOptions.value = await getDictFunc('Taskpriority')
    servenrityOptions.value = await getDictFunc('servenrity')
    ApplicationscenarioOptions.value = await getDictFunc('Applicationscenario')
    TaskstatusOptions.value = await getDictFunc('Taskstatus')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTaskFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const taskNames = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          taskNames.push(item.taskName)
        })
      const res = await deleteTaskByIds({ taskNames })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === taskNames.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTaskFunc = async(row) => {
    const res = await findTask({ taskName: row.taskName })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}

// 删除行
const deleteTaskFunc = async (row) => {
    const res = await deleteTask({ taskName: row.taskName })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        dateRange: [],
        dateType: '每天',
        pollingDuration: '',
        pollingUnit: '',
        pollingType: '实时',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTask({ taskName: row.taskName })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}

// 新增：tab切换、筛选项、表单项等相关数据
const activeTab = ref('target')
const onTabChange = () => { getTableData() }
const taskTypeOptions = [
  { label: '目标检测', value: 'target' },
  { label: '姿势检测', value: 'pose' },
  { label: '人脸识别', value: 'face' },
  { label: '目标统计', value: 'stat' },
  { label: '文字检测', value: 'text' }
]
const cameraOptions = [
  { label: '惠州市惠城区西湖', value: '惠州市惠城区西湖' },
  { label: '新沂新悦广场', value: '新沂新悦广场' },
  { label: '杭州市西湖', value: '杭州市西湖' },
  { label: '山东泰安泰山', value: '山东泰安泰山' },
  { label: '桂林阳桥', value: '桂林阳桥' },
  { label: '深圳光明电脑汇', value: '深圳光明电脑汇' },
  { label: '淄博市沂源县', value: '淄博市沂源县' }
]
// 表单项补充
formData.value = {
  ...formData.value,
  dateRange: [],
  dateType: '每天',
  pollingDuration: '',
  pollingUnit: '',
  pollingType: '实时',
}

</script>

<style>
.form-two-col .el-form-item {
  margin-bottom: 18px;
}
.form-two-col .el-col-12 {
  padding-right: 10px;
}
@keyframes spin {
  100% { transform: rotate(360deg); }
}

/* 状态和告警级别背景色样式 */
.status-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.status-started {
  color: #67c23a;
  background-color: #f0f9ff;
  border: 1px solid #b3d8ff;
}

.status-stopped {
  color: #909399;
  background-color: #f4f4f5;
  border: 1px solid #d3d4d6;
}

.warn-level-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
  color: #fff;
}

.warn-level-高 {
  background-color: #f56c6c;
}

.warn-level-中 {
  background-color: #e6a23c;
}

.warn-level-低 {
  background-color: #909399;
}
</style>
