<template>
  <el-card class="feedback">
    <el-row class="title">意见反馈</el-row>
    <el-input
        v-model="feedbackCreateFormData.content"
        type="textarea"
        :rows="3"
        maxlength="100"
        show-word-limit
        placeholder="请输入反馈建议"
    />
    <div class="actions">
      <span class="meta-label">tip：登录后可反馈</span>
      <el-button type="primary" @click="submitForm">提交</el-button>
    </div>

    <div class="divider"></div>

    <div class="feedback-list">
      <div v-for="(item, i) in feedbackInfoList" :key="i" class="feedback-item">
        <p class="content">{{ item.content }}</p>
        <span class="time">{{ item.time }}</span>
        <p v-if="item.reply" class="reply">回复：{{ item.reply }}</p>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue";
import {feedbackCreate, type FeedbackCreateRequest, feedbackNew} from "@/api/feedback";

const feedbackCreateFormData = reactive<FeedbackCreateRequest>({content: ''})

interface FeedbackItem {
  content: string;
  reply: string;
  time: string;
}

const feedbackInfoList = ref<FeedbackItem[]>([])

const getFeedbackNew = async () => {
  feedbackInfoList.value = []
  const res = await feedbackNew()
  if (res.code === 0) {
    res.data.forEach(value => {
      feedbackInfoList.value.push({
        content: value.content,
        reply: value.reply,
        time: new Date(value.created_at).toLocaleString(),
      })
    })
  }
}

getFeedbackNew()

const submitForm = async () => {
  if (!feedbackCreateFormData.content.trim()) {
    ElMessage.warning('请输入反馈内容')
    return
  }
  const res = await feedbackCreate(feedbackCreateFormData)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    feedbackCreateFormData.content = ''
    getFeedbackNew()
  }
}
</script>

<style scoped lang="scss">
.feedback {
  margin-top: 20px;

  .title {
    font-size: 24px;
    margin-bottom: 20px;
  }

  .actions {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 12px;
  }

  .divider {
    border-top: 1px solid var(--line-strong);
    margin: 24px 0;
  }

  .feedback-list {
    display: flex;
    flex-direction: column;

    .feedback-item {
      padding: 14px 0;
      border-bottom: 1px solid var(--line);

      &:last-child {
        border-bottom: none;
      }

      .content {
        font-family: var(--font-mono);
        font-size: 14px;
        color: var(--ink);
        margin: 0 0 8px;
      }

      .time {
        font-family: var(--font-mono);
        font-size: 12px;
        color: var(--ink-3);
      }

      .reply {
        font-family: var(--font-serif);
        font-size: 14px;
        color: var(--ink-2);
        border-left: 3px solid var(--accent);
        padding-left: 12px;
        margin: 8px 0 0;
      }
    }
  }
}
</style>
