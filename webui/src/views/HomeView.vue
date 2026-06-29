<script>
export default {
	data() {
		return {
			conversationPreviews: [],
			errorMessage: null,
			loading: false,
		}
	},
	methods: {
		async listChats() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}`, 
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				this.conversationPreviews = response.data || [];
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
		this.listChats();
	}
}
</script>

<template>
	<div class="main-box">
		<h1>Chats</h1>
		<div class="chat-list-box">
			<ul v-if="conversationPreviews.length > 0">
				<li 
					v-for="conversationPreview in conversationPreviews"
						:key="conversationPreview.id"
				>
					<strong>{{ conversationPreview.id }}</strong>
				</li>
			</ul>
			<div v-else>
				No user found
			</div>
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
