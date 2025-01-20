<script setup>
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
const route = useRoute()
const design = ref([])
onMounted(() => {
  const id = route.params.id
  fetch(`http://localhost:3000/design/${id}`)
    .then(d => d.json().then(data => {
      design.value = data
    }))
    .catch(e => console.error(e))

})
</script>
<template>
  <div>
    <h2>{{ design.name }}</h2>
    <h3>Requirements: </h3>
    <div v-for="(f, i) in design.requirements" :key=i>
      <router-link :to="`/material/${f.id}`">{{ f.name }}</router-link>
      : {{ f.quantity }}
    </div>

  </div>
</template>
