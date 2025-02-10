<script setup lang="ts">
import { ref, onMounted, inject } from "vue"
import { useRoute, useRouter } from "vue-router"
import type { Api, Design } from "@/types"
const api = inject<Api>("api")
const route = useRoute()
const router = useRouter()
const design = ref<Design | null>(null)
onMounted(() => {
  const id = route.params.id
  api?.get(`design/${id}`)
    .then(data => {
      design.value = data
    })
    .catch(e => console.error(e))

})
const navigate = (whereTo: number | undefined) => {
  if (whereTo) {
    router.push(`/material/${whereTo}`)
  }
}
</script>
<template>
  <div class="flex flex-column align-items-center" v-if="design">
    <h2 class="my-5">{{ design.name }}</h2>

    <Card class="md:w-5">
      <template #title>
        <div class="text-center">
          Requirements:
        </div>
      </template>
      <template #content>
        <div v-for="(f, i) in design.requirements" :key=i>
          <span class="text-primary" @click="navigate(f.id)">
            {{ f.name }}
          </span>
          <span>
            : {{ f.quantity }}
          </span>
        </div>
      </template>
    </Card>
  </div>
</template>
