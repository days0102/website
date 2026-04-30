<script setup lang="ts">
const skills = [
  { category: 'Frontend', items: ['Vue 3', 'TypeScript', 'Vite', 'Tailwind CSS'] },
  { category: 'Backend', items: ['Go', 'Node.js', 'PostgreSQL', 'Redis'] },
  { category: 'DevOps', items: ['Docker', 'Nginx', 'GitHub Actions', 'Linux'] }
];
</script>

<template>
  <section id="about" class="section">
    <h1>About Me</h1>
    <div class="about-grid">
      <div class="bio">
        <p>I am a passionate software engineer based in China, dedicated to building clean, efficient, and user-centric web applications.</p>
        <p>With a strong foundation in both frontend and backend technologies, I enjoy solving complex problems and turning ideas into reality through code.</p>
      </div>
      <div class="skills">
        <h3>Technical Skills</h3>
        <div class="skill-groups">
          <div v-for="skill in skills" :key="skill.category" class="skill-group">
            <h4>{{ skill.category }}</h4>
            <div class="tags">
              <span v-for="item in skill.items" :key="item" class="tag">{{ item }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about-grid {
  display: grid;
  grid-template-columns: 1.2fr 1fr;
  gap: 4rem;
  align-items: start;
}
.bio p {
  font-size: 1.2rem;
  color: var(--text-muted);
  margin-bottom: 1.5rem;
}
h3 {
  margin-top: 0;
  margin-bottom: 2rem;
}
.skill-group {
  margin-bottom: 1.5rem;
}
h4 {
  margin: 0 0 0.75rem;
  font-size: 0.9rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--text-muted);
}
.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}
.tag {
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 0.4rem 0.8rem;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: 500;
  transition: all 0.2s;
}
.tag:hover {
  border-color: var(--primary);
  color: var(--primary);
  transform: translateY(-1px);
}

@media (max-width: 900px) {
  .about-grid {
    grid-template-columns: 1fr;
    gap: 3rem;
  }
}
</style>
