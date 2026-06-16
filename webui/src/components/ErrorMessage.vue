<script>
export default {
	props: ['message']
}
</script>

<!--<template>
	<div class="alert-box" role="alert">
		<span class="alert-content">{{ message }}</span>
	</div>
</template> -->

<template>
	<div class="alert alert-danger" role="alert">
		{{ message }}
	</div>
</template>

<style scoped>
.alert-box {
	display: flex;
    align-items: center;
	padding: 5px;
	margin: 10px;
	border: 1px solid #c10d0d;
    border-radius: 5px;
	background-color:  #fc6767;
	color: black;
	font-size: 16px;
}

.alert-content {
    flex: 1;
    word-break: break-word;
}
</style>
