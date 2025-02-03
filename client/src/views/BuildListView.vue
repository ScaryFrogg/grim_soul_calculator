<script setup lang="ts">
import { ref } from "vue"
import type { Design, Requirement } from "@/types"

const designs = ref<Design[]>([]);
const requirements = ref<Record<string, number>>({});
const selectedDesigns = ref<Record<number, { design: Design; count: number }>>({});
fetch(`http://localhost:3000/designs`)
  .then(d => d.json().then(data => {
    designs.value = data
  }))
  .catch(e => console.error(e))

const getRequirements = async (design: Design) => {
  const data = await fetchRequirements(design.id)
  if (!data) return
  data.forEach((req: Requirement) => {
    if (req.name in requirements.value) {
      requirements.value[req.name] = requirements.value[req.name] + req.quantity
    }
    else {
      requirements.value[req.name] = req.quantity
    }
  })
  design.requirements = data
  addDesign(design)
}

const fetchRequirements = async (planId: number): Promise<Requirement[] | null> => {
  return fetch(`http://localhost:3000/requirement/${planId}`)
    .then(d => d.json().then((data: Requirement[]) => {
      return data
    }))
    .catch(e => {
      console.error(e)
      return null
    })
}

const addDesign = (design: Design) => {
  if (selectedDesigns.value[design.id]) {
    selectedDesigns.value[design.id].count++;
  } else {
    selectedDesigns.value[design.id] = { design, count: 1 };
  }
};

const removeDesign = (design: Design) => {
  if (selectedDesigns.value[design.id].count != 1) {
    selectedDesigns.value[design.id].count--;
  } else {
    delete selectedDesigns.value[design.id];
  }
  design.requirements.forEach(req => {
    requirements.value[req.name] = requirements.value[req.name] - req.quantity
    if (requirements.value[req.name] == 0) {
      delete requirements.value[req.name];
    }
  });
};

</script>
<template>
  <div id="build-list">
    <div id="designs-container">
      <div v-for="(design, i) in designs" :key=i>
        <div @click="getRequirements(design)">{{ design.name }}</div>
      </div>
    </div>
    <div id="wrapperr">
      <div id="selected-designs-container">
        <p class="m-0">Selected</p>

        <div v-for="(pair, i) in selectedDesigns" :key=i>
          <div>{{ pair.design.name }} :: {{ pair.count }} :: <div @click="removeDesign(pair.design)">-</div>
          </div>
        </div>
      </div>

      <div id="requirements-container">
        <p class="m-0 m-b-1">Total Materials</p>
        <div v-for="(i, f) in requirements" :key=f>
          <div>{{ f }}{{ i }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
#build-list {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100vh;
}

#designs-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  padding: 10px;
  width: 100%;
  height: 33vh;
  overflow-y: auto;
}

#designs-container>div {
  padding: 10px;
  border: 1px solid #ccc;
  flex: 1 1 calc(25% - 20px);
  box-sizing: border-box;
  text-align: center;
}

#wrapperr {
  display: flex;
  flex-direction: row;
  width: 100%;
  padding: 10px;
  gap: 10px;
  box-sizing: border-box;
}

#selected-designs-container,
#requirements-container {
  flex: 1;
  padding: 10px;
  overflow-y: auto;
  border: 1px solid #ccc;
  box-sizing: border-box;
}
</style>
