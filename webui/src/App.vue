<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	name: 'App',
	data() {
		return {
			isLoggedIn: false
		};
	},
	methods: {
		handleLogin() {
			const token = sessionStorage.getItem('token');
			if ((token === null) || (token.trim() === '')) {
				this.isLoggedIn = false;
				this.$router.push("/login");
				return;
			}
			this.isLoggedIn = true;
			// this.$router.push("/home");
		}
	}
}
</script>

<template>

	<div class="page">
        <aside class="sidebar" v-if="isLoggedIn">
            <nav>
				<div>
					<a href="/" class="logo">Wt</a>
					<ul>
						<li>
							<RouterLink to="/home">
								<svg class="icon"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
							</RouterLink>
						</li>
						<li>
							<RouterLink to="/users">
								<svg class="icon"><use href="/feather-sprite-v4.29.0.svg#users"/></svg>
							</RouterLink>
						</li>
						<li>
							<RouterLink to="/settings">
								<svg class="icon"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg>
							</RouterLink>
						</li>
					</ul>
				</div>
            </nav>
        </aside>

        <main class="main">
            <RouterView @login-success="handleLogin"/>
        </main>

    </div>
</template>

<style scoped>
/* PAGE */
.page {
	display: flex;
  	width: 100vw;
  	height: 100vh;
	margin: 0;
	background-color: var(--color-2);
}

/* SIDEBAR */
.logo {
    display: block;
    padding: 20px 0;
	font-size: 23px;
    font-weight: bold;
    color: var(--color-1);
	text-align: center;
    text-decoration: none;
}

.sidebar {
	width: 65px;
  	background-color: var(--color-3);
}

.sidebar ul {
	list-style: none;
	padding: 0;
	margin: 0;
}

.sidebar ul li {
	margin-bottom: 15px;
	display: flex;
    justify-content: center;
}

.sidebar ul li a {
	color: var(--color-2);
	text-decoration: none;
	font-size: 16px;
	display: flex;
	justify-content: center;
    align-items: center;
    padding: 0;
	height: 48px;
	width: 48px;
	border-radius: 50%;
	transition: all 0.3s ease;
}

.sidebar ul li a:hover, 
.sidebar ul li a.router-link-exact-active {
	background-color: var(--color-5);
}

.icon {
	width: 28px;
	height: 28px;
	max-width: 100%;
	stroke: var(--color-1);
	stroke-width: 2;
	stroke-linecap: round;
	stroke-linejoin: round;
	fill: none;
}

/* MAIN */
.main {
	flex: 1;
	background-color: var(--color-2);
}
</style>
