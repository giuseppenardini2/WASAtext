<script>
export default {  
  data() {
    return {
      username: "",
      users: [],
      isLoading: false,
      error: null,
    };
  },
  methods: {
	async fetchUsers() {
    	this.loading = true;
		this.errorMessage = "";
		try {
			const response = await this.$axios.get('/users', {
				params: {
					username: this.username
				}
			});
			this.users = response.data || [];
		}
		catch(error) {
			this.errorMessage = error.toString();
		}
		finally {
			this.loading = false;
		}
    },
    selectUser(user) {
		this.$emit('user-selected', user);
		
		this.username = "";
		this.users = [];
    }
  }
};
</script>

<template>
  <div class="search-bar-container">
    <input
      type="text"
      v-model="username"
      @input="fetchUsers"
      placeholder="Search username..."
      class="search-input"
    />
    
    <div v-if="isLoading" class="loading-indicator">Loading...</div>
    <div v-if="error" class="error-message">{{ error }}</div>

    <ul v-if="users.length > 0 && !isLoading" class="results-list">
      <li
        v-for="user in users"
        :key="user.id"
        @click="selectUser(user)"
        class="result-item"
      >
        {{ user.username }}
      </li>
    </ul>
  </div>
</template>

<style scoped>
.search-bar-container {
  position: relative;
  width: 300px;
  font-family: sans-serif;
}

.search-input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.results-list {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  margin: 4px 0 0;
  padding: 0;
  list-style: none;
  background: white;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  z-index: 1000;
  max-height: 200px;
  overflow-y: auto;
}

.result-item {
  padding: 10px;
  cursor: pointer;
  border-bottom: 1px solid #eee;
}

.result-item:last-child {
  border-bottom: none;
}

.result-item:hover {
  background-color: #f5f5f5;
}

.loading-indicator, .error-message {
  font-size: 0.85em;
  color: #666;
  margin-top: 5px;
}
.error-message {
  color: #d9534f;
}
</style>