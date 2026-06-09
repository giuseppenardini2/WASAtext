<script>
export default {
	data() {
		return {
			username: "",
			picture: "",
			loading: false,
			errorMessage: null,
		}
	},
	methods: {
		login() {
			this.loading = true;

			if (this.username.length < 3 || this.username.length > 16) throw "Invalid username"

			const response = await this.$axios.post('/session', {
          		username: this.username,
        	});

			sessionStorage.setItem('userID', response.data.userID);
			sessionStorage.setItem('username', response.data.username);
			sessionStorage.setItem('picture', response.data.picture);

			this.loading = false;
			// this.$router.push("");
		}
	},
	mounted() {
		if (sessionStorage.userID) {
			// this.$router.push("");
		}
	}
}
</script>

<template>
	<div>
		<div>
			<h1>Login</h1>
			<p>Username</p>
			<input type="text" v-model="username" placeholder="Enter your username">
			<button type="submit" @click="login">Submit</button>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
