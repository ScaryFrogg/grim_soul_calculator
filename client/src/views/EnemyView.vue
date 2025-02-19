<script setup lang="ts">
import { ref, onMounted, inject } from "vue"
import { useRoute } from "vue-router"
import type { Api, Enemy } from "@/types"
import EnemyComponent from '../components/EnemyComponent.vue';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
const api = inject<Api>("api")
const route = useRoute()
const enemyData = ref<Enemy[] | null>(null)
onMounted(() => {
  const id = route.params.id
  api?.get(`enemy/${id}`)
    .then(data => {
      enemyData.value = data
    })
    .catch(e => console.error(e))

})
</script>
<template>
  <div class="flex" v-if="enemyData">
    <TabView v-if="enemyData.length > 1">
      <TabPanel v-for="k of enemyData" :key="k.difficulty" :header="k.difficulty">
        <EnemyComponent :enemy="k" />
      </TabPanel>
    </TabView>
    <EnemyComponent v-else :enemy="enemyData[0]" />
  </div>
</template>
