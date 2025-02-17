<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavigationComponent from "./components/NavigationComponent.vue"
</script>
<template>
  <NavigationComponent />
  <RouterView />
  <footer>
    <p class="text-xs">
      This site is independent from brickworks games and is completely maintained by the community.
    </p>
  </footer>
</template>
<style>
html {
  background-image: url(/banner.png);
  background-position: center top
}

body {
  min-height: 100vh;
  margin: 0;
  background-color: rgba(0, 0, 0, 0.75);
}

footer {
  padding: 5px;
  margin-top: auto;
  background-color: #18181b;
}
</style>
