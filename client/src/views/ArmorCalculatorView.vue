<script setup lang="ts">
import { ref, onMounted, inject, watch, computed } from "vue"
import Dropdown from "primevue/dropdown"
import type { Api, ArmorInfo, Shield } from "@/types"
const api = inject<Api>("api")
const sets = ref<ArmorInfo[]>()
const selectedSet = ref<ArmorInfo>()
const loadout = ref<ArmorInfo[]>(Array(5))
const shields = ref<Shield[]>([])
const shield = ref<Shield | null>(null);
const armorPerSlot = ref(Array(5))
onMounted(() => {
  api?.get(`armor`)
    .then(data => {
      sets.value = data
    })
    .catch(e => console.error(e))
  api?.get(`shields`)
    .then(data => {
      shields.value = data
    })
    .catch(e => console.error(e))
  api?.get(`armorPerSlot`)
    .then(data => {
      armorPerSlot.value = data
    })
    .catch(e => console.error(e))
})
const dmgReduction = computed(() => {
  let total = 0
  if (shield.value) {
    total += shield.value.armor
  }
  if (loadout.value) {
    total += loadout.value.reduce((acc, y) => acc + y.armor, 0)
  }
  let reduction = total / (165 + total) * 100
  return Math.round((reduction + Number.EPSILON) * 100) / 100
})
watch(selectedSet, async (newS) => {
  if (newS && selectedSet.value) {
    await api?.get(`armor/set/${newS.id}`)
      .then(data => {
        loadout.value = data
      })
  }
})
</script>
<template>
  <div class="flex flex-column lg:flex-row">
    <Card class="w-full lg:w-4">
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

    <div class="p-3 w-full lg:w-5 flex flex-column">
      <div class="flex">
        <img src="/head_slot.png" />
        <div>
          <Dropdown v-model="loadout[0]" data-key="id" :options="armorPerSlot[0]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex">
        <img src="/chest_slot.png" />
        <div>
          <Dropdown v-model="loadout[1]" data-key="id" :options="armorPerSlot[1]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex">
        <img src="/hands_slot.png" />
        <div>
          <Dropdown v-model="loadout[2]" data-key="id" :options="armorPerSlot[2]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex">
        <img src="/legs_slot.png" />
        <div>
          <Dropdown v-model="loadout[3]" data-key="id" :options="armorPerSlot[3]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex">
        <img src="/boots_slot.png" />
        <div>
          <Dropdown v-model="loadout[4]" data-key="id" :options="armorPerSlot[4]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex align-items-center">
        <img src="/shield_slot.png" />
        <div>
          <Dropdown v-model="shield" :options="shields" optionLabel="name" placeholder="Shields" />
        </div>
      </div>
    </div>
    <div class="p-3 w-full lg:w-3">
      <h2>
        Total Dmg Reduction: <span class="text-primary">{{ dmgReduction }}</span>%
      </h2>
    </div>
  </div>
</template>
<style scoped>
img {
  max-width: 100px;
  height: auto;
  cursor: pointer;
}
</style>
