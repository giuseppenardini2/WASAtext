<script>
export default {
	data() {
		return {
			errorMessage: null,
			loading: false,
			some_data: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errorMessage = null;
			try {
				let response = await this.$router.push("/home");
				this.some_data = response.data;
			} catch (e) {
				this.errorMessage = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div>
			Hello world!
		</div>
		<ErrorMessage v-if="errorMessage" :message="errorMessage"></ErrorMessage>
	</div>
</template>

<style>
</style>
