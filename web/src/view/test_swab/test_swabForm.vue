<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="检测机构:">
          <el-select v-model="formData.facility" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in facilityOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="采样时间:">
          <el-date-picker v-model="formData.sample_date" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="检测时间:">
          <el-date-picker v-model="formData.test_date" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="检测结果:">
          <el-select v-model="formData.test_result" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in test_resultOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="截图链接:">
          <el-input v-model="formData.screenshot_url" clearable placeholder="请输入" />
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
  name: 'Test_swab'
}
</script>

<script setup>
import {
  createTest_swab,
  updateTest_swab,
  findTest_swab
} from '@/api/test_swab'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const facilityOptions = ref([])
const test_resultOptions = ref([])
const formData = ref({
        student_id: '',
        facility: undefined,
        sample_date: new Date(),
        test_date: new Date(),
        test_result: undefined,
        screenshot_url: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTest_swab({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retest_swab
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    facilityOptions.value = await getDictFunc('facility')
    test_resultOptions.value = await getDictFunc('test_result')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createTest_swab(formData.value)
          break
        case 'update':
          res = await updateTest_swab(formData.value)
          break
        default:
          res = await createTest_swab(formData.value)
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
