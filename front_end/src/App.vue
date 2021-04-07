<template>
  <div>
    <el-menu
      :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
    >
      <el-menu-item :router="false"><img src="./assets/logo.png" width="30" height="30">EasyVideo</el-menu-item>
      <el-menu-item index="1"><i class="el-icon-s-home"></i>首页</el-menu-item>
      <el-menu-item index="2"><i class="el-icon-collection-tag"></i>标签</el-menu-item>
      <el-menu-item index="3"><i class="el-icon-upload"></i>上传</el-menu-item>
    </el-menu>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'App',
  setup () {
    const router = useRouter()
    const activeIndex = ref('1')
    const handleSelect = (key: string) => {
      if (key === '1') {
        router.push('/')
      } else if (key === '2') {
        router.push('/tags')
      } else if (key === '3') {
        router.push('/upload')
      }
    }
    return {
      activeIndex,
      handleSelect
    }
  }
})
</script>

<style>
  body {
    margin: 0;
  }
</style>
