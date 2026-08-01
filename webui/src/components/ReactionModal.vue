<script>
export default {
  props: {
    conversationID: {
      type: Number,
      required: true,
    },
    conversationParticipants: {
      type: Array,
      required: true,
    },
    message: {
      type: Object,
      required: true,
    }
  },
  data() {
    return {
      hasUserReacted: false,
      reactions: [],
      emojis: ["🙂", "🙁", "😭", "😍", "😎", "😏"],

      loading: false,
      errorMessage: ""
    }
  },
  methods: {
    async listReactions() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/conversations/${this.conversationID}/messages/${this.message.id}/reactions`, 
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				
				this.reactions = response.data?.map(reaction => ({
					...reaction,
					username: ""
				})) || [];

        this.reactions.forEach(reaction => {
					this.fetchReactionUsername(reaction);
				});
        
        this.setHasUserReacted();
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
    async addReaction(emoji) {
      if(this.hasUserReacted) return;

			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/conversations/${this.conversationID}/messages/${this.message.id}/reactions`,
        {
          emoji: emoji
        }, 
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
				
				let reaction = {
          id: response.data.id,
          userID: response.data.userID,
          username: "",
          conversationID: response.data.conversationID,
          messageID: response.data.messageID,
          emoji: response.data.emoji
        }
        this.fetchReactionUsername(reaction);

        this.reactions.push(reaction);
        this.hasUserReacted = true;
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
    async removeReaction() {
      if(!this.hasUserReacted) return;

			this.loading = true;
			this.errorMessage = "";
			try {
        const currentUserReaction = this.getCurrentUserReaction();
        if(!currentUserReaction) throw "Current user reaction error";
				await this.$axios.delete(`/users/${sessionStorage.getItem("userID")}/conversations/${this.conversationID}/messages/${this.message.id}/reactions/${currentUserReaction.id}`, 
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

        const index = this.reactions.findIndex(reaction => reaction.id === currentUserReaction.id);
        if (index !== -1) {
          this.reactions.splice(index, 1);
        }
        this.hasUserReacted = false;
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
    setHasUserReacted() {
      const currentUserID = sessionStorage.getItem("userID");
      for(let reaction of this.reactions) {
        if(reaction.userID == currentUserID) {
          this.hasUserReacted = true;
          return;
        }
      }
      this.hasUserReacted = false;
    },
    getCurrentUserReaction() {
      const currentUserID = sessionStorage.getItem("userID");
      for(let reaction of this.reactions) {
        if(reaction.userID == currentUserID) {
          return reaction;
        }
      }
      return null;
    },
    fetchReactionUsername(reaction){
      if(reaction.userID == -1) return "";

      for(let participant of this.conversationParticipants) {
        if(participant.id == reaction.userID) {
          reaction.username = participant.username;
          return;
        }
      }
      reaction.username = "";
    },
    closeModal() {
      this.$emit('close-modal');
    }
  },
  mounted() {
    this.listReactions();
  }
}
</script>

<template>
  <div class="modal-overlay" @click.self="closeModal">
    <div class="modal-content">
      <h2>Reactions</h2>
      
      <ul class="reaction-list">
        <li 
          v-for="reaction in reactions" 
          :key="reaction.id" 
        >
          {{ reaction.username }}: {{ reaction.emoji }}
        </li>
      </ul>

      <div class="emoji-box">
        <span
          v-for="(emoji, index) in emojis"
          :key="index"
          class="emoji-item"
          @click="addReaction(emoji)"
        >
          {{ emoji }}
        </span>
      </div>
      <div class="buttons-box">
        <button class="close-button" @click="removeReaction" :disabled="loading || !hasUserReacted">Delete reaction</button>
        <button class="close-button" @click="closeModal" :disabled="loading">Close</button>
      </div>
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
  background-color: var(--color-2-a);
  padding: 24px;
  border-radius: 12px;
  width: 300px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.2);
  text-align: center;
}

.modal-content button:disabled {
	background-color: var(--color-7);
	cursor: not-allowed;
	opacity: 0.8;
}

.reaction-list {
  list-style-type: none;
  padding: 0;
  margin: 20px 0;
  max-height: 150px;
  overflow-y: auto;
}

.reaction-list li {
  padding: 10px;
  margin-bottom: 5px;
  background-color: var(--color-2);
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.reaction-list li:hover {
  background-color: var(--color-5);
}

.emoji-box {
  display: flex;
  flex-direction: row; 
  flex-wrap: wrap;
  justify-content: center;
  gap: 15px;
  margin-bottom: 20px;
  font-size: 12px;
}

.emoji-item {
  font-size: 2.5rem;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.buttons-box {
  display: flex;
  gap: 15px;
}

.buttons-box button {
  padding: 8px 16px;
  background-color: var(--color-6-a);
  color: var(--text-color-4);
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.buttons-box button:not(:disabled):hover {
  background-color: var(--color-6-b);
}
</style>