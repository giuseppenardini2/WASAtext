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
				let response = await this.$axios.get("/home");
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
		<ErrorMsg v-if="errorMessage" :msg="errorMessage"></ErrorMsg>
	</div>
</template>

<style>
</style>
