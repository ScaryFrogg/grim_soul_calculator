import type { App } from 'vue'
import PrimeVue from "primevue/config"
import "primevue/resources/primevue.min.css"
import "primevue/resources/themes/aura-dark-purple/theme.css"
import "../../../node_modules/primeflex/primeflex.css"

import Ripple from "primevue/ripple"

import Button from "primevue/button"
import Card from "primevue/card"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import Listbox from "primevue/listbox"

export default {
	install: (app: App) => {
		app.use(PrimeVue)
		app.component('Button', Button);
		app.component('Card', Card);
		app.component('DataTable', DataTable);
		app.component('Column', Column);
		app.component('Listbox', Listbox);

		app.directive("ripple", Ripple)
	}
}
