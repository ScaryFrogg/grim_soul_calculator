<script setup lang="ts">
import { defineProps } from "vue"
import type { Enemy } from "@/types"
const props = defineProps<{
  enemy: Enemy
}>()
const enemy = props.enemy
</script>
<template>
  <div>
    <h4>{{ enemy.name }}</h4>
    <div>
      <h5>Base</h5>
      <div>Health: {{ enemy.health }}</div>
      <div>Attack: {{ enemy.attack }}</div>
      <div>Armor: {{ enemy.armor }}</div>
      <div>Elemental Attack: <span :class="enemy.elemAttType">{{ enemy.elemAtt ?? 0 }}</span></div>
      <div>Specila Ability:
        <ul>
          <li v-for="str of enemy.special.split(';')" :key="str">
            {{ str }}
          </li>
        </ul>

      </div>
    </div>
    <div>
      <h5>Protection Modifiers</h5>
      <div>Fire: {{ enemy.fire }}</div>
      <div>Cold: {{ enemy.cold }}</div>
      <div>Pierce: {{ enemy.pierce }}</div>
      <div>Reap: TODO</div>
      <div>Poison: TODO</div>
      <div>Bleed: TODO</div>
      <div>Frostbite: TODO</div>
      <div>Ignite: TODO</div>

    </div>
    <div>
      <h5>CC</h5>
      <div>Fear: TODO</div>
      <div>Stun: TODO</div>
      <div>Freeze: TODO</div>
      <div>Blind: TODO</div>
      <div>Icy Slow: TODO</div>
      <div>Shrieking Slow: TODO</div>
      <div>EoE: TODO</div>
      <div>SoC(1,2,3): TODO</div>
      <div>Fracture Armor: {{ enemy.fractureArmor }}</div>
      <div>Fracture Weapon: {{ enemy.fractureWeapon }}</div>
    </div>
    <div>
      <h5>Details</h5>
      <div>Soul Eater Charge: {{ enemy.soulEater }}</div>
      <div>Human/Monster: {{ enemy.horM }}</div>
      <div>Insta Killable: {{ enemy.instakill }}</div>
      <div>Thunder God's Knuckledusters: {{ enemy.tgk }}</div>
    </div>
  </div>
</template>
