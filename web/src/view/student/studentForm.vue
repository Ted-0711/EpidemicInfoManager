<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学院:">
          <el-select v-model="formData.dept_name" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in dept_nameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.student_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-select v-model="formData.student_gender" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系方式:">
          <el-input v-model="formData.student_phone_number" clearable placeholder="请输入" />
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
  name: 'Student'
}
</script>

<script setup>
import {
  createStudent,
  updateStudent,
  findStudent
} from '@/api/student'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const dept_nameOptions = ref([])
const genderOptions = ref([])
const formData = ref({
        student_id: '',
        dept_name: undefined,
        student_name: '',
        student_gender: undefined,
        student_phone_number: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findStudent({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.restudent
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    dept_nameOptions.value = await getDictFunc('dept_name')
    genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createStudent(formData.value)
          break
        case 'update':
          res = await updateStudent(formData.value)
          break
        default:
          res = await createStudent(formData.value)
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
