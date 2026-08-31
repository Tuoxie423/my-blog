<template>
  <div class="article">
    <web-navbar :noScroll="true"/>
    <el-container class="main-content">
      <div class="container">
        <el-main>
          <div class="info rise">
            <h1 class="title">{{ articleInfo.title }}</h1>
            <div class="meta">
              <span class="category">{{ articleInfo.category }}</span>
              <span>发布 {{ articleInfo.created_at }}</span>
              <span>更新 {{ articleInfo.updated_at }}</span>
            </div>
            <div class="tags">
              <el-tag v-for="item in articleInfo.tags" :key="item" effect="plain">{{ item }}</el-tag>
            </div>
            <div class="abstract">{{ articleInfo.abstract }}</div>
          </div>
          <MdPreview :id="mdID" :modelValue="articleInfo.content"/>
          <comment-section :article-id="articleID"/>
        </el-main>

        <el-aside>
          <div class="aside-content">
            <div class="catalog">
              <el-row class="title">目录</el-row>
              <MdCatalog :editorId="mdID" :scrollElement="scrollElement" :offsetTop="100" :scrollElementOffsetTop="80"/>
            </div>
            <div class="status">
              <el-icon size="24">
                <component is="View"/>
              </el-icon>
              {{ articleInfo.views }}
              <el-icon size="24" @click="handleLike" :color="isLike ? 'var(--accent)' : 'inherit'">
                <component is="Star" :fill="isLike ? 'currentColor' : 'none'"/>
              </el-icon>
              {{ articleInfo.likes }}
            </div>
          </div>
        </el-aside>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {useRoute} from "vue-router";
import {type Article, articleInfoByID} from "@/api/article";
import router from "@/router";
import {computed, ref} from "vue";
import {MdPreview, MdCatalog} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import 'md-editor-v3/lib/preview.css';
import WebNavbar from "@/components/layout/WebNavbar.vue";
import {articleIsLike, articleLike, type ArticleLikeRequest} from "@/api/article";
import {useUserStore} from "@/stores/user";
import CommentSection from "@/components/common/CommentSection.vue";

const mdID = "md-id"

const articleInfo = ref<Article>({
  created_at: '',
  updated_at: '',
  cover: '',
  title: '',
  keyword: '',
  category: '',
  tags: [],
  abstract: '',
  content: '',
  comments: 0,
  views: 0,
  likes: 0,
})

const scrollElement = document.documentElement

const route = useRoute()

const articleID = computed(() => route.params.id as string)

const getArticleInfo = async () => {
  const res = await articleInfoByID(articleID.value as string)
  if (res.code === 0) {
    articleInfo.value = res.data
  } else {
    await router.push({name: "404"})
  }
}

getArticleInfo()

const isLike = ref(false)

const getIsLikeInfo = async () => {
  const req: ArticleLikeRequest = {
    article_id: articleID.value as string
  }
  const res = await articleIsLike(req)
  if (res.code === 0) {
    isLike.value = res.data
  }
}

if (useUserStore().state.userInfo.role_id !== 0) {
  getIsLikeInfo()
}

const handleLike = async () => {
  const req: ArticleLikeRequest = {
    article_id: articleID.value as string
  }
  const res = await articleLike(req)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    articleInfo.value.likes += isLike.value ? -1 : 1
    isLike.value = !isLike.value
  }
}
</script>

<style scoped lang="scss">
.article {
  .main-content {
    margin-top: 70px;
    display: flex;
    justify-content: center;

    .container {
      display: flex;
      max-width: 1400px;
      width: 100%;

      .el-main {
        width: 70%;

        .info {
          border-bottom: 1px solid var(--line);
          padding-bottom: 28px;
          margin-bottom: 28px;

          .title {
            font-size: clamp(28px, 4vw, 44px);
            line-height: 1.15;
            margin: 0 0 20px;
          }

          .meta {
            display: flex;
            flex-wrap: wrap;
            gap: 20px;
            font-family: var(--font-mono);
            font-size: 13px;
            color: var(--ink-3);
            margin-bottom: 16px;

            .category {
              color: var(--accent);
              letter-spacing: 0.06em;
              text-transform: uppercase;
            }
          }

          .tags {
            display: flex;
            flex-wrap: wrap;
            gap: 8px;
            margin-bottom: 16px;
          }

          .abstract {
            font-family: var(--font-serif);
            font-size: 18px;
            line-height: 1.6;
            color: var(--ink-2);
            border-left: 3px solid var(--accent);
            padding-left: 16px;
          }
        }
      }

      .el-aside {
        width: 30%;
        padding-top: 20px;

        .aside-content {
          position: fixed;

          .catalog {
            width: 420px;
            height: 70vh;
            overflow: auto;
            padding: 20px;
            border: 1px solid var(--line);

            .title {
              font-size: 24px;
              margin-bottom: 10px;
            }
          }

          .status {
            justify-content: center;
            display: flex;
            align-items: center;
            padding: 20px;
            border-left: 1px solid var(--line);
            border-right: 1px solid var(--line);
            border-bottom: 1px solid var(--line);
            font-family: var(--font-mono);
            color: var(--ink-2);

            .el-icon {
              margin-left: 20px;
              margin-right: 8px;
              cursor: pointer;
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .article .main-content {
    padding: 0 16px;

    .container {
      flex-direction: column;

      .el-main {
        width: 100%;
      }

      .el-aside {
        display: none;
      }
    }
  }
}
</style>
