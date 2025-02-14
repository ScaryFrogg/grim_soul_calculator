<script setup lang="ts">
import { ref, watchEffect, inject } from "vue"
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';
import type { Api, Design, Requirement } from "@/types"
const api = inject<Api>("api")
const toast = useToast();
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

const copyToClipboard = async () => {
  try {
    const text = Object.entries(requirements.value)
      .map(([key, value]) => `${key}: ${value}`)
      .join('\n');

    await navigator.clipboard.writeText(text)
    toast.add({ severity: 'info', summary: 'Copied', life: 3000 })
  } catch (err) {
    console.error('Failed to copy text: ', err)
  }
}
</script>
<template>
  <Toast />
  <div class="flex flex-column xl:flex-row justify-content-evenly align-items-center w-full" id="build-list">
    <div class="w-9 md:w-5">
      <p>Select Designs to Build</p>
      <Listbox v-model="selectedDesigns" :options="designs" filter optionLabel="name"
        listStyle="max-height:500px; text-align:center" striped multiple />
    </div>
    <div class="w-9 md:w-5">
      {{ }}
      <div class="flex justify-content-between">
        <span class="p-2">Total Materials</span>
        <Button v-if="Object.keys(requirements).length > 0" icon="pi pi-copy" severity="secondary" aria-label="Save"
          @click=copyToClipboard />
      </div>
      <Listbox emptyMessage="Select Designs To Get List of Materials" :options="Object.keys(requirements)"
        listStyle="max-height:500px" striped>
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
