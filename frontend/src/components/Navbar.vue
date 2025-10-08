<script setup>
import { RouterLink } from 'vue-router'

</script>

<template>
	<nav class="glass-card border-b border-white/20 sticky top-0 z-50 ">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="flex justify-center items-center h-25 gap-20">
				<!-- Logo/Brand -->
				<div class="flex items-center space-x-2">
					<div class="w-16 h-16 rounded-lg  flex items-center justify-center">
						<img src="/project-manager-logo-nobg.png" alt="project manager">
					</div>
					<span class="text-2xl font-bold text-gray-800">Project Manager</span>
				</div>

				<!-- Navigation Links -->
				<div class="flex items-center space-x-8">
					<RouterLink 
						to="/" 
						class="text-gray-600 hover:text-indigo-600 font-medium transition-colors duration-200 flex items-center space-x-2"
						active-class="text-indigo-600"
					>
						<i class="pi pi-home"></i>
						<span>Home</span>
					</RouterLink>
					<RouterLink 
						to="/boards" 
						class="text-gray-600 hover:text-indigo-600 font-medium transition-colors duration-200 flex items-center space-x-2"
						active-class="text-indigo-600"
					>
						<i class="pi pi-th-large"></i>
						<span>Boards</span>
					</RouterLink>
				</div>
			</div>
		</div>
	</nav>
</template>


