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
              <div class="body">
                <p class="content">{{ m.content }}</p>
                <div class="actions">
                  <button
                      class="agree-btn"
                      :class="{ agreed: m.is_liked }"
                      :disabled="pending.has(m.id)"
                      :aria-pressed="m.is_liked"
                      :title="m.is_liked ? '取消赞同' : '赞同'"
                      @click="handleAgree(m)"
                  >
                    <component is="ThumbsUp" :size="16" :fill="m.is_liked ? 'currentColor' : 'none'"/>
                    <span>赞同</span>
                    <span class="count">{{ m.likes }}</span>
                  </button>
                </div>
              </div>
            </li>
          </ol>
        </el-main>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {murmurAll, murmurLike, murmurLiked, type MurmurWithLiked} from "@/api/murmur";
import {useUserStore} from "@/stores/user";
import {useLayoutStore} from "@/stores/layout";
import {ref, watch} from "vue";

const userStore = useUserStore()
const layoutStore = useLayoutStore()
const murmurs = ref<MurmurWithLiked[]>([])
const pending = ref<Set<number>>(new Set())

const isLoggedIn = () => userStore.state.userInfo.id !== 0

const mergeLiked = (ids: number[]) => {
  const likedSet = new Set(ids)
  murmurs.value.forEach(m => {
    m.is_liked = likedSet.has(m.id)
  })
}

const getMurmurs = async () => {
  // 列表与"我赞同过的"并行请求；未登录不请求后者
  const [listRes, likedRes] = await Promise.all([
    murmurAll(),
    isLoggedIn() ? murmurLiked() : Promise.resolve(null),
  ])

  if (listRes.code === 0) {
    murmurs.value = listRes.data.map(m => ({
      ...m,
      is_liked: false,
    }))
    if (likedRes && likedRes.code === 0) {
      mergeLiked(likedRes.data)
    }
  }
}

getMurmurs()

// 登录 / 退出后同步赞同状态
watch(() => userStore.state.userInfo.id, async (id) => {
  if (id === 0) {
    murmurs.value.forEach(m => {
      m.is_liked = false
    })
    return
  }
  const res = await murmurLiked()
  if (res.code === 0) {
    mergeLiked(res.data)
  }
})

const handleAgree = async (m: MurmurWithLiked) => {
  if (!isLoggedIn()) {
    layoutStore.state.popoverVisible = true
    layoutStore.state.loginVisible = true
    return
  }

  if (pending.value.has(m.id)) {
    return
  }

  pending.value.add(m.id)
  try {
    const res = await murmurLike({murmur_id: m.id})
    if (res.code === 0) {
      // 以后端返回的最终状态为准
      m.is_liked = res.data.is_liked
      m.likes = res.data.likes
    }
  } finally {
    pending.value.delete(m.id)
  }
}

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

            .body {
              display: flex;
              flex-direction: column;
              min-width: 0;

              .content {
                font-family: var(--font-serif);
                font-size: 20px;
                line-height: 1.7;
                color: var(--ink);
                margin: 0;
              }

              .actions {
                margin-top: 12px;
                display: flex;
                justify-content: flex-end;
              }

              .agree-btn {
                display: inline-flex;
                align-items: center;
                gap: 6px;
                padding: 5px 12px;
                border: 1px solid var(--line);
                border-radius: 999px;
                background: transparent;
                color: var(--ink-3);
                font-family: var(--font-mono);
                font-size: 13px;
                line-height: 1;
                cursor: pointer;
                transition: color .2s, border-color .2s;

                .count {
                  min-width: 1ch;
                  font-variant-numeric: tabular-nums;
                }

                &:hover:not(:disabled) {
                  color: var(--accent);
                  border-color: var(--accent);
                }

                &.agreed {
                  color: var(--accent);
                  border-color: var(--accent);
                }

                &:disabled {
                  opacity: .6;
                  cursor: default;
                }
              }
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
