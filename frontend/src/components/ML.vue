<template>
  <div>
    <label for="Img">Img:</label>
    <input type="file" id="Img" @change="onFileChange(0, $event)" />
    <img :src="imageURLs[0]" v-if="imageURLs && imageURLs[0]" style="max-width: 200px; max-height: 200px" />

    <label for="GT">GT:</label>
    <input type="file" id="GT" @change="onFileChange(1, $event)" />
    <img :src="imageURLs[1]" v-if="imageURLs && imageURLs[1]" style="max-width: 200px; max-height: 200px" />

    <button @click="confirmUpload">Confirm</button>
    <img :src="newImageURL" v-if="newImageURL" style="max-width: 400px; max-height: 400px" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const imageURLs = ref<[string | null, string | null]>([null, null])
const newImageURL = ref<string | null>(null)

function onFileChange(index: number, e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files[0]) {
    const reader = new FileReader()
    reader.onload = () => {
      if (reader.result) {
        imageURLs.value[index] = reader.result.toString()
      }
    }
    reader.readAsDataURL(input.files[0])
  }
}

async function confirmUpload() {
  const formData = new FormData()
  formData.append('Img', imageURLs.value[0] || '')
  formData.append('GT', imageURLs.value[1] || '')

  try {
    const response = await axios.post<{ imageURL: string }>('http://127.0.0.1:8080/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    newImageURL.value = response.data.imageURL
  } catch (error) {
    console.error('fail to upload file', error)
  }
}

onMounted(() => {
  imageURLs.value = [null, null];
})

</script>
