<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" disabled />
        </el-form-item>
        <el-form-item label="打卡日期:">
          <el-date-picker v-model="formData.clock_in_date" type="date" placeholder="选择日期" clearable disabled></el-date-picker>
        </el-form-item>
        <el-form-item label="所在区域:">
          <el-input v-model="formData.area_name" clearable placeholder="请输入" disabled />
        </el-form-item>
        <el-form-item label="体温:">
          <el-input-number v-model="formData.temperature" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="不适症状:">
          <el-input v-model="formData.symptom" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">提交</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
          <!-- <el-button size="mini" type="primary" @click="this.getMapCenter">获取中心点</el-button> -->
        </el-form-item>
      </el-form>
      <div id="map"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Clock_in',
  "BMap": "BMap",
  data() {
    return {
      map: undefined,
    }
  },
  methods: {
    //创建地图实例
    createMap() {
      var map = new BMap.Map("map");
      this.map = map;
      var geolocation = new BMap.Geolocation();
      //调用百度地图api 中的获取当前位置接口
      geolocation.getCurrentPosition(function (r) {
        if (this.getStatus() == BMAP_STATUS_SUCCESS) {
          //获取当前位置经纬度
          let myGeo = new BMap.Geocoder();
          myGeo.getLocation(new BMap.Point(r.point.lng, r.point.lat), function (result) {
            if (result) {
              // console.log(result);
              // 初始化地图,设置中心点坐标和地图级别
              map.centerAndZoom(new BMap.Point(result.point.lng, result.point.lat), 20);
              //开启鼠标滚轮缩放,默认关闭
              map.enableScrollWheelZoom(false)
              //添加缩略图控件
              map.addControl(new BMap.OverviewMapControl({anchor:BMAP_ANCHOR_BOTTOM_RIGHT}));
              //添加缩放平移控件
              map.addControl(new BMap.NavigationControl());
              //添加比例尺控件
              map.addControl(new BMap.ScaleControl());
              //添加地图类型控件
              map.addControl(new BMap.MapTypeControl());
              // 添加城市列表控件
              map.addControl(new BMap.CityListControl({offset: new BMap.Size(70, 15)}));
              //设置地图标记点的位置（坐标）
              map.addOverlay(new BMap.Marker(new BMap.Point(result.point.lng, result.point.lat)));
            }
          });
        }
        // else {
        //   alert('failed'+this.getStatus());
        // } 
      });
    },
    // getMapCenter() {
    //   var cen = this.map.getCenter(); // 获取地图中心点
    //   alert('地图中心点: (' + cen.lng.toFixed(5) + ', ' + cen.lat.toFixed(5) + ')');
    // }
  },
  mounted() {
    this.createMap();
  }
}
</script>

<script setup>
import {
  createClock_in,
  updateClock_in,
  findClock_in,
  getClock_inList,
} from '@/api/clock_in'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
const route = useRoute()
const router = useRouter()
const type = ref('')
const userStore = useUserStore()
const searchInfo = ref({})
const formData = ref({
        student_id: '',
        clock_in_date: new Date(),
        area_name: '',
        temperature: 36.80,
        symptom: '',
        })
userStore.GetUserInfo().then((res) => {
  // console.log(res['data']['userInfo']['userName'])
  formData._value['student_id'] = res['data']['userInfo']['userName']
  searchInfo._value['student_id'] = res['data']['userInfo']['userName']
  // console.log(formData)
})

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findClock_in({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reclock_in
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
    // 避免重复打卡
    let year = formData.value.clock_in_date.getFullYear()
    let month = formData.value.clock_in_date.getMonth()
    let date = formData.value.clock_in_date.getDate()
    const table = await getClock_inList({ page: 1, pageSize: 10, ...searchInfo.value })
    // console.log(table.data.list.length)
    if (table.code === 0 && table.data.list.length > 0) {
      let d = new Date(table.data.list[table.data.list.length-1].clock_in_date)
      if (d.getFullYear() === year && d.getMonth() === month && d.getDate() === date) {
        ElMessage({
          type: 'error',
          message: '请勿重复打卡'
        })
        return
      }
    }
    let res
    switch (type.value) {
      case 'create':
        res = await createClock_in(formData.value)
        break
      case 'update':
        res = await updateClock_in(formData.value)
        break
      default:
        res = await createClock_in(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '打卡成功'
      })
    }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

var geolocation = new BMap.Geolocation();
geolocation.getCurrentPosition(function (r) {
  if (this.getStatus() == BMAP_STATUS_SUCCESS) {
    //获取当前位置经纬度
    let myGeo = new BMap.Geocoder();
    myGeo.getLocation(new BMap.Point(r.point.lng, r.point.lat), function (result) {
      if (result) {
        // console.log(result.address.substr(0, result.address.indexOf('区')+1));
        formData._value['area_name'] = result.address.substr(0, result.address.indexOf('区')+1);
      }
    });
  }
  // else {
  //   alert('failed'+this.getStatus());
  // }
});

defineExpose({
  formData
});

</script>

<style>
#map {
  margin-top: 20px;
  height: 700px;
}
</style>
