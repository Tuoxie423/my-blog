import {defineStore} from "pinia";
import {ref, watch} from 'vue';

function initState() {
    const savedIsCollapse = localStorage.getItem('isCollapse');
    return {
        isCollapse: savedIsCollapse !== null ? savedIsCollapse === 'true' : window.innerWidth < 768,
        popoverVisible: false,
        loginVisible: false,
        registerVisible: false,
        forgotPasswordVisible: false,
        passwordResetVisible:false,
        shouldRefreshUserTable:false,
        shouldRefreshImageTable:false,
        articleCreateVisible:false,
        articleUpdateVisible:false,
        shouldRefreshArticleTable:false,
        feedbackReplyVisible:false,
        shouldRefreshFeedbackTable:false,
        shouldRefreshCommentTable:false,
        shouldRefreshCommentList:false,
    };
}

export const useLayoutStore = defineStore('layout', () => {
    const state = ref(initState());

    // 监视 isCollapse 的变化并更新 localStorage
    watch(() => state.value.isCollapse, (newIsCollapse) => {
        localStorage.setItem('isCollapse', String(newIsCollapse));
    });

    return {state};
});
