<script>
export default {
	data() {
		return {
			id: -1,
			name: null,
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
			}
			catch(error) {
				this.errorMessage = error.toString();
			}
			finally {
				this.loading = false;
			}
		}
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
		async changePicture() {
			if (!this.selectedFile) return;
			
			this.loading = true;
			this.errorMessage = "";
			try {
				const payload = new FormData();
				payload.append("image", this.selectedFile);

				const response = await this.$axios.put(`/users/${sessionStorage.getItem("userID")}/picture`, payload,
				{
					headers: { "Authorization": `Bearer ${sessionStorage.getItem("token")}` } 
				});

				this.pictureURI = response.data.imageURI;
				sessionStorage.setItem("pictureURI", response.data.imageURI);

				this.getUserPicture();
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
		
  	}
}
</script>

<template>
	<div class="main-box">
		<h1>Group</h1>
		<form @submit.prevent="changePicture">
			<div class="group-picture">
				<img v-if="pictureURI === ''" src="@/assets/images/default-group-picture.jpg" alt="Group picture">
				<img v-else :src="localPictureURL" alt="Group picture">
			</div>
			<div>
				<input type="file" accept="image/*" @change="handleFileSelection">
				<button type="submit" :disabled="!selectedFile || loading">Change</button>
			</div>
		</form>
		<form @submit.prevent="changeUsername">
			<div class="user-username-bar">
				<input type="text" v-model="username" placeholder="Enter username">
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
