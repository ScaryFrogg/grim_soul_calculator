import PrimeVue from "primevue/config"
import "primevue/resources/primevue.min.css"
import "primevue/resources/themes/aura-dark-purple/theme.css"
import "../node_modules/primeflex/primeflex.css"

import Ripple from "primevue/ripple"

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Button from "primevue/button"


const app = createApp(App);

app.use(router)
app.use(PrimeVue);

//PrimeVue reused stuff //TODO extract later
app.component('Button', Button);
app.directive("ripple", Ripple)


app.mount('#app')
