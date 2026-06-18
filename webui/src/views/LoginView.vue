<script>
export default {
	data() {
		return {
			username: "",
			loading: false,
			errorMessage: null,
		}
	},
	methods: {
		async doLogin() {
			this.loading = true;
			try {
				if (this.username.length < 3 || this.username.length > 16) throw "Invalid username"

				const response = await this.$axios.post('/session', {
					username: this.username,
				});

				sessionStorage.setItem('userID', response.data.userID);
				sessionStorage.setItem('username', response.data.username);autorizzazione
				sessionStorage.setItem('picture', response.data.picture);
				sessionStorage.setItem('token', response.data.userID);

				this.$router.push('/home');
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
		if (sessionStorage.userID) {
			this.$router.push("/home");
			return
		}
		sessionStorage.clear()
	}
}
</script>

<template>
	<div class="main-box">
		<div class="login-box">
			<h1>Login</h1>
			<form @submit.prevent="doLogin">
				<label for="username-in">Username</label>
				<input type="text" id="username-in" v-model="username" placeholder="Enter your username">
				<button type="submit" :disabled="loading">Submit</button>
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
  	justify-content: center;
  	align-items: center;
}

.login-box {
    background: var(--color-5);
    padding: 30px;
	border-radius: 15px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    width: 300px;
    text-align: center;
}

.login-box h1 {
    margin-bottom: 20px;
    color: #333;
}

.login-box form {
	text-align: left;
}

.login-box label {
    display: block;
    margin-bottom: 5px;
}

.login-box input {
    width: 100%;
    padding: 10px;
    margin-bottom: 15px;
    border: 1px solid #ccc;
    border-radius: 8px;
    font-size: 16px;
}

.login-box button {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 8px;
	background-color: #0078ff;
    color: white;
    font-size: 16px;
	cursor: pointer;
}

.login-box button:hover {
    background-color: #0062d1;
}

.login-box button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  opacity: 0.8;
}
</style>
