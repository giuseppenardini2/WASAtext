<script>
export default {
	data() {
		return {
			id: -1,
			name: "",
			pictureURI: null,
			userCreatorID: -1,
			partecipants: [],

			selectedFile: null,
			localPictureURL: null,

			errorMessage: null,
			loading: false,
		}
	},
	methods: {
		async fetchGroup() {
			this.loading = true;
			this.errorMessage = "";
			try {
				const response = await this.$axios.get(`/users/${sessionStorage.getItem("userID")}/groups`,
				{
					params: {
						peerid: this.$route.params.groupid
					},
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				this.id = response.data.id
				this.name = response.data.name
				this.pictureURI = response.data.pictureURI
				this.userCreatorID = response.data.userCreatorID
				this.partecipants = response.data.partecipants

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
				payload.append("image", this.selectedFile); ///users/{userid}/groups/{groupid}/picture

				const response = await this.$axios.put(`/users/${sessionStorage.getItem("userID")}/groups/${this.id}/picture`, payload,
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				this.pictureURI = response.data.imageURI;
				sessionStorage.setItem("pictureURI", response.data.imageURI);

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
			<div class="user-username-bar">
				<input type="text" v-model="name" placeholder="Enter username">
				<button type="submit" :disabled="loading">Change</button>
			</div>
		</form>
		<div class="group">
			
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
</style>
