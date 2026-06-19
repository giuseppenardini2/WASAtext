import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import UsersView from '../views/UsersView.vue'
import SettingsView from '../views/SettingsView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/users', component: UsersView},
		{path: '/settings', component: SettingsView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
