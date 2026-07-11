<script>
export default {
	data() {
		return {
			selectedFile: null,
			errorMessage: null,
			loading: false,
		}
	},
	methods: {
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
		
	},
	beforeUnmount() {
		
  	}
}
</script>

<template>
	<div class="main-box">
		<h1>Group</h1>
		<form @submit.prevent="changePicture">
			<div class="user-picture">
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
    color: white;
    font-size: 14px;
	cursor: pointer;
}

.options-box button:hover {
    background-color: var(--color-3-a);
}

.options-box button:disabled {
  background-color: #6c757d;
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
	margin-left: auto;
}

</style>
