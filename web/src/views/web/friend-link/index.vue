<template>
  <div class="tags-page">
    <el-container class="main-content">
      <div class="container">
        <el-main>
          <div class="header rise">
            <p class="meta-label">All Tags</p>
            <h1 class="title">文章标签</h1>
            <p class="sub">共 {{ tags.length }} 个标签 · 字号越大代表文章越多</p>
          </div>

          <div class="cloud">
            <span
                v-for="tag in tags"
                :key="tag.tag"
                class="tag"
                :style="{fontSize: getFontSize(tag.number)}"
                :title="tag.number + ' 篇文章'"
                @click="handleTagClick(tag.tag)"
            >{{ tag.tag }}</span>
          </div>
        </el-main>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {articleTags, type ArticleTag} from "@/api/article";
import {ref} from "vue";

const tags = ref<ArticleTag[]>([])

const getTags = async () => {
  const res = await articleTags()
  if (res.code === 0) {
    tags.value = res.data
  }
}

getTags()

const MIN_SIZE = 14
const MAX_SIZE = 60

const getFontSize = (number: number): string => {
  if (tags.value.length === 0) return `${MIN_SIZE}px`
  const nums = tags.value.map(t => t.number)
  const max = Math.max(...nums)
  const min = Math.min(...nums)
  if (max === min) return `${Math.round((MIN_SIZE + MAX_SIZE) / 2)}px`
  const ratio = (number - min) / (max - min)
  return `${Math.round(MIN_SIZE + ratio * (MAX_SIZE - MIN_SIZE))}px`
}

const handleTagClick = (tag: string) => {
  window.open("/search?tag=" + encodeURIComponent(tag))
}
</script>

<style scoped lang="scss">
.tags-page {
  .main-content {
    margin-top: 70px;
    display: flex;
    justify-content: center;

    .container {
      max-width: 1200px;
      width: 100%;

      .el-main {
        .header {
          margin-bottom: 48px;
          padding-bottom: 32px;
          border-bottom: 1px solid var(--line-strong);

          .meta-label {
            margin-bottom: 16px;

            &::before {
              content: '';
              display: inline-block;
              width: 24px;
              height: 1px;
              background: var(--accent);
              margin-right: 12px;
              vertical-align: middle;
            }
          }

          .title {
            font-size: clamp(40px, 6vw, 72px);
            line-height: 1;
            margin: 0 0 16px;
          }

          .sub {
            font-family: var(--font-mono);
            font-size: 14px;
            color: var(--ink-3);
            margin: 0;
          }
        }

        .cloud {
          display: flex;
          flex-wrap: wrap;
          align-items: baseline;
          gap: 8px 22px;
          padding: 8px 0;

          .tag {
            font-family: var(--font-serif);
            line-height: 1.25;
            color: var(--ink);
            cursor: pointer;
            padding: 2px;
            transition: color 0.2s ease;

            &:hover {
              color: var(--accent);
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .tags-page .main-content {
    padding: 0 16px;
  }
}
</style>
