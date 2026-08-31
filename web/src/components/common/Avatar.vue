<template>
  <img
      v-if="!showInitial"
      class="avatar-image"
      :style="boxStyle"
      :src="currentSrc"
      :alt="name"
      @error="handleError"
  />
  <span v-else class="avatar-initial" :style="boxStyle">{{ name.slice(0, 1) }}</span>
</template>

<script setup lang="ts">
import {computed, ref, watch} from "vue";

const DEFAULT_AVATAR = '/image/avatar.png'

const props = withDefaults(defineProps<{
  src?: string
  name: string
  size?: number
}>(), {
  src: '',
  size: 40,
})

// useDefault：自定义头像加载失败后改用默认头像
const useDefault = ref(false)
// showInitial：默认头像也加载失败，回退到首字母块
const showInitial = ref(false)

const currentSrc = computed(() => {
  if (!props.src || useDefault.value) return DEFAULT_AVATAR
  return props.src
})

watch(() => props.src, () => {
  useDefault.value = false
  showInitial.value = false
})

const boxStyle = computed(() => ({
  width: props.size + 'px',
  height: props.size + 'px',
  fontSize: Math.round(props.size * 0.4) + 'px',
}))

const handleError = () => {
  if (props.src && !useDefault.value && props.src !== DEFAULT_AVATAR) {
    useDefault.value = true
  } else {
    showInitial.value = true
  }
}
</script>

<style scoped>
.avatar-image {
  object-fit: cover;
  border-radius: 0;
  flex-shrink: 0;
}
</style>
