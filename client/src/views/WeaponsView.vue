<script setup lang="ts">
import { onMounted, ref } from "vue"
import type { WeaponInfo } from "@/types"
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
  <h2 class="text-center">Weapons</h2>
  <DataTable size="small" :filters="filters" :value="weapons" tableStyle="max-width: 100vw;"
    :globalFilterFields="['name', 'damage']" removableSort stripedRows :pt="{ root: { 'style': 'overflow:scroll' } }">
    <template #header>
      <div class="flex justify-content-center">
        <InputText v-model="filters['global'].value" placeholder="Search" />
      </div>
    </template>
    <Column field="name" header="Weapon" sortable headerClass="px-1 md:px-3" bodyClass="px-1 md:px-5"></Column>
    <Column field="damage" header="Dmg" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
    <Column field="attackSpeed" header="AS" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
    <Column field="s1" header="s+1" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
    <Column field="s2" header="s+2" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
    <Column field="s3" header="s+3" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
    <Column field="s4" header="s+4" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
    <Column field="s5" header="s+5" sortable headerClass="px-1 md:px-3" bodyClass="px-0 md:px-5"></Column>
  </DataTable>
</template>
