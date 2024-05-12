<template>
  <div id="page" class="flex flex-col w-screen h-screen">
    <div id="topbar" class="flex flex-row justify-between h-2/7 p-4 bg-[#014a95]">
      <div id="icon" class="">
        <NImage src="/public/xy_logo.png"></NImage>
      </div>
      <div id="info" class="flex flex-col flex-wrap content-between">
        <label class="text-[#fff] font-mono text-lg">计算机学院</label>
        <br />
        <label class="text-[#fff] font-mono text-lg">张磊刚</label>
      </div>
    </div>
    <div
      id="content"
      class="flex flex-row flex-wrap justify-between content-stretch h-5/7 p-2 bg-gray-100"
    >
      <div id="left" class="flex flex-col content-between w-2/5">
        <div id="upload-image" class="w-full h-1/2">
          <NUpload
            class="w-full bg-sky-100"
            accept="image/*"
            list-type="image-card"
            action="http://127.0.0.1:8080/uploadv2"
            :default-file-list="fileList"
            @change="changeHandler"
          />
          <NImage
            v-if="previewImageUrlRef1.length > 0"
            :src="previewImageUrlRef1"
            :alt="`Uploaded Image 1`"
          />
        </div>
        <NDivider></NDivider>
      </div>
      <div id="button" class="w-1/5 flex justify-center content-center">
        <NButton type="primary" :loading="loading" @click="confirmUpload">生成结果</NButton>
      </div>
      <div id="preview" class="w-2/5">
        <NImage
          v-if="newImageContent"
          :src="newImageContent"
          :width="100"
          :alt="`Generated Image`"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { NUpload, NButton, NImage, NDivider, type UploadFileInfo } from 'naive-ui'

const imageURLs = ref<[File | null, File | null]>([null, null])
const newImageContent = ref<string | null>(null)
const fileList = ref<UploadFileInfo[]>([])
const loading = ref<boolean>(false)

function changeHandler(options: {
  file: UploadFileInfo
  fileList: Array<UploadFileInfo>
  event?: Event
}): void {
  console.log(options)
  fileList.value.push(options.file)
  console.log(fileList.value)
  console.log(options.fileList)
}

function finishHandler(file: UploadFileInfo): UploadFileInfo {
  file.status = 'finished'
  fileList.value.push(file)
  return file
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
