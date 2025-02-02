<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useRouter } from "vue-router"
import type { Trade } from "@/types"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import InputText from "primevue/inputtext"
import { FilterMatchMode } from "primevue/api";
const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS },
});
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
  <div class="flex justify-content-center">
    <DataTable :value="trades" tableStyle="max-width: 100vw; width: 50rem" removableSort :filters="filters"
      :globalFilterFields="['giveName', 'getName']" stripedRows>
      <template #header>
        <div class="flex justify-content-center">
          <InputText v-model="filters['global'].value" placeholder="Search" />
        </div>
      </template>
      <Column field="give" header="Give" sortField="giveName" sortable bodyClass="px-0 md:px-5">
        <template #body="{ data }">
          <span @click="navigate(data.giveId)">{{ data.giveName }}</span>
          X {{ data.giveQuantity }}
        </template>
      </Column>
      <Column field="split" bodyClass="p-0">
        <template #body>
          ->
        </template>
      </Column>
      <Column field="get" header="Get" sortField="getName" sortable bodyClass="px-0 md:px-5">
        <template #body="{ data }">
          <span @click="navigate(data.getId)">{{ data.getName }}</span>
          X {{ data.getQuantity }}
        </template>
      </Column>
    </DataTable>
  </div>
</template>
