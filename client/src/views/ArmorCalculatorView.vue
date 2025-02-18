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
const dmgReductionData = [
  { dmgr: 90, armor: 1485 },
  { dmgr: 85, armor: 935 },
  { dmgr: 80, armor: 660 },
  { dmgr: 75, armor: 495 },
  { dmgr: 70, armor: 385 },
  { dmgr: 65, armor: 306.5 },
  { dmgr: 60, armor: 247.5 },
  { dmgr: 55, armor: 202 },
  { dmgr: 50, armor: 165 },
  { dmgr: 45, armor: 135 },
  { dmgr: 40, armor: 110 },
  { dmgr: 35, armor: 89 },
  { dmgr: 30, armor: 71 },
  { dmgr: 25, armor: 55 },
  { dmgr: 20, armor: 41.25 },
  { dmgr: 15, armor: 30 },
  { dmgr: 10, armor: 18 },
  { dmgr: 5, armor: 9 },
]
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
function getReductionFromArmor(armor: number) {
  let reduction = armor / (165 + armor) * 100
  return Math.round((reduction + Number.EPSILON) * 100) / 100
}
const dmgReduction = computed(() => {
  let total = 0
  if (shield.value) {
    total += shield.value.armor
  }
  if (loadout.value) {
    total += loadout.value.reduce((acc, y) => acc + y.armor, 0)
  }
  return getReductionFromArmor(total)
})
const protection = computed(() => {
  let total = new Map()
  // if (shield.value) {
  //   total += shield.value.
  // }
  loadout.value.filter(x => x.protection).forEach((l) => {
    if (total.has(l.protectionType)) {
      total.set(l.protectionType, total.get(l.protectionType) + l.protection)
    } else {
      total.set(l.protectionType, l.protection)
    }
  })
  for (let [k, v] of total) {
    total.set(k, getReductionFromArmor(v))
  }
  return Array.from(total)
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
      <div class="flex align-items-center">
        <img src="/head_slot.png" />
        <div>
          <Dropdown v-model="loadout[0]" data-key="id" :options="armorPerSlot[0]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex align-items-center">
        <img src="/chest_slot.png" />
        <div>
          <Dropdown v-model="loadout[1]" data-key="id" :options="armorPerSlot[1]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex align-items-center">
        <img src="/hands_slot.png" />
        <div>
          <Dropdown v-model="loadout[2]" data-key="id" :options="armorPerSlot[2]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex align-items-center">
        <img src="/legs_slot.png" />
        <div>
          <Dropdown v-model="loadout[3]" data-key="id" :options="armorPerSlot[3]" optionLabel="name"
            placeholder="armorPerSlot" />
        </div>
      </div>
      <div class="flex align-items-center">
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
    <div class="p-3 w-full lg:w-3 flex flex-column justify-content-center">
      <h2>
        Dmg Reduction: <span class="text-primary">{{ dmgReduction }}</span>%
      </h2>
      <h2 v-if="protection[0]">
        Elemental Reduction:
        <p v-for="k, i of protection" :key="i">
          {{ k[0] }}
          <span :class="k[0]">
            :{{ k[1] }}
          </span>
          %
        </p>
      </h2>
    </div>
  </div>
  <div class="w-full flex flex-column align-items-center">
    <h2>Table of Armor required for Damage Reduction</h2>
    <Card class="w-full lg:w-4">
      <template #content>
        <div class="p-4">
          <DataTable size="small" :value="dmgReductionData">
            <Column sortable field="armor" header="Armor"></Column>
            <Column sortable field="dmgr" header="Dmg Reduction (%)"></Column>
          </DataTable>
        </div>
      </template>
    </Card>
  </div>
</template>
<style scoped>
img {
  max-width: 100px;
  height: auto;
  cursor: pointer;
}

.fire,
.cold,
.decay {
  text-transform: capitalize;
}

.fire {
  color: orange;
}

.cold {
  color: #9abdff;
}

.decay {
  color: #b3ff3f;
  ;
}
</style>
