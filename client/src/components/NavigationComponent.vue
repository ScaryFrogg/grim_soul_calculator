<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import Menubar from "primevue/menubar"

const router = useRouter()
const navigate = (whereTo: string | undefined) => {
  if (whereTo) {
    router.push(whereTo)
  }
}
const items = ref([
  { label: 'Home', url: '/', },
  { label: 'Armor-DmgReduction', url: '/armorcalculator', },
  { label: 'Crafting List', url: '/buildList', },
  { separator: true },
  { label: 'Designs', url: '/designs', },
  { label: 'Trades', url: '/trades', },
  { label: 'Cook', url: '/cook', },
  { label: 'Weapons', url: '/weapons', },
  { label: 'Armor Sets', url: '/armor/', },

]);

//  label: 'Data', items: [{ label: 'Designs', url: '/designs', }, { label: 'Trades', url: '/trades', }, { label: 'Materials', url: '/materials', }, { separator: true }, { label: 'Weapons', url: '/weapons', },      /* {        label: 'Grindstone',        url: '/grindstone',      }, */      { label: 'Enemies', url: '/enemies', },]
</script>
<template>
  <Menubar :model="items" :pt="{ root: { 'style': 'border-radius: 0px' } }">
    <template #item="{ item, props, hasSubmenu, root }">
      <a v-ripple class="flex align-items-center" v-bind="props.action" @click="navigate(item.url)">
        <span class="ml-2">{{ item.label }}</span>
        <i v-if="hasSubmenu"
          :class="['pi pi-angle-down', { 'pi-angle-down ml-2': root, 'pi-angle-right ml-auto': !root }]"></i>
      </a>
    </template>
  </Menubar>
</template>
