<template>
  <div class="logo">
    <img
        v-if="currentLogo"
        class="logo-image"
        :src="currentLogo"
        alt="logo"
        @error="handleError"
    />
    <el-icon v-else class="logo-icon">
      <component is="Feather"/>
    </el-icon>
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";
import {useWebsiteStore} from "@/stores/website";

const websiteStore = useWebsiteStore()

// 图片加载失败时回退到默认图标
const failed = ref(false)

const currentLogo = computed(() => {
  if (failed.value) return ''
  return websiteStore.state.websiteInfo.logo || ''
})

const handleError = () => {
  failed.value = true
}
</script>

<style scoped lang="scss">
.logo {
  display: flex;
  align-items: center;
  padding: 10px;

  .logo-icon {
    font-size: 28px;
    color: var(--accent);
  }

  .logo-image {
    height: 40px;
    width: auto;
    object-fit: contain;
  }
}

/* 前台导航栏 */
.web-navbar .logo {
  padding: 5px 0 5px 20px;

  .logo-icon {
    font-size: 22px;
  }

  .logo-image {
    height: 36px;
  }
}

/* 侧边栏折叠时图标居中 */
.collapsed .logo {
  justify-content: center;
}
</style>
