<script setup lang="ts">
import { ref, onMounted, inject } from "vue"
import { useRoute } from "vue-router"
import type { Api, MaterialInfo, Trade } from "@/types"

const route = useRoute()
const material = ref<MaterialInfo | null>(null)
const trades = ref<Trade[]>([])
const api = inject<Api>("api")

onMounted(() => {
  const id = route.params.id
  api?.get(`item/${id}`)
    .then(data => {
      material.value = data
    })
    .catch(e => console.error(e))

  api?.get(`trades/${id}`)
    .then(data => {
      trades.value = data
    })
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
