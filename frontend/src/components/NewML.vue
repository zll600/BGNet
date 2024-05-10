<template>
  <div class="flex flex-col max-w-100 item-center">
    <div class="flex flex-row item-center justify-between">
        <div>
            icon
        </div>
        <div>
            leigang
        </div>
    </div>
    <div class="flex flex-row">
      <div class="flex flex-col max-w-768">
        left
        <div>
          upload image
          <NUpload
            v-model="imageURLs[0]"
            accept="image/*"
            :show-preview="true"
            :drag-drop="true"
            @change="onFileChange(0)"
          />
          <NImage v-if="imageURLs[0]" :src="imageURLs[0]" :alt="`Uploaded Image 1`" />
        </div>
        upload gt
        <div>
          <NUpload
            v-model="imageURLs[1]"
            accept="image/*"
            :show-preview="true"
            :drag-drop="true"
            @change="onFileChange(1)"
          />
          <NImage v-if="imageURLs[1]" :src="imageURLs[1]" :alt="`Uploaded Image 2`" />
        </div>
      </div>
      <div>
        // middle
        <NButton type="primary" @click="confirmUpload">生成结果</NButton>
      </div>
      <div>
        result
        <NImage :src="newImageContent" :alt="`Generated Image`" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { NUpload, NButton, NImage } from 'naive-ui'

const imageURLs = ref<[File | null, File | null]>([null, null])
const newImageContent = ref<string | null>(null)

function onFileChange(index: number) {
  // Handle file change logic here (optional)
}

async function confirmUpload() {
  const formData = new FormData()
  formData.append('upload[]', imageURLs.value[0] || '')
  formData.append('upload[]', imageURLs.value[1] || '')

  try {
    const response = await axios.post('http://127.0.0.1:8080/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      },
      responseType: 'arraybuffer'
    })

    const arrayBuffer: Uint8Array = new Uint8Array(response.data)
    const blob: Blob = new Blob([arrayBuffer], { type: 'image/png' })

    newImageContent.value = URL.createObjectURL(blob)
  } catch (error) {
    console.error('fail to upload file', error)
  }
}

onMounted(() => {
  imageURLs.value = [null, null]
})
</script>
