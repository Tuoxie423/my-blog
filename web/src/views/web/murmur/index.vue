<template>
  <div class="murmur">
    <el-container class="main-content">
      <div class="container">
        <el-main>
          <div class="header rise">
            <p class="meta-label">Murmurs</p>
            <h1 class="title">碎碎念</h1>
            <p class="sub">共 {{ murmurs.length }} 条 · 按时间倒序</p>
          </div>

          <ol class="timeline">
            <li v-for="m in murmurs" :key="m.id" class="entry">
              <time class="date">{{ formatDate(m.created_at) }}</time>
              <p class="content">{{ m.content }}</p>
            </li>
          </ol>
        </el-main>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {murmurAll, type Murmur} from "@/api/murmur";
import {ref} from "vue";

const murmurs = ref<Murmur[]>([])

const getMurmurs = async () => {
  const res = await murmurAll()
  if (res.code === 0) {
    murmurs.value = res.data
  }
}

getMurmurs()

const formatDate = (date: string): string => {
  return date.slice(0, 10).replace(/-/g, '.')
}
</script>

<style scoped lang="scss">
.murmur {
  .main-content {
    margin-top: 70px;
    display: flex;
    justify-content: center;

    .container {
      max-width: 720px;
      width: 100%;

      .el-main {
        .header {
          margin-bottom: 40px;
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

        .timeline {
          list-style: none;
          margin: 0;
          padding: 0;

          .entry {
            display: grid;
            grid-template-columns: 96px 1fr;
            gap: 24px;
            padding: 26px 0;
            border-bottom: 1px solid var(--line);

            &:last-child {
              border-bottom: none;
            }

            .date {
              font-family: var(--font-mono);
              font-size: 13px;
              color: var(--ink-3);
              padding-top: 7px;
            }

            .content {
              font-family: var(--font-serif);
              font-size: 20px;
              line-height: 1.7;
              color: var(--ink);
              margin: 0;
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .murmur .main-content {
    padding: 0 16px;

    .timeline .entry {
      grid-template-columns: 1fr;
      gap: 8px;

      .date {
        padding-top: 0;
      }
    }
  }
}
</style>
