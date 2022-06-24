<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">早安，{{ userStore.userInfo.nickName }}，请开始一天的工作吧</div>
          <div class="gva-top-card-left-dot">上海市嘉定区，今日{{ this.weather['weather'] }}，气温{{ this.weather['temp_low'] }}℃ - {{ this.weather['temp_high'] }}℃，当前湿度{{ this.weather['humidity'] }}，{{ this.weather['wind'] }} {{ this.weather['winp'] }}。</div>
          <div class="gva-top-card-left-rows">
            <el-row v-auth="888">
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <sort />
                  </el-icon>
                  今日流量 (1231231)
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <avatar />
                  </el-icon>
                  总用户数 (24001)
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <comment />
                  </el-icon>
                  好评率 (99%)
                </div>
              </el-col>
            </el-row>
          </div>
          <div>
            <div class="gva-top-card-left-item">
              仓库地址：
              <a
                style="color:#409EFF"
                target="view_window"
                href="https://github.com/Ted-0711/EpidemicInfoManager/"
              >https://github.com/Ted-0711/EpidemicInfoManager</a>
            </div>
            <div class="gva-top-card-left-item">
              个人主页：
              <a
                style="color:#409EFF"
                target="view_window"
                href="http://www.tedxu.xyz/"
              >http://www.tedxu.xyz</a>
            </div>
          </div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt>
      </div>
    </div>
    <div class="gva-card-box">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>快捷入口</span>
          </div>
        </template>
        <el-row :gutter="20">
          <el-col
            v-for="(card, key) in this.toolCards"
            :key="key"
            :span="4"
            :xs="8"
            class="quick-entrance-items"
            @click="toTarget(card.name)"
          >
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
    <!-- <div class="quick-entrance-title"></div> -->
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>数据统计</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <echarts-line />
            </el-col>
            <el-col :xs="24" :sm="6">
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import echartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
import dashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { getWeather } from '@/api/getWeather.js'

const toolCards = ref([
  {
    label: '用户管理',
    icon: 'monitor',
    name: 'user',
    color: '#ff9c6e',
    bg: 'rgba(255, 156, 110,.3)'
  },
  {
    label: '角色管理',
    icon: 'setting',
    name: 'authority',
    color: '#69c0ff',
    bg: 'rgba(105, 192, 255,.3)'
  },
  {
    label: '菜单管理',
    icon: 'menu',
    name: 'menu',
    color: '#b37feb',
    bg: 'rgba(179, 127, 235,.3)'
  },
  {
    label: '代码生成器',
    icon: 'cpu',
    name: 'autoCode',
    color: '#ffd666',
    bg: 'rgba(255, 214, 102,.3)'
  },
  {
    label: '表单生成器',
    icon: 'document-checked',
    name: 'formCreate',
    color: '#ff85c0',
    bg: 'rgba(255, 133, 192,.3)'
  },
  {
    label: '关于我们',
    icon: 'user',
    name: 'about',
    color: '#5cdbd3',
    bg: 'rgba(92, 219, 211,.3)'
  }
])

const router = useRouter()

const toTarget = (name) => {
  router.push({ name })
}

