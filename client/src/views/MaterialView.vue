<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
const route = useRoute()
const material = ref([])
const trades = ref([])
onMounted(() => {
  const id = route.params.id
  console.log(id)
  fetch(`http://localhost:3000/requirement/${id}`)
    .then(d => d.json().then(data => {
      console.log(data)
      material.value = data
    }))
    .catch(e => console.error(e))

  fetch(`http://localhost:3000/trades/${id}`)
    .then(d => d.json().then(data => {
      console.log(data)
      trades.value = data
    }))
    .catch(e => console.error(e))

})
</script>
<template>
  <div>
    <h2>{{ material.name }}</h2>
    <Card>
      <h3>Used For:</h3>
      <div v-for="(f, i) in material" :key=i>
        <router-link :to="`/design/${f.id}`">{{ f.name }}</router-link>
        : {{ f.quantity }}
      </div>
    </Card>
    <Card>
      <h3>Trades</h3>
      <div v-for="(f, i) in trades" :key=i>
        {{ f.giveName }} X {{ f.giveQuantity }} -> {{ f.getName }} X {{ f.getQuantity }}
      </div>
    </Card>
  </div>
</template>
