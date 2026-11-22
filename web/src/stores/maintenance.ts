import {defineStore} from 'pinia'
import {ref} from 'vue'

// 自动重试间隔（秒）与单次健康检查超时（毫秒）
const RETRY_SECONDS = 60
const CHECK_TIMEOUT = 3000

export const useMaintenanceStore = defineStore('maintenance', () => {
    // 后端不可用时置为 true，App.vue 据此展示兜底页
    const offline = ref(false)
    const checking = ref(false)
    const countdown = ref(0)

    let countdownTimer: ReturnType<typeof setInterval> | null = null
    let checkingLock = false

    const healthUrl = () => {
        const base = import.meta.env.VITE_BASE_API || '/api'
        return `${base.replace(/\/$/, '')}/website/info`
    }

    const clearCountdown = () => {
        if (countdownTimer !== null) {
            clearInterval(countdownTimer)
            countdownTimer = null
        }
    }

    // 主动探测后端是否恢复；恢复后整页刷新，重新走一遍初始化
    const checkNow = async () => {
        if (checkingLock) return
        checkingLock = true
        checking.value = true
        clearCountdown()

        const controller = new AbortController()
        const abortTimer = setTimeout(() => controller.abort(), CHECK_TIMEOUT)
        let recovered = false
        try {
            const res = await fetch(healthUrl(), {cache: 'no-store', signal: controller.signal})
            if (res.ok) {
                const data = await res.json()
                recovered = data?.code === 0
            }
        } catch {
            recovered = false
        } finally {
            clearTimeout(abortTimer)
            checkingLock = false
            checking.value = false
        }

        if (recovered) {
            offline.value = false
            window.location.reload()
            return
        }
        startCountdown()
    }

    const startCountdown = () => {
        clearCountdown()
        countdown.value = RETRY_SECONDS
        countdownTimer = setInterval(() => {
            countdown.value -= 1
            if (countdown.value <= 0) {
                clearCountdown()
                void checkNow()
            }
        }, 1000)
    }

    // 请求层检测到后端不可用时调用
    const markOffline = () => {
        if (offline.value) return
        offline.value = true
        startCountdown()
    }

    const retryNow = () => {
        void checkNow()
    }

    return {offline, checking, countdown, markOffline, retryNow}
})
