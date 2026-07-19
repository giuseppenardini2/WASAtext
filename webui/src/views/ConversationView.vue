<script>
export default {
	data() {
		return {
			id: -1,
			type: -1,
			recipientID: -1,
			recipientName: null,
			recipientPictureURI: "",
			messages: [],

			localRecipientPictureURL: null,

			errorMessage: null,
			loading: false,
		}
	},
	methods: {
		async fecthConversation() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/conversations`, {
					params: {
						recipientid: this.$route.params.conversationid
					},
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				
				this.id = response.data.id;
				this.type = response.data.type;
				this.recipientID = response.data.recipientID;
				this.recipientName = response.data.recipientName;
				this.recipientPictureURI = response.data.recipientPictureURI;

				this.messages = response.data.messages?.map(message => ({
					...message,
					localPictureURL: null
				})) || [];

				this.messages.forEach(message => {
					if (message.pictureURI) {
						this.fetchMessageImage(message);
					}
				});

				this.fetchConversationImage();
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async fetchConversationImage() {
			if (!this.recipientPictureURI) return;

			let response;
			switch(this.type) {
				case 0:
					this.loading = true;
					this.errorMessage = "";
					try {
						response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/picture`, {
							headers: {
								"Authorization": `Bearer ${sessionStorage.getItem("token")}` 
							},
							params: {
								peerid: this.recipientID
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
						response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/groups/${this.recipientID}/picture`, {
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

			if (this.localRecipientPictureURL) {
				URL.revokeObjectURL(this.localRecipientPictureURL);
			}
			this.localRecipientPictureURL = URL.createObjectURL(response.data);
		},
		async fetchMessageImage(message) {
			if(!message.pictureURI) return;

			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/conversations/${this.id}/messages/${message.id}/image`, {
					headers: {
						"Authorization": `Bearer ${sessionStorage.getItem("token")}` 
					},
					responseType: "blob"
				});

				if (message.localPictureURL) {
					URL.revokeObjectURL(message.localPictureURL);
				}
				message.localPictureURL = URL.createObjectURL(response.data);
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		redirectToGroupInfo() {
			this.$router.push(`/groups/${this.id}`);
		}
	},
	mounted() {
		this.fecthConversation();
	},
	beforeUnmount() {
		this.messages.forEach(message => {
			if (message.localPictureURL) {
				URL.revokeObjectURL(message.localPictureURL);
			}
		});
		if(this.localRecipientPictureURL) {
			URL.revokeObjectURL(this.localRecipientPictureURL);
		}
  	}
}
</script>

<template>
	<div class="main-box">
		<div class="conversation-box">
			<div class="recipient-box">
				<div class="recipient-picture">
					<img v-if="recipientPictureURI === '' && type === 0" src="@/assets/images/default-user-picture.jpg" alt="User picture">
					<img v-else-if="recipientPictureURI === '' && type === 1" src="@/assets/images/default-group-picture.jpg" alt="Group picture">
					<img v-else :src="localRecipientPictureURL" alt="Profile picture">
				</div>
				<div class="recipient-username">
					<strong>{{ recipientName }}</strong>
				</div>
				<div v-if="type === 1" class="recipient-button">
					<button type="button" @click="redirectToGroupInfo">Edit</button>
				</div>
			</div>

			<div>

			</div>
			
			<div>
				<ErrorMessage v-if="errorMessage" :message="errorMessage"></ErrorMessage>
			</div>
		</div>
	</div>
</template>

<style scoped>
/* MAIN */
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

/* CONVERSATION */
.conversation-box button:hover {
    background-color: var(--color-3-a);
}

.conversation-box button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  opacity: 0.8;
}

/* RECIPIENT */
.recipient-box {
	display: flex;
	flex-direction: row;
	align-items: center;
	border: 1px solid var(--color-5);
	background-color: var(--color-2-a);
	border-radius: 5px;
	padding: 10px;
	margin: 0.5rem;
}

.recipient-button button {
    width: 60px;
    padding: 10px;
    border: none;
    border-radius: 12px;
	background-color: var(--color-3);
    color: white;
    font-size: 14px;
	cursor: pointer;
}

.recipient-picture {
	display: flex;
	align-items: center;
}

.recipient-picture img {
	width: 50px;
	height: 50px;
	border-radius: 50%;
	display: block;
	object-fit: cover;
}

.recipient-username {
	display: flex;
	flex-direction: column;
	margin-left: 1rem;
	align-self: flex-start;
}

.recipient-button {
	margin-left: auto;
}

/* MESSAGE LIST */

/* MESSAGE WRITING */


</style>
