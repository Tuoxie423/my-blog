<template>
  <div class="comment-section">
    <div class="header">
      <el-row class="title">评论<span class="count">{{ totalCount }}</span></el-row>
    </div>

    <div class="editor">
      <template v-if="userStore.isLoggedIn">
        <el-input
            v-model="content"
            type="textarea"
            :rows="3"
            maxlength="320"
            show-word-limit
            placeholder="写下你的评论..."
        />
        <div class="actions">
          <el-button type="primary" @click="submitComment">发表评论</el-button>
        </div>
      </template>
      <span v-else class="login-tip">登录后参与评论</span>
    </div>

    <div class="list">
      <comment-item v-if="comments.length" :comments="comments"/>
      <div v-else class="empty">还没有评论，来抢沙发吧</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch} from "vue";
import {commentCreate, commentInfoByArticleID, type Comment} from "@/api/comment";
import {useUserStore} from "@/stores/user";
import {useLayoutStore} from "@/stores/layout";
import CommentItem from "@/components/common/CommentItem.vue";

const props = defineProps<{
  articleId: string
}>()

const userStore = useUserStore()
const layoutStore = useLayoutStore()

const comments = ref<Comment[]>([])
const content = ref('')

const getComments = async () => {
  const res = await commentInfoByArticleID(props.articleId)
  if (res.code === 0) {
    comments.value = res.data
  }
}

getComments()

// 递归统计评论总数（含回复）
const countComments = (list: Comment[]): number =>
    list.reduce((sum, c) => sum + 1 + countComments(c.children || []), 0)

const totalCount = computed(() => countComments(comments.value))

const submitComment = async () => {
  if (!content.value.trim()) {
    ElMessage.warning('请输入评论内容')
    return
  }
  const res = await commentCreate({
    article_id: props.articleId,
    p_id: null,
    content: content.value,
  })
  if (res.code === 0) {
    ElMessage.success(res.msg)
    content.value = ''
    getComments()
  }
}

watch(() => layoutStore.state.shouldRefreshCommentList, (newVal) => {
  if (newVal) {
    getComments()
    layoutStore.state.shouldRefreshCommentList = false
  }
})
</script>

<style scoped lang="scss">
.comment-section {
  margin-top: 48px;
  border-top: 1px solid var(--line-strong);
  padding-top: 28px;

  .header {
    margin-bottom: 24px;

    .title {
      font-family: var(--font-serif);
      font-size: 28px;
      display: flex;
      align-items: center;
      gap: 12px;

      &::before {
        content: '';
        flex-shrink: 0;
        width: 20px;
        height: 2px;
        background: var(--accent);
      }

      .count {
        font-family: var(--font-mono);
        font-size: 13px;
        color: var(--ink-3);
      }
    }
  }

  .editor {
    margin-bottom: 32px;

    .actions {
      display: flex;
      justify-content: flex-end;
      margin-top: 12px;
    }

    .login-tip {
      display: block;
      padding: 16px;
      border: 1px dashed var(--line);
      color: var(--ink-3);
      font-family: var(--font-mono);
      font-size: 13px;
    }
  }

  .list {
    .empty {
      padding: 40px 0;
      text-align: center;
      color: var(--ink-3);
      font-family: var(--font-mono);
      font-size: 14px;
    }
  }
}
</style>
