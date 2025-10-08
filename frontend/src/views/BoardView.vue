<script setup>
import ListColumn from '@/components/ListColumn.vue';
import { nextTick, reactive, computed, ref, onMounted } from 'vue';
import { useRoute, RouterLink } from 'vue-router';
import axios from 'axios';
import api from "@/api/api.js"
import CreateTaskOverlay from '@/components/CreateTaskOverlay.vue'
import TaskCard from '@/components/TaskCard.vue'

const route = useRoute();
const boardID = route.params.id;

const show = ref(false)
const showCreateList = ref(false)
const newListTitle = ref('')
const boardTitle = ref('')

const createNewTask = () => {
	return
}
const state = reactive({
	lists: [],
	isLoading: true,
	boardTitle: 'Loading...',
	draggedTask: null,
	draggedFromList: null
})

const listLength = ref(state.lists.length)

// Drag and drop handlers
const handleDragStart = (event, task, listId) => {
	state.draggedTask = task
	state.draggedFromList = listId
	event.dataTransfer.effectAllowed = 'move'
	event.dataTransfer.setData('text/html', event.target.outerHTML)
	event.target.style.opacity = '0.5'
}

const handleDragEnd = (event) => {
	event.target.style.opacity = '1'
	state.draggedTask = null
	state.draggedFromList = null
}

const handleDragOver = (event) => {
	event.preventDefault()
	event.dataTransfer.dropEffect = 'move'
}

const handleDrop = async (event, targetListId) => {
	event.preventDefault()

	if (state.draggedTask && state.draggedFromList !== targetListId) {
		try {
			await axios.put(`/api/v1/tasks/${state.draggedTask.id}/move`, {
				list_id: targetListId
			})

			await fetchLists()
		} catch (err) {
			console.error('Error moving task:', err)
		}
	}
}

const createNewList = async () => {
	if (!newListTitle.value.trim()) return

	try {
		await axios.post(`/api/v1/lists`, {
			board_id: boardID,
			title: newListTitle.value
		})
		await fetchLists()
		newListTitle.value = ''
		showCreateList.value = false
	} catch (err) {
		console.error('Error creating list:', err)
	}
}


const fetchLists = async () => {
	try {
		state.isLoading = true
		const resp = await axios.get(`/api/v1/lists/${boardID}`)
		state.lists = resp.data
		listLength.value = state.lists.length
		state.lists = Object.values(
			resp.data.reduce((acc, item) => {
				if (!acc[item.list_id]) {
					acc[item.list_id] = {
						list_id: item.list_id,
						list_title: item.list_title,
						tasks: []
					};
				}
				const { list_id, list_title, ...task } = item
				acc[item.list_id].tasks.push(task);
				return acc;
			}, {})
		);

		const resp1 = await axios.get(`/api/v1/boards/${boardID}`)
		boardTitle.value = resp1.data
	} catch (err) {
		console.log("Error fetching lists", err)
	} finally {
		state.isLoading = false
	}
}

onMounted(fetchLists)
</script>

<template>
	<div class="min-h-screen gradient-bg">
		<!-- Board Header -->
		<div class="glass-card border-b border-white/20 sticky top-16 z-40">
			<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
				<div class="flex items-center justify-between">
					<div class="flex items-center space-x-4">
						<RouterLink to="/boards" class="text-gray-600 hover:text-indigo-600 transition-colors">
							<i class="pi pi-arrow-left text-xl"></i>
						</RouterLink>
						<div>
							<h1 class="text-2xl font-bold text-gray-800">{{ boardTitle.board_title }}</h1>
							<p class="text-gray-600">{{ state.lists.length }} lists</p>
						</div>
					</div>

					<div class="flex items-center space-x-3">
						<button @click="showCreateList = !showCreateList" class="btn-primary flex items-center space-x-2">
							<i class="pi pi-plus"></i>
							<span>Add List</span>
						</button>
						<button class="btn-secondary flex items-center space-x-2">
							<i class="pi pi-cog"></i>
							<span>Settings</span>
						</button>
					</div>
				</div>
			</div>
		</div>

		<!-- Create List Form -->
		<div v-if="showCreateList" class="glass-card mx-4 mt-4 rounded-2xl p-6">
			<form @submit.prevent="createNewList" class="space-y-4">
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-2">
						<input v-model="newListTitle" type="text" required
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
							placeholder="Enter list title...">
					</label>
				</div>
				<div class="flex space-x-3">
					<button type="submit" class="btn-primary">
						<i class="pi pi-check"></i>
						Create List
					</button>
					<button type="button" @click="showCreateList = false" class="btn-secondary">
						<i class="pi pi-times"></i>
						Cancel
					</button>
				</div>
			</form>
		</div>

		<div class="p-4">
			<div class="max-w-7xl mx-auto">
				<div class="flex space-x-6 overflow-x-auto pb-6">
					<div v-if="listLength === 0" class="red font-bold w-full h-full mt-40 text-center text-3xl opacity-50">No list yet. Create One?</div>
					<div v-else v-for="list in state.lists" :key="list.list_id" class="flex-shrink-0">
						<ListColumn v-if="list" :list="list" :boardID="boardID" />
					</div>
				</div>
			</div>
		</div>

		<!-- Loading State -->
		<div v-if="state.isLoading" class="flex justify-center items-center py-12">
			<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
		</div>
	</div>
</template>


