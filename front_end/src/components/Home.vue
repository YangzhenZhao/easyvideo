<template>
  <div>
    <div class="tags-selected">
      <span>已选中标签:</span>
      <el-tag
        :key="tag"
        v-for="tag in selectedTags"
        closable
        @close="handleClose(tag)"
        :disable-transitions="false"
      >
        {{ tag }}
      </el-tag>
      <el-autocomplete
        v-if="inputVisible"
        class="input-new-tag"
        v-model="inputValue"
        ref="saveTagInput"
        size="small"
        :fetch-suggestions="querySearch"
        placeholder="请输入标签名称"
        @select="handleSelect"
      >
      </el-autocomplete>
      <el-button v-else class="button-new-tag" size="small" @click="showInput">添加标签</el-button>
    </div>
    <el-row>
      <el-col :span="8" v-for="item in videoList" :key="item.name">
        <el-card :body-style="{ padding: '0px' }">
          <img :src="`${nowServerAddress}/cover_picture/${item.name}`" class="image">
          <div style="padding: 5px;">
            <div class="video-header">
              <span class="video-name">{{ item.name }}</span> <span class="video-size">{{ bytesFormat(item.bytesSize) }}</span>
            </div>
            <div class="video-tags">
              <el-tag v-for="tagName in item.tags" :key="tagName">{{ tagName }}</el-tag>
            </div>
            <div class='donwload-play'>
              <el-button icon="el-icon-download" type="primary" @click="downloadVideo(item.name)">下载</el-button>
              <el-button icon="el-icon-video-play" type="primary" @click="playVideo(item.name)">在线播放</el-button>
              <el-button icon="el-icon-delete" type="primary" @click="deleteVideo(item.name)">删除</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import axios from 'axios'
import { defineComponent, onMounted, ref, Ref, watch, computed } from 'vue'
import * as utils from '../hooks/utils'
import { ElMessage } from 'element-plus'

interface videoModel {
  name: string,
  bytesSize: number,
  tags: string[],
}
interface videoList {
  data: videoModel[]
}
interface tagList {
  data: string[]
}

export default defineComponent({
  name: 'Home',
  setup () {
    const nowServerAddress = computed(() => {
      return utils.serverAddress()
    })
    const currentDate = new Date()
    const videoList = ref()
    const selectedTags: Ref<string[]> = ref([])
    const inputVisible = ref(false)
    const inputValue = ref('')
    const allTags: Ref<any[]> = ref([])
    const downloadVideo = (videoName: string) => {
      window.open(`${nowServerAddress.value}/download/${videoName}`, '__blank')
    }
    const playVideo = (videoName: string) => {
      window.open(`play-video/${videoName}`, '__blank')
    }
    const getVideos = async () => {
      const res: videoList = await axios.get(`${nowServerAddress.value}/videos`)
      videoList.value = res.data
    }
    const deleteVideo = async (videoName: string) => {
      const response = await axios.delete(`${nowServerAddress.value}/video/${videoName}`)
      console.log('---- delete response: ', response)
    }
    const bytesFormat = (bytesSize: number) => {
      return utils.bytesFormat(bytesSize)
    }
    onMounted(async () => {
      await getVideos()
    })
    const showInput = () => {
      inputVisible.value = true
    }
    const querySearch = (queryString: string, cb: any) => {
      var results = queryString
        ? allTags.value.filter(createFilter(queryString))
        : allTags.value
      cb(results)
    }
    const createFilter = (queryString: string) => {
      return (tagName: any) => {
        return (
          tagName.value.toLowerCase().indexOf(queryString.toLowerCase()) ===
          0
        )
      }
    }
    const getTags = async () => {
      const res: tagList = await axios.get(`${nowServerAddress.value}/tags`)
      for (let i = 0; i < res.data.length; i++) {
        allTags.value.push({
          value: res.data[i]
        })
      }
    }
    onMounted(async () => {
      await getTags()
    })
    const handleSelect = (item: any) => {
      inputVisible.value = false
      if (selectedTags.value.indexOf(item.value) !== -1) {
        ElMessage.error('标签已存在!')
      } else if (item.value !== '') {
        selectedTags.value.push(item.value)
      }
      inputValue.value = ''
    }
    const handleClose = (tag: string) => {
      selectedTags.value.splice(selectedTags.value.indexOf(tag), 1)
    }
    watch(selectedTags.value, async () => {
      if (selectedTags.value.length === 0) {
        const res: videoList = await axios.get(`${nowServerAddress.value}/videos`)
        videoList.value = res.data
      } else {
        const res: videoList = await axios.get(`${nowServerAddress.value}/tags_videos/${selectedTags.value.join(',')}`)
        videoList.value = res.data
      }
    })
    return {
      currentDate,
      downloadVideo,
      videoList,
      bytesFormat,
      selectedTags,
      inputVisible,
      inputValue,
      showInput,
      querySearch,
      handleSelect,
      handleClose,
      nowServerAddress,
      playVideo,
      deleteVideo
    }
  }
})
</script>

<style scoped>
  .image {
    width: 100%;
    height: 18em;
  }
  .el-col {
    padding: 10px;
    min-width: 320px;
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
  .donwload-play {
    display: flex;
  }
  .donwload-play .el-button {
    flex: 1;
  }
  .tags-selected {
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
    vertical-align: middle;
  }
  .el-tag + .el-tag {
    margin-left: 10px;
    margin-bottom: 4px;
  }
  .button-new-tag {
    margin-left: 10px;
  }
  .input-new-tag {
    width: 90px;
    margin-left: 10px;
  }
</style>
