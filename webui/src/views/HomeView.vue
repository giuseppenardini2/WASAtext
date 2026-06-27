<script>
export default {
	data() {
		return {
			chatList: [],
			errorMessage: null,
			loading: false,
		}
	},
	methods: {
		listChats() {
			this.loading = true;
			this.loading = false;
		},
	}
}
</script>

<template>
	<div class="main-box">
		<h1>Chats</h1>
		<div class="chat-list-box">
			
		</div>
		<div>
			<ErrorMessage v-if="errorMessage" :message="errorMessage"></ErrorMessage>
		</div>
	</div>
</template>

<style>
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
</style>
