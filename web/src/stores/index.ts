import {createPinia} from 'pinia'

const pinia = createPinia()

// 注册插件
pinia.use(({ store }) => {
    if (store.$id === 'website') {
        // 确保在 store 创建时调用；失败时由请求拦截器触发全局维护页
        store.initializeWebsite().catch(() => {
        })
    }
})

export {
    pinia
}
