import { createMemoryHistory, createRouter } from 'vue-router'

import DesignsView from './components/DesignsView.vue'
import DesignView from './components/DesignView.vue'
import MaterialView from './components/MaterialView.vue'

const routes = [
  { path: '/', component: DesignsView },
  { path: '/design/:id', component: DesignView },
  { path: '/material/:id', component: MaterialView },
]

export default createRouter({
  history: createMemoryHistory(),
  routes,
})
