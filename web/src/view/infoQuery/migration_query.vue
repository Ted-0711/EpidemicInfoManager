<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="学号">
          <el-input v-model="searchInfo.student_id" placeholder="搜索条件" :disabled="studentInfo.isStudent" />
        </el-form-item>
        <el-form-item label="起点区域">
          <el-input v-model="searchInfo.start_area" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="终点区域">
          <el-input v-model="searchInfo.des_area" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="截图链接">
          <el-input v-model="searchInfo.screenshot_url" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="审核状态">
          <el-input v-model="searchInfo.audit_status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="学号" prop="student_id" width="120" />
        <el-table-column align="left" label="起点区域" prop="start_area" width="120" />
        <el-table-column align="left" label="终点区域" prop="des_area" width="120" />
        <el-table-column align="left" label="出发时间" prop="mig_time" width="120" />
        <el-table-column align="left" label="交通方式" prop="vehicle_type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.vehicle_type,vehicle_typeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="截图链接" prop="screenshot_url" width="120" />
        <el-table-column align="left" label="票务信息" prop="screenshot_url" width="120">
            <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.screenshot_url"/>
            </template>
        </el-table-column>
        <el-table-column align="left" label="审核状态" prop="audit_status" width="120">
            <template #default="scope">
              <el-tag :type="scope.row.audit_status == 0 ? 'warning' : (scope.row.audit_status == 1 ? 'success' : 'error')">{{ filterDict(scope.row.audit_status,audit_statusOptions) }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateMigrationFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.student_id" clearable placeholder="请输入" :disabled="studentInfo.isStudent" />
        </el-form-item>
        <el-form-item label="起点区域:">
          <el-input v-model="formData.start_area" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="终点区域:">
          <el-input v-model="formData.des_area" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出发时间:">
          <el-date-picker v-model="formData.mig_time" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="交通方式:">
          <el-select v-model="formData.vehicle_type" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in vehicle_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="截图链接:">
          <el-input v-model="formData.screenshot_url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态:">
          <el-select v-model="formData.audit_status" placeholder="请选择" style="width:100%" clearable :disabled="studentInfo.isStudent">
            <el-option v-for="(item,key) in audit_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteMigration,
  deleteMigrationByIds,
  updateMigration,
  findMigration,
  getMigrationList
} from '@/api/migration'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import CustomPic from '@/components/customPic/index.vue'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const studentInfo = {
  isStudent: false,
  studentId: ''
}
userStore.GetUserInfo().then((res) => {
  if (res['data']['userInfo']['authorityId'] == 9990) {
    studentInfo.isStudent = true
    studentInfo.studentId = res['data']['userInfo']['userName']
    searchInfo._value['student_id'] = studentInfo.studentId
    formData._value['student_id'] = studentInfo.studentId
    onSubmit()
    // console.log('Is Student, id: ', studentId)
  }
})

// 自动化生成的字典（可能为空）以及字段
const vehicle_typeOptions = ref([])
const audit_statusOptions = ref([])
const formData = ref({
        student_id: '',
        start_area: '',
        des_area: '',
        mig_time: new Date(),
        vehicle_type: undefined,
        screenshot_url: '',
        audit_status: undefined,
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  if (studentInfo.isStudent) {
    searchInfo._value['student_id'] = studentInfo.studentId
  }
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMigrationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    vehicle_typeOptions.value = await getDictFunc('vehicle_type')
    audit_statusOptions.value = await getDictFunc('audit_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMigrationFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteMigrationByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMigrationFunc = async(row) => {
    const res = await findMigration({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.remigration
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMigrationFunc = async (row) => {
    const res = await deleteMigration({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        student_id: '',
        start_area: '',
        des_area: '',
        mig_time: new Date(),
        vehicle_type: undefined,
        screenshot_url: '',
        audit_status: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
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
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
