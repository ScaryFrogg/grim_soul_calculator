<script setup lang="ts">
import { onMounted, ref } from "vue"
import type { WeaponInfo } from "@/types"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import InputText from "primevue/inputtext"
import { FilterMatchMode } from "primevue/api";
const weapons = ref<WeaponInfo[]>([])
const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS },
});

const apiRoot = "http://localhost:3000/"
onMounted(() => {
  fetch(`${apiRoot}weapons`)
    .then(d => d.json().then(data => {
      weapons.value = data
    }))
    .catch(e => console.error(e))
})

</script>
<template>
  <div class="card">
    <DataTable :filters="filters" :value="weapons" tableStyle="min-width: 50rem"
      :globalFilterFields="['name', 'damage']">
      <template #header>
        <div class="flex justify-content-center">
          <InputText v-model="filters['global'].value" placeholder="Search" />
        </div>
      </template>
      <Column field="name" header="Weapon" sortable style="width: 23%"></Column>
      <Column field="damage" header="Dmg" sortable style="width: 11%"></Column>
      <Column field="attackSpeed" header="AS" sortable style="width: 11%"></Column>
      <Column field="s1" header="s+1" sortable style="width: 11%"></Column>
      <Column field="s2" header="s+2" sortable style="width: 11%"></Column>
      <Column field="s3" header="s+3" sortable style="width: 11%"></Column>
      <Column field="s4" header="s+4" sortable style="width: 11%"></Column>
      <Column field="s5" header="s+5" style="width: 11%"></Column>
    </DataTable>
  </div>
</template>