const userStore = useUserStore()
</script>
<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      weather: {
        aqi: "50",
        cityid: "101020500",
        citynm: "嘉定",
        cityno: "jiading",
        days: "2022-04-27",
        humi_high: "0",
        humi_low: "0",
        humidity: "37%",
        temp_curr: "19",
        temp_high: "15",
        temp_low: "12",
        temperature: "12℃/12℃",
        temperature_curr: "19℃",
        weaid: "38",
        weather: "晴",
        weather_curr: "多云",
        weather_icon: "http://api.k780.com/upload/weather/d/1.gif",
        weather_icon1: "",
        weather_iconid: "1",
        weatid: "2",
        weatid1: "",
        week: "星期三",
        wind: "东北风",
        windid: "1",
        winp: "2级",
        winpid: "2",
      },
      toolCards: [
        {
          label: '用户管理',
          icon: 'monitor',
          name: 'user',
          color: '#ff9c6e',
          bg: 'rgba(255, 156, 110,.3)'
        },
        {
          label: '角色管理',
          icon: 'setting',
          name: 'authority',
          color: '#69c0ff',
          bg: 'rgba(105, 192, 255,.3)'
        },
        {
          label: '菜单管理',
          icon: 'menu',
          name: 'menu',
          color: '#b37feb',
          bg: 'rgba(179, 127, 235,.3)'
        },
        {
          label: '代码生成器',
          icon: 'cpu',
          name: 'autoCode',
          color: '#ffd666',
          bg: 'rgba(255, 214, 102,.3)'
        },
        {
          label: '表单生成器',
          icon: 'document-checked',
          name: 'formCreate',
          color: '#ff85c0',
          bg: 'rgba(255, 133, 192,.3)'
        },
        {
          label: '关于我们',
          icon: 'user',
          name: 'about',
          color: '#5cdbd3',
          bg: 'rgba(92, 219, 211,.3)'
        }
      ],
    }
  },
  created() {
    getWeather().then((res) => {
      this.weather = JSON.parse(res['data']['list'])['result']
      console.log(this.weather)
    })
    const userStore = useUserStore()
    userStore.GetUserInfo().then((res) => {
      let authorityId = res['data']['userInfo']['authority']['authorityId']
      console.log(authorityId)
      if (authorityId === '888') {
        this.toolCards = [
          {
            label: '用户管理',
            icon: 'monitor',
            name: 'user',
            color: '#ff9c6e',
            bg: 'rgba(255, 156, 110,.3)'
          },
          {
            label: '角色管理',
            icon: 'setting',
            name: 'authority',
            color: '#69c0ff',
            bg: 'rgba(105, 192, 255,.3)'
          },
          {
            label: '菜单管理',
            icon: 'menu',
            name: 'menu',
            color: '#b37feb',
            bg: 'rgba(179, 127, 235,.3)'
          },
          {
            label: '代码生成器',
            icon: 'cpu',
            name: 'autoCode',
            color: '#ffd666',
            bg: 'rgba(255, 214, 102,.3)'
          },
          {
            label: '表单生成器',
            icon: 'document-checked',
            name: 'formCreate',
            color: '#ff85c0',
            bg: 'rgba(255, 133, 192,.3)'
          },
          {
            label: '关于我们',
            icon: 'user',
            name: 'about',
            color: '#5cdbd3',
            bg: 'rgba(92, 219, 211,.3)'
          }
        ]
      }
      else if (authorityId === '9990') {
        this.toolCards = [
          {
            label: '每日打卡',
            icon: 'alarm-clock',
            name: 'clock_in_report',
            color: '#ff9c6e',
            bg: 'rgba(255, 156, 110,.3)'
          },
          {
            label: '疫苗接种',
            icon: 'arrow-down-bold',
            name: 'vaccine_report',
            color: '#69c0ff',
            bg: 'rgba(105, 192, 255,.3)'
          },
          {
            label: '核酸检测',
            icon: 'warning-filled',
            name: 'test_swab_report',
            color: '#b37feb',
            bg: 'rgba(179, 127, 235,.3)'
          },
          {
            label: '隔离记录',
            icon: 'key',
            name: 'quarantine_report',
            color: '#ffd666',
            bg: 'rgba(255, 214, 102,.3)'
          },
          {
            label: '人口流动',
            icon: 'bicycle',
            name: 'migration_report',
            color: '#ff85c0',
            bg: 'rgba(255, 133, 192,.3)'
          },
          {
            label: '问卷信息',
            icon: 'question-filled',
            name: 'qtnInfoStu',
            color: '#5cdbd3',
            bg: 'rgba(92, 219, 211,.3)'
          }
        ]
      }
      else if (authorityId === '9991') {
        this.toolCards = [
          {
            label: '打卡记录',
            icon: 'alarm-clock',
            name: 'clock_in_query',
            color: '#ff9c6e',
            bg: 'rgba(255, 156, 110,.3)'
          },
          {
            label: '疫苗接种',
            icon: 'arrow-down-bold',
            name: 'vaccine_query',
            color: '#69c0ff',
            bg: 'rgba(105, 192, 255,.3)'
          },
          {
            label: '核酸检测',
            icon: 'warning-filled',
            name: 'test_swab_query',
            color: '#b37feb',
            bg: 'rgba(179, 127, 235,.3)'
          },
          {
            label: '隔离记录',
            icon: 'key',
            name: 'quarantine_query',
            color: '#ffd666',
            bg: 'rgba(255, 214, 102,.3)'
          },
          {
            label: '人口流动',
            icon: 'bicycle',
            name: 'migration_query',
            color: '#ff85c0',
            bg: 'rgba(255, 133, 192,.3)'
          },
          {
            label: '问卷信息',
            icon: 'question-filled',
            name: 'qtnInfo',
            color: '#5cdbd3',
            bg: 'rgba(92, 219, 211,.3)'
          }
        ]
      }
    })
  }
}
</script>

