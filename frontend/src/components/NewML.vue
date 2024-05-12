<template>
  <div id="page" class="flex flex-col flex-wrap w-screen h-screen content-stretch">
    <div id="topbar" class="flex flex-row justify-between h-1/6 p-4 bg-[#014a95]">
      <div id="icon" class="">
        <NImage src="/public/xy_logo.png"></NImage>
      </div>
      <div id="info" class="flex flex-col flex-wrap content-between">
        <label class="text-[#fff] font-mono text-lg">计算机学院</label>
        <br />
        <label class="text-[#fff] font-mono text-lg">张磊刚</label>
      </div>
    </div>
    <div id="content" class="flex flex-row flex-wrap justify-between content-stretch h-5/6 p-2">
      <div id="left" class="flex flex-col justify-between items-center content-between w-2/5">
        <label class="font-mono text-lg">上传两张照片，原图是 jpg 格式，GT 是 png 格式</label>
        <div id="upload-image" class="w-full h-1/2 flex flex-col items-center content-center">
          <NUpload
            class="w-full bg-sky-100"
            accept="image/*"
            list-type="image-card"
            action="http://127.0.0.1:8080/uploadv2"
            v-model:file-list="fileList"
            @change="handleUploadChange"
            @remove="handleRemove"
            @update:file-list="handleFileListChange"
          />
          <NImage
            v-if="previewImageUrlRef1.length > 0"
            :src="previewImageUrlRef1"
            :alt="`Uploaded Image 1`"
          />
        </div>
      </div>
      <div id="button" class="w-1/5 flex justify-center items-center content-center">
        <NButton type="primary" :loading="loading" @click="confirmUpload">生成结果</NButton>
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
import { NUpload, NButton, NImage, useMessage, type UploadFileInfo } from 'naive-ui'

const imageURLs = ref<[File | null, File | null]>([null, null])
const newImageContent = ref<string | null>(null)
const fileList = ref<UploadFileInfo[]>([])
const loading = ref<boolean>(false)
const message = useMessage()

function handleUploadChange(options: {
  file: UploadFileInfo
  fileList: Array<UploadFileInfo>
  event?: Event
}): void {
  console.log(options)
  fileList.value = options.fileList
  console.log(fileList.value)
}

const handleFileListChange = () => {
  message.info('是的，file-list 的值变了')
}

const handleRemove = (data: { file: UploadFileInfo; fileList: UploadFileInfo[] }) => {
  message.info('删除照片 ' + data.file.id)
}

const previewImageUrlRef1 = ref('')

async function confirmUpload() {
  loading.value = true
  setTimeout(() => {
    loading.value = false
  }, 60 * 1000)
  try {
    const response = await axios.get('http://127.0.0.1:8080/generate', {
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
