<script>
export default {
	data() {
		return {
			username: sessionStorage.getItem("username"),
			picture: sessionStorage.getItem("picture"),
			loading: false,
			errorMessage: null,
		}
	},
	methods: {
		async changeUsername() {
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
	}
}
</script>

<template>
	<div class="main-box">
		<h1>Settings</h1>
		<div class="profile-box">
			<form @submit.prevent="changeUsername">
				<div class="user-username-bar">
					<input type="text" v-model="username" placeholder="Enter username">
					<button type="submit" :disabled="loading">Change</button>
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

.profile-box {
	margin: 2px;
	padding: 10px;
	display: flex;
	flex-direction: column;
	align-items: center;
}

.user-username-bar {
	display: flex;
	flex-direction: row;
	margin-bottom: 10px;
	align-items: center;
}

.user-username-bar input {
    width: 300px;
	padding: 10px;
    margin-right: 2px;
    border: 1px solid #ccc;
    border-radius: 8px;
	background-color: var(--color-2-a);
    font-size: 16px;
	color: var(--text-color-1)
}

.user-username-bar button {
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

.user-username-bar button:hover {
    background-color: var(--color-3-a);
}

.user-username-bar button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  opacity: 0.8;
}
</style>
