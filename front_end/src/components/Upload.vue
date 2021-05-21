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
    <div>
      <h3>标签</h3>
      <el-tag
        :key="tag"
        v-for="tag in nowUseTags"
        closable
        :disable-transitions="false"
        @close="handleClose(tag)"
      >
        {{tag}}
      </el-tag>
      <el-input
        v-if="inputVisible"
        class="input-new-tag"
        v-model="inputValue"
        ref="saveTagInput"
        size="small"
        @blur="handleInputConfirm"
      >
      </el-input>
      <el-button v-else class="button-new-tag" size="small" @click="showInput">添加标签</el-button>
    </div>
    <div>
      <h3>所有标签</h3>
      <el-check-tag v-for="tagName in allTags" @change="onChangeChecked(tagName)" :checked="isChecked(tagName)" :key="tagName">{{ tagName }}</el-check-tag>
    </div>
    <el-button class="save-video" @click="saveVideo" type="primary">立即保存</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, Ref } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

interface tagList {
  data: string[]
}

export default defineComponent({
  name: 'Upload',
  setup () {
    const title = ref('')
    const coverPictureName = ref('')
    const videoName = ref('')
    const videoSize = ref(-1)
    const allTags = ref()
    const nowUseTags: Ref<string[]> = ref([])
    const inputVisible = ref(false)
    const inputValue = ref('')
    const getTags = async () => {
      const res: tagList = await axios.get('http://127.0.0.1:5000/tags')
      allTags.value = res.data
      console.log(allTags.value)
    }
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
          cover_picture_file_name: coverPictureName.value,
          tags: JSON.stringify(nowUseTags.value)
        }
      }).then(function (response) {
        console.log(response)
      })
    }
    onMounted(async () => {
      await getTags()
    })
    const showInput = () => {
      inputVisible.value = true
    }
    const handleInputConfirm = () => {
      inputVisible.value = false
      console.log(inputValue.value)
      nowUseTags.value.push(inputValue.value)
      inputValue.value = ''
    }
    const handleClose = (tag: string) => {
      nowUseTags.value.splice(nowUseTags.value.indexOf(tag), 1)
    }
    const isChecked = (tag: string) => {
      if (nowUseTags.value.indexOf(tag) >= 0) {
        return false
      }
      return true
    }
    const onChangeChecked = (tag: string) => {
      if (nowUseTags.value.indexOf(tag) >= 0) {
        nowUseTags.value.splice(nowUseTags.value.indexOf(tag), 1)
      } else {
        nowUseTags.value.push(tag)
      }
    }
    return {
      title,
      checkPicUpload,
      checkVideoUpload,
      saveVideo,
      allTags,
      nowUseTags,
      inputVisible,
      showInput,
      handleInputConfirm,
      inputValue,
      handleClose,
      isChecked,
      onChangeChecked
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
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .el-check-tag + .el-check-tag {
    margin-left: 10px;
  }
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 90px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>
