<template>
  <div class="container">
    <div>
      <h3>{{ videoName }}</h3>
    </div>
    <div id="dplayer"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, computed } from 'vue'
import { useStore } from 'vuex'
import DPlayer from 'dplayer'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'Manage',
  setup () {
    const store = useStore()
    const nowServerAddress = computed(() => store.state.serverAddress)
    const route = useRoute()
    const videoName = route.params.video_name
    onMounted(() => {
      const dp = new DPlayer({
        container: document.getElementById('dplayer'),
        video: {
          url: `${nowServerAddress.value}/download/${videoName}`,
          hotkey: true
        }
      })
    })
    return {
      nowServerAddress,
      videoName
    }
  }
})
</script>

<style scoped>
  .el-header, .el-footer {
    color: #333;
    line-height: 60px;
  }
  h3 {
    margin-left: 30px;
  }
</style>
