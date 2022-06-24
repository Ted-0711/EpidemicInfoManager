<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="标题">
          <el-input v-model="searchInfo.qtn_title" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="问题1">
          <el-input v-model="searchInfo.qtn_q1" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="问题2">
          <el-input v-model="searchInfo.qtn_q2" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="问题3">
          <el-input v-model="searchInfo.qtn_q3" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="问题4">
          <el-input v-model="searchInfo.qtn_q4" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="问题5">
          <el-input v-model="searchInfo.qtn_q5" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
          <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div v-for="(value, key, index) in tableData" :key="index">
      <div class="gva-table-box">
        <h1>{{ value['qtn_title'] }}</h1>
        <br>
        <h2>发布时间：{{ formatDate(value['CreatedAt']) }} &nbsp; 更新时间：{{ formatDate(value['UpdatedAt']) }} &nbsp; 截止时间：{{ formatDate(value['qtn_deadline']) }}</h2>
        <br>
        <p v-if="value['qtn_q1'].length > 0">1. {{ value['qtn_q1'] }}</p>
        <br v-if="value['qtn_q2'].length > 0">
        <p v-if="value['qtn_q2'].length > 0">2. {{ value['qtn_q2'] }}</p>
        <br v-if="value['qtn_q3'].length > 0">
        <p v-if="value['qtn_q3'].length > 0">3. {{ value['qtn_q3'] }}</p>
        <br v-if="value['qtn_q4'].length > 0">
        <p v-if="value['qtn_q4'].length > 0">4. {{ value['qtn_q4'] }}</p>
        <br v-if="value['qtn_q5'].length > 0">
        <p v-if="value['qtn_q5'].length > 0">5. {{ value['qtn_q5'] }}</p>
        <br>
        <el-button type="text" icon="finished" size="small" class="table-button" @click="searchQuestionnaireFunc(value)">填写情况</el-button>
        <el-button type="text" icon="edit" size="small" class="table-button" @click="updateQuestionnaireFunc(value)">变更</el-button>
        <el-button type="text" icon="delete" size="small" @click="deleteRow(value)">删除</el-button>
      </div>
      <br>
    </div>
    <!-- <div class="gva-table-box">
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
        <el-table-column align="left" label="标题" prop="qtn_title" width="120" />
        <el-table-column align="left" label="截止时间" prop="qtn_deadline" width="120" />
        <el-table-column align="left" label="问题1" prop="qtn_q1" width="240" />
        <el-table-column align="left" label="问题2" prop="qtn_q2" width="240" />
        <el-table-column align="left" label="问题3" prop="qtn_q3" width="240" />
        <el-table-column align="left" label="问题4" prop="qtn_q4" width="240" />
        <el-table-column align="left" label="问题5" prop="qtn_q5" width="240" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="finished" size="small" class="table-button" @click="searchQuestionnaireFunc(scope.row)">填写情况</el-button>
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateQuestionnaireFunc(scope.row)">变更</el-button>
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
    </div> -->
		<el-dialog v-model="fTableVisible" :before-close="closeTable" :title="'问卷填写情况 —— ' + formData.qtn_title">
      <div class="gva-search-box">
				<el-form :inline="true" :model="fSearchInfo" class="demo-form-inline">
					<el-form-item label="学号">
						<el-input v-model="fSearchInfo.student_id" placeholder="搜索条件" />
					</el-form-item>
					<el-form-item label="问卷编号">
						<el-input v-model="fSearchInfo.qtn_id" placeholder="搜索条件" disabled />
					</el-form-item>
          <el-popover
            v-if="formData.qtn_q1.length > 0"
            placement="top-start"
            title="问题1:"
            width="200"
            trigger="hover"
            :content="formData.qtn_q1">
            <template v-slot:reference>
              <el-form-item label="回答1">
                <el-input v-model="fSearchInfo.fill_in_a1" placeholder="搜索条件" />
              </el-form-item>
            </template>
          </el-popover>
          <el-popover
            v-if="formData.qtn_q2.length > 0"
            placement="top-start"
            title="问题2:"
            width="200"
            trigger="hover"
            :content="formData.qtn_q2">
            <template v-slot:reference>
              <el-form-item label="回答2">
                <el-input v-model="fSearchInfo.fill_in_a2" placeholder="搜索条件" />
              </el-form-item>
            </template>
          </el-popover>
          <el-popover
            v-if="formData.qtn_q3.length > 0"
            placement="top-start"
            title="问题3:"
            width="200"
            trigger="hover"
            :content="formData.qtn_q3">
            <template v-slot:reference>
              <el-form-item label="回答3">
                <el-input v-model="fSearchInfo.fill_in_a3" placeholder="搜索条件" />
              </el-form-item>
            </template>
          </el-popover>
          <el-popover
            v-if="formData.qtn_q4.length > 0"
            placement="top-start"
            title="问题4:"
            width="200"
            trigger="hover"
            :content="formData.qtn_q4">
            <template v-slot:reference>
              <el-form-item label="回答4">
                <el-input v-model="fSearchInfo.fill_in_a4" placeholder="搜索条件" />
              </el-form-item>
            </template>
          </el-popover>
          <el-popover
            v-if="formData.qtn_q5.length > 0"
            placement="top-start"
            title="问题5:"
            width="200"
            trigger="hover"
            :content="formData.qtn_q5">
            <template v-slot:reference>
              <el-form-item label="回答5">
                <el-input v-model="fSearchInfo.fill_in_a5" placeholder="搜索条件" />
              </el-form-item>
            </template>
          </el-popover>
					<!-- <el-form-item label="回答2">
						<el-input v-model="fSearchInfo.fill_in_a2" placeholder="搜索条件" />
					</el-form-item>
					<el-form-item label="回答3">
						<el-input v-model="fSearchInfo.fill_in_a3" placeholder="搜索条件" />
					</el-form-item>
					<el-form-item label="回答4">
						<el-input v-model="fSearchInfo.fill_in_a4" placeholder="搜索条件" />
					</el-form-item>
					<el-form-item label="回答5">
						<el-input v-model="fSearchInfo.fill_in_a5" placeholder="搜索条件" />
					</el-form-item> -->
					<el-form-item>
						<el-button size="small" type="primary" icon="search" @click="fOnSubmit">查询</el-button>
						<el-button size="small" icon="refresh" @click="fOnReset">重置</el-button>
					</el-form-item>
				</el-form>
			</div>
			<div class="gva-table-box">
					<el-table
					ref="multipleTable"
					style="width: 100%"
					tooltip-effect="dark"
					:data="fTableData"
					row-key="ID"
					>
					<el-table-column align="left" label="日期" width="180">
							<template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
					</el-table-column>
					<el-table-column align="left" label="学号" prop="student_id" width="120" />
					<el-table-column align="left" label="问卷编号" prop="qtn_id" width="120" />
          <el-table-column v-if="formData.qtn_q1.length > 0" align="left" label="回答1" prop="fill_in_a1" width="120" />
					<el-table-column v-if="formData.qtn_q2.length > 0" align="left" label="回答2" prop="fill_in_a2" width="120" />
					<el-table-column v-if="formData.qtn_q3.length > 0" align="left" label="回答3" prop="fill_in_a3" width="120" />
					<el-table-column v-if="formData.qtn_q4.length > 0" align="left" label="回答4" prop="fill_in_a4" width="120" />
					<el-table-column v-if="formData.qtn_q5.length > 0" align="left" label="回答5" prop="fill_in_a5" width="120" />
					</el-table>
					<div class="gva-pagination">
							<el-pagination
							layout="total, sizes, prev, pager, next, jumper"
							:current-page="fPage"
							:page-size="fPageSize"
							:page-sizes="[10, 30, 50, 100]"
							:total="total"
							@current-change="handleCurrentChange"
							@size-change="handleSizeChange"
							/>
					</div>
			</div>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeTable">取 消</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="创建/更改问卷">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="标题:">
          <el-input v-model="formData.qtn_title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="截止时间:">
          <el-date-picker v-model="formData.qtn_deadline" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="问题1:">
          <el-input v-model="formData.qtn_q1" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题2:">
          <el-input v-model="formData.qtn_q2" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题3:">
          <el-input v-model="formData.qtn_q3" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题4:">
          <el-input v-model="formData.qtn_q4" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问题5:">
          <el-input v-model="formData.qtn_q5" clearable placeholder="请输入" />
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
  name: 'Questionnaire'
}
</script>

