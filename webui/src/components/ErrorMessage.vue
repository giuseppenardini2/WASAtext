<script>
export default {
	props: ['message']
}
</script>

<template>
	<div class="alert-box">
		{{ message }}
	</div>
</template>

<style scoped>
.alert-box {
	display: flex;
	font-size: 12px;
	border-style: solid;
	border-width: 2px;
	border-color: #c10d0d;
	background-color:  #fc6767;

}
</style>
