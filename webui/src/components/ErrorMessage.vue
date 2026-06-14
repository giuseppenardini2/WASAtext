<script>
export default {
	props: ['message']
}
</script>

<template>
	<div class="alert-box">
		<span class="alert-content">{{ message }}</span>
	</div>
</template>

<style scoped>
.alert-box {
	display: flex;
    align-items: center;

	border: 1px solid #c10d0d;
    border-radius: 2px;
	
	background-color:  #fc6767;
	color: black;

	font-size: 12px;
}

.alert-content {
    flex: 1;
    word-break: break-word;
}
</style>
