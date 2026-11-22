<template>
  <div id="app">
    <el-config-provider :locale="zhCn">
      <MaintenancePage v-if="maintenance.offline"/>
      <router-view v-else/>
    </el-config-provider>
  </div>
</template>

<script setup lang="ts">
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import MaintenancePage from '@/components/common/MaintenancePage.vue'
import {useMaintenanceStore} from '@/stores/maintenance'

const maintenance = useMaintenanceStore()
</script>

<style scoped>
</style>
