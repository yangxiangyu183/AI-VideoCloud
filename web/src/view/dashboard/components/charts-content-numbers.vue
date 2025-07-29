<!--
    本组件参考 arco-pro 的实现 将 ts 改为 js 写法
    https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/views/dashboard/workplace/components/content-chart.vue
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div class="chart-wrapper">
    <Chart :height="height" :option="chartOption" />
  </div>
</template>

<script setup>
  import Chart from '@/components/charts/index.vue'
  import useChartOption from '@/hooks/charts'
  import { graphic } from 'echarts'
  import { computed, ref, watch, nextTick } from 'vue'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { config } = storeToRefs(appStore)
  const props = defineProps({
    height: {
      type: String,
      default: '128px'
    },
    timeRange: {
      type: String,
      default: '24h'
    }
  })
  const dotColor = computed(() => {
    return appStore.isDark ? '#333' : '#E5E8EF'
  })
  const graphicFactory = (side) => {
    return {
      type: 'text',
      bottom: '8',
      ...side,
      style: {
        text: '',
        textAlign: 'center',
        fill: '#4E5969',
        fontSize: 12
      }
    }
  }
  // 不同时间范围的数据配置
  const timeRangeConfig = {
    '24h': {
      xAxis: ['00:00', '03:00', '06:00', '09:00', '12:00', '15:00', '18:00', '21:00'],
      data: [15, 18, 25, 35, 42, 38, 45, 32]
    },
    '7d': {
      xAxis: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
      data: [32, 28, 45, 52, 48, 65, 58]
    },
    '30d': {
      xAxis: ['2024-2', '2024-3', '2024-4', '2024-5', '2024-6', '2024-7'],
      data: [20, 25, 35, 42, 65, 78]
    }
  }

  // 响应式数据
  const xAxis = computed(() => timeRangeConfig[props.timeRange]?.xAxis || timeRangeConfig['24h'].xAxis)
  const chartsData = computed(() => timeRangeConfig[props.timeRange]?.data || timeRangeConfig['24h'].data)
  const graphicElements = ref([
    graphicFactory({ left: '5%' }),
    graphicFactory({ right: 0 })
  ])
  const { chartOption } = useChartOption(() => {
    return {
      backgroundColor: '#ffffff',
      grid: {
        left: '40',
        right: '0',
        top: '10',
        bottom: '30'
      },
      xAxis: {
        type: 'category',
        offset: 2,
        data: xAxis.value,
        boundaryGap: false,
        axisLabel: {
          color: '#4E5969',
          formatter(value, idx) {
            if (idx === 0) return ''
            if (idx === xAxis.value.length - 1) return ''
            return `${value}`
          }
        },
        axisLine: {
          show: false
        },
        axisTick: {
          show: false
        },
        splitLine: {
          show: true,
          interval: (idx) => {
            if (idx === 0) return false
            if (idx === xAxis.value.length - 1) return false
            return true
          },
          lineStyle: {
            color: dotColor.value
          }
        },
        axisPointer: {
          show: true,
          lineStyle: {
            color: `${config.value.primaryColor}FF`,
            width: 2
          }
        }
      },
      yAxis: {
        type: 'value',
        axisLine: {
          show: false
        },
        axisLabel: {
          formatter(value, idx) {
            if (idx === 0) return value
            return `${value}k`
          }
        },
        splitLine: {
          show: true,
          lineStyle: {
            type: 'dashed',
            color: dotColor.value
          }
        }
      },
      tooltip: {
        trigger: 'axis',
        formatter(params) {
          const [firstElement] = params
          const timeRangeText = {
            '24h': '小时',
            '7d': '天',
            '30d': '月'
          }
          const unit = timeRangeText[props.timeRange] || '时间段'
          return `<div>
            <p class="tooltip-title">${firstElement.axisValueLabel}</p>
            <div class="content-panel"><span>预警数量</span><span class="tooltip-value">${firstElement.value}k</span></div>
          </div>`
        },
        className: 'echarts-tooltip-diy'
      },
      graphic: {
        elements: graphicElements.value
      },
      animation: true,
      animationDuration: 1200,
      animationEasing: 'cubicInOut',
      animationDurationUpdate: 800,
      animationEasingUpdate: 'cubicInOut',
      series: [
        {
          data: chartsData.value,
          type: 'line',
          smooth: true,
          // symbol: 'circle',
          symbolSize: 12,
          emphasis: {
            focus: 'series',
            itemStyle: {
              borderWidth: 2
            }
          },
          lineStyle: {
            width: 3,
            color: new graphic.LinearGradient(0, 0, 1, 0, [
              {
                offset: 0,
                color: `${config.value.primaryColor}80`
              },
              {
                offset: 0.5,
                color: `${config.value.primaryColor}92`
              },
              {
                offset: 1,
                color: `${config.value.primaryColor}FF`
              }
            ])
          },
          showSymbol: false,
          areaStyle: {
            opacity: 0.8,
            color: new graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: `${config.value.primaryColor}20`
              },
              {
                offset: 1,
                color: `${config.value.primaryColor}08`
              }
            ])
          },
          animationDelay: (idx) => idx * 100
        }
      ]
    }
  })

  // 监听时间范围变化，添加动画效果
  watch(() => props.timeRange, (newRange, oldRange) => {
    if (newRange !== oldRange) {
      nextTick(() => {
        // 图表会自动重新渲染，因为数据是响应式的
        // ECharts会自动处理数据变化的动画
      })
    }
  })
</script>

<style scoped lang="scss">
.chart-wrapper {
  width: 100%;
  height: 100%;
  transition: all 0.3s ease;
}
</style>
