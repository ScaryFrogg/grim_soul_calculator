<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavigationComponent from "./components/NavigationComponent.vue"
</script>
<template>
  <NavigationComponent />
  <RouterView />
  <footer>
    <p class="text-xs text-center">
      This site is independent from brickworks games and is completely maintained by the community.
      If you wish you can contribute data <a class="text-primary"
        href="https://docs.google.com/spreadsheets/d/1sfATWVr7ePNWx7XmvqM0djCUoAi2OmRBWqqhjLBzXMU"
        target="_blank">here</a>.
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
  background-color: rgba(0, 0, 0, 0.8);
}

footer {
  padding: 5px;
  margin-top: auto;
  background-color: #18181b;
}

a:visited {
  color: inherit;
}
</style>
