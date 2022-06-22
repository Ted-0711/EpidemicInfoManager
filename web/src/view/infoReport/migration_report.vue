<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" disabled />
        </el-form-item>
        <el-form-item label="起点区域:">
          <el-input v-model="formData.start_area" clearable placeholder="请输入" disabled />
        </el-form-item>
        <el-form-item label="终点区域:">
          <el-input v-model="formData.des_area" clearable placeholder="请输入" disabled />
        </el-form-item>
        <el-form-item label="出发时间:">
          <el-date-picker v-model="formData.mig_time" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="交通方式:">
          <el-select v-model="formData.vehicle_type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in vehicle_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="截图链接:">
          <el-input v-model="formData.screenshot_url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="票务信息:">
          <CustomPic pic-type="file" :pic-src="formData.screenshot_url"/>
          <upload-image
              v-model:imageUrl="imageUrl"
              :file-size="512"
              :max-w-h="1080"
              class="upload-btn"
              @on-success="imgtest"
          />
        </el-form-item>
        <el-form-item label="审核状态:">
          <el-select v-model="formData.audit_status" placeholder="请选择" clearable disabled>
            <el-option v-for="(item,key) in audit_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
          <el-button size="mini" type="primary" @click="this.getMapCenter(0)">获取起点</el-button>
          <el-button size="mini" type="primary" @click="this.getMapCenter(1)">获取终点</el-button>
        </el-form-item>
      </el-form>
      <div id="map" class="map"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Migration',
  "BMap": "BMap",
  data() {
    return {
      map: {},
    }
  },
  methods: {
    //创建地图实例
    createMap(id="map") {
      var map = new BMap.Map(id);
      this.map[id] = map;
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
    getMapCenter(type=0) {
      var cen = this.map["map"].getCenter(); // 获取地图中心点
      let myGeo = new BMap.Geocoder();
      myGeo.getLocation(cen, (result) => {
        ElMessage({
          type: 'success',
          message: '地图中心点: ' + result.address.substr(0, result.address.indexOf('区')+1)
        });
        if (type == 0) {
          this.formData.start_area = result.address.substr(0, result.address.indexOf('区')+1);
        }
        else {
          this.formData.des_area = result.address.substr(0, result.address.indexOf('区')+1);
        }
      });
    }
  },
  mounted() {
    this.createMap();
  }
}
</script>

<script setup>
import {
  createMigration,
  updateMigration,
  findMigration
} from '@/api/migration'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
const route = useRoute()
const router = useRouter()
const type = ref('')
const vehicle_typeOptions = ref([])
const audit_statusOptions = ref([])
const userStore = useUserStore()
const formData = ref({
        student_id: '',
        start_area: '',
        des_area: '',
        mig_time: new Date(),
        vehicle_type: undefined,
        screenshot_url: '',
        audit_status: 0,
        })
userStore.GetUserInfo().then((res) => {
  console.log(res['data']['userInfo']['userName'])
  formData._value['student_id'] = res['data']['userInfo']['userName']
  console.log(formData)
})

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMigration({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remigration
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    vehicle_typeOptions.value = await getDictFunc('vehicle_type')
    audit_statusOptions.value = await getDictFunc('audit_status')
}

init()
// 保存按钮
const save = async() => {
      console.log(formData)
      let res
      switch (type.value) {
        case 'create':
          res = await createMigration(formData.value)
          break
        case 'update':
          res = await updateMigration(formData.value)
          break
        default:
          res = await createMigration(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

const imageUrl = ref('')

import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'

// 上传截图
const imgtest = async(url) => {
  ElMessage({
    type: 'success',
    message: '上传成功'
  })
  formData._value['screenshot_url'] = url
}

</script>

<style>
/* .gva-form-box {
  overflow: hidden;
} */
.map {
  margin-top: 20px;
  height: 700px;
  width: 100%;
  /* float: left; */
}
</style>