<script setup>
import {
  createQuestionnaire,
  deleteQuestionnaire,
  deleteQuestionnaireByIds,
  updateQuestionnaire,
  findQuestionnaire,
  getQuestionnaireList
} from '@/api/questionnaire'

import {
  getFill_inList
} from '@/api/fill_in'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        qtn_title: '',
        qtn_deadline: new Date(),
        qtn_q1: '',
        qtn_q2: '',
        qtn_q3: '',
        qtn_q4: '',
        qtn_q5: '',
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const fPage = ref(1)
const fTotal = ref(0)
const fPageSize = ref(10)
const fTableData = ref([])
const fSearchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
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
  const table = await getQuestionnaireList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    // 按照更新时间由近到远排列
    table.data.list.sort((a, b) => {
      return new Date(b.UpdatedAt.substr(0, 22)) - new Date(a.UpdatedAt.substr(0, 22))
    })
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 重置
const fOnReset = () => {
  fSearchInfo.value = {}
}

// 搜索
const fOnSubmit = () => {
  fPage.value = 1
  fPageSize.value = 10
  fGetTableData()
}

// 分页
const fHandleSizeChange = (val) => {
  fPageSize.value = val
  fGetTableData()
}

// 修改页面容量
const fHandleCurrentChange = (val) => {
  fPage.value = val
  fGetTableData()
}

