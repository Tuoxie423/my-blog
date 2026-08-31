<template>
  <footer class="web-footer">
    <div class="container">
      <div class="footer-top">
        <div class="brand">
          <span class="name">{{ websiteStore.state.websiteInfo.title }}</span>
          <span class="description">{{ websiteStore.state.websiteInfo.description }}</span>
        </div>
        <nav class="footer-links">
          <el-link v-if="websiteStore.state.websiteInfo.github_url" :href="websiteStore.state.websiteInfo.github_url" target="_blank" :underline="false">
            GitHub
          </el-link>
          <el-link v-if="websiteStore.state.websiteInfo.gitee_url" :href="websiteStore.state.websiteInfo.gitee_url" target="_blank" :underline="false">
            Gitee
          </el-link>
          <el-link v-for="item in footerLinkList" :key="item.title" :href="item.link" :underline="false">
            {{ item.title }}
          </el-link>
        </nav>
      </div>

      <div class="footer-bottom">
        <el-link class="filing" href="https://beian.miit.gov.cn/#/Integrated/index" :underline="false">
          {{ websiteStore.state.websiteInfo.icp_filing }}
        </el-link>
        <span class="filing">{{ websiteStore.state.websiteInfo.public_security_filing }}</span>
        <span class="meta-label">建站 {{ websiteStore.state.websiteInfo.created_at }} · 已运行 {{ elapsedTime }}</span>
        <span class="version">{{ websiteStore.state.websiteInfo.version }}</span>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import {useWebsiteStore} from "@/stores/website";
import {ref} from "vue";
import {onUnmounted} from "vue";
import {type FooterLink, websiteFooterLink} from "@/api/website";

const footerLinkList = ref<FooterLink[]>([])

const getFooterLinkList = async () => {
  const res = await websiteFooterLink()
  if (res.code === 0) {
    footerLinkList.value = res.data
  }
}

getFooterLinkList()

const websiteStore = useWebsiteStore()

let timerId: number | null = null;
const elapsedTime = ref("");

function updateElapsedTime() {
  let creationDate = websiteStore.state.websiteInfo.created_at;
  if (creationDate) {
    let creationTimestamp = new Date(creationDate).getTime();
    let currentTimestamp = new Date().getTime();
    let totalDays = (currentTimestamp - creationTimestamp) / 1000 / (60 * 60 * 24);
    let daysPassed = Math.floor(totalDays);
    let hoursRemaining = Math.floor((totalDays - daysPassed) * 24);
    let minutesRemaining = Math.floor((totalDays - daysPassed - (hoursRemaining / 24)) * 24 * 60);
    let secondsRemaining = Math.floor((totalDays - daysPassed - (hoursRemaining / 24) - (minutesRemaining / 24 / 60)) * 24 * 60 * 60);
    elapsedTime.value = `${daysPassed}天${hoursRemaining}时${minutesRemaining}分${secondsRemaining}秒`;
  }
}

function initializeTimer() {
  updateElapsedTime();
  timerId = setInterval(updateElapsedTime, 1000);
}

onUnmounted(() => {
  clearInterval(timerId as number);
});

initializeTimer();
</script>

<style scoped lang="scss">
.web-footer {
  display: flex;
  justify-content: center;
  padding: 48px 24px;
  border-top: 1px solid var(--line-strong);

  .container {
    display: flex;
    flex-direction: column;
    gap: 32px;
    max-width: 1400px;
    width: 100%;
  }

  .footer-top {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 32px;

    .brand {
      display: flex;
      flex-direction: column;
      gap: 8px;

      .name {
        font-family: var(--font-serif);
        font-size: 28px;
      }

      .description {
        font-family: var(--font-mono);
        font-size: 13px;
        color: var(--ink-3);
      }
    }

    .footer-links {
      display: flex;
      flex-wrap: wrap;
      gap: 20px;

      .el-link {
        font-family: var(--font-mono);
        font-size: 13px;
        letter-spacing: 0.06em;
        text-transform: uppercase;
        color: var(--ink-2);
      }
    }
  }

  .footer-bottom {
    display: flex;
    flex-wrap: wrap;
    gap: 24px;
    align-items: center;
    padding-top: 24px;
    border-top: 1px solid var(--line);

    .filing {
      font-family: var(--font-mono);
      font-size: 12px;
      color: var(--ink-3);
    }

    .version {
      margin-left: auto;
      font-family: var(--font-mono);
      font-size: 12px;
      color: var(--ink);
      border: 1px solid var(--line-strong);
      padding: 2px 8px;
    }
  }
}

@media (max-width: 768px) {
  .web-footer {
    .container {
      gap: 24px;
    }

    .footer-top {
      flex-direction: column;
      gap: 24px;
    }

    .footer-bottom {
      .version {
        margin-left: 0;
      }
    }
  }
}
</style>
