<template>
  <el-row>
    <el-col :span="12" v-for="item in videoList" :key="item.name">
      <el-card :body-style="{ padding: '0px' }">
        <img :src="`http://127.0.0.1:5000/cover_picture/${item.name}`" class="image">
        <div style="padding: 5px;">
          <div class="video-header">
            <span class="video-name">{{ item.name }}</span> <span class="video-size">{{ bytesFormat(item.bytesSize) }}</span>
          </div>
          <div class="video-tags">
            <el-tag v-for="tagName in item.tags" :key="tagName">标签一</el-tag>
          </div>
          <div>
            <el-button icon="el-icon-download" type="primary" @click="downloadVideo(item.name)" class="button">下载</el-button>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import axios from 'axios'
import { defineComponent, onMounted, ref } from 'vue'
import * as utils from '../hooks/utils'

interface videoModel {
  name: string,
  bytesSize: number,
  tags: string[],
}
interface videoList {
  data: videoModel[]
}

export default defineComponent({
  name: 'Home',
  setup () {
    const currentDate = new Date()
    const videoList = ref()
    const downloadVideo = (videoName: string) => {
      window.open(`http://127.0.0.1:5000/download/${videoName}`, '__blank')
    }
    const getVideos = async () => {
      const res: videoList = await axios.get('http://127.0.0.1:5000/videos')
      videoList.value = res.data
    }
    const bytesFormat = (bytesSize: number) => {
      return utils.bytesFormat(bytesSize)
    }
    onMounted(async () => {
      await getVideos()
    })
    return {
      currentDate,
      downloadVideo,
      videoList,
      bytesFormat
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
