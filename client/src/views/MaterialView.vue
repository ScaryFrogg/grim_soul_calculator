<script setup lang="ts">
import { ref, onMounted, inject } from "vue"
import { useRoute, useRouter } from "vue-router"
import type { Api, MaterialInfo, Trade } from "@/types"

const route = useRoute()
const router = useRouter()
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

const navigate = (whereTo: number | undefined) => {
  if (whereTo) {
    router.push(`/design/${whereTo}`)
  }
}
</script>
<template>
  <div v-if="material != null">
    <h2>{{ material.name }}</h2>
    <Card>
      <template #title>
        Used For:
      </template>
      <template #content>
        <p class="text-primary" v-for="(f, i) in material.designs" :key=i @click="navigate(f.designId)">
          {{ f.design }}
          : {{ f.quantity }}
        </p>
      </template>
    </Card>
    <Card v-if="trades.length > 0">
      <template #title>
        Trades
      </template>
      <template #content>
        <p class="text-primary" v-for="(f, i) in trades" :key=i>
          {{ f.giveName }} X {{ f.giveQuantity }} -> {{ f.getName }} X {{ f.getQuantity }}
        </p>
      </template>
    </Card>
  </div>
</template>
