<script setup>
import {useRoute, useRouter } from 'vue-router'
import {onMounted, reactive, ref} from 'vue'
import api from "@/api/api.js"
const route = useRoute()
const id = route.params.id
const boardID = route.params.boardID
const router = useRouter()

const state = reactive({
	task : {},
})
	
const formatted = ref("")
const handleUpdate = async () => {
	try{
		await api.put(`/api/v1/tasks/edit/${id}`, {
			task_title: state.task.task_title,
			description: state.task.description,
			due_date: formatted.value,
			assigned_to: state.task.assigned_to,
			status: state.task.status,
			labels: state.task.labels[0],
		})
		router.push(`/boards/${boardID}`)
	} catch(err){
		console.log(err)
	}

}

onMounted(async () => {
	const resp = await api.get(`/api/v1/task/${id}`)
	state.task = resp.data
	const iso = state.task.due_date
	const date = new Date(iso)
	formatted.value = date.toISOString().split("T")[0]
})

</script>
<template>
		<!-- Modal Card -->
		<div class="glass-card mx-auto mt-30 rounded-2xl w-full max-w-lg p-6 text-black">
				<h3 class="text-xl font-semibold text-gray-800 text-center">Edit Task</h3>

			<form @submit.prevent="handleUpdate" class="space-y-4">
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Title</label>
					<input v-model="state.task.task_title" type="text" required placeholder="Enter task title"
						class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
				</div>

				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Description</label>
					<textarea v-model="state.task.description" rows="3" placeholder="Describe the task..."
						class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"></textarea>
				</div>

				<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Due Date</label>
						<input v-model="formatted" type="date"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
					</div>
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Assigned To</label>
						<input v-model="state.task.assigned_to" type="text" placeholder="Assignee name"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
					</div>
				</div>

				<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Status</label>
						<select v-model="state.task.status"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent">
							<option value="todo">To Do</option>
							<option value="progress">In Progress</option>
							<option value="done">Done</option>
							<option value="blocked">Blocked</option>
						</select>
					</div>
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Labels</label>
						<input v-model="state.task.labels" type="text" placeholder="e.g. ui, bug"
							class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent" />
					</div>
				</div>
			<button type="submit" class="btn btn-primary mx-auto block">Update</button>

			</form>
		</div>
</template>
