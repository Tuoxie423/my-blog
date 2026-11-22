<template>
  <div class="maintenance">
    <div class="card">
      <button
          class="theme-toggle"
          type="button"
          :title="themeStore.isDark ? '切换到浅色模式' : '切换到深色模式'"
          @click="themeStore.toggle()"
      >
        <el-icon>
          <component :is="themeStore.isDark ? 'Sun' : 'Moon'"/>
        </el-icon>
      </button>

      <img class="mark" src="/favicon.svg" alt=""/>

      <h1 class="title">服务暂时不可用</h1>
      <p class="desc">
        站点可能正在维护，或出现了临时故障。<br/>
        服务恢复后页面会自动刷新，请稍候。
      </p>

      <div v-if="email" class="contact-row">
        <span class="hint">长时间无法访问？欢迎发邮件提醒我：</span>
        <div class="contact-line">
          <button class="contact" type="button" title="点击发邮件" @click="openMail">
            {{ obfuscatedEmail }}
          </button>
          <button class="copy" type="button" @click="copyMail">{{ copied ? '已复制' : '复制' }}</button>
        </div>
      </div>

      <div class="actions">
        <button class="retry" type="button" :disabled="checking" @click="retryNow">
          {{ checking ? '正在连接…' : '立即重试' }}
        </button>
        <span class="countdown">
          {{ checking ? '正在检测服务状态' : `${countdown} 秒后自动重试` }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue'
import {useMaintenanceStore} from '@/stores/maintenance'
import {useThemeStore} from '@/stores/theme'

const maintenance = useMaintenanceStore()
const themeStore = useThemeStore()

// 联系方式在构建时以 base64 形式写入，运行时解码，避免源码/仓库里出现明文邮箱
const email = computed(() => {
  try {
    return atob(import.meta.env.VITE_CONTACT_EMAIL_B64 || '')
  } catch {
    return ''
  }
})

// 展示时进一步混淆，降低被爬虫抓取的概率
const obfuscatedEmail = computed(() => {
  const [user, domain] = email.value.split('@')
  if (!user || !domain) return ''
  return `${user} (at) ${domain.replace(/\./g, ' (dot) ')}`
})

const openMail = () => {
  if (email.value) {
    window.location.href = `mailto:${email.value}?subject=${encodeURIComponent('站点无法访问')}`
  }
}

const copied = ref(false)
const copyMail = async () => {
  if (!email.value) return
  try {
    await navigator.clipboard.writeText(email.value)
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  } catch {
    // 非安全上下文或没有剪贴板权限时，退化为打开邮件客户端
    openMail()
  }
}

const checking = computed(() => maintenance.checking)
const countdown = computed(() => maintenance.countdown)
const retryNow = () => maintenance.retryNow()
</script>

<style scoped lang="scss">
.maintenance {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 24px;
  background: var(--bg);
  color: var(--ink);
}

.card {
  position: relative;
  width: 100%;
  max-width: 560px;
  padding: 56px 40px 40px;
  border: 1px solid var(--line-strong);
  background: var(--bg-panel);
  text-align: center;
}

.theme-toggle {
  position: absolute;
  top: 12px;
  right: 12px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border: 1px solid var(--line);
  background: transparent;
  color: var(--ink-3);
  cursor: pointer;
  transition: color 0.15s, border-color 0.15s;

  &:hover {
    color: var(--ink);
    border-color: var(--ink-3);
  }
}

.mark {
  width: 56px;
  height: 56px;
  margin-bottom: 24px;
}

.title {
  margin: 0 0 16px;
  font-family: var(--font-serif);
  font-size: 32px;
  font-weight: 400;
  letter-spacing: 0.02em;
}

.desc {
  margin: 0;
  font-family: var(--font-mono);
  font-size: 14px;
  line-height: 1.9;
  color: var(--ink-2);
}

.contact-row {
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid var(--line);
}

.hint {
  display: block;
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--ink-3);
}

.contact-line {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-top: 10px;
  flex-wrap: wrap;
}

.contact {
  padding: 0;
  border: none;
  background: none;
  font-family: var(--font-mono);
  font-size: 14px;
  color: var(--accent);
  text-decoration: underline;
  text-underline-offset: 3px;
  cursor: pointer;
}

.copy {
  padding: 2px 8px;
  border: 1px solid var(--line);
  background: transparent;
  font-family: var(--font-mono);
  font-size: 11px;
  color: var(--ink-3);
  cursor: pointer;
  transition: color 0.15s, border-color 0.15s;

  &:hover {
    color: var(--ink);
    border-color: var(--ink-3);
  }
}

.actions {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
  margin-top: 40px;
  flex-wrap: wrap;
}

.retry {
  padding: 10px 28px;
  border: 1px solid var(--accent);
  background: var(--accent);
  font-family: var(--font-mono);
  font-size: 13px;
  letter-spacing: 0.06em;
  color: #fbfaf7;
  cursor: pointer;
  transition: opacity 0.15s;

  &:hover {
    opacity: 0.88;
  }

  &:disabled {
    cursor: default;
    opacity: 0.6;
  }
}

.countdown {
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--ink-3);
}

@media (max-width: 560px) {
  .card {
    padding: 40px 24px 32px;
  }

  .title {
    font-size: 26px;
  }
}
</style>
