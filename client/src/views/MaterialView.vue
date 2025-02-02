<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
import type { MaterialInfo, Trade } from "@/types"

const route = useRoute()
const material = ref<MaterialInfo | null>(null)
const trades = ref<Trade[]>([])

onMounted(() => {
  const id = route.params.id
  fetch(`http://localhost:3000/item/${id}`)
    .then(d => d.json().then(data => {
      material.value = data
    }))
    .catch(e => console.error(e))

  fetch(`http://localhost:3000/trades/${id}`)
    .then(d => d.json().then(data => {
      trades.value = data
    }))
    .catch(e => console.error(e))
})

</script>
<template>
  <div v-if="material != null">
    <h2>{{ material.name }}</h2>
    <Card>
      <template #content>
        <h3>Used For:</h3>
        <p class="m-0">
        <div v-for="(f, i) in material.designs" :key=i>
          <router-link :to="`/design/${f.designId}`">{{ f.design }}</router-link>
          : {{ f.quantity }}
        </div>
        </p>
      </template>
    </Card>
    <Card>
      <template #content>
        <h3>Trades</h3>
        <div v-for="(f, i) in trades" :key=i>
          {{ f.giveName }} X {{ f.giveQuantity }} -> {{ f.getName }} X {{ f.getQuantity }}
        </div>
      </template>
    </Card>
  </div>
</template>