<style lang="scss" scoped>
@mixin flex-center {
    display: flex;
    align-items: center;
}
.page {
    background: #f0f2f5;
    padding: 0;
    .gva-card-box{
      padding: 12px 16px;
      &+.gva-card-box{
        padding-top: 0px;
      }
    }
    .gva-card {
      box-sizing: border-box;
        background-color: #fff;
        border-radius: 2px;
        height: auto;
        padding: 26px 30px;
        overflow: hidden;
        box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
    }
    .gva-top-card {
        height: 260px;
        @include flex-center;
        justify-content: space-between;
        color: #777;
        &-left {
          height: 100%;
          display: flex;
          flex-direction: column;
            &-title {
                font-size: 22px;
                color: #343844;
            }
            &-dot {
                font-size: 14px;
                color: #6B7687;
                margin-top: 24px;
            }
            &-rows {
                // margin-top: 15px;
                margin-top: 18px;
                color: #6B7687;
                width: 600px;
                align-items: center;
            }
            &-item{
              +.gva-top-card-left-item{
                margin-top: 24px;
              }
              margin-top: 14px;
            }
        }
        &-right {
            height: 600px;
            width: 600px;
            margin-top: 28px;
        }
    }
     ::v-deep(.el-card__header){
          padding:0;
          border-bottom: none;
        }
        .card-header{
          padding-bottom: 20px;
          border-bottom: 1px solid #e8e8e8;
        }
    .quick-entrance-title {
        height: 30px;
        font-size: 22px;
        color: #333;
        width: 100%;
        border-bottom: 1px solid #eee;
    }
    .quick-entrance-items {
        @include flex-center;
        justify-content: center;
        text-align: center;
        color: #333;
        .quick-entrance-item {
          padding: 16px 28px;
          margin-top: -16px;
          margin-bottom: -16px;
          border-radius: 4px;
          transition: all 0.2s;
          &:hover{
            box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
          }
            cursor: pointer;
            height: auto;
            text-align: center;
            // align-items: center;
            &-icon {
                width: 50px;
                height: 50px !important;
                border-radius: 8px;
                @include flex-center;
                justify-content: center;
                margin: 0 auto;
                i {
                    font-size: 24px;
                }
            }
            p {
                margin-top: 10px;
            }
        }
    }
    .echart-box{
      padding: 14px;
    }
}
.dasboard-icon {
    font-size: 20px;
    color: rgb(85, 160, 248);
    width: 30px;
    height: 30px;
    margin-right: 10px;
    @include flex-center;
}
.flex-center {
    @include flex-center;
}

//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
    .gva-card {
        padding: 20px 10px !important;
        .gva-top-card {
            height: auto;
            &-left {
                &-title {
                    font-size: 20px !important;
                }
                &-rows {
                    margin-top: 15px;
                    align-items: center;
                }
            }
            &-right {
                display: none;
            }
        }
        .gva-middle-card {
            &-item {
                line-height: 20px;
            }
        }
        .dasboard-icon {
            font-size: 18px;
        }
    }
}
</style>
