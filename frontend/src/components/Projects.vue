<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface Project {
  id: number;
  title: string;
  description: string;
  link: string;
}

const projects = ref<Project[]>([]);

onMounted(async () => {
  try {
    const response = await fetch('/api/projects');
    projects.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch projects:', error);
  }
});
</script>

<template>
  <section id="projects" class="section">
    <h1>Projects</h1>
    <div class="grid">
      <div v-for="project in projects" :key="project.id" class="card">
        <div class="card-content">
          <h3>{{ project.title }}</h3>
          <p>{{ project.description }}</p>
          <a :href="project.link" target="_blank" class="project-link">
            View Project
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M7 17l9.2-9.2M17 17V7H7"/></svg>
          </a>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
}
.card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  overflow: hidden;
  transition: all 0.3s ease;
}
.card:hover {
  transform: translateY(-8px);
  border-color: var(--primary);
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}
.card-content {
  padding: 2rem;
}
h3 {
  margin: 0 0 1rem;
  font-size: 1.25rem;
  font-weight: 700;
}
p {
  color: var(--text-muted);
  margin-bottom: 2rem;
  font-size: 0.95rem;
  min-height: 3em;
}
.project-link {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--primary);
  text-decoration: none;
  font-weight: 600;
  font-size: 0.9rem;
}
.project-link svg {
  transition: transform 0.2s;
}
.project-link:hover svg {
  transform: translate(2px, -2px);
}
</style>
