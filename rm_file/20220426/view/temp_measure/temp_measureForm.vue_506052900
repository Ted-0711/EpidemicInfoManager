<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测温时间:">
          <el-date-picker v-model="formData.temp_measure_time" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="仪器编号:">
          <el-input v-model.number="formData.therm_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="温度:">
          <el-input-number v-model="formData.temperature" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Temp_measure'
}
</script>

<script setup>
import {
  createTemp_measure,
  updateTemp_measure,
  findTemp_measure
} from '@/api/temp_measure'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        student_id: '',
        temp_measure_time: new Date(),
        therm_id: 0,
        temperature: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTemp_measure({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retemp_measure
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
      switch (type.value) {
        case 'create':
          res = await createTemp_measure(formData.value)
          break
        case 'update':
          res = await updateTemp_measure(formData.value)
          break
        default:
          res = await createTemp_measure(formData.value)
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
</style>
