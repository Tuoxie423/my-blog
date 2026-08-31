<template>
  <div class="murmur-list">
    <div class="title">
      <el-row>碎碎念管理</el-row>
    </div>

    <div class="murmur-create">
      <el-input
          v-model="content"
          type="textarea"
          :rows="3"
          maxlength="100"
          show-word-limit
          placeholder="写一条碎碎念..."
      />
      <div class="actions">
        <el-button type="primary" @click="handleCreate">发布</el-button>
      </div>
    </div>

    <el-table :data="murmurs">
      <el-table-column prop="id" label="ID" width="80"/>
      <el-table-column prop="content" label="内容"/>
      <el-table-column prop="created_at" label="发布时间" width="180"/>
      <el-table-column label="操作" width="100">
        <template #default="scope:{ row: Murmur, column: any, $index: number }">
          <el-button type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import {murmurAll, murmurCreate, murmurDelete, type Murmur} from "@/api/murmur";
import {ref} from "vue";
import {ElMessage} from "element-plus";

const murmurs = ref<Murmur[]>([])
const content = ref('')

const getMurmurs = async () => {
  const res = await murmurAll()
  if (res.code === 0) {
    murmurs.value = res.data
  }
}

getMurmurs()

const handleCreate = async () => {
  if (!content.value.trim()) {
    ElMessage.warning('请输入内容')
    return
  }
  const res = await murmurCreate({content: content.value})
  if (res.code === 0) {
    ElMessage.success(res.msg)
    content.value = ''
    getMurmurs()
  }
}

const handleDelete = async (id: number) => {
  const res = await murmurDelete({ids: [id]})
  if (res.code === 0) {
    ElMessage.success(res.msg)
    getMurmurs()
  }
}
</script>

<style scoped lang="scss">
.murmur-list {
  .title {
    display: flex;

    .el-row {
      font-size: 24px;
    }
  }

  .murmur-create {
    border: 1px solid var(--line);
    padding: 20px;
    margin-top: 20px;
    margin-bottom: 20px;

    .actions {
      display: flex;
      justify-content: flex-end;
      margin-top: 12px;
    }
  }

  .el-table {
    border: 1px solid var(--line);
  }
}
</style>
