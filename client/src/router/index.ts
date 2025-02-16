import { createRouter, createWebHistory } from "vue-router"
import HomeView from "../views/HomeView.vue"
import DesignsView from "../views/DesignsView.vue"
import DesignView from "../views/DesignView.vue"
import MaterialView from "../views/MaterialView.vue"
import BuildListView from "../views/BuildListView.vue"
import TradesView from "../views/TradesView.vue"
import WeaponsView from "../views/WeaponsView.vue"
import CookView from "../views/CookView.vue"
import ArmorView from "../views/ArmorView.vue"
import ArmorCalculatorView from "../views/ArmorCalculatorView.vue"

const routes = [
  { path: "/", component: HomeView },
  { path: "/design/:id", component: DesignView },
  { path: "/designs", component: DesignsView },
  { path: "/material/:id", component: MaterialView },
  { path: "/buildList/", component: BuildListView },
  { path: "/trades", component: TradesView },
  { path: "/weapons", component: WeaponsView },
  { path: "/cook", component: CookView },
  { path: "/armor/:id?", component: ArmorView },
  { path: "/armorCalculator", component: ArmorCalculatorView },
]


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

export default router
