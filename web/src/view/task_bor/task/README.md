# 任务管理模块修改说明

## 修改内容

### 1. 监控点位服务集成
- **修改文件**: `task.vue`
- **修改内容**: 将监控点位选择从静态数据改为调用视频接入中的监控设备API
- **API接口**: `getMonitorDeviceList` (来自 `/api/device/monitorDevice`)

### 2. 主要变更

#### 2.1 导入监控设备API
```javascript
// 引入监控设备API
import { getMonitorDeviceList } from '@/api/device/monitorDevice'
```

#### 2.2 动态获取监控点位数据
```javascript
// 监控点位选项数据，从API动态获取
const cameraOptions = ref([])

// 获取监控设备列表
const getCameraOptions = async () => {
  try {
    const res = await getMonitorDeviceList({ page: 1, pageSize: 1000 })
    if (res.code === 0 && res.data && res.data.list) {
      cameraOptions.value = res.data.list.map(item => ({
        label: item.deviceName || item.deviceLocation || `设备${item.id}`,
        value: item.deviceName || item.deviceLocation || `设备${item.id}`,
        deviceId: item.id,
        deviceInfo: item
      }))
    }
  } catch (error) {
    // 错误处理，使用默认数据
  }
}
```

#### 2.3 UI改进
- 添加了设备ID显示
- 添加了刷新按钮，用户可以手动刷新监控点位列表
- 改进了应用场景选项，与图片中的模块环境保持一致
- 添加了用户反馈消息

#### 2.4 错误处理
- API调用失败时自动回退到默认数据
- 提供用户友好的错误提示信息
- 控制台错误日志记录

### 3. 功能特性

1. **自动加载**: 页面初始化时自动获取监控设备列表
2. **手动刷新**: 用户可以点击刷新按钮重新获取设备列表
3. **容错机制**: API失败时使用默认数据，确保功能可用
4. **用户反馈**: 提供加载状态和结果反馈
5. **设备信息**: 显示设备ID和详细信息

### 4. 使用说明

1. 打开任务管理页面
2. 点击"添加任务"按钮
3. 在"监控点位"字段中选择设备
4. 如需刷新设备列表，点击选择框旁边的刷新按钮

### 5. 技术细节

- **API端点**: `/monitorDevice/getMonitorDeviceList`
- **数据映射**: 设备名称/位置 -> 选项标签和值
- **分页**: 获取前1000个设备（可根据需要调整）
- **响应式**: 使用Vue 3的ref进行状态管理

### 6. 后续优化建议

1. 添加设备搜索功能
2. 支持设备分组显示
3. 添加设备状态指示器
4. 实现设备详情预览
5. 支持批量选择设备