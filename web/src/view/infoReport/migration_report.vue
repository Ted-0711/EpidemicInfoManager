<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" disabled />
        </el-form-item>
        <el-form-item label="起点编号:">
          <el-input v-model.number="formData.start_area_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="终点编号:">
          <el-input v-model.number="formData.des_area_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交通方式:">
          <el-select v-model="formData.vehicle_type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in vehicle_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="班次信息:">
          <el-input v-model="formData.vehicle_info" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出发时间:">
          <el-date-picker v-model="formData.mig_time" type="date" placeholder="选择日期" clearable></el-date-picker>
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
  name: 'Migration'
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
const userStore = useUserStore()
const formData = ref({
        student_id: '',
        start_area_id: 0,
        des_area_id: 0,
        vehicle_type: undefined,
        vehicle_info: '',
        mig_time: new Date(),
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
}

init()
// 保存按钮
const save = async() => {
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

</script>

<style>
</style>
