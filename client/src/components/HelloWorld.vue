<script setup>
import { ref } from "vue"
const designs = ref([])
fetch(`http://localhost:3000/designs`)
  .then(d => d.json().then(data => {
    designs.value = data
  }))
  .catch(e => console.error(e))

const getMaterial = id => {
  fetch(`http://localhost:3000/material/${id}`)
    .then(d => console.log(d.json()))
    .catch(e => console.error(e))
}
const health = () => {
  fetch(`http://localhost:3000/health`)
    .then(d => console.log(d.json()))
    .catch(e => console.error(e))
}
</script>
<template>
  <button @click="getMaterial(2)">MATERIAL</button>
  <button @click="health()">Health</button>
  <div>
    <div v-for="(f, i) in designs" :key=i>
      <a>{{ f.name }}</a>
    </div>

    {{ designs }}
  </div>
</template>
