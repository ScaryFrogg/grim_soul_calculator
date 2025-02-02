<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRouter } from "vue-router"
import type { Trade } from "@/types"
const router = useRouter()
const trades = ref<Trade[]>([])

const apiRoot = "http://localhost:3000/"
onMounted(() => {
  fetch(`${apiRoot}trades/`)
    .then(d => d.json().then(data => {
      trades.value = data
    }))
    .catch(e => console.error(e))
})
const navigate = (whereTo: number | undefined) => {
  if (whereTo) {
    router.push(`material/${whereTo}`)
  }
}
</script>
<template>
  <h2>Trades</h2>
  <Card>
    <template #content>
      <div v-for="(f, i) in trades" :key=i>
        <span @click="navigate(f.giveId)">{{ f.giveName }}</span>
        X {{ f.giveQuantity }} ->
        <span @click="navigate(f.getId)">{{ f.getName }}</span>
        X {{ f.getQuantity }}
      </div>
    </template>
  </Card>
</template>
