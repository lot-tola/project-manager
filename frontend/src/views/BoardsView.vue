<script setup>
import BoardCard from '../components/BoardCard.vue'
import api from "@/api/api.js"
import { onMounted, ref } from 'vue'

const boards = ref([])
const title = ref("")
const error = ref(null)
const showCreateForm = ref(false)

const handleSubmit = async () => {
	if (!title.value.trim()) return

	try {
		await api.post("/v1/boards", {
			title: title.value
		})
		await fetchBoards()
		showCreateForm.value = false
	} catch (err) {
		console.log("Error creating boards: ", err)
		error.value = "Failed to create board"
	} finally {
		title.value = ''
	}
}

const fetchBoards = async () => {
	try {
		error.value = null
		const resp = await api.get("/v1/boards")
		boards.value = resp.data.filter(item => typeof item === 'object')
	} catch (err) {
		console.error("Error fetching boards:", err)
		error.value = err.message
		boards.value = []
	} 
}

onMounted(fetchBoards)

</script>

<template>
	<div class="min-h-screen gradient-bg py-8">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<!-- Header -->
			<div class="mb-8">
				<h1 class="text-4xl font-bold text-white mb-4">Your Boards</h1>
				<p class="text-white/80 text-lg">Organize your projects and tasks efficiently</p>
			</div>

			<!-- Create Board Form -->
			<div class="glass-card p-6 rounded-2xl mb-8">
				<div class="flex items-center justify-between mb-4">
					<h2 class="text-xl font-semibold text-gray-800">Create New Board</h2>
					<button @click="showCreateForm = !showCreateForm" class="btn-primary flex items-center space-x-2">
						<i :class="showCreateForm ? 'pi pi-times' : 'pi pi-plus'"></i>
						<span>{{ showCreateForm ? 'Cancel' : 'New Board' }}</span>
					</button>
				</div>

				<form v-if="showCreateForm" @submit.prevent="handleSubmit" class="text-gray-800 space-y-4">
					<div>
						<label for="newboard" class="block text-sm font-medium text-gray-700 mb-2">
						<input v-model="title" type="text" id="newboard" name="title" required
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all duration-200"
							placeholder="Enter board title...">
						</label>
					</div>
					<div class="flex space-x-3">
						<button type="submit" class="btn-primary">
							<i class="pi pi-check"></i>
							Create Board
						</button>
						<button type="button" @click="showCreateForm = false" class="btn-secondary">
							<i class="pi pi-times"></i>
							Cancel
						</button>
					</div>
				</form>
			</div>

			<!-- Error Message -->
			<div v-if="error" class="glass-card p-4 rounded-lg mb-6 border-l-4 border-red-500 bg-red-50">
				<div class="flex items-center">
					<i class="pi pi-exclamation-triangle text-red-500 mr-3"></i>
					<span class="text-red-700">{{ error }}</span>
				</div>
			</div>

			<!-- Boards Grid -->
			<div v-else-if="boards.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
				<BoardCard v-for="board in boards" :key="board.id" :board="board" class="fade-in" />
			</div>

			<!-- Empty State -->
			<div v-else class="text-center py-12">
				<div class="glass-card p-12 rounded-2xl">
					<div class="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-6">
						<i class="pi pi-th-large text-3xl text-gray-400"></i>
					</div>
					<h3 class="text-2xl font-semibold text-gray-800 mb-4">No boards yet</h3>
					<p class="text-gray-600 mb-6">Create your first board to start organizing your projects</p>
					<button @click="showCreateForm = true" class="btn-primary">
						<i class="pi pi-plus"></i>
						Create Your First Board
					</button>
				</div>
			</div>
		</div>
	</div>
</template>


