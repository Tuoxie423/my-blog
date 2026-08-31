<template>
  <div :class="{'web-navbar': true, show: isShow}">
    <div class="container">
      <logo/>
      <div class="web-menu">
        <el-menu mode="horizontal" :ellipsis="false" :router="true" :default-active="$route.path">
          <template v-for="item in menuList">
            <el-menu-item :index="item.name"><span>{{ item.title }}</span></el-menu-item>
          </template>
        </el-menu>
      </div>
      <div class="right-actions">
        <button class="mobile-toggle" @click="mobileMenuVisible = !mobileMenuVisible" title="菜单">
          <el-icon>
            <component :is="mobileMenuVisible ? 'X' : 'Menu'"/>
          </el-icon>
        </button>
        <button class="theme-toggle" @click="themeStore.toggle()" title="切换深色模式">
          <el-icon>
            <component :is="themeStore.isDark ? 'Sun' : 'Moon'"/>
          </el-icon>
        </button>
        <auth-popover/>
      </div>
    </div>

    <transition name="mobile-menu-fade">
      <nav v-if="mobileMenuVisible" class="mobile-menu">
        <router-link
            v-for="item in menuList"
            :key="item.name"
            :to="item.name"
            class="mobile-menu-item"
            @click="mobileMenuVisible = false"
        >{{ item.title }}</router-link>
      </nav>
    </transition>
  </div>
</template>

<script setup lang="ts">
import AuthPopover from "@/components/common/AuthPopover.vue";
import Logo from "@/components/widgets/Logo.vue";
import {ref, watch} from "vue";
import {onUnmounted} from "vue";
import {useThemeStore} from "@/stores/theme";

const themeStore = useThemeStore()

const isShow = ref(true)
const mobileMenuVisible = ref(false)

const props = defineProps<{
  noScroll?: boolean
}>()

function scroll() {
  let top = document.documentElement.scrollTop
  isShow.value = top >= 100;
}

watch(() => props.noScroll, (noScroll) => {
  if (noScroll) {
    isShow.value = true
    window.removeEventListener("scroll", scroll)
  } else {
    window.addEventListener("scroll", scroll)
    scroll()
  }
}, {immediate: true})

onUnmounted(() => {
  window.removeEventListener("scroll", scroll)
})

interface MenuItem {
  title: string;
  name: string;
}

const menuList: MenuItem[] = [
  {
    title: "首页",
    name: "/",
  },
  {
    title: "搜索",
    name: "/search",
  },
  {
    title: "热榜",
    name: "/news",
  },
  {
    title: "标签",
    name: "/friend-link",
  },
  {
    title: "碎碎念",
    name: "/murmur",
  },
  {
    title: "关于",
    name: "/about",
  }
]

</script>


<style scoped lang="scss">
.web-navbar {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  position: fixed;
  z-index: 6;
  color: var(--ink-3);
  --el-menu-text-color: var(--ink-3);
  --color: var(--ink-3);

  &.show {
    top: 0;
    background-color: var(--bg-panel);
    color: var(--ink);
    --el-menu-text-color: var(--ink);
    --color: var(--ink);

    .container {
      margin-top: 8px;
    }
  }

  .container {
    display: flex;
    max-width: 1400px;
    width: 100%;

    .logo {
      height: 60px;
      width: auto;
    }

    .web-menu {
      margin-left: 20px;

      .el-menu {
        background-color: transparent;
        border-bottom: none;
        --el-menu-item-font-size: 20px;

        .el-menu-item {
          border-bottom: none;
          background-color: transparent;
        }
      }
    }

    .right-actions {
      margin-left: auto;
      display: flex;
      align-items: center;
      gap: 12px;
      padding-right: 20px;

      .auth-popover {
        margin-top: auto;
        margin-bottom: auto;
      }

      .mobile-toggle,
      .theme-toggle {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 36px;
        height: 36px;
        border: 1px solid var(--line);
        background: transparent;
        color: var(--ink-2);
        cursor: pointer;
        transition: border-color 0.2s ease, color 0.2s ease;

        &:hover {
          border-color: var(--ink);
          color: var(--accent);
        }
      }

      .mobile-toggle {
        display: none;
      }
    }
  }

  .mobile-menu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    display: flex;
    flex-direction: column;
    background: var(--bg-panel);
    border-bottom: 1px solid var(--line-strong);

    .mobile-menu-item {
      padding: 16px 24px;
      font-family: var(--font-serif);
      font-size: 20px;
      color: var(--ink);
      text-decoration: none;
      border-bottom: 1px solid var(--line);

      &:active {
        color: var(--accent);
      }
    }
  }
}

.mobile-menu-fade-enter-active,
.mobile-menu-fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.mobile-menu-fade-enter-from,
.mobile-menu-fade-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}

/* 移动端 */
@media (max-width: 768px) {
  .web-navbar {
    .container {
      .logo {
        width: auto;
        padding-left: 12px;
      }

      .web-menu {
        display: none;
      }

      .right-actions {
        gap: 8px;
        padding-right: 12px;

        .mobile-toggle {
          display: flex;
        }
      }
    }
  }
}
</style>
