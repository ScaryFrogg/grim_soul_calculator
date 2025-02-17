<script setup lang="ts">
import { ref, inject } from "vue"
import type { Ref } from "vue"
import type { Api, Design } from "@/types"
import { useRouter } from "vue-router"
const api = inject<Api>("api")
const router = useRouter()

const designs: Ref<Design[]> = ref([])
api?.get("designs")
  .then(data => {
    designs.value = data
  })
  .catch(e => console.error(e))

const navigate = (whereTo: Design) => {
  router.push(`design/${whereTo.id}`)
}
</script>
<template>

  <div class="flex flex-column align-items-center">
    <h2 class="text-center">Designs</h2>
    <Listbox :options="designs" striped optionLabel="name" @update:modelValue="navigate" filter :filterFields="['name']"
      class="w-9 md:w-5" list-style="max-height:80vh" />
  </div>
</template>
