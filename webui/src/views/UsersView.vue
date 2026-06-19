<script>
export default {
	data() {
		return {
			username: "",
			users: [],
			loading: false,
			errorMessage: null,
		}
	},
	methods: {
		async handleSearch() {
			this.loading = true;
			try {
				const response = await this.$axios.get('/users', {
					params: {
						username: this.username
					}
				});
				this.users = response.data || [];
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		}
	},
	mounted() {
		this.handleSearch();
	}
}
</script>

<template>
	<div class="main-box">
		<h1>Users</h1>
		<div class="users-box">
			<form @submit.prevent="handleSearch">
				<div class="user-search-bar">
					<input type="text" id="username-in" v-model="username" placeholder="Enter username">
					<button type="submit" :disabled="loading">Submit</button>
				</div>
				<ul v-if="users.length > 0">
      				<li 
        				v-for="user in users"
        					:key="user.id" 
        				class="user-item"
     				>
        				<span>Nickname:</span> 
        				<strong>{{ user.username }}</strong>
     	 			</li>
    			</ul>
				<div v-else>
					No user found
				</div>
			</form>
		</div>
		<div>
			<ErrorMessage v-if="errorMessage" :message="errorMessage"></ErrorMessage>
		</div>
	</div>
</template>

<style scoped>
.main-box {
	width: 100%;
  	height: 100%;
  	display: flex;
	flex-direction: column;
	padding: 10px;
}

.main-box h1 {
	color: var(--text-color-1);
}

.users-box {
	margin: 2px;
	padding: 10px;
	background-color: var(--color-5);
	border-radius: 5px;
	display: flex;
	flex-direction: column;
}

.user-search-bar {
	display: flex;
	flex-direction: row;
	margin-bottom: 10px;
}

.user-search-bar input {
    width: 300px;
	padding: 10px;
    margin-right: 2px;
    border: 1px solid #ccc;
    border-radius: 8px;
	background-color: var(--color-2-a);
    font-size: 16px;
	color: var(--text-color-1)
}

.user-search-bar button {
    width: 100px;
	padding: 10px;
	margin-left: 2px;
    border: none;
    border-radius: 8px;
	background-color: var(--color-3);
    color: white;
    font-size: 16px;
	cursor: pointer;
}

.user-search-bar button:hover {
    background-color: var(--color-3-a);
}

.user-search-bar button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  opacity: 0.8;
}
</style>
