<script setup lang="ts">
import { ref } from "vue"
import type { Ref } from "vue"
import type { Design } from "@/types"

const designs: Ref<Design[]> = ref([])
fetch(`http://localhost:3000/designs`)
  .then(d => d.json().then(data => {
    designs.value = data
  }))
  .catch(e => console.error(e))

</script>
<template>
  <div>
    <router-link to="buildList/">make a requirement list</router-link>
    <div v-for="(f, i) in designs" :key=i>
      <router-link :to="`design/${f.id}`">{{ f.name }}</router-link>
    </div>

    {{ designs }}
  </div>
</template>
