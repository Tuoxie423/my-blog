<template>
  <web-navbar :noScroll="true"/>
  <div class="error rise">
    <p class="meta-label">Error 404</p>
    <h1 class="code">404</h1>
    <p class="message">页面不存在或已被移除</p>
    <el-link class="back" href="/" :underline="false">返回首页 ↗</el-link>
  </div>
  <web-footer/>
</template>

<script setup lang="ts">
import WebNavbar from "@/components/layout/WebNavbar.vue";
import WebFooter from "@/components/layout/WebFooter.vue";
</script>

<style scoped lang="scss">
.error {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 70vh;
  padding: 120px 24px;

  .meta-label {
    margin-bottom: 24px;
  }

  .code {
    font-family: var(--font-serif);
    font-weight: 400;
    font-size: clamp(120px, 28vw, 320px);
    line-height: 0.85;
    margin: 0 0 24px;
    color: var(--ink);
  }

  .message {
    font-family: var(--font-mono);
    font-size: 15px;
    color: var(--ink-2);
    margin: 0 0 32px;
  }

  .back {
    font-family: var(--font-mono);
    font-size: 14px;
    color: var(--accent);
  }
}
</style>
