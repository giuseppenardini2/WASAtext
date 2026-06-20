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
				if (this.username.length < 3 || this.username.length > 16) throw "Invalid username"

				const response = await this.$axios.put(`/users/${sessionStorage.getItem("userID")}/username`, {
					username: this.username
				},
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				sessionStorage.setItem("username", this.username)
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
			<form @submit.prevent="">
				<div class="user-picture-bar">
					<img src="@/assets/images/default-user-picture.jpg" alt="User picture">
					<button type="submit" :disabled="loading">Change</button>
				</div>
			</form>
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
	gap: 20px;
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

.profile-box input {
    width: 300px;
	padding: 10px;
    border: 1px solid #ccc;
    border-radius: 8px;
	background-color: var(--color-2-a);
    font-size: 16px;
	color: var(--text-color-1)
}

.profile-box button {
    width: 100px;
	padding: 10px;
    border: none;
    border-radius: 8px;
	background-color: var(--color-3);
    color: white;
    font-size: 16px;
	cursor: pointer;
}

.profile-box button:hover {
    background-color: var(--color-3-a);
}

.profile-box button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  opacity: 0.8;
}

.user-username-bar {
	display: flex;
	flex-direction: row;
	gap: 5px;
	margin-bottom: 10px;
	align-items: center;
}

.user-picture-bar {
	display: flex;
	flex-direction: column;
	gap: 5px;
	margin-bottom: 10px;
	align-items: center;
}

.user-picture-bar img {
	width: 150px;
	height: 150px;
	object-fit: cover;
	border-radius: 50%;
	box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}
</style>
