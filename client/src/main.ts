import PrimeVue from "primevue/config"
import "primevue/resources/primevue.min.css"
import "primevue/resources/themes/aura-dark-purple/theme.css"
import "../node_modules/primeflex/primeflex.css"

import Ripple from "primevue/ripple"

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Button from "primevue/button"
import Card from "primevue/card"
import DataTable from "primevue/datatable"
import Column from "primevue/column"



const app = createApp(App);

app.use(router)
app.use(PrimeVue);

//PrimeVue reused stuff //TODO extract later
app.component('Button', Button);
app.component('Card', Card);
app.component('DataTable', DataTable);
app.component('Column', Column);

app.directive("ripple", Ripple)


app.mount('#app')
