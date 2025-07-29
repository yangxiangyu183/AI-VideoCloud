# 告警筛选条件复选框功能设计文档

## 概述

本设计文档描述了告警页面筛选条件复选框功能的技术实现方案。该功能将增强现有的筛选系统，实现实时响应的复选框筛选，提供流畅的用户交互体验。

## 架构

### 整体架构
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   筛选条件组件   │────│   数据筛选引擎    │────│   结果展示组件   │
│  (Filter Panel) │    │ (Filter Engine)  │    │ (Result Display)│
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   状态管理器     │    │   缓存管理器      │    │   分页管理器     │
│ (State Manager) │    │ (Cache Manager)  │    │(Pagination Mgr) │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

### 数据流架构
```
用户操作 → 筛选条件更新 → 数据筛选 → 结果缓存 → UI更新 → 分页重置
```

## 组件和接口

### 1. 筛选条件管理器 (FilterManager)

**职责：** 管理所有筛选条件的状态和逻辑

**接口：**
```javascript
class FilterManager {
  // 更新设备状态筛选
  updateDeviceStatusFilter(statusList: string[]): void
  
  // 更新预警类型筛选  
  updateAlertTypeFilter(typeList: string[]): void
  
  // 更新摄像头点位筛选
  updateCameraAddressFilter(addressList: string[]): void
  
  // 更新时间范围筛选
  updateTimeRangeFilter(rangeList: string[]): void
  
  // 获取当前筛选条件
  getCurrentFilters(): FilterConditions
  
  // 重置所有筛选条件
  resetAllFilters(): void
}
```

### 2. 数据筛选引擎 (DataFilterEngine)

**职责：** 根据筛选条件对数据进行过滤

**接口：**
```javascript
class DataFilterEngine {
  // 应用筛选条件
  applyFilters(data: AlertData[], filters: FilterConditions): AlertData[]
  
  // 按设备状态筛选
  filterByDeviceStatus(data: AlertData[], statusList: string[]): AlertData[]
  
  // 按预警类型筛选
  filterByAlertType(data: AlertData[], typeList: string[]): AlertData[]
  
  // 按摄像头点位筛选
  filterByCameraAddress(data: AlertData[], addressList: string[]): AlertData[]
  
  // 按时间范围筛选
  filterByTimeRange(data: AlertData[], timeRange: DateRange): AlertData[]
}
```

### 3. 复选框组件 (CheckboxGroup)

**职责：** 渲染和管理复选框组的交互

**接口：**
```javascript
interface CheckboxGroupProps {
  options: CheckboxOption[]
  selectedValues: string[]
  onSelectionChange: (selectedValues: string[]) => void
  showSelectAll: boolean
  title: string
}

interface CheckboxOption {
  value: string
  label: string
  checked: boolean
  disabled?: boolean
}
```

## 数据模型

### 筛选条件模型
```javascript
interface FilterConditions {
  deviceStatus: {
    all: boolean
    selected: string[]
  }
  alertType: {
    all: boolean
    selected: string[]
  }
  cameraAddress: {
    all: boolean
    selected: string[]
  }
  timeRange: {
    all: boolean
    selected: string[]
    customRange?: DateRange
  }
}

interface DateRange {
  startTime: Date
  endTime: Date
}
```

### 告警数据模型
```javascript
interface AlertData {
  ID: number
  video: string
  cameraAddress: string
  alertType: string
  alertTime: string
  deviceName: string
  deviceStatus: string
  resolution: string
  streamUrl: string
}
```

### 选项数据模型
```javascript
interface FilterOption {
  value: string
  label: string
  checked: boolean
  color?: string // 用于设备状态的颜色标识
}
```

## 核心算法设计

### 1. 实时筛选算法

