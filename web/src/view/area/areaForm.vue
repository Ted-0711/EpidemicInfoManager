<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.area_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="风险等级:">
          <el-select v-model="formData.area_risk_level" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in area_risk_levelOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'Area'
}
</script>

<script setup>
import {
  createArea,
  updateArea,
  findArea
} from '@/api/area'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const area_risk_levelOptions = ref([])
const formData = ref({
        area_name: '',
        area_risk_level: undefined,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findArea({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rearea
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    area_risk_levelOptions.value = await getDictFunc('area_risk_level')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createArea(formData.value)
          break
        case 'update':
          res = await updateArea(formData.value)
          break
        default:
          res = await createArea(formData.value)
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
