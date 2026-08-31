<template>
  <div class="article-update-form">
    <el-form
        :model="articleUpdateFormData"
        :validate-on-rule-change="false"
    >
      <el-form-item label="文章标题" prop="title">
        <el-input
            v-model="articleUpdateFormData.title"
            size="large"
            placeholder="请输入文章标题"
        />
      </el-form-item>
      <el-form-item label="文章标签" prop="tags">
        <el-tag v-for="tag in articleUpdateFormData.tags"
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
            v-model="articleUpdateFormData.abstract"
            type="textarea"
            placeholder="请输入文章简介"
        />
      </el-form-item>

      <el-form-item label="文章内容" prop="content">
        <el-button @click="drawer = true" icon="EditPen">编辑内容</el-button>
          <el-drawer v-model="drawer" :direction="direction" size="80%">
            <template #header>
              编辑内容
            </template>
            <template #default>
              <MdEditor v-model="articleUpdateFormData.content" @onUploadImg="onUploadImg"/>
            </template>
            <template #footer>
              <el-text>点击上方X或外部任意区域即可退出编辑</el-text>
            </template>
          </el-drawer>
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
              @click="layoutStore.state.articleUpdateVisible = false"
          >取消
          </el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import {defineProps, nextTick, reactive, ref} from "vue";
import {type DrawerProps, ElMessage, type InputInstance} from "element-plus";
import {type Article, articleUpdate, type ArticleUpdateRequest} from "@/api/article";
import {imageUpload} from "@/api/image";
import {useLayoutStore} from "@/stores/layout";
import type {Hit} from "@/api/common";
import {MdEditor} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const props = defineProps<{
  article: Hit<Article>;
}>();

const layoutStore = useLayoutStore()

const articleUpdateFormData = reactive<ArticleUpdateRequest>({
  id: props.article._id,
  title: props.article._source.title,
  tags: props.article._source.tags,
  abstract: props.article._source.abstract,
  content: props.article._source.content,
})

const inputValue = ref('')
const inputVisible = ref(false)
const InputRef = ref<InputInstance>()

const handleClose = (tag: string) => {
  articleUpdateFormData.tags.splice(articleUpdateFormData.tags.indexOf(tag), 1)
}

const showInput = () => {
  inputVisible.value = true
  nextTick(() => {
    InputRef.value!.input!.focus()
  })
}

const handleInputConfirm = () => {
  if (inputValue.value) {
    articleUpdateFormData.tags.push(inputValue.value)
  }
  inputVisible.value = false
  inputValue.value = ''
}

const drawer = ref(false)
const direction = ref<DrawerProps['direction']>('rtl')

const onUploadImg = async (files: File[], callback: (urls: string[]) => void): Promise<void> => {
  const res = await Promise.all(files.map((file) => imageUpload(file)));
  callback(res.map((item) => item.data.url));
};

const submitForm = async () => {
  const res = await articleUpdate(articleUpdateFormData)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    layoutStore.state.articleUpdateVisible = false
  }
}
</script>

<style scoped lang="scss">
.article-update-form {
  .el-form {
    .el-form-item {
      .button-group {
        margin-left: auto;
      }
    }
  }
}
</style>

<style lang="scss">
.el-drawer{
  .md-editor .md-editor-toolbar-wrapper .md-editor-toolbar svg.md-editor-icon {
    height: 24px;
    width: 24px;
  }
}
</style>
