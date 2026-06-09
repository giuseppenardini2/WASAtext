import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', component: LoginView},
		{path: '/link1', component: HomeView},
		{path: '/home', component: HomeView},
	]
})

export default router
