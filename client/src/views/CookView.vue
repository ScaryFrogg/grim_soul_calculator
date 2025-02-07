<script setup lang="ts">
import { ref, inject } from "vue"
import type { Ref } from "vue"
import type { Api, Recipe } from "@/types"
const api = inject<Api>("api")

const recipes: Ref<Recipe[]> = ref([])
api?.get("cook")
  .then(data => {
    recipes.value = data
  })
  .catch(e => console.error(e))
</script>
<template>
  <div class="flex flex-column align-items-center">
    <h2 class="text-center">Cook Recipes</h2>
    <div class="flex justify-content-between cook-box w-full xl:w-7" v-for="v, i in recipes" :key="i">
      <div class="w-6">
        <span>
          {{ v.ing1 }}x {{ v.ing1Quantity }}
        </span>
        <span v-if="v.ing2" class="ml-4">
          <span class="mr-4">+</span>
          {{ v.ing2 }}x {{ v.ing2Quantity }}
        </span>
      </div>
      <div class="flex justify-content-between w-4">
        <div>=</div>
        <span>
          {{ v.result }}x {{ v.resultQuantity }}
        </span>
      </div>
    </div>
  </div>
</template>
<style scoped>
.cook-box {
  padding: 10px;
  border: 1px solid white
}
</style>
