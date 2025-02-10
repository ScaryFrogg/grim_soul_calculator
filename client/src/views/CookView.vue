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
          {{ v.ing1Quantity }}
          <span class="text-primary">
            {{ v.ing1 }}
          </span>
        </span>
        <span v-if="v.ing2">
          &nbsp;
          +
          {{ v.ing2Quantity }}
          <span class="text-primary">
            {{ v.ing2 }}
          </span>
        </span>
      </div>
      <div class="flex justify-content-between w-5">
        <div>=</div>
        <div>
          <span>
            &nbsp; {{ v.resultQuantity }} &nbsp;
          </span>
          <span class="text-primary">
            {{ v.result }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.cook-box {
  padding: 8px;
  border: 1px solid silver
}
</style>
