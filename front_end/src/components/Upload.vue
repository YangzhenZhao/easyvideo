<template>
  <div class="container">
    <el-input v-model="title" placeholder="请输入视频标题"></el-input>
    <el-upload
      class="upload-cover-pic"
      :action="`http://localhost:5000/upload/${title}`"
      :limit=1
      :before-upload="checkPicUpload"
      list-type="picture">
      <el-button size="small" type="primary">上传封面图片</el-button>
    </el-upload>
    <el-upload
      class="upload-video"
      :action="`http://localhost:5000/upload/${title}`"
      :limit=1
      :before-upload="checkVideoUpload"
      list-type="picture">
      <el-button size="small" type="primary">上传视频</el-button>
    </el-upload>
    <h3>标签</h3>
    <el-button type="info" icon="el-icon-plus" circle></el-button><br>
    <el-button class="save-video" @click="saveVideo" type="primary">立即保存</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

export default defineComponent({
  name: 'Upload',
  setup () {
    const title = ref('')
    const coverPictureName = ref('')
    const videoName = ref('')
    const videoSize = ref(-1)
    const checkPicUpload = (file: File) => {
      if (title.value === '') {
        ElMessage.error('视频标题不可为空!')
        return false
      }
      coverPictureName.value = file.name
    }
    const checkVideoUpload = (file: File) => {
      if (title.value === '') {
        ElMessage.error('视频标题不可为空!')
        return false
      }
      videoName.value = file.name
      videoSize.value = file.size
    }
    const saveVideo = () => {
      console.log(coverPictureName.value, videoName.value)
      axios({
        method: 'post',
        url: 'http://127.0.0.1:5000/save_video',
        params: {
          name: title.value,
          bytes_size: videoSize.value,
          video_file_name: videoName.value,
          cover_picture_file_name: coverPictureName.value
        }
      }).then(function (response) {
        console.log(response)
      })
    }
    return {
      title,
      checkPicUpload,
      checkVideoUpload,
      saveVideo
    }
  }
})
</script>

<style scoped>
  .container {
    margin-left: 50px;
  }
  .el-input {
    margin-top: 20px;
    margin-bottom: 20px;
    width: 360px;
  }
  .upload-cover-pic {
    padding-right: 50px;
  }
  .upload-video {
    margin-top: 20px;
    padding-right: 50px;
  }
  .save-video {
    margin-top: 20px;
  }
</style>
