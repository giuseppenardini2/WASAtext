<script>
export default {
	data() {
		return {
			id: -1,
			name: "",
			pictureURI: null,
			userCreatorID: -1,
			participants: [],

			selectedFile: null,
			localPictureURL: null,

			selectedUser: null,

			errorMessage: null,
			loading: false,
		}
	},
	methods: {
		async fetchGroup() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/groups`, {
					params: {
						peerid: this.$route.params.groupid
					},
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				this.id = response.data.id
				this.name = response.data.name
				this.pictureURI = response.data.pictureURI
				this.userCreatorID = response.data.userCreatorID
				this.participants = response.data.participants

				this.getGroupPicture();
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
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
		async changeGroupPicture() {
			if (!this.selectedFile) return;
			
			this.loading = true;
			this.errorMessage = "";
			try {
				const payload = new FormData();
				payload.append("image", this.selectedFile);

				const response = await this.$axios.put(`/users/${sessionStorage.getItem("userID")}/groups/${this.id}/picture`, payload,
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				this.pictureURI = response.data.imageURI;

				this.getGroupPicture();
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async getGroupPicture() {
			if (!this.pictureURI) return;

			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/groups/${this.id}/picture`, {
					headers: { 
						"Authorization": `Bearer ${sessionStorage.getItem("token")}` 
					},
					responseType: "blob"
				});

				if (this.localPictureURL) {
          			URL.revokeObjectURL(this.localPictureURL);
        		}
				this.localPictureURL = URL.createObjectURL(response.data);
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async changeGroupName() {
			this.loading = true;
			this.errorMessage = "";
			try {
				if (this.name.length < 3 || this.name.length > 20) throw "Invalid group name"

				const response = await this.$axios.put(`/users/${sessionStorage.getItem("userID")}/groups/${this.id}/name`, {
					name: this.name
				},
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		handleUserSelection(user) {
			this.selectedUser = user;
		},
		async addParticipant() {
			this.loading = true;
			this.errorMessage = "";
			try {
				if (!this.selectedUser) throw "Select a user"

				const response = await this.$axios.put(`/users/${sessionStorage.getItem("userID")}/groups/${this.id}/participants`,
				{
					id: this.selectedUser.id,
					username: this.selectedUser.username,
					pictureURI: this.selectedUser.pictureURI
				}, 
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` }
				});

				this.participants.push(this.selectedUser);
				this.selectedUser = null;
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		},
		async leaveGroup() {
			this.loading = true;
			this.errorMessage = "";
			try {
				await this.$axios.delete(`/users/${sessionStorage.getItem("userID")}/groups/${this.id}/participants`,
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
		}
	},
	mounted() {
		this.fetchGroup();
	},
	beforeUnmount() {
		if (this.localPictureURL) {
      		URL.revokeObjectURL(this.localPictureURL);
    	}
  	}
}
</script>

<template>
	<div class="main-box">
		<h1>Group</h1>
		<div class="profile-box">
			<form @submit.prevent="changeGroupPicture">
				<div class="group-picture">
					<img v-if="pictureURI === ''" src="@/assets/images/default-group-picture.jpg" alt="Group picture">
					<img v-else :src="localPictureURL" alt="Group picture">
				</div>
				<div>
					<input type="file" accept="image/*" @change="handleFileSelection">
					<button type="submit" :disabled="!selectedFile || loading">Change</button>
				</div>
			</form>
			<form @submit.prevent="changeGroupName">
				<div class="group-name-bar">
					<input type="text" v-model="name" placeholder="Enter group name">
					<button type="submit" :disabled="loading">Change</button>
				</div>
			</form>
		</div>
		<div>
			<div v-if="participants.length > 0">
				<strong>Participants</strong>
				<div
				v-for="participant in participants"
					:key="participant.id"
				>
				{{ participant.username }}
				</div>
			</div>
			<div v-else>
				No participants
			</div>
			<div>
				<div class="add-participant-box">
					<UserSearchBar @user-selected="handleUserSelection" />
					<div v-if="selectedUser">
						<p>Hai selezionato: <strong>{{ selectedUser.username }}</strong> (ID: {{ selectedUser.id }})</p>
					</div>
					<button type="button" @click="addParticipant" :disabled="!selectedUser || loading">Add participant</button>
				</div>
			</div>
		</div>
		<div>
			<button type="button" @click="leaveGroup" :disabled="loading">Leave</button>
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

.main-box button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  opacity: 0.8;
}

.profile-box {
	margin: 2px;
	padding: 10px;
	display: flex;
	flex-direction: column;
	align-items: center;
}

.profile-box form {
	margin-bottom: 20px;
}

.profile-box button {
    width: 100px;
	padding: 10px;
    border: none;
    border-radius: 8px;
	background-color: var(--color-3);
    color: white;
    font-size: 16px;
	cursor: pointer;
}

.profile-box button:hover {
    background-color: var(--color-3-a);
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

.group-picture {
	display: flex;
	flex-direction: column;
	gap: 5px;
	margin-bottom: 10px;
	align-items: center;
}

.group-picture img {
	width: 150px;
	height: 150px;
	object-fit: cover;
	border-radius: 50%;
	box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.participants-box {
	display: flex;
	flex-direction: row;
	gap: 5px;
	margin-bottom: 10px;
	align-items: center;
}

.add-participants-box {
	display: flex;
	flex-direction: row;
	gap: 5px;
	margin-bottom: 10px;
	align-items: center;
}
</style>
