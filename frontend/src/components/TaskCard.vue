<script setup>
import { computed } from 'vue'

const props = defineProps({
	task: {}
})

// Helpers to read values from either backend (sql.Null*) or demo JSON
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
	<div class="glass-card rounded-xl p-4 hover:shadow-lg transition-all duration-200">
		<!-- Title and menu -->
		<div class="flex items-start justify-between mb-2">
			<h4 class="text-gray-800 font-semibold leading-snug">{{ title || 'Untitled task' }}</h4>
			<button class="text-gray-400 hover:text-indigo-600 transition-colors">
				<i class="pi pi-ellipsis-h"></i>
			</button>
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


