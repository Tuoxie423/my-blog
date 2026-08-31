<template>
  <div class="article-create-form">
    <el-form
        :model="articleCreateFormData"
        :validate-on-rule-change="false"
    >
      <el-form-item label="文章标题" prop="title">
        <el-input
            v-model="articleCreateFormData.title"
            size="large"
            placeholder="请输入文章标题"
        />
      </el-form-item>
      <el-form-item label="文章标签" prop="tags">
        <el-tag v-for="tag in articleCreateFormData.tags"
                :key="tag"
                closable
                :disable-transitions="false"
                size="large"
                @close="handleClose(tag)">
          {{ tag }}
        </el-tag>
        <el-input
            v-if="inputVisible"
            ref="InputRef"
            v-model="inputValue"
            style="width: 80px"
            @keyup.enter="handleInputConfirm"
            @blur="handleInputConfirm"
        />
        <el-button v-else @click="showInput">+ 新建标签</el-button>
      </el-form-item>
      <el-form-item label="文章简介" prop="abstract">
        <el-input
            v-model="articleCreateFormData.abstract"
            type="textarea"
            placeholder="请输入文章简介"
        />
      </el-form-item>
      <el-form-item>
        <div class="button-group">
          <el-button
              type="primary"
              size="large"
              @click="submitForm"
          >确定
          </el-button>
          <el-button
              size="large"
              @click="layoutStore.state.articleCreateVisible = false"
          >取消
          </el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import {defineProps, nextTick, reactive, ref} from "vue";
import {ElMessage, type InputInstance} from "element-plus";
import {articleCreate, type ArticleCreateRequest} from "@/api/article";
import {useLayoutStore} from "@/stores/layout";

const props = defineProps<{
  title: string;
  content: string;
}>();

const layoutStore = useLayoutStore()

const articleCreateFormData = reactive<ArticleCreateRequest>({
  title: props.title,
  tags: [],
  abstract: '',
  content: props.content,
})

const inputValue = ref('')
const inputVisible = ref(false)
const InputRef = ref<InputInstance>()

const handleClose = (tag: string) => {
  articleCreateFormData.tags.splice(articleCreateFormData.tags.indexOf(tag), 1)
}

const showInput = () => {
  inputVisible.value = true
  nextTick(() => {
    InputRef.value!.input!.focus()
  })
}

const handleInputConfirm = () => {
  if (inputValue.value) {
    articleCreateFormData.tags.push(inputValue.value)
  }
  inputVisible.value = false
  inputValue.value = ''
}

const submitForm = async () => {
  const res = await articleCreate(articleCreateFormData)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    layoutStore.state.articleCreateVisible = false
  }
}
</script>

<style scoped lang="scss">
.article-create-form {
  .el-form {
    .el-form-item {
      .button-group {
        margin-left: auto;
      }
    }
  }
}
</style>
