/// <reference types="vite/client" />

import 'vue-router';

export interface ImportMetaEnv {
    VITE_SERVER_URL: string
    VITE_BASE_API: string
    VITE_CONTACT_EMAIL_B64: string
}


declare module 'vue-router' {
    interface RouteMeta {
        requiresAuth?: boolean
        requiresAdmin?: boolean
        title?: string
    }
}
