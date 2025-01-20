<script setup>
import { ref } from "vue"
const designs = ref([])
fetch(`http://localhost:3000/designs`)
  .then(d => d.json().then(data => {
    designs.value = data
  }))
  .catch(e => console.error(e))

</script>
<template>
  <div>
    <div v-for="(f, i) in designs" :key=i>
      <router-link :to="`design/${f.id}`">{{ f.name }}</router-link>
    </div>

    {{ designs }}
  </div>
</template>
