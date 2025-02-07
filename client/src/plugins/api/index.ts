import type { App } from 'vue'
class Api {
	constructor() { }
	baseUrl = import.meta.env.VITE_API_PATH || "http://172.104.249.149/gscapi/"
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
