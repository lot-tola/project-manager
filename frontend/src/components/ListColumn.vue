<script setup>
import TaskCard from '@/components/TaskCard.vue'
import CreateTaskOverlay from '@/components/CreateTaskOverlay.vue'
import { ref, onMounted } from 'vue'
import axios from 'axios'

const showCreateTask = ref(false)
const props = defineProps({
	list: {
		type: Object,
		required: true
	}
})
console.log(props.list)

function handleCloseCreateTaskOverlay() {
	showCreateTask.value = false
}
function handleOpenCreateTaskOverlay() {
	showCreateTask.value = true
}

// onMounted(fetchTasks)
</script>

<template>
	<div class="w-80 flex-shrink-0">
		<div v-if="showCreateTask">
			<CreateTaskOverlay @close="handleCloseCreateTaskOverlay" :list_id="list.list_id" />
		</div>

		<div class="glass-card rounded-2xl p-4">
			<div class="flex items-center justify-between mb-3">
				<h3 class="text-lg font-semibold text-gray-800">{{ list.list_title }}</h3>
				<span class="text-xs px-2 py-1 rounded-full bg-gray-100 text-gray-600">{{ list.tasks?.length || 0 }}</span>
			</div>
			<div class="space-y-3">
				<div v-for="task in list.tasks" :key="task.id">
					<TaskCard :task="task" />
				</div>
			</div>
			<div class="mt-4">
				<button @click="handleOpenCreateTaskOverlay" class="btn-secondary w-full flex items-center justify-center space-x-2">
					<i class="pi pi-plus"></i>
					<span>Add task</span>
				</button>
			</div>
		</div>
	</div>
</template>


