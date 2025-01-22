import { createRouter, createWebHistory } from 'vue-router'
import DesignsView from '../views/DesignsView.vue'
import DesignView from '../views/DesignView.vue'
import MaterialView from '../views/MaterialView.vue'
import BuildListView from '../views/BuildListView.vue'

const routes = [
  { path: '/', component: DesignsView },
  { path: '/design/:id', component: DesignView },
  { path: '/material/:id', component: MaterialView },
  { path: '/buildList/', component: BuildListView },
]


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

export default router
