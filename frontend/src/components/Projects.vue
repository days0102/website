<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface Project {
  id: number;
  title: string;
  description: string;
  link: string;
}

const projects = ref<Project[]>([]);
const isVisible = ref(false)
const sectionRef = ref<HTMLElement | null>(null)

onMounted(async () => {
  try {
    const response = await fetch('/api/projects');
    projects.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch projects:', error);
  }

  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      isVisible.value = true
    }
  }, { threshold: 0.1 })

  if (sectionRef.value) observer.observe(sectionRef.value)
});
</script>

<template>
  <section id="projects" ref="sectionRef" class="section">
    <h1 :class="{ 'reveal-text': isVisible }">Projects</h1>
    <div class="grid">
      <div 
        v-for="(project, index) in projects" 
        :key="project.id" 
        class="card"
        :class="{ 'fade-in-up': isVisible }"
        :style="{ animationDelay: `${index * 100 + 300}ms` }"
      >
        <div class="card-glow"></div>
        <div class="card-content">
          <div class="project-tag">Case Study</div>
          <h3>{{ project.title }}</h3>
          <p>{{ project.description }}</p>
          <a :href="project.link" target="_blank" class="project-link">
            View Project
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M7 17l9.2-9.2M17 17V7H7"/></svg>
          </a>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 2.5rem;
}

.card {
  background: white;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  position: relative;
  overflow: hidden;
  transition: var(--transition-smooth);
  opacity: 0;
}

.card:hover {
  transform: translateY(-12px) scale(1.02);
  border-color: var(--primary);
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.1);
}

.card-glow {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at top right, var(--primary-glow), transparent 70%);
  opacity: 0;
  transition: opacity 0.5s ease;
  pointer-events: none;
}

.card:hover .card-glow {
  opacity: 1;
}

.card-content {
  padding: 2.5rem;
  position: relative;
  z-index: 1;
}

.project-tag {
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  color: var(--primary);
  margin-bottom: 1rem;
}

h3 {
  margin: 0 0 1.25rem;
  font-size: 1.4rem;
  font-weight: 800;
  color: var(--text-main);
  letter-spacing: -0.02em;
}

p {
  color: var(--text-muted);
  margin-bottom: 2.5rem;
  font-size: 1rem;
  line-height: 1.7;
  min-height: 4.5em;
}

.project-link {
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
  color: var(--text-main);
  text-decoration: none;
  font-weight: 700;
  font-size: 0.95rem;
  transition: var(--transition-smooth);
}

.project-link svg {
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.project-link:hover {
  color: var(--primary);
}

.project-link:hover svg {
  transform: translate(3px, -3px);
}

@media (max-width: 640px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
</style>
