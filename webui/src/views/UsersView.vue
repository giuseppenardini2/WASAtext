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
			this.errorMessage = "";
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
		},
		async startNewConversation(recipientID) {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/conversations`, {
					secondUserID: recipientID,
				},
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				this.$router.push("/home");
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
					<input type="text" v-model="username" placeholder="Enter username">
					<button type="submit" :disabled="loading">Submit</button>
				</div>
				<div v-if="users.length > 0">
      				<div 
        				v-for="user in users"
        					:key="user.id" 
        				class="user-item"
     				>
        				<strong>{{ user.username }}</strong>
						<button type="button" @click="startNewConversation(user.id)" :disabled="loading">New chat</button>
					</div>
    			</div>
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

.users-box button:disabled {
	background-color: #6c757d;
	cursor: not-allowed;
	opacity: 0.8;
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
    color: var(--text-color-4);
    font-size: 16px;
	cursor: pointer;
}

.user-search-bar button:hover {
    background-color: var(--color-3-a);
}

.user-item {
	display: flex;
	gap: 10px;
	margin: 0.5rem;
}

.user-item button {
	width: 80px;
	padding: 10px;
	margin-left: 2px;
    border: none;
    border-radius: 8px;
	background-color: var(--color-3);
    color: var(--text-color-4);
    font-size: 10px;
	cursor: pointer;
}

.user-search-bar button:hover {
    background-color: var(--color-3-a);
}
</style>
