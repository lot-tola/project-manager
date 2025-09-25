<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
	<div class="min-h-screen gradient-bg flex items-center justify-center py-16">
		<div class="glass-card p-10 rounded-2xl text-center max-w-xl">
			<h1 class="text-4xl font-bold text-gray-800 mb-4">Welcome to Project Manager</h1>
			<p class="text-gray-600 mb-6">Organize your boards, lists, and tasks beautifully.</p>
			<RouterLink to="/boards" class="btn-primary inline-flex items-center space-x-2">
				<i class="pi pi-th-large"></i>
				<span>Go to Boards</span>
			</RouterLink>
		</div>
	</div>
</template>


