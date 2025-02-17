<script setup lang="ts">

</script>
<template>
  <div class="mt-8 flex justify-content-evenly">
    <div>
      <h2>
        Learn About the game
      </h2>
      <ul>
        <li>
          <RouterLink to="/weapons">Weapons</RouterLink>
        </li>
        <li>
          <RouterLink to="/armor">Armor Sets</RouterLink>
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
      </ul>
    </div>
    <div>
      <h2>
        Calculators
      </h2>
      <ul>
        <li>
          <RouterLink to="/armorCalculator">Armor-DmgReduction Calculator</RouterLink>
        </li>
        <li>
          <RouterLink to="/buildList">Building Calculator</RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>
