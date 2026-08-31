<template>
  <div class="dashboard">
    <el-container>
      <el-aside :class="{collapsed: isCollapse}">
        <logo/>
        <dashboard-menu/>
      </el-aside>
      <el-container>
        <el-header>
          <div class="header-top">
            <breadcrumb/>
            <div class="header-top-right">
              <button class="theme-toggle" @click="themeStore.toggle()" title="切换深色模式">
                <el-icon>
                  <component :is="themeStore.isDark ? 'Sun' : 'Moon'"/>
                </el-icon>
              </button>
              <auth-popover/>
            </div>
          </div>
          <dashboard-tag/>
        </el-header>
        <el-main>
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import Logo from "@/components/widgets/Logo.vue";
import Breadcrumb from "@/components/layout/Breadcrumb.vue";
import DashboardMenu from "@/components/layout/DashboardMenu.vue";
import {computed} from "vue";
import DashboardTag from "@/components/layout/DashboardTag.vue";
import AuthPopover from "@/components/common/AuthPopover.vue";
import {useLayoutStore} from "@/stores/layout";
import {useThemeStore} from "@/stores/theme";

const store = useLayoutStore()
const themeStore = useThemeStore()
const isCollapse = computed(() => store.state.isCollapse)
</script>

<style scoped lang="scss">
.dashboard {
  display: flex;

  .el-aside {
    border: 1px solid var(--line);
    width: 240px;
    height: 100vh;
    &::-webkit-scrollbar{
      display: none;
    }
  }

  .el-aside.collapsed {
    width: 64px;
  }

  .el-header {
    height: auto;
    border: 1px solid var(--line);

    .header-top {
      display: flex;
      border-bottom: 1px solid var(--line);

      .header-top-right {
        margin-left: auto;
        margin-top: auto;
        margin-bottom: auto;
        display: flex;
        align-items: center;
        gap: 12px;
        padding-right: 20px;

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
      }
    }
  }

  .el-main{
    height: calc(100vh - 100px);
    &::-webkit-scrollbar{
      display: none;
    }
  }
}
</style>
