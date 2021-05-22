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
    >
    </el-input>
    <el-button v-else class="button-new-tag" size="small" @click="showInput">{{ nowServerAddress }}</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from 'vue'
import { useStore } from 'vuex'

export default defineComponent({
  name: 'Manage',
  setup () {
    const store = useStore()
    const nowServerAddress = computed(() => store.state.serverAddress)
    const inputVisible = ref(false)
    const inputValue = ref(store.state.serverAddress)
    const handleInputConfirm = () => {
      inputVisible.value = false
      store.commit('setServerAddress', inputValue.value)
      inputValue.value = store.state.serverAddress
    }
    const showInput = () => {
      inputVisible.value = true
    }
    return {
      nowServerAddress,
      inputVisible,
      inputValue,
      handleInputConfirm,
      showInput
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
