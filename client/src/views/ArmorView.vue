<script setup lang="ts">
import { ref, onMounted, inject, watch, nextTick } from "vue"
import { useRoute } from "vue-router"
import type { Api, ArmorInfo } from "@/types"
const api = inject<Api>("api")
const route = useRoute()
const sets = ref<ArmorInfo[]>()
const selectedSet = ref<ArmorInfo>()
const setGridData = ref<ArmorInfo[]>()
onMounted(() => {
  const id = route.params.id
  console.log(id)
  api?.get(`armor`)
    .then(data => {
      sets.value = data
    })
    .catch(e => console.error(e))
})
const selectedSetContainer = ref<HTMLDivElement | null>(null);
watch(selectedSet, async (newS) => {
  if (newS && selectedSet.value) {
    await api?.get(`armor/set/${newS.id}`)
      .then(data => {
        setGridData.value = data
        setGridData.value?.unshift(selectedSet.value!)
      })
    await nextTick()
    const div = selectedSetContainer.value!;
    div.scrollIntoView({ behavior: "smooth", block: "start" });
  }
})
</script>
<template>
  <div class="flex flex-column lg:flex-row">
    <Card class="w-full lg:w-3">
      <template #title>
        <div class="text-center">
          Sets:
        </div>
      </template>
      <template #content>
        <Listbox v-model="selectedSet" :options="sets" filter optionLabel="name"
          listStyle="max-height:400px; text-align:center" striped />
      </template>
    </Card>

    <div ref="selectedSetContainer" class="p-3 w-full lg:w-9" v-if="selectedSet">
      <h2>{{ selectedSet.name }}</h2>
      <DataTable size="small" :value="setGridData" removableSort stripedRows
        :pt="{ root: { 'style': 'overflow:scroll' } }">
        <Column field="name" header="Name"></Column>
        <Column field="armor" header="Armor"></Column>
        <Column field="protection" header="Protection"></Column>
        <Column field="durability" header="Durability"></Column>
        <Column field="effect" header="Bonus"></Column>
        <Column field="crafting" header="Craft Requirements"></Column>
      </DataTable>
    </div>
  </div>
</template>
