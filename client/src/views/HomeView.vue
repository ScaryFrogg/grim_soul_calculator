<script setup lang="ts">

</script>
<template>
  <div>
    <Card>
      <template #title>
        Learn About the game
      </template>
      <template #content>
        <ul>
          <li>
            <RouterLink to="/weapons">Weapons</RouterLink>
          </li>
          <li>
            <RouterLink to="/trades">Trades</RouterLink>
          </li>
          <li>
            <RouterLink to="/designs">Designs</RouterLink>
          </li>
          <li>
            <RouterLink to="/cook">Cook Recipes</RouterLink>
          </li>
          <li>
            <RouterLink to="/items">Items</RouterLink>
          </li>
        </ul>
      </template>
    </Card>
    <Card>
      <template #title>
        Calculators
      </template>
      <template #content>
        <ul>
          <li>
            <RouterLink to="/buildList">Building Calculator</RouterLink>
          </li>
        </ul>
      </template>
    </Card>
  </div>
</template>
