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
		async listConversations() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}`, 
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				
				this.conversationPreviews = response.data?.map(conversationPreview => ({
					...conversationPreview,
					localPictureURL: null,
					lastMessageSentDate: this.formatMessageDate(conversationPreview.lastMessageSentDate)
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
			switch(conversationPreview.type) {
				case 0:
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
					break;
				case 1:
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
					break;
				default:
					return;
					break;
			}

			if (conversationPreview.localPictureURL) {
				URL.revokeObjectURL(conversationPreview.localPictureURL);
			}
			conversationPreview.localPictureURL = URL.createObjectURL(response.data);
		},
		formatMessageDate(dateString) {
			if(!dateString) return "";

			const date = new Date(dateString);

			return new Intl.DateTimeFormat('en-US', {
				day: "numeric",
				month: "short",
				hour: "2-digit",
				minute: "2-digit"
			}).format(date);
		},
		redirectToNewChatView() {
			this.$router.push("/users");
		},
		redirectToNewGroupView() {
			this.$router.push("/groups");
		},
		redirectToSelectedChat(conversationID) {
			this.$router.push(`/conversations/${conversationID}`);
		}
	},
	mounted() {
		this.listConversations();
		setInterval(this.listConversations, 30000); // Auto update home view every 30 seconds
	},
	beforeUnmount() {
		this.conversationPreviews.forEach(conversationPreview => {
			if (conversationPreview.localPictureURL) {
				URL.revokeObjectURL(conversationPreview.localPictureURL);
			}
		});
  	}
}
</script>

<template>
	<div class="main-box">
		<h1>Conversations</h1>
		<div class="options-box">
			<button type="button" @click="redirectToNewChatView" :disabled="loading">New chat</button>
			<button type="button" @click="redirectToNewGroupView" :disabled="loading">New group</button>
		</div>
		<div class="conversations-list-box">
			<div v-if="conversationPreviews.length > 0">
				<div
					v-for="conversationPreview in conversationPreviews"
						:key="conversationPreview.id"
					class="conversation-row" @click="redirectToSelectedChat(conversationPreview.id)"
				>
					<div class="conversation-picture">
						<img v-if="conversationPreview.recipientPictureURI === '' && conversationPreview.type === 0" src="@/assets/images/default-user-picture.jpg" alt="User picture">
						<img v-else-if="conversationPreview.recipientPictureURI === '' && conversationPreview.type === 1" src="@/assets/images/default-group-picture.jpg" alt="Group picture">
						<img v-else :src="conversationPreview.localPictureURL" alt="Profile picture">
					</div>
					<div class="conversation-text">
						<strong>{{ conversationPreview.recipientName }} </strong>
						<p v-if="conversationPreview.lastMessageID != -1 && !conversationPreview.lastMessagePictureURI">{{ conversationPreview.lastMessageBody }} </p>
						<p v-else-if="conversationPreview.lastMessageID != -1 && conversationPreview.lastMessagePictureURI">📷 Photo</p>
					</div>
					<div v-if="conversationPreview.lastMessageID != -1" class="conversation-date">
						<i>{{ conversationPreview.lastMessageSentDate }}</i>
					</div>
				</div>
			</div>
			<div v-else>
				No conversation
			</div>
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

.options-box {
	display: flex;
	gap: 10px;
}

.options-box button {
    width: 100px;
    padding: 10px;
    border: none;
    border-radius: 8px;
	background-color: var(--color-3);
    color: var(--text-color-4);
    font-size: 14px;
	cursor: pointer;
}

.options-box button:not(:disabled):hover {
    background-color: var(--color-3-a);
}

.options-box button:disabled {
  background-color: var(--color-7);
  cursor: not-allowed;
  opacity: 0.8;
}

.conversation-row {
	display: flex; 
	align-items: center;
	border: 1px solid var(--color-5);
	background-color: var(--color-2-a);
	border-radius: 5px;
	padding: 0.5rem;
	margin: 0.5rem;
}

.conversation-row:hover {
    background-color: var(--color-5);
    cursor: pointer;
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
	align-self: flex-start;
}

.conversation-date {
	font-size: 12px;
	margin-left: auto;
}

</style>
