<template>
  <el-card class="profile-card">
    <el-row class="title">个人名片</el-row>
    <div class="content">
      <div class="field">
        <span class="meta-label">昵称</span>
        <span class="value">{{ websiteStore.state.websiteInfo.name }}</span>
      </div>
      <div class="field">
        <span class="meta-label">住址</span>
        <span class="value">{{ websiteStore.state.websiteInfo.address }}</span>
      </div>
      <div class="field">
        <span class="meta-label">邮箱</span>
        <span class="value">{{ websiteStore.state.websiteInfo.email }}</span>
      </div>
    </div>

    <div class="divider"></div>

    <div class="tags-section">
      <span class="meta-label">文章标签</span>
      <div class="tags">
        <span v-for="tag in tagCloudArray" :key="tag.tag" class="tag" @click="handleTagClick(tag.tag)">
          {{ tag.tag }}<span class="num">{{ tag.number }}</span>
        </span>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import {useWebsiteStore} from "@/stores/website";
import {articleTags, type ArticleTag} from "@/api/article";
import {ref} from "vue";

const websiteStore = useWebsiteStore()

const tagCloudArray = ref<ArticleTag[]>([])

const getTags = async () => {
  const res = await articleTags()
  if (res.code === 0) {
    tagCloudArray.value = res.data
  }
}

getTags()

const handleTagClick = (tag: string) => {
  window.open("/search?tag=" + encodeURIComponent(tag))
}
</script>

<style scoped lang="scss">
.profile-card {
  .title {
    font-size: 24px;
    margin-bottom: 20px;
  }

  .content {
    display: flex;
    flex-direction: column;

    .field {
      display: flex;
      justify-content: space-between;
      align-items: baseline;
      padding: 12px 0;
      border-bottom: 1px solid var(--line);

      &:last-child {
        border-bottom: none;
      }

      .meta-label {
        flex-shrink: 0;
        width: 48px;
      }

      .value {
        font-family: var(--font-mono);
        font-size: 14px;
        color: var(--ink);
        text-align: right;
        word-break: break-all;
      }
    }
  }

  .divider {
    border-top: 1px solid var(--line-strong);
    margin: 24px 0;
  }

  .tags-section {
    .meta-label {
      display: block;
      margin-bottom: 16px;
    }

    .tags {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;

      .tag {
        font-family: var(--font-mono);
        font-size: 13px;
        color: var(--ink-2);
        border: 1px solid var(--line);
        padding: 4px 10px;
        cursor: pointer;
        transition: border-color 0.2s ease, color 0.2s ease;

        .num {
          margin-left: 4px;
          color: var(--ink-3);
        }

        &:hover {
          border-color: var(--ink);
          color: var(--accent);

          .num {
            color: var(--accent);
          }
        }
      }
    }
  }
}
</style>
