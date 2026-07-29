<script>
export default {
  data() {
    return {
      conversations: [],
      loading: false,
      errorMessage: ""
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
				
				this.conversations = response.data || [];
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
    selectConversation(selectedConversation) {
      this.$emit('conversation-selected', selectedConversation);
      this.closeModal();
    },
    closeModal() {
      this.$emit('close-modal');
    }
  },
  mounted() {
    this.listConversations();
  }
}
</script>

<template>
  <div class="modal-overlay" @click.self="closeModal">
    <div class="modal-content">
      <h2>Select a conversation</h2>
      
      <ul class="conversation-list">
        <li 
          v-for="conversation in conversations" 
          :key="conversation.id" 
          @click="selectConversation(conversation)"
        >
          {{ conversation.recipientName }}
        </li>
      </ul>

      <button class="close-button" @click="closeModal">Abort</button>
      <div>
				<ErrorMessage v-if="errorMessage" :message="errorMessage"></ErrorMessage>
			</div>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 24px;
  border-radius: 12px;
  width: 300px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.2);
  text-align: center;
}

.conversation-list {
  list-style-type: none;
  padding: 0;
  margin: 20px 0;
}

.conversation-list li {
  padding: 10px;
  margin-bottom: 5px;
  background-color: #f4f4f9;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.conversation-list li:hover {
  background-color: #e0e0f0;
}

.close-button {
  padding: 8px 16px;
  background-color: #ff4d4f;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.close-button:hover {
  background-color: #d9363e;
}
</style>