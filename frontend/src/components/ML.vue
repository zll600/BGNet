<template>
  <div>
    <label for="Img">Img:</label>
    <input type="file" id="Img" @change="onFileChange(0, $event)" />
    <img
      :src="imageContents[0]"
      v-if="imageContents && imageContents[0]"
      style="max-width: 200px; max-height: 200px"
    />

    <label for="GT">GT:</label>
    <input type="file" id="GT" @change="onFileChange(1, $event)" />
    <img
      :src="imageContents[1]"
      v-if="imageContents && imageContents[1]"
      style="max-width: 200px; max-height: 200px"
    />

    <button @click="confirmUpload">Confirm</button>
    <img
      :src="newImageContent"
      v-if="newImageContent"
      style="max-width: 400px; max-height: 400px"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const imageURLs = ref<[File | null, File | null]>([null, null])
const imageContents = ref<[string | null, string | null]>([null, null])
const newImageContent = ref<string | null>(null)

function onFileChange(index: number, e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files[0]) {
    imageURLs.value[index] = input.files[0]
    const reader = new FileReader()
    reader.onload = () => {
      if (reader.result) {
        imageContents.value[index] = reader.result.toString()
      }
    }
    reader.readAsDataURL(input.files[0])
  }
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
