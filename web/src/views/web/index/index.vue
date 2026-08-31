<template>
  <div class="index">
    <el-container class="main-content">
      <div class="container">
        <el-main>
          <section class="daily rise">
            <div class="quote-block">
              <p class="meta-label">今日一言</p>
              <blockquote class="quote-text">「{{ quote }}」</blockquote>
            </div>
            <router-link class="hot-link" to="/news">
              查看热榜 <span class="arrow">↗</span>
            </router-link>
          </section>

          <article-list/>
        </el-main>

        <el-aside>
          <profile-card/>
          <feedback/>
        </el-aside>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import ArticleList from "@/components/pages/ArticleList.vue";
import ProfileCard from "@/components/pages/ProfileCard.vue";
import Feedback from "@/components/pages/Feedback.vue";
import {websiteYiyan} from "@/api/website";
import {ref} from "vue";

const quote = ref('')

const getQuote = async () => {
  const res = await websiteYiyan()
  if (res.code === 0) {
    quote.value = res.data
  }
}

getQuote()
</script>

<style scoped lang="scss">
.index {
  .main-content {
    display: flex;
    justify-content: center;
    padding-top: 120px;

    .container {
      display: flex;
      max-width: 1400px;
      width: 100%;

      .el-main {
        width: 70%;

        .daily {
          margin-bottom: 32px;
          padding-bottom: 28px;
          border-bottom: 1px solid var(--line-strong);

          .quote-block {
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

            .quote-text {
              font-family: var(--font-serif);
              font-size: clamp(22px, 3vw, 32px);
              line-height: 1.5;
              margin: 0 0 16px;
              border-left: 3px solid var(--accent);
              padding-left: 20px;
            }

            .quote-author {
              font-family: var(--font-mono);
              font-size: 13px;
              color: var(--ink-3);
              margin: 0;
            }
          }

          .hot-link {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            margin-top: 20px;
            padding: 10px 20px;
            border: 1px solid var(--ink);
            font-family: var(--font-mono);
            font-size: 14px;
            color: var(--ink);
            text-decoration: none;
            transition: background-color 0.2s ease, color 0.2s ease;

            .arrow {
              color: var(--accent);
            }

            &:hover {
              background-color: var(--ink);
              color: var(--bg-panel);
            }
          }
        }
      }

      .el-aside {
        width: 30%;
        padding: 20px;
      }
    }
  }
}

@media (max-width: 768px) {
  .index .main-content {
    padding-top: 80px;

    .container {
      flex-direction: column;
      padding: 0 16px;

      .el-main {
        width: 100%;
      }

      .el-aside {
        width: 100%;
        padding: 20px 0;
      }
    }
  }
}
</style>
