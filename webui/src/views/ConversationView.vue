<script>
export default {
	data() {
		return {
			id: -1,
			type: -1,
			recipientID: -1,
			recipientName: null,
			recipientPictureURI: "",
			participants: [],
			messages: [],

			localRecipientPictureURL: null,
			newMessageText: "",
			
			selectedFile: null,
			newMessagePictureURI: "",

			selectedReplyMessage: null,

			activeConversationMessageID: null,
			activeReactionMessageID: null,

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
				this.participants = response.data.participants;

				this.messages = response.data.messages?.map(message => ({
					...message,
					localPictureURL: null,
					sentDate: this.formatMessageDate(message.sentDate),
					replyMessagePreview: ""
				})) || [];

				this.messages.forEach(message => {
					if (message.pictureURI) {
						this.fetchMessageImage(message);
					}
					if (message.replyMessageID != -1) {
						this.fetchReplyMessagePreview(message);
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
		},
		handleFileSelection(event) {
			const file = event.target.files[0];
     		if (!file) return;
			if (file.type !== "image/jpeg" && file.type !== "image/png") {
                this.errorMessage = "File type is wrong";
                return
            }
            if (file.size > 3145728) {
                this.errorMessage = "Max file size is 3 MB";
                return
            }
			this.selectedFile = file;
		},
		handleMessageSelection(message){
			this.selectedReplyMessage = message;
		},
		async sendMessage() {
			if(!this.newMessageText && !this.selectedFile) return;

			this.loading = true;
			this.errorMessage = "";
			try {
				await this.uploadMessageImage();

				let response;
				if(this.selectedReplyMessage) {
					response = await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/conversations/${this.id}/messages`, 
					{
						body: this.newMessageText,
						pictureURI: this.newMessagePictureURI
					},
					{	
						params: {
							replymessageid: this.selectedReplyMessage.id
						},
						headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
					});
				}
				else {
					response = await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/conversations/${this.id}/messages`, 
					{
						body: this.newMessageText,
						pictureURI: this.newMessagePictureURI
					},
					{
						headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
					});
				}

				let message = {
					id: response.data.id,
					conversationID: response.data.conversationID,
					body: response.data.body,
					sentDate: this.formatMessageDate(response.data.sentDate),
					pictureURI: response.data.pictureURI,
					localPictureURL: "",
					status: response.data.status,
					senderID: response.data.senderID,
					replyMessageID: response.data.replyMessageID,
					replyMessageConversationID: response.data.replyMessageConversationID,
					replyMessagePreview: "",
					forwardedMessageID: response.data.forwardedMessageID,
					forwardedMessageConversationID: response.data.forwardedMessageConversationID
				};
				await this.fetchMessageImage(message);
				this.fetchReplyMessagePreview(message);
				this.messages.push(message);

				this.newMessageText = "";
				this.newMessagePictureURI = "";
				this.selectedFile = null;
				this.selectedReplyMessage = null;
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async uploadMessageImage() {
			if(!this.selectedFile) {
				this.newMessagePictureURI="";
				return;
			}

			this.loading = true;
			this.errorMessage = "";
			try {
				const payload = new FormData();
				payload.append("image", this.selectedFile);

				const response = await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/conversations/${this.id}/uploads/images`, payload,
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				this.newMessagePictureURI = response.data.imageURI;
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async deleteMessage(message) {
			this.loading = true;
			this.errorMessage = "";
			try {
				await this.$axios.delete(`/users/${sessionStorage.getItem("userID")}/conversations/${this.id}/messages/${message.id}`,
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				const messageIndex = this.messages.indexOf(message);
				if(messageIndex != -1) {
					this.messages.splice(messageIndex, 1);
				}
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async forwardMessage(conversationPreview, message) {
			this.loading = true;
			this.errorMessage = "";
			try {
				await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/conversations/${this.id}/messages/${message.id}`, {},
				{	
					params: {
						recipientid: conversationPreview.id
					},
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				if(conversationPreview.id == this.id) {
					this.fecthConversation();
				}
				this.$router.push(`/conversations/${conversationPreview.id}`);
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
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
		isMessageMine(message) {
			return message.senderID == sessionStorage.getItem("userID");
		},
		getMessageRecipientName(message) {
			for(let user of this.participants) {
				if (message.senderID == user.id) {
					return user.username;
				}
			}
			return "";
		},
		fetchReplyMessagePreview(message) {
			if(!message || message.replyMessageID === -1) return "";

			let replyMessage;
			for(let m of this.messages){
				if(m.id === message.replyMessageID) {
					replyMessage = m;
					break;
				}
			}

			if(!replyMessage) return;

			if(replyMessage.pictureURI) {
				message.replyMessagePreview = "📷 Photo";
				return;
			}

			let replyMessagePreview = replyMessage.body.trim();
			if(replyMessagePreview.length > 10)
				replyMessagePreview = replyMessagePreview.substring(0, 9) + "...";
			message.replyMessagePreview = replyMessagePreview;
		},
		openConversationModal(message) {
			this.activeConversationMessageID = message.id;
		},
		closeConversationModal() {
			this.activeConversationMessageID = null;
		},
		openReactionModal(message) {
			this.activeReactionMessageID = message.id;
		},
		closeReactionModal() {
			this.activeReactionMessageID = null;
		}
	},
	mounted() {
		this.fecthConversation();
		setInterval(this.fecthConversation, 30000); // Auto update conversation view every 30 seconds
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

			<div class="messages-box">
				<div v-if="messages.length > 0">
					<div
						v-for="message in messages"
							:key="message.id"
						class="message-row" 
						:class="{'message-row-mine': isMessageMine(message), 'message-row-other': !isMessageMine(message)}"
					>
						<MessageMenu 
                            v-if="isMessageMine(message)" 
							:isMine="true"
                            @delete="deleteMessage(message)" 
                            @reply="handleMessageSelection(message)"
							@react="openReactionModal(message)" 
                            @forward="openConversationModal(message)"
                        />
						<div class="message-bubble">
							<div v-if="type === 1" class="message-recipient-name">
								<label>{{getMessageRecipientName(message)}}</label>
							</div>
							<div v-if="message.forwardedMessageID != -1" class="forwarded-message">
								<label>Forwarded</label>
							</div>
							<div v-if="message.replyMessageID != -1 && message.forwardedMessageID === -1" class="reply-message-preview">
								{{ message.replyMessagePreview }}
							</div>
							<div>
								<img v-if="message.pictureURI" :src="message.localPictureURL" class="message-image">
								{{ message.body }}
							</div>
							<div class="message-footer">
								<label class="message-date">{{ message.sentDate }}</label>
								<label v-if="isMessageMine(message) && message.status === 0" class="message-read-checkmark">⏲️</label>
								<label v-if="isMessageMine(message) && message.status === 1" class="message-read-checkmark">✔️</label>
								<label v-if="isMessageMine(message) && message.status === 2" class="message-read-checkmark">✔️✔️</label>
							</div>
							<ConversationModal 
								v-if="activeConversationMessageID === message.id"
								@close-modal="closeConversationModal"
								@conversation-selected="forwardMessage($event, message)"
							/>
							<ReactionModal 
								v-if="activeReactionMessageID === message.id"
								:conversationID="id"
								:conversationParticipants="participants"
								:message="message"
								@close-modal="closeReactionModal"
							/>
						</div>
						<MessageMenu 
                            v-if="!isMessageMine(message)"
							:isMine="false"
                            @reply="handleMessageSelection(message)" 
                            @forward="openConversationModal(message)" 
							@react="openReactionModal(message)"
                        />
					</div>
				</div>
				<div v-else>
					No messages
				</div>
			</div>

			<div class="writing-box">
				<div v-if="selectedReplyMessage" class="message-reply-warning">
					You are replying to a message
				</div>
				<form @submit.prevent="sendMessage">
						<input type="file" accept="image/*" @change="handleFileSelection">
						<textarea v-model="newMessageText" placeholder="Insert message" rows="4"></textarea>
						<button type="submit" :disabled="(!newMessageText && !selectedFile) || loading">Send</button>
				</form>
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
  	height: 100vh;
  	display: flex;
	flex-direction: column;
	gap: 20px;
	padding: 10px;
}

.main-box h1 {
	color: var(--text-color-1);
}

/* CONVERSATION */
.conversation-box {
	display: flex;
    flex-direction: column;
    height: 100%;
}
.conversation-box button:not(:disabled):hover {
    background-color: var(--color-3-a);
}

.conversation-box button:disabled {
	background-color: var(--color-7);
	cursor: not-allowed;
	opacity: 0.8;
}

/* RECIPIENT */
.recipient-box {
	display: flex;
	flex-direction: row;
	align-items: center;
	flex-shrink: 0;
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
    color: var(--text-color-4);
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
.messages-box {
	display: flex;
  	flex-direction: column;
	flex: 1;
  	gap: 12px;
  	padding: 20px;
  	overflow-y: auto;
}

.message-row {
	display: flex;
	width: 100%;
	align-items: center;
	gap: 8px;
}

.message-row-mine {
	justify-content: flex-end;
}

.message-row-other {
	justify-content: flex-start;
}

.message-row-mine .message-bubble {
	background-color: #dcf8c6;
	color: var(--text-color-2);
}

.message-row-other .message-bubble {
	background-color: #ffffff;
	color: var(--text-color-2);
	border: 1px solid #ddd;
}

.message-bubble {
	max-width: 70%;
	height: fit-content;
	padding: 10px 15px;
	border-radius: 8px;
	word-wrap: break-word;
	line-height: 1.4;
}

.message-bubble .message-recipient-name {
	font-size: 11px;
	font-weight: bold;
	user-select: none;
}

.message-bubble .forwarded-message {
	font-size: 10px;
	font-style: italic;
	user-select: none;
}

.message-bubble .reply-message-preview {
	margin-bottom: 2px;
	padding: 3px;
	border-radius: 6px;
	background-color: #bfbbbb;
	color: #565656;
	font-size: 12px;
}

.message-bubble .message-image {
	display: block;
	max-height: 300px;
}

.message-bubble .message-footer {
	display: flex;
    margin-top: 4px;
}

.message-bubble .message-footer .message-date {
  	font-size: 9px;
	font-style: italic;
	user-select: none;
}

.message-bubble .message-footer .message-read-checkmark {
	font-size: 9px;
	margin-left: auto;
	user-select: none;
}

.message-reply-warning {
	margin-bottom: 5px;
	padding: 10px;
	border-radius: 6px;
	background-color: #bfbbbb;
	color: #565656;
}

/* MESSAGE WRITING */
.writing-box {
    flex-shrink: 0;
    padding: 10px 0;
    width: 100%;
    box-sizing: border-box;
}

.writing-box form {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 12px;
    width: 100%;
}

.writing-box input[type="file"] {
    max-width: 200px;
}

.writing-box textarea {
	flex: 1;
    height: 60px;
	resize: none; 
	overflow-y: auto; 
	padding: 12px;
	box-sizing: border-box;
	border: 1px solid #ccc;
	border-radius: 4px;
	font-family: inherit;
	line-height: 1.5;
}

.writing-box button {
    width: 60px;
    padding: 10px;
    border: none;
    border-radius: 12px;
	background-color: var(--color-3);
    color: var(--text-color-4);
    font-size: 14px;
	cursor: pointer;
	flex-shrink: 0;
}
</style>
