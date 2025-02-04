import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import prime from './plugins/primeVue'
import api from './plugins/api'

const app = createApp(App);

app.use(router)
app.use(prime)
app.use(api)


app.mount('#app')
