<template>
  <div class="web">
    <web-navbar :noScroll="!isHome"/>
    <el-container>
      <router-view/>
      <el-footer>
        <web-footer/>
      </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import WebFooter from "@/components/layout/WebFooter.vue";
import WebNavbar from "@/components/layout/WebNavbar.vue";
import {computed} from "vue";
import {useRoute} from "vue-router";

const route = useRoute()
const isHome = computed(() => route.name === 'index')
</script>

<style scoped lang="scss">
.web {
  .el-footer {
    height: auto;
    border: 1px solid var(--line);
  }
}
</style>
