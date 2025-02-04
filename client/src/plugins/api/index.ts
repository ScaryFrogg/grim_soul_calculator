import type { App } from 'vue'
class Api {
	constructor() { }
	baseUrl = "http://localhost:3000/"
	async get(url: string) {
		const res = await fetch(this.baseUrl + url, {
			method: 'GET',
		});
		return await res.json();
	}
}
//TODO this can be improved further for other projects to accept options
export default {
	install: (app: App) => {
		const api = new Api()
		app.provide('api', api)
	}
}
