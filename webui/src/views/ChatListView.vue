<script>
export default {
	data() {
		return {
			username: "",
			pictureURI: "",
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
			sessionStorage.setItem('pictureURI', response.data.pictureURI);

			this.loading = false;
			this.$router.push("/home");
		}
	},
	mounted() {
		if (sessionStorage.userID) {
			this.$router.push("/home");
			return
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

		<ErrorMessage v-if="errorMessage" :message="errorMessage"></ErrorMessage>
	</div>
</template>

<style>
</style>
