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
				let response = await this.$axios.get("/");
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

		</div>
		<ErrorMsg v-if="errorMessage" :msg="errorMessage"></ErrorMsg>
	</div>
</template>

<style>
</style>
