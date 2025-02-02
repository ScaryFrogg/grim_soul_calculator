import { createRouter, createWebHistory } from "vue-router"
import DesignsView from "../views/DesignsView.vue"
import DesignView from "../views/DesignView.vue"
import MaterialView from "../views/MaterialView.vue"
import BuildListView from "../views/BuildListView.vue"
import TradesView from "../views/TradesView.vue"
import WeaponsView from "../views/WeaponsView.vue"

const routes = [
  { path: "/", component: DesignsView },
  { path: "/design/:id", component: DesignView },
  { path: "/designs", component: DesignsView },
  { path: "/material/:id", component: MaterialView },
  { path: "/buildList/", component: BuildListView },
  { path: "/trades/", component: TradesView },
  { path: "/weapons/", component: WeaponsView },
]


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

export default router
