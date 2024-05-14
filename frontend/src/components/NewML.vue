<template>
  <div id="page" class="flex flex-col flex-wrap w-screen h-screen content-stretch">
    <div id="topbar" class="flex flex-row justify-between h-1/6 p-4 bg-[#014a95]">
      <div id="icon" class="">
        <NImage src="/xy_logo.png"></NImage>
      </div>
      <div id="info" class="flex flex-col flex-wrap content-between">
        <label class="text-[#fff] font-mono text-lg">计算机学院</label>
        <br />
        <label class="text-[#fff] font-mono text-lg">张磊刚</label>
      </div>
    </div>
    <div id="content" class="flex flex-row flex-wrap justify-between content-stretch h-5/6 p-2">
      <div id="left" class="flex flex-col justify-evenly items-center content-between w-2/5">
        <label class="font-mono text-lg">上传两张照片，原图是 jpg 格式，GT 是 png 格式</label>
        <div
          id="upload-image"
          class="w-full h-1/2 flex flex-col justify-center items-center content-center"
        >
          <NUpload
            accept="image/*"
            list-type="image-card"
            :max="1"
            action="http://127.0.0.1:8080/uploadv2"
            v-model:file-list="fileListImg"
            @change="handleUploadImgChange"
            @remove="handleRemove"
            @update:file-list="handleFileListChange"
          />
          <NDivider></NDivider>
          <NUpload
            accept="image/*"
            list-type="image-card"
            :max="1"
            action="http://127.0.0.1:8080/uploadv2"
            v-model:file-list="fileListGT"
            @change="handleUploadGTChange"
            @remove="handleRemove"
            @update:file-list="handleFileListChange"
          />
        </div>
      </div>
      <div id="button" class="w-1/5 flex justify-center items-center content-center">
        <NButton type="primary" round size="large" :loading="loading" @click="confirmUpload"
          >生成结果</NButton
        >
      </div>
      <div id="preview" class="w-2/5 flex flex-col justify-center items-center content-between">
        <label class="font-mono text-lg">预览结果：</label>
        <NImage
          v-if="newImageContent"
          :src="newImageContent"
          :height="416"
          :width="416"
          :alt="`Generated Image`"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { NUpload, NButton, NImage, NDivider, useMessage, type UploadFileInfo } from 'naive-ui'

const imageURLs = ref<[File | null, File | null]>([null, null])
const newImageContent = ref<string | null>(null)
const fileListImg = ref<UploadFileInfo[]>([])
const fileListGT = ref<UploadFileInfo[]>([])
const loading = ref<boolean>(false)
const message = useMessage()

function handleUploadImgChange(options: {
  file: UploadFileInfo
  fileList: Array<UploadFileInfo>
  event?: Event
}): void {
  console.log(options)
  fileListImg.value = options.fileList
  console.log(fileListImg.value)
}

const handleFileListChange = () => {
  message.info('是的，file-list 的值变了')
}

const handleRemove = (data: { file: UploadFileInfo; fileList: UploadFileInfo[] }) => {
  message.info('删除照片 ' + data.file.id)
}

function handleUploadGTChange(options: {
  file: UploadFileInfo
  fileList: Array<UploadFileInfo>
  event?: Event
}): void {
  console.log(options)
  fileListGT.value = options.fileList
  console.log(fileListGT.value)
}

async function confirmUpload() {
  loading.value = true
  setTimeout(() => {
    loading.value = false
  }, 60 * 1000)
  try {
    const response = await axios.post('http://127.0.0.1:8080/generate', {
        userName: 'andy',
    }, {
        headers: {
            'Content-Type': 'application/json'
        },
      responseType: 'arraybuffer'
    })

    const arrayBuffer: Uint8Array = new Uint8Array(response.data)
    const blob: Blob = new Blob([arrayBuffer], { type: 'image/png' })

    newImageContent.value = URL.createObjectURL(blob)
  } catch (error) {
    console.error('fail to upload file', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  imageURLs.value = [null, null]
})
</script>

<style scoped>
.n-upload-file-list .n-upload-file.n-upload-file--image-card-type {
  position: relative;
  width: 416px;
  height: 416px;
  border: var(--n-item-border-image-card);
  border-radius: var(--n-border-radius);
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition:
    border-color 0.3s var(--n-bezier),
    background-color 0.3s var(--n-bezier);
  border-radius: var(--n-border-radius);
  overflow: hidden;
}
</style>
