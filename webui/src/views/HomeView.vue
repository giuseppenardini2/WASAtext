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
				
				this.conversationPreviews = response.data?.map(conversationPreview => ({
					...conversationPreview,
					localPictureURL: null
				})) || [];

				this.conversationPreviews.forEach(conversationPreview => {
					if (conversationPreview.recipientPictureURI) {
						this.fetchConversationImage(conversationPreview);
					}
				});
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async fetchConversationImage(conversationPreview) {
			if (!conversationPreview.recipientPictureURI) return;

			let response;
			console.log(typeof conversationPreview.type)
			if (conversationPreview.type === 0) {
				this.loading = true;
				this.errorMessage = "";
				try {
					response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/picture`, {
						headers: {
							"Authorization": `Bearer ${sessionStorage.getItem("token")}` 
						},
						params: {
							peerid: conversationPreview.recipientID
						},
						responseType: "blob"
					});
				}
				catch(error) {
					this.errorMessage = error.toString();
				}
				finally {
					this.loading = false;
				}
			}
			else if (conversationPreview.type === 1) {
				this.loading = true;
				this.errorMessage = "";
				try {
					response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/groups/${conversationPreview.recipientID}/picture`, {
						headers: {
							"Authorization": `Bearer ${sessionStorage.getItem("token")}` 
						},
						responseType: "blob"
					});
				}
				catch(error) {
					this.errorMessage = error.toString();
				}
				finally {
					this.loading = false;
				}
			}
			else {
				return;
			}

			if (conversationPreview.localPictureURL) {
				URL.revokeObjectURL(conversationPreview.localPictureURL);
			}
			conversationPreview.localPictureURL = URL.createObjectURL(response.data);
		}
	},
	mounted() {
		this.listChats();
	},
	beforeUnmount() {
		this.conversationPreviews.forEach(conversationPreview => {
			if (conversationPreview.localPictureURL) {
				URL.revokeObjectURL(conversationPreview.localPictureURL);
			}
		});
  },
}
</script>

<template>
	<div class="main-box">
		<h1>Conversations</h1>
		<div class="conversations-list-box">
			<div v-if="conversationPreviews.length > 0">
				<div
					v-for="conversationPreview in conversationPreviews"
						:key="conversationPreview.id"
					class="conversation-raw"
				>
					<div class="conversation-picture">
						<img v-if="conversationPreview.recipientPictureURI === ''" src="@/assets/images/default-user-picture.jpg" alt="User picture">
						<img v-else :src="conversationPreview.localPictureURL" alt="User picture">
					</div>
					<div class="conversation-text">
						<strong>{{ conversationPreview.recipientName }} </strong>
						<strong v-if="conversationPreview.lastMessageID != -1 && !conversationPreview.lastMessagePictureURI">{{ conversationPreview.lastMessageBody }} </strong>
						<strong v-else-if="conversationPreview.lastMessageID != -1 && conversationPreview.lastMessagePictureURI">📷 Photo</strong>
					</div>
					<div v-if="conversationPreview.lastMessageID != -1" class="conversation-date">
						<strong>{{ conversationPreview.lastMessageDate }} </strong>
					</div>
				</div>
			</div>
			<div v-else>
				No convarsation
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

.conversation-row {
	display: flex;
	justify-content: space-between; 
	align-items: center; 
	width: 100%;

	border-bottom: 1px solid #eee;
	padding-bottom: 0.5rem;
}

.conversation-picture {
	display: flex;
	align-items: center;
}

.conversation-picture img {
	width: 50px;
	height: 50px;
	border-radius: 50%;
	display: block;
	object-fit: cover;
}

.conversation-text {
	display: flex;
	flex-direction: column;
	margin-left: 1rem; 
}

.conversation-date {
	margin-left: auto;
}

</style>
