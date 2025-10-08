<script setup>
import { ref } from 'vue'
import axios from 'axios'
import api from "@/api/api.js"

const props = defineProps({
	list_id: {
		type: [String, Number],
		required: true
	}
})

const emit = defineEmits(['close', 'created'])

const task_title = ref('')
const description = ref('')
const due_date = ref('') 
const assigned_to = ref('')
const status = ref('todo')
const labels = ref('') 
const loading = ref(false)
const error = ref('')

async function handleSubmit() {
	if (!task_title.value.trim()) return
	loading.value = true
	error.value = ''
	try {
		await axios.post('/api/v1/tasks', {
			list_id: String(props.list_id),
			task_title: task_title.value,
			description: description.value,
			due_date: due_date.value,
			assigned_to: assigned_to.value,
			status: status.value,
			labels: labels.value
		})
		emit('created')
		emit('close')
	} catch (e) {
		error.value = 'Failed to create task'
		console.error('Create task error:', e)
	} finally {
		loading.value = false
	}
}

function close() {
	emit('close')
}
</script>

<template>
	<!-- Backdrop -->
	<div class="fixed inset-0 bg-black/40 flex items-center justify-center z-50 text-black">
		<!-- Modal Card -->
		<div class="glass-card rounded-2xl w-full max-w-lg p-6">
			<div class="flex items-start justify-between mb-4">
				<h3 class="text-xl font-semibold text-gray-800">Create Task</h3>
				<button @click="close" class="text-gray-400 hover:text-gray-600"><i class="pi pi-times text-xl"></i></button>
			</div>

			<form @submit="handleSubmit" class="space-y-4">
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Title</label>
					<input v-model="task_title" type="text" required placeholder="Enter task title"
						class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
				</div>

				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
					<textarea v-model="description" rows="3" placeholder="Describe the task..."
						class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"></textarea>
				</div>

				<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Due Date</label>
						<input v-model="due_date" type="date"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
					</div>
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Assigned To</label>
						<input v-model="assigned_to" type="text" placeholder="Assignee name"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
					</div>
				</div>

				<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Status</label>
						<select v-model="status"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent">
							<option value="todo">To Do</option>
							<option value="progress">In Progress</option>
							<option value="done">Done</option>
							<option value="blocked">Blocked</option>
						</select>
					</div>
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Labels</label>
						<input v-model="labels" type="text" placeholder="e.g. ui, bug"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
					</div>
				</div>

				<div v-if="error" class="text-red-600 text-sm">{{ error }}</div>

				<div class="flex items-center justify-end space-x-3 pt-2">
					<button type="button" @click="close" class="btn-secondary">Cancel</button>
					<button type="submit" class="btn-primary" :disabled="loading">
						<span v-if="!loading"><i class="pi pi-check"></i> Create Task</span>
						<span v-else>Saving...</span>
					</button>
				</div>
			</form>
		</div>
	</div>
</template>


