<template>
  <div class="top-container">
    <div class="upload-section">
      <div class="left-section">
        <label for="Img">上传图片:</label>
        <input type="file" id="Img" @change="onFileChange(0, $event)" />
        <img :src="imageContents[0]" v-if="imageContents && imageContents[0]" />
      </div>

      <div class="right-section">
        <label for="GT">上传 GT:</label>
        <input type="file" id="GT" @change="onFileChange(1, $event)" />
        <img :src="imageContents[1]" v-if="imageContents && imageContents[1]" />
      </div>
    </div>

    <button class="btn" @click="confirmUpload">生成结果</button>

    <div class="preview-section">
      <label for="preview">结果预览:</label>
      <img :src="newImageContent" v-if="newImageContent" />
    </div>
  </div>

  <div class="bg-[#020817] h-[72px] p-4 flex items-center justify-between sticky top-0 z-[999]">
    <a class="cursor-pointer flex items-center space-x-4" href="/">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="32"
        height="32"
        viewBox="0 0 32 32"
        fill="none"
      >
        <path
          d="M21.332 14.666L23.594 16.7017C25.6507 18.5528 26.6791 19.4784 26.6791 20.666C26.6791 21.8536 25.6507 22.7792 23.594 24.6303L21.332 26.666"
          stroke="white"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        ></path>
        <path
          d="M18.6478 6.66602L15.9899 16.666L13.332 26.666"
          stroke="white"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        ></path>
        <path
          d="M10.6674 6.43945L8.40542 8.47521C6.34867 10.3263 5.32031 11.2518 5.32031 12.4395C5.32031 13.6271 6.34867 14.5527 8.40542 16.4037L10.6674 18.4395"
          stroke="white"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        ></path>
      </svg>
      <h1 class="text-xl font-semibold text-white dark:text-zinc-50">JS2TS</h1>
    </a>
    <img src="./xiyou.ico" alt="menu" class="lg:hidden w-6 filter invert" />
    <div class="hidden lg:flex gap-8 pt-8">
      <div
        class="text-base xs:font-cairo hover:text-primary transition-colors duration-300 pb-8 cursor-pointer group flex items-center"
      >
        <a href="/">Home</a>
      </div>
      <div
        class="text-base xs:font-cairo hover:text-primary transition-colors duration-300 pb-8 cursor-pointer group flex items-center"
      ></div>
      <div
        class="text-base xs:font-cairo hover:text-primary transition-colors duration-300 pb-8 cursor-pointer group flex items-center"
      >
        <a class="px-3 py-2 border border-white rounded-md text-white" href="/contact-us">
          Contact Us
        </a>
      </div>
    </div>
  </div>
  <div class="upload-container">
    <div class="upload-section">
      <n-upload
        v-model="imageURLs[0]"
        accept="image/*"
        :show-preview="true"
        :drag-drop="true"
        @change="onFileChange(0)"
      />
      <n-image
        v-if="imageURLs[0]"
        :src="imageURLs[0]"
        :alt="`Uploaded Image 1`"
        class="upload-preview"
      />
    </div>
    <div class="upload-section">
      <n-upload
        v-model="imageURLs[1]"
        accept="image/*"
        :show-preview="true"
        :drag-drop="true"
        @change="onFileChange(1)"
      />
      <n-image
        v-if="imageURLs[1]"
        :src="imageURLs[1]"
        :alt="`Uploaded Image 2`"
        class="upload-preview"
      />
    </div>

    <n-button type="primary" @click="confirmUpload" class="upload-button">生成结果</n-button>

    <div v-if="newImageContent">
      <n-image :src="newImageContent" :alt="`Generated Image`" class="generated-image" />
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
  /* padding: 20px; */
  margin-top: 20px;
  /* background-color: #ffffff;  */
  /* White background for upload section */
}

.left-section {
  width: 50%;
  height: 100%;
  padding: 20px 10px 0px 20px;
  margin-left: 10px;
  border-radius: 10px; /* Rounded corners */
  outline-style: groove;
}

.right-section {
  width: 50%;
  height: 100%;
  padding: 20px 20px 0px 10px;
  margin-right: 10px;
  border-radius: 10px;
  outline-style: groove;
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
  /* margin-top: 20px; */
  cursor: pointer; /* Change cursor to pointer on labels */
  padding: 8px 12px; /* Add padding to labels */
  background-color: #4caf50; /* Green background color */
  color: white; /* White text color */
  border-radius: 4px; /* Rounded corners */
  display: block;
}

.preview-section {
  padding: 20px;
  margin: 10px 10px 10px 10px;
  width: 100%;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  /* background-color: #ffffff;  */
  /* White background for preview section */
  border-radius: 10px;
  outline-style: groove;
}

/* Add styles for input elements */
input[type='file'] {
  display: none; /* Hide the default file input */
}

label[for='Img'],
label[for='GT'] {
  cursor: pointer; /* Change cursor to pointer on labels */
  padding: 8px 12px; /* Add padding to labels */
  background-color: #4caf50; /* Green background color */
  color: white; /* White text color */
  border-radius: 4px; /* Rounded corners */
  display: block;
}

label[for='Img']:hover,
label[for='GT']:hover {
  background-color: #45a049; /* Darker green on hover */
}

input[type='file'] + label {
  margin-top: 10px; /* Add top margin between input and label */
  /* display: inline-block; */
}

input[type='file'] + label::before {
  content: 'Select File'; /* Custom text before the label */
}

input[type='file'] + label:hover {
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
