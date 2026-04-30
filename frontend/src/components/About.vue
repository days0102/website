<script setup lang="ts">
const skills = [
  { category: 'Systems & HPC', items: ['HPC', 'Parallel File Systems', 'Linux Kernel', 'C/C++'] },
  { category: 'Storage & DB', items: ['Storage Systems', 'DBMS Kernels', 'Distributed Systems', 'Blockchain'] },
  { category: 'Full-stack', items: ['Go', 'Vue 3', 'TypeScript', 'PostgreSQL'] }
];
</script>

<template>
  <section id="about" class="section">
    <h1>About Me</h1>
    <div class="about-grid">
      <div class="bio">
        <p>I am a researcher and software engineer with a deep fascination for the underlying mechanics of modern computing. My work focuses on <strong>High Performance Computing (HPC)</strong>, where I explore the intricacies of <strong>Parallel File Systems</strong> and next-generation <strong>Storage Systems</strong>.</p>
        <p>As a systems enthusiast, I spend my time diving into the core of <strong>Operating Systems</strong> and <strong>Database Management Systems</strong>. I believe that understanding the kernel is the key to building truly optimized and robust software.</p>
        <p>Beyond systems research, I am a <strong>Full-stack Developer</strong> and an avid <strong>Cryptocurrency Enthusiast</strong>. I love turning complex architectural concepts into functional, high-performance digital experiences.</p>
      </div>
      <div class="skills">
        <h3>Technical Expertise</h3>
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
  font-size: 1.15rem;
  color: var(--text-muted);
  margin-bottom: 1.5rem;
  line-height: 1.8;
}
.bio strong {
  color: var(--text-main);
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
