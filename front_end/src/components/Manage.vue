<template>
  <div class="container">
    <h3>当前服务器地址</h3>
    <el-input
      class="input-new-tag"
      v-if="inputVisible"
      v-model="inputValue"
      ref="saveTagInput"
      size="small"
      @blur="handleInputConfirm"
      @keyup.enter="handleInputConfirm"
    >
    </el-input>
    <el-button v-else class="button-new-tag" size="small" @click="showInput">{{ nowServerAddress }}</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from 'vue'
import { useStore } from 'vuex'
import * as utils from '../hooks/utils'

export default defineComponent({
  name: 'Manage',
  setup () {
    const store = useStore()
    const nowServerAddress = computed(() => {
      return utils.serverAddress()
    })
    const inputValue = ref(utils.serverAddress())
    const inputVisible = ref(false)
    const handleInputConfirm = () => {
      inputVisible.value = false
      store.commit('setServerAddress', inputValue.value)
      inputValue.value = store.state.serverAddress
    }
    const showInput = () => {
      inputVisible.value = true
    }
    return {
      inputVisible,
      inputValue,
      handleInputConfirm,
      showInput,
      nowServerAddress
    }
  }
})
</script>

<style scoped>
  .container {
    margin-top: 15px;
    margin-left: 10px;
  }
  .el-tag + .el-tag {
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
    width: 240px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>
