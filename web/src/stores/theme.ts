import {defineStore} from 'pinia';
import {ref} from 'vue';

export const useThemeStore = defineStore('theme', () => {
    const isDark = ref(document.documentElement.getAttribute('data-theme') === 'dark');

    const toggle = () => {
        isDark.value = !isDark.value;
        const theme = isDark.value ? 'dark' : 'light';
        document.documentElement.setAttribute('data-theme', theme);
        localStorage.setItem('theme', theme);
    };

    return {isDark, toggle};
});
