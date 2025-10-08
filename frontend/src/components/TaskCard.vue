<script setup>
import { computed } from 'vue'
import axios from 'axios'
import api from "@/api/api.js"
import { RouterLink } from 'vue-router'

const props = defineProps({
	task: {},
	boardID: {
	type: String,
	},
})

const readString = (obj, keyCandidates) => {
	for (const key of keyCandidates) {
		const val = obj?.[key]
		if (val === undefined || val === null) continue
		if (typeof val === 'object' && 'String' in val) return val.String || ''
		if (typeof val === 'string') return val
	}
	return ''
}

const readDateISO = (obj, keyCandidates) => {
	for (const key of keyCandidates) {
		const val = obj?.[key]
		if (!val) continue
		if (typeof val === 'object' && 'Time' in val) return val.Time
		if (typeof val === 'string' || val instanceof Date) return val
	}
	return ''
}

const handleDeleteTask = () => {
	const ok = confirm("Are you sure? This action can not be undone.")
	if (!ok){
		return 
	}
	axios.delete(`/api/v1/tasks/${props.task.task_id}`)
}

const title = computed(() => readString(props.task, ['task_title', 'title']))
const description = computed(() => readString(props.task, ['description']))
const assignedTo = computed(() => readString(props.task, ['assigned_to', 'assignedTo']))
const status = computed(() => readString(props.task, ['status']))
const dueDateRaw = computed(() => readDateISO(props.task, ['due_date', 'dueDate']))
const dueDate = computed(() => {
	if (!dueDateRaw.value) return ''
	const d = new Date(dueDateRaw.value)
	return isNaN(d.getTime()) ? '' : d.toLocaleDateString()
})

const labels = computed(() => Array.isArray(props.task?.labels) ? props.task.labels : [])
const id = computed(() => props.task?.id || props.task?.task_id?.String || '')

const statusClass = computed(() => {
	const s = status.value.toLowerCase()
	if (s.includes('done') || s === 'completed') return 'bg-green-100 text-green-700'
	if (s.includes('progress') || s === 'doing') return 'bg-amber-100 text-amber-700'
	if (s.includes('blocked')) return 'bg-red-100 text-red-700'
	return 'bg-gray-100 text-gray-700'
})
</script>

<template>
	<div v-if="props.task.task_id === null" class="font-bold opacity-50 text-center text-red-400">
		No task yet. Create new task?
	</div>
	<div v-else class="glass-card rounded-xl p-4 hover:shadow-lg transition-all duration-200">
		<!-- Title and menu -->
		<div class="flex items-start justify-between mb-2">
			<h4 class="text-gray-800 font-semibold leading-snug">{{ title || 'Untitled task' }}</h4>
				<div class="dropdown">
					<i tabindex="0" role="button" class="bg-white text-gray-800 h-5 rounded-lg border-none shadow-none pi pi-ellipsis-h btn m-1"></i>
					<ul tabindex="0" class="dropdown-content menu bg-[#f6f6f6] text-black opacity-50 rounded-box z-1 w-52 p-2 shadow-sm ">
						<form @submit="handleDeleteTask">
							<li class="hover:bg-[#ececec] hover:rounded-md">
								<button type="submit">Delete</button>
							</li>
						</form>
							<li class="hover:bg-[#ececec] hover:rounded-md">
						<RouterLink :to="`/boards/edit/${props.task.task_id}/${props.boardID}`">Edit</RouterLink>
							</li>
					</ul>
				</div>
		</div>

		<!-- Meta row -->
		<div class="flex flex-wrap items-center gap-2 mb-2">
			<span v-if="status" class="text-xs px-2 py-1 rounded-full" :class="statusClass">{{ status }}</span>
			<span v-if="dueDate" class="text-xs px-2 py-1 rounded-full bg-indigo-50 text-indigo-700">
				<i class="pi pi-calendar mr-1"></i>{{ dueDate }}
			</span>
			<span v-if="assignedTo" class="text-xs px-2 py-1 rounded-full bg-sky-50 text-sky-700">
				<i class="pi pi-user mr-1"></i>{{ assignedTo }}
			</span>
			<span v-if="id" class="text-xs text-gray-400">#{{ id }}</span>
		</div>

		<!-- Description -->
		<p v-if="description" class="text-sm text-gray-600 mb-3 line-clamp-3">{{ description }}</p>

		<!-- Labels -->
		<div v-if="labels.length" class="flex flex-wrap gap-2">
			<span v-for="(label, idx) in labels" :key="idx" class="text-xs px-2 py-1 rounded-full bg-gray-100 text-gray-700">{{ label }}</span>
		</div>
	</div>
</template>


