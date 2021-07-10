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
import DPlayer from 'dplayer'
import { useRoute } from 'vue-router'
import * as utils from '../hooks/utils'

export default defineComponent({
  name: 'Manage',
  setup () {
    const nowServerAddress = computed(() => {
      return utils.serverAddress()
    })
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
