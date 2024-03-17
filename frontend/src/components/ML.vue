<template>
  <div class="top-container">
    <div class="upload-section">
      <div class="left-section">
        <label for="Img">Upload Img:</label>
        <input type="file" id="Img" @change="onFileChange(0, $event)" />
        <img :src="imageContents[0]" v-if="imageContents && imageContents[0]" />
      </div>

      <div class="right-section">
        <label for="GT">Upload GT:</label>
        <input type="file" id="GT" @change="onFileChange(1, $event)" />
        <img :src="imageContents[1]" v-if="imageContents && imageContents[1]" />
      </div>
    </div>

    <button class="btn" @click="confirmUpload">Confirm</button>


    <div class="preview-section">
      <label>Preview Image:</label>
      <img :src="newImageContent" v-if="newImageContent" />
    </div>
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

<style scoped>

.top-container {
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  padding: 0;
  margin: 0;
  height: 100vh;
  width: 100vw;
  background-color: #f4f4f4; /* Light gray background */
}

.upload-section {
  width: 100%;
  height: 50%;
  display: flex;
  justify-content: space-around;
  flex-direction: row;
  padding: 20px;
  margin: 20px;
  background-color: #ffffff; /* White background for upload section */
}

.left-section {
  width: 50%;
  height: 100%;
  padding: 20px;
}

.right-section {
  width: 50%;
  height: 100%;
  padding: 20px;
}

.left-section,
.right-section {
  display: flex;
  flex-direction: column;
}

.btn {
  width: 200px;
  height: 50px;
  align-self: center;
  margin-top: 20px;
}

.preview-section {
  padding: 20px;
  width: 100%;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #ffffff; /* White background for preview section */
}

/* Add styles for input elements */
input[type="file"] {
  display: none; /* Hide the default file input */
}

label[for="Img"],
label[for="GT"] {
  cursor: pointer; /* Change cursor to pointer on labels */
  padding: 8px 12px; /* Add padding to labels */
  background-color: #4CAF50; /* Green background color */
  color: white; /* White text color */
  border-radius: 4px; /* Rounded corners */
  display: block;
}

label[for="Img"]:hover,
label[for="GT"]:hover {
  background-color: #45a049; /* Darker green on hover */
}

input[type="file"] + label {
  margin-top: 10px; /* Add top margin between input and label */
  /* display: inline-block; */
}

input[type="file"] + label::before {
  content: 'Select File'; /* Custom text before the label */
}

input[type="file"] + label:hover {
  background-color: #45a049; /* Darker green on hover */
}


.preview-section label {
  align-self: center;
  color: #333333;
}

.upload-section img,
.preview-section img {
  max-width: 100%;
  max-height: 400px;
}

</style>