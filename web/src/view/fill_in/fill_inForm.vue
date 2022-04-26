<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="问卷编号:">
          <el-input v-model.number="formData.qtn_id" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="时间:">
          <el-date-picker v-model="formData.fill_in_time" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="内容:">
          <el-input v-model="formData.fill_in_content" clearable placeholder="请输入" />
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
  name: 'Fill_in'
}
</script>

<script setup>
import {
  createFill_in,
  updateFill_in,
  findFill_in
} from '@/api/fill_in'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        qtn_id: 0,
        fill_in_time: new Date(),
        fill_in_content: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFill_in({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refill_in
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
          res = await createFill_in(formData.value)
          break
        case 'update':
          res = await updateFill_in(formData.value)
          break
        default:
          res = await createFill_in(formData.value)
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
