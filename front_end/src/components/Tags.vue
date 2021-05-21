<template>
  <div class="container">
    <h3>所有标签</h3>
    <el-row>
      <el-col :span="6" v-for="tagName in allTags" :key="tagName">
        <el-button type="primary" plain>{{ tagName }}</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import axios from 'axios'

interface tagList {
  data: string[]
}

export default defineComponent({
  name: 'Tags',
  setup () {
    const allTags = ref()
    const getTags = async () => {
      const res: tagList = await axios.get('http://127.0.0.1:5000/tags')
      allTags.value = res.data
      console.log(allTags.value)
    }
    onMounted(async () => {
      await getTags()
    })
    return {
      allTags
    }
  }
})
</script>

<style scoped>
  .container {
    margin-top: 18px;
    margin-left: 15px;
    margin-bottom: 15px;
  }
  .el-button {
    margin-bottom: 8px;
    width: 300px;
  }
</style>
