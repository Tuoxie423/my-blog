<template>
  <div class="hot">
    <el-container class="main-content">
      <div class="container">
        <el-main>
          <!-- 概览：各平台热榜容器 -->
          <template v-if="!currentType">
            <div class="header rise">
              <p class="meta-label">Hot Ranking</p>
              <h1 class="title">热榜</h1>
              <p class="sub">{{ platforms.length }} 个平台 · 点击进入查看完整热搜</p>
            </div>

            <div class="platforms">
              <div v-for="p in platforms" :key="p.type" class="platform-card" @click="openPlatform(p)">
                <div class="card-header">
                  <div class="brand">
                    <img class="icon" :src="iconUrl(p.icon)" alt="" @error="onIconError"/>
                    <span class="name">{{ p.name }}</span>
                  </div>
                  <span class="arrow">↗</span>
                </div>
                <ol class="preview">
                  <li v-for="item in p.list.slice(0, 3)" :key="item.index">
                    <span class="rank">{{ String(item.index).padStart(2, '0') }}</span>
                    <span class="title">{{ item.title }}</span>
                  </li>
                </ol>
              </div>
            </div>
          </template>

          <!-- 详情：某平台热搜 -->
          <template v-else>
            <div class="detail-header rise">
              <button class="back" @click="back">← 返回热榜</button>
              <h2 class="title">{{ currentName }}</h2>
              <p class="sub">共 {{ currentList.length }} 条热搜</p>
            </div>

            <ol class="hot-list">
              <li v-for="(item, i) in currentList" :key="i" class="hot-item" @click="handleItemClick(item)">
                <span class="rank" :class="{top: i < 3}">{{ String(item.index).padStart(2, '0') }}</span>
                <span class="title">{{ item.title }}</span>
                <span class="hot-value">{{ item.hot_value }}</span>
              </li>
            </ol>
          </template>
        </el-main>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import {hotAll, type HotPlatform, type HotItem} from "@/api/hot";
import {ref} from "vue";

const platforms = ref<HotPlatform[]>([])
const currentType = ref('')
const currentName = ref('')
const currentList = ref<HotItem[]>([])

const iconUrl = (icon: string): string => {
  if (icon.startsWith('http') || icon.startsWith('/')) return icon
  return '/image/' + icon
}

// 后端返回的图标加载失败时的兜底图标
const DEFAULT_ICON = '/image/hot.png'
const onIconError = (e: Event) => {
  const img = e.target as HTMLImageElement
  if (img.getAttribute('src') !== DEFAULT_ICON) {
    img.setAttribute('src', DEFAULT_ICON)
  }
}

const getHot = async () => {
  const res = await hotAll()
  if (res.code === 0) {
    platforms.value = res.data
  }
}

getHot()

const openPlatform = (p: HotPlatform) => {
  currentType.value = p.type
  currentName.value = p.name
  currentList.value = p.list
}

const back = () => {
  currentType.value = ''
  currentName.value = ''
  currentList.value = []
}

const handleItemClick = (item: HotItem) => {
  window.open(item.url)
}
</script>

<style scoped lang="scss">
.hot {
  .main-content {
    margin-top: 70px;
    display: flex;
    justify-content: center;

    .container {
      max-width: 1100px;
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

        .platforms {
          display: grid;
          grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
          gap: 0;

          .platform-card {
            padding: 28px 24px;
            border: 1px solid var(--line);
            margin: -1px 0 0 -1px;
            cursor: pointer;
            transition: border-color 0.2s ease;

            .card-header {
              display: flex;
              justify-content: space-between;
              align-items: center;
              margin-bottom: 20px;

              .brand {
                display: flex;
                align-items: center;
                gap: 12px;
                min-width: 0;

                .icon {
                  width: 30px;
                  height: 30px;
                  flex-shrink: 0;
                  border-radius: 8px;
                  object-fit: contain;
                }

                .name {
                  font-family: var(--font-serif);
                  font-size: 26px;
                  transition: color 0.2s ease;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  white-space: nowrap;
                }
              }

              .arrow {
                font-family: var(--font-mono);
                color: var(--ink-3);
              }
            }

            .preview {
              list-style: none;
              margin: 0;
              padding: 0;

              li {
                display: flex;
                align-items: baseline;
                gap: 12px;
                padding: 6px 0;
                border-bottom: 1px solid var(--line);

                &:last-child {
                  border-bottom: none;
                }

                .rank {
                  font-family: var(--font-mono);
                  font-size: 12px;
                  color: var(--ink-3);
                  flex-shrink: 0;
                }

                .title {
                  font-family: var(--font-mono);
                  font-size: 13px;
                  color: var(--ink-2);
                  overflow: hidden;
                  text-overflow: ellipsis;
                  white-space: nowrap;
                }
              }
            }

            &:hover {
              border-color: var(--line-strong);
              z-index: 1;

              .name {
                color: var(--accent);
              }
            }
          }
        }

        .detail-header {
          margin-bottom: 40px;
          padding-bottom: 28px;
          border-bottom: 1px solid var(--line-strong);

          .back {
            font-family: var(--font-mono);
            font-size: 13px;
            color: var(--ink-3);
            background: none;
            border: 1px solid var(--line);
            padding: 8px 14px;
            cursor: pointer;
            margin-bottom: 24px;
            transition: border-color 0.2s ease, color 0.2s ease;

            &:hover {
              border-color: var(--ink);
              color: var(--accent);
            }
          }

          .title {
            font-size: clamp(32px, 5vw, 56px);
            margin: 0 0 12px;
          }

          .sub {
            font-family: var(--font-mono);
            font-size: 14px;
            color: var(--ink-3);
            margin: 0;
          }
        }

        .hot-list {
          list-style: none;
          margin: 0;
          padding: 0;

          .hot-item {
            display: flex;
            align-items: baseline;
            gap: 20px;
            padding: 20px 8px;
            border-bottom: 1px solid var(--line);
            cursor: pointer;
            transition: background-color 0.2s ease;

            .rank {
              font-family: var(--font-mono);
              font-size: 16px;
              color: var(--ink-3);
              flex-shrink: 0;
              width: 28px;

              &.top {
                color: var(--accent);
                font-weight: 500;
              }
            }

            .title {
              font-family: var(--font-serif);
              font-size: 19px;
              line-height: 1.4;
              flex: 1;
              transition: color 0.2s ease;
            }

            .hot-value {
              font-family: var(--font-mono);
              font-size: 13px;
              color: var(--ink-3);
              flex-shrink: 0;
            }

            &:hover {
              background-color: var(--accent-soft);

              .title {
                color: var(--accent);
              }
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .hot .main-content {
    padding: 0 16px;

    .hot-list .hot-item {
      gap: 12px;

      .hot-value {
        display: none;
      }
    }
  }
}
</style>
