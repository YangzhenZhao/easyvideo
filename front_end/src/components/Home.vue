<template>
  <el-row>
    <el-col :span="12" v-for="o in 5" :key="o">
      <el-card :body-style="{ padding: '0px' }">
        <img src="http://192.168.0.102:5000/cover_picture/功夫" class="image">
        <div style="padding: 5px;">
          <div class="video-header">
            <span class="video-name">好吃的汉堡</span> <span class="video-size">3.5 GB</span>
          </div>
          <div class="video-tags">
            <el-tag v-for="i in 20" :key="i">标签一</el-tag>
          </div>
          <div>
            <el-button icon="el-icon-download" type="primary" @click="downloadVideo" class="button">下载</el-button>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Home',
  setup () {
    const currentDate = new Date()
    const downloadVideo = () => {
      window.open('http://192.168.0.102:5000/download/功夫', '__blank')
    }
    return {
      currentDate,
      downloadVideo
    }
  }
})
</script>

<style scoped>
  .image {
    width: 100%;
    height: 486px;
  }
  .el-col {
    padding: 10px;
  }
  .video-header {
    margin-top: 5px;
    margin-bottom: 5px;
  }
  .video-tags {
    margin-bottom: 5px;
  }
  .video-name {
    color: #3273dc;
  }
  .video-size {
    color: #7a7a7a;
  }
  .button {
    width: 100%;
  }
  .el-tag {
    margin-right: 5px;
    margin-bottom: 3px;
  }
</style>