```javascript
// 防抖处理的筛选函数
const debouncedFilter = debounce((filters: FilterConditions) => {
  const filteredData = applyAllFilters(originalData, filters)
  updateDisplayData(filteredData)
  updateStatistics(filteredData.length)
  resetPagination()
}, 300)

// 应用所有筛选条件
function applyAllFilters(data: AlertData[], filters: FilterConditions): AlertData[] {
  let result = [...data]
  
  // 按设备状态筛选
  if (!filters.deviceStatus.all && filters.deviceStatus.selected.length > 0) {
    result = result.filter(item => 
      filters.deviceStatus.selected.includes(item.deviceStatus)
    )
  }
  
  // 按预警类型筛选
  if (!filters.alertType.all && filters.alertType.selected.length > 0) {
    result = result.filter(item => 
      filters.alertType.selected.includes(item.alertType)
    )
  }
  
  // 按摄像头点位筛选
  if (!filters.cameraAddress.all && filters.cameraAddress.selected.length > 0) {
    result = result.filter(item => 
      filters.cameraAddress.selected.includes(item.cameraAddress)
    )
  }
  
  // 按时间范围筛选
  if (filters.timeRange.customRange) {
    result = result.filter(item => {
      const itemTime = new Date(item.alertTime)
      return itemTime >= filters.timeRange.customRange.startTime && 
             itemTime <= filters.timeRange.customRange.endTime
    })
  }
  
  return result
}
```

### 2. 全选逻辑算法

```javascript
// 处理全选状态变化
function handleSelectAllChange(category: string, isSelectAll: boolean) {
  const options = getOptionsByCategory(category)
  
  // 更新所有选项的选中状态
  options.forEach(option => {
    option.checked = isSelectAll
  })
  
  // 更新筛选条件
  const selectedValues = isSelectAll ? options.map(opt => opt.value) : []
  updateFilterCondition(category, selectedValues, isSelectAll)
  
  // 触发筛选
  triggerFilter()
}

// 处理单个选项变化
function handleSingleOptionChange(category: string, optionValue: string, isChecked: boolean) {
  const options = getOptionsByCategory(category)
  const targetOption = options.find(opt => opt.value === optionValue)
  
  if (targetOption) {
    targetOption.checked = isChecked
  }
  
  // 检查是否需要更新全选状态
  const checkedCount = options.filter(opt => opt.checked).length
  const isAllSelected = checkedCount === options.length
  const isNoneSelected = checkedCount === 0
  
  updateSelectAllState(category, isAllSelected)
  
  // 更新筛选条件
  const selectedValues = options.filter(opt => opt.checked).map(opt => opt.value)
  updateFilterCondition(category, selectedValues, isAllSelected)
  
  // 触发筛选
  triggerFilter()
}
```

## 错误处理

### 1. 数据异常处理
- 当筛选结果为空时，显示友好的空状态提示
- 当数据加载失败时，显示错误信息并提供重试选项
- 当筛选条件冲突时，以最后操作为准

### 2. 性能优化处理
- 使用防抖技术避免频繁筛选操作
- 实现虚拟滚动处理大量数据
- 缓存筛选结果避免重复计算

### 3. 用户体验优化
- 筛选过程中显示加载状态
- 筛选条件变化时平滑过渡动画
- 保持筛选状态的持久化

## 测试策略

### 1. 单元测试
- 筛选算法的正确性测试
- 全选逻辑的边界条件测试
- 数据模型的验证测试

### 2. 集成测试
- 筛选条件与UI同步测试
- 多筛选条件组合测试
- 分页与筛选结合测试

### 3. 用户体验测试
- 筛选响应时间测试
- 大数据量下的性能测试
- 不同浏览器的兼容性测试

## 性能考虑

### 1. 筛选性能优化
- 使用索引优化筛选算法
- 实现增量筛选减少计算量
- 缓存常用筛选结果

### 2. 渲染性能优化
- 使用虚拟列表处理大量数据
- 实现懒加载减少初始渲染时间
- 优化DOM操作减少重排重绘

### 3. 内存管理
- 及时清理不需要的筛选结果缓存
- 避免内存泄漏的事件监听器
- 合理控制数据缓存大小