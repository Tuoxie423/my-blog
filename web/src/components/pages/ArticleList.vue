<template>
  <el-card class="article-list">
    <el-row class="title">文章列表<span v-if="total > 0" class="count">共 {{ total }} 篇</span></el-row>

    <el-table :data="articleTableData" :show-header="false" :row-style="{height: '150px'}">
      <el-table-column label="index" width="64">
        <template #default="scope:{ row: any, column: any, $index: number }">
          <span class="index">{{ String(scope.$index + 1).padStart(2, '0') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="description">
        <template #default="scope:{ row: Hit<Article>, column: any, $index: number }">
          <div class="description" @click="handleArticleJumps(scope.row._id)">
            <el-row class="title">{{ scope.row._source.title }}</el-row>
            <el-text class="abstract" size="large">{{ scope.row._source.abstract }}</el-text>
            <el-text class="footer">
              <div class="tags">
                <el-tag v-for="item in scope.row._source.tags">{{ item }}</el-tag>
              </div>
              <div class="status">
                {{ scope.row._source.created_at }}
                <el-icon>
                  <component is="Star"/>
                </el-icon>
                {{ scope.row._source.likes }}
              </div>
            </el-text>
          </div>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="page_size"
        :page-sizes="[10, 30, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
    />
  </el-card>
</template>

<script setup lang="ts">
import type {Hit} from "@/api/common";
import {type Article, articleSearch, type ArticleSearchRequest} from "@/api/article";
import {reactive, ref} from "vue";

const articleSearchRequest = reactive<ArticleSearchRequest>({
  query: "",
  tag: "",
  page: 1,
  page_size: 10,
})

const page = ref(1)
const page_size = ref(10)
const total = ref(0)
const articleTableData = ref<Hit<Article>[]>()

const getArticleSearchTableData = async () => {
  articleSearchRequest.page = page.value;
  articleSearchRequest.page_size = page_size.value;

  const table = await articleSearch(articleSearchRequest)

  if (table.code === 0) {
    articleTableData.value = table.data.list;
    total.value = table.data.total;
  }
}

getArticleSearchTableData()

const handleArticleJumps = (id: string) => {
  window.open("/article/" + id)
}

const handleSizeChange = (val: number) => {
  page_size.value = val
  getArticleSearchTableData()
}

const handleCurrentChange = (val: number) => {
  page.value = val
  getArticleSearchTableData()
}
</script>

<style scoped lang="scss">
.article-list {
  .title {
    font-size: 24px;
    margin-bottom: 20px;

    .count {
      margin-left: 14px;
      font-family: var(--font-mono);
      font-size: 13px;
      font-weight: 400;
      color: var(--ink-3);
    }
  }

  .el-table {
    .index {
      font-family: var(--font-mono);
      font-size: 14px;
      color: var(--ink-3);
    }

    .description {
      height: 120px;
      display: flex;
      flex-direction: column;
      cursor: pointer;

      .title {
        font-size: 24px;
        margin-bottom: 10px;
        transition: color 0.2s ease;
      }

      &:hover .title {
        color: var(--accent);
      }

      .abstract {
        margin-right: auto;
        color: var(--ink-2);
      }

      .footer {
        margin-top: auto;
        display: flex;
        width: 100%;

        .tags {
          margin-right: auto;

          .el-tag {
            margin-right: 10px;
          }
        }

        .status {
          margin-left: auto;
          font-family: var(--font-mono);
          font-size: 13px;
          color: var(--ink-3);
        }
      }
    }
  }

  .el-pagination {
    margin-top: 10px;
    display: flex;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .article-list .el-table .description .footer {
    flex-direction: column;
    gap: 8px;

    .tags {
      margin-right: 0;
    }

    .status {
      margin-left: 0;
    }
  }
}
</style>