// 查询
const fGetTableData = async() => {
  const table = await getFill_inList({ page: fPage.value, pageSize: fPageSize.value, ...fSearchInfo.value })
  if (table.code === 0) {
    fTableData.value = table.data.list
    fTotal.value = table.data.total
    fPage.value = table.data.page
    fPageSize.value = table.data.pageSize
  }
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
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
            deleteQuestionnaireFunc(row)
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
      const res = await deleteQuestionnaireByIds({ ids })
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

// 查看行详细信息
const searchQuestionnaireFunc = async(row) => {
    const res = await findQuestionnaire({ ID: row.ID })
    if (res.code === 0) {
        console.log(formData)
        formData.value = res.data.requestionnaire
				fSearchInfo.value.qtn_id = row.ID
				// console.log(fSearchInfo)
				fGetTableData()
				fTableVisible.value = true
		}
}

// 更新行
const updateQuestionnaireFunc = async(row) => {
    const res = await findQuestionnaire({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.requestionnaire
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteQuestionnaireFunc = async (row) => {
    const res = await deleteQuestionnaire({ ID: row.ID })
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
        qtn_title: '',
        qtn_deadline: new Date(),
        qtn_q1: '',
        qtn_q2: '',
        qtn_q3: '',
        qtn_q4: '',
        qtn_q5: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createQuestionnaire(formData.value)
          break
        case 'update':
          res = await updateQuestionnaire(formData.value)
          break
        default:
          res = await createQuestionnaire(formData.value)
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

const fTableVisible = ref(false)

// 关闭弹窗
const closeTable = () => {
    fTableVisible.value = false
    formData.value = {
        qtn_title: '',
        qtn_deadline: new Date(),
        qtn_q1: '',
        qtn_q2: '',
        qtn_q3: '',
        qtn_q4: '',
        qtn_q5: '',
        }
}
</script>

<style>
h1 {
  font-size: 20px;
  color: #303133;
}
h2 {
  font-size: 14px;
  color: #606266;
}
p {
  font-size: 14px;
  color: #606266;
}
</style>
