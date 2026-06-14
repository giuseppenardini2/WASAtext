<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>

	<div class="page">
        <aside class="sidebar">
            <nav>
				<div>
					<a href="/" class="logo">WASAtext</a>
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
            <RouterView />
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
	background-color: var(--color2);
}

/* SIDEBAR */
.logo {
    display: block;
    padding: 20px 0;
    color: var(--color1);
	text-align: center;
    text-decoration: none;
    font-weight: bold;
}

.sidebar {
	width: 100px;
  	background-color: var(--color3);
  	color: #fff;
}

.sidebar ul {
	list-style: none;
	padding: 0;
	margin: 0;
}

.sidebar ul li {
	margin-bottom: 15px;
	padding: 0 10px;
}

.sidebar ul li a {
	color: var(--color2);
	text-decoration: none;
	font-size: 16px;
	display: flex;
	justify-content: center;
    align-items: center;
    padding: 8px;
	border-radius: 6px;
	transition: all 0.3s ease;
}

.sidebar ul li a:hover {
	background-color: #334155;
	color: #fff;
}

.sidebar ul li a.router-link-exact-active {
	background-color: #334155;
	color: #fff;
}

.icon {
	width: 28px;
	height: 28px;
	max-width: 100%;
	stroke: var(--color1);
	stroke-width: 2;
	stroke-linecap: round;
	stroke-linejoin: round;
	fill: none;
}

/* MAIN */
.main {
	flex: 1;
	background-color: var(--color2);
	color: #333;
}
</style>
