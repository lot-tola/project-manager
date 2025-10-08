<script setup>
import { RouterLink } from 'vue-router'
import api from "@/api/api.js"
import { ref, onMounted } from 'vue'

const props = defineProps({
	board: Object,
})
const deleteResp = ref("")
const succeedDeleteMsg = ref("")
const handleDelete = async () => {
	const ok = confirm("Are you sure you want to delete this? This action can not be undone.")
	if (!ok){
		return 
	}
		const resp = await api.delete(`/api/v1/boards/${props.board.id}`)
}

const iso = props.board.created_at
const date = new Date(iso)
const formatted = date.toLocaleString("en-US", {
	dateStyle: "long",
	timeStyle: "medium",
	timeZone: "UTC",
})
onMounted(() => {
})
</script>

<template>
	<div class="glass-card p-6 rounded-2xl hover:-translate-y-1 hover:shadow-xl transition-all duration-300 group cursor-pointer h-70 flex flex-col justify-between">
		<RouterLink :to="`/boards/${board.id}`" class="block h-full">
			<div class="h-full flex flex-col justify-between">
				<!-- Board Header -->
				<div class="flex items-start justify-between mb-4">
					<div
						class="w-12 h-12 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-xl flex items-center justify-center">
						<i class="pi pi-kanban text-white text-lg"></i>
					</div>
					<div class="opacity-0 group-hover:opacity-100 transition-opacity duration-200">
						<i class="pi pi-arrow-right text-indigo-600 text-lg"></i>
					</div>
				</div>

				<!-- Board Title -->
				<div class="mb-4">
					<h3
						class="text-xl font-semibold text-gray-800 mb-2 group-hover:text-indigo-600 transition-colors duration-200">
						{{ props.board.board_title }}
					</h3>
					<p class="text-sm text-gray-500">
						{{ formatted }}
					</p>
				</div>
			</div>
				</RouterLink>

				<!-- Board Stats -->
				<div class="flex items-center justify-between text-sm text-gray-500">
					<!-- <div class="flex items-center space-x-4"> -->
					<!-- 	<span class="flex items-center space-x-1"> -->
					<!-- 		<i class="pi pi-list"></i> -->
					<!-- 		<span>0 lists</span> -->
					<!-- 	</span> -->
					<!-- 	<span class="flex items-center space-x-1"> -->
					<!-- 		<i class="pi pi-check-circle"></i> -->
					<!-- 		<span>0 tasks</span> -->
					<!-- 	</span> -->
					<!-- </div> -->
					<div class="opacity-0 group-hover:opacity-100 transition-opacity duration-200">
						<span class="text-indigo-600 font-medium">Open →</span>
					</div>
			<form @submit="handleDelete">

				<button class="btn btn-error " type="submit">X</button>
			</form>
				</div>
</div>
</template>


