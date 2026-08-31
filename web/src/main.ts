import {createApp} from 'vue'
import '@fontsource/instrument-serif/400.css'
import '@fontsource/instrument-serif/400-italic.css'
import '@fontsource/jetbrains-mono/400.css'
import '@fontsource/jetbrains-mono/500.css'
import '@fontsource/jetbrains-mono/700.css'
import '@/assets/base.css'
import '@/assets/theme.css'
import {
    Search, Plus, Trash2, Eye, Star, ArrowDown, ArrowUp, User, PenLine,
    Upload, Home, Monitor, Contact, MessageSquare, MessageCircle, UserCog,
    FileText, Newspaper, Files, Image, Settings, Link, ScrollText, Megaphone,
    PanelLeftClose, PanelLeftOpen, LockKeyholeOpen, Mail, Sun, Moon, Menu, X, Feather
} from 'lucide-vue-next'
import App from './App.vue'
import router from './router'
import {pinia} from "@/stores";

const app = createApp(App)

// 用 Lucide 图标替换 Element Plus 图标；保留原有组件名，模板无需改动
const icons: Record<string, any> = {
    Search, Plus, Delete: Trash2, View: Eye, Star, SortDown: ArrowDown, SortUp: ArrowUp,
    User, Edit: PenLine, Unlock: LockKeyholeOpen, UploadFilled: Upload, Mail,
    House: Home, Monitor, Postcard: Contact, Message: MessageSquare, ChatDotRound: MessageCircle,
    SetUp: UserCog, Document: FileText, Collection: Newspaper, ChatLineRound: MessageSquare,
    DocumentCopy: Files, Picture: Image, PictureRounded: Image, Coin: Settings,
    Position: MessageSquare, Connection: Megaphone, Link, memo: ScrollText, setting: Settings,
    Fold: PanelLeftClose, Expand: PanelLeftOpen, Sun, Moon, Menu, X, Feather,
}

for (const [name, component] of Object.entries(icons)) {
    app.component(name, component)
}
app.use(pinia).use(router)

app.mount('#app')
