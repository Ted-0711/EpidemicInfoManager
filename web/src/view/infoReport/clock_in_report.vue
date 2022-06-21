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
        <el-form-item label="区域编号:">
          <el-input v-model.number="formData.area_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="体温:">
          <el-input-number v-model="formData.temperature" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="不适症状:">
          <el-input v-model="formData.symptom" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
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
  methods: {
    //创建地图实例
    createMap() {
      var map = new BMap.Map("map");
      var point = new BMap.Point(116.404, 39.925);
      map.centerAndZoom(point, 15)
      map.enableScrollWheelZoom(true)
    }
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
  findClock_in
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
const formData = ref({
        student_id: '',
        clock_in_date: new Date(),
        area_id: 1,
        temperature: 36.80,
        symptom: '',
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
      let res
      console.log(formData.value)
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
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
#map {
      margin-top: 20px;
      height: 700px;
}
</style>
