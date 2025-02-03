<script setup lang="ts">
import Listbox from "primevue/listbox"
import { ref, watchEffect } from "vue"
import type { Design, Requirement } from "@/types"
const fetchedRequirements = new Map<number, Requirement[]>()

const test = ref<Design[]>([]);
const designs = ref<Design[]>([]);
const requirements = ref<Record<string, number>>({});
fetch(`http://localhost:3000/designs`)
  .then(d => d.json().then(data => {
    designs.value = data
  }))
  .catch(e => console.error(e))

watchEffect(async () => {
  for (let design of test.value) {
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
    const reqs = await fetchRequirements(planId)
    fetchedRequirements.set(planId, reqs)
  }
  return fetchedRequirements.get(planId)
}
const fetchRequirements = async (planId: number): Promise<Requirement[]> => {
  return fetch(`http://localhost:3000/requirement/${planId}`)
    .then(d => d.json().then((data: Requirement[]) => {
      return data
    }))
    .catch(e => {
      console.error(e)
      return []
    })
}

</script>
<template>
  <div class="flex flex-column align-items-center w-full" id="build-list">
    <p>Select Designs to Build</p>
    <Listbox v-model="test" :options="designs" filter optionLabel="name" class="w-9 md:w-5"
      listStyle="height:250px; text-align:center" striped multiple />
    <p>Total Materials</p>
    <Listbox :options="Object.keys(requirements)" class="w-9 md:w-5" listStyle="height:250px" striped>
      <template #option="row">
        <div class="flex justify-content-between">
          <span>{{ row.option }}</span>
          <span>{{ requirements[row.option] }}</span>
        </div>

      </template>
    </Listbox>
  </div>
</template>
<style scoped>
#build-list {
  height: 100vh;
}

#designs-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  padding: 10px;
  width: 100%;
  height: 33vh;
  overflow-y: auto;
}

#designs-container>div {
  padding: 10px;
  border: 1px solid #ccc;
  flex: 1 1 calc(25% - 20px);
  box-sizing: border-box;
  text-align: center;
}

#wrapperr {
  display: flex;
  flex-direction: row;
  width: 100%;
  padding: 10px;
  gap: 10px;
  box-sizing: border-box;
}

#selected-designs-container,
#requirements-container {
  flex: 1;
  padding: 10px;
  overflow-y: auto;
  box-sizing: border-box;
}

#selected-designs-container>div,
#requirements-container>div {
  border: 1px solid #ccc;
}
</style>
