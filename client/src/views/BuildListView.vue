<script setup lang="ts">
import Listbox from "primevue/listbox"
import { ref, watchEffect, inject } from "vue"
import type { Api, Design, Requirement } from "@/types"
const api = inject<Api>("api")
const fetchedRequirements = new Map<number, Requirement[]>()
const selectedDesigns = ref<Design[]>([]);
const designs = ref<Design[]>([]);
const requirements = ref<Record<string, number>>({});

api?.get("designs")
  .then(data => {
    designs.value = data
  }).catch(e => console.error(e))

watchEffect(async () => {
  requirements.value = {}
  for (let design of selectedDesigns.value) {
    const reqs = await getReq(design.id)
    if (!reqs) return
    reqs.forEach((req: Requirement) => {
      if (req.name in requirements.value) {
        requirements.value[req.name] = requirements.value[req.name] + req.quantity
      }
      else {
        requirements.value[req.name] = req.quantity
      }
    })
  }
})
const getReq = async (planId: number) => {
  if (!fetchedRequirements.has(planId)) {
    const reqs = await api?.get(`requirement/${planId}`)
    fetchedRequirements.set(planId, reqs)
  }
  return fetchedRequirements.get(planId)
}

</script>
<template>
  <div class="flex flex-column xl:flex-row justify-content-evenly align-items-center w-full" id="build-list">
    <div class="w-9 md:w-5">
      <p>Select Designs to Build</p>
      <Listbox v-model="selectedDesigns" :options="designs" filter optionLabel="name"
        listStyle="height:250px; text-align:center" striped multiple />
    </div>
    <div class="w-9 md:w-5">
      <p>Total Materials</p>
      <Listbox emptyMessage="Select Designs To Get List of Materials" :options="Object.keys(requirements)"
        listStyle="height:250px" striped>
        <template #option="row">
          <div class="flex justify-content-between">
            <span>{{ row.option }}</span>
            <span>{{ requirements[row.option] }}</span>
          </div>
        </template>
      </Listbox>
    </div>
  </div>
</template>
