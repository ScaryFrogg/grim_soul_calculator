<script setup>
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
const route = useRoute()
const material = ref([])
onMounted(() => {
  const id = route.params.id
  console.log(id)
  fetch(`http://localhost:3000/material/${id}`)
    .then(d => d.json().then(data => {
      material.value = data
    }))
    .catch(e => console.error(e))

})
</script>
<template>
  <div>
    <h2>{{ material.name }}</h2>
    <h3>Used For:</h3>
    <div v-for="(f, i) in material.designs" :key=i>
      <router-link :to="`/design/${f.designId}`">{{ f.design }}</router-link>
      : {{ f.quantity }}
    </div>
  </div>
</template>
