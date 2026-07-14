<script>
export default {
	data() {
		return {
			name: "",
			users: [],
			usersToAdd: [],
			errorMessage: null,
			loading: false,
		}
	},
	methods: {
		async createGroup() {
			this.loading = true;
			this.errorMessage = "";
			try {
				if (this.name.length < 3 || this.name.length > 16) throw "Invalid group name"

				const response = await this.$axios.post(`/users/${sessionStorage.getItem("userID")}/groups`, {
					name: this.name,
					partecipants: this.usersToAdd
				},
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` }
				});
				this.$router.push("/home");
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async listUsers() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get('/users', {
					params: {
						username: ""
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
		}
	},
	mounted() {
		this.listUsers();
	}
}
</script>

<template>
	<div class="main-box">
		<h1>New group</h1>
		<form @submit.prevent="createGroup" class="group-box">
			<div class="group-name-bar">
				<input type="text" v-model="name" placeholder="Enter group name">
			</div>
			<div class="group-users-box">
				<p>Users to add</p>
				<div v-if="!loading && users.length > 0">
					<div 
						v-for="user in users"
							:key="user.id" 
						class="group-user-item"
					>
						<label>
							<input type="checkbox" :value="user" v-model="usersToAdd">
							{{ user.username }}
						</label>
					</div>
				</div>
				<div v-else>
					No user found
				</div>
			</div>
			<div class="group-buttons">
				<button type="submit" :disabled="loading">Create</button>
			</div>
		</form>
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

.group-box button {
    width: 100px;
    padding: 10px;
    border: none;
    border-radius: 8px;
	background-color: var(--color-3);
    color: white;
    font-size: 14px;
	cursor: pointer;
}

.group-box button:hover {
    background-color: var(--color-3-a);
}

.group-box button:disabled {
	background-color: #6c757d;
	cursor: not-allowed;
	opacity: 0.8;
}

.group-name-bar {
	display: flex;
	flex-direction: row;
	gap: 5px;
	margin-bottom: 10px;
	align-items: center;
}

.group-name-bar input {
    width: 300px;
	padding: 10px;
    border: 1px solid #ccc;
    border-radius: 8px;
	background-color: var(--color-2-a);
    font-size: 16px;
	color: var(--text-color-1)
}

.group-users-box {
	margin: 2px;
	padding: 10px;
	background-color: var(--color-5);
	border-radius: 5px;
	display: flex;
	flex-direction: column;
}

.group-user-item {
	display: flex;
	gap: 10px;
	margin: 0.5rem;
}

.group-buttons {
	display: flex;
	gap: 10px;
	margin-top: 10px;
}

</style>
