<script setup lang="ts">
import { ref, onMounted } from 'vue'

const skills = [
  { category: 'Systems & HPC', items: ['HPC', 'Parallel File Systems', 'Linux Kernel', 'C/C++'] },
  { category: 'Storage & DB', items: ['Storage Systems', 'DBMS Kernels', 'Distributed Systems', 'Blockchain'] },
  { category: 'Full-stack', items: ['Go', 'Vue 3', 'TypeScript', 'PostgreSQL'] }
];

const isVisible = ref(false)
const aboutRef = ref<HTMLElement | null>(null)

onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      isVisible.value = true
    }
  }, { threshold: 0.2 })

  if (aboutRef.value) observer.observe(aboutRef.value)
})
</script>

<template>
  <section id="about" ref="aboutRef" class="section">
    <h1 :class="{ 'reveal-text': isVisible }">About Me</h1>
    <div class="about-grid">
      <div class="bio" :class="{ 'fade-in-up': isVisible }">
        <p>I am a researcher and software engineer with a deep fascination for the underlying mechanics of modern computing. My work focuses on <strong>High Performance Computing (HPC)</strong>, where I explore the intricacies of <strong>Parallel File Systems</strong> and next-generation <strong>Storage Systems</strong>.</p>
        <p>As a systems enthusiast, I spend my time diving into the core of <strong>Operating Systems</strong> and <strong>Database Management Systems</strong>. I believe that understanding the kernel is the key to building truly optimized and robust software.</p>
        <p>Beyond systems research, I am a <strong>Full-stack Developer</strong> and an avid <strong>Cryptocurrency Enthusiast</strong>. I love turning complex architectural concepts into functional, high-performance digital experiences.</p>
      </div>
      <div class="skills" :class="{ 'fade-in-up-delay': isVisible }">
        <h3>Technical Expertise</h3>
        <div class="skill-groups">
          <div v-for="(skill, gIndex) in skills" :key="skill.category" class="skill-group">
            <h4>{{ skill.category }}</h4>
            <div class="tags">
              <span 
                v-for="(item, iIndex) in skill.items" 
                :key="item" 
                class="tag"
                :style="{ animationDelay: `${(gIndex * 4 + iIndex) * 50 + 500}ms` }"
                :class="{ 'pop-in': isVisible }"
              >
                {{ item }}
              </span>
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
  gap: 5rem;
  align-items: start;
}

.bio p {
  font-size: 1.15rem;
  color: var(--text-muted);
  margin-bottom: 2rem;
  line-height: 1.8;
}

.bio strong {
  color: var(--text-main);
  position: relative;
  z-index: 1;
}

.bio strong::after {
  content: "";
  position: absolute;
  bottom: 2px;
  left: 0;
  width: 100%;
  height: 8px;
  background: var(--primary-glow);
  z-index: -1;
  border-radius: 2px;
}

h3 {
  margin-top: 0;
  margin-bottom: 2.5rem;
  font-size: 1.5rem;
  font-weight: 700;
}

.skill-group {
  margin-bottom: 2rem;
}

h4 {
  margin: 0 0 1rem;
  font-size: 0.85rem;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  color: var(--text-muted);
  font-weight: 600;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.tag {
  background: white;
  border: 1px solid var(--border);
  padding: 0.5rem 1rem;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 600;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  color: var(--text-main);
  opacity: 0;
}

.tag.pop-in {
  animation: popIn 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;
}

@keyframes popIn {
  from { transform: scale(0.8); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}

.tag:hover {
  border-color: var(--primary);
  color: var(--primary);
  transform: translateY(-3px);
  box-shadow: 0 5px 15px var(--primary-glow);
}

@media (max-width: 900px) {
  .about-grid {
    grid-template-columns: 1fr;
    gap: 4rem;
  }
}
</style>
