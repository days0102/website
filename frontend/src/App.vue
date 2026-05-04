<script setup lang="ts">
import { ref, onMounted } from 'vue'
import About from './components/About.vue'
import Projects from './components/Projects.vue'
import Blog from './components/Blog.vue'
import Contact from './components/Contact.vue'

const activeSection = ref('about')

onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        activeSection.value = entry.target.id
      }
    })
  }, { threshold: 0.3 })

  document.querySelectorAll('section').forEach(section => observer.observe(section))
})
</script>

<template>
  <div class="app-container">
    <header class="header">
      <nav class="container">
        <div class="logo">LWZ <span>.</span></div>
        <ul>
          <li><a href="#about" :class="{ active: activeSection === 'about' }">About</a></li>
          <li><a href="#projects" :class="{ active: activeSection === 'projects' }">Work</a></li>
          <li><a href="#blog" :class="{ active: activeSection === 'blog' }">Writing</a></li>
          <li><a href="#contact" class="nav-cta">Say Hello</a></li>
        </ul>
      </nav>
    </header>

    <main>
      <div class="hero">
        <div class="container">
          <div class="hero-content">
            <h1 class="reveal-text">Exploring the <br/><span class="gradient-text">depths of systems.</span></h1>
            <p class="subtitle fade-in-up">HPC Researcher & Full-stack Engineer. Dedicated to Parallel File Systems, Storage Architectures, and Kernel Internals.</p>
            <div class="hero-actions fade-in-up-delay">
              <a href="#projects" class="btn-primary">View My Work</a>
              <a href="#blog" class="btn-secondary">Read Blog</a>
            </div>
          </div>
        </div>
      </div>
      <div class="container main-content">
        <About />
        <Projects />
        <Blog />
        <Contact />
      </div>
    </main>

    <footer class="footer">
      <div class="container">
        <div class="footer-info">
          <p>&copy; 2026 Designed & Built by LWZ</p>
          <a href="https://beian.miit.gov.cn" target="_blank" class="icp-link">粤ICP备2026049174号-1</a>
        </div>
        <div class="footer-links">
          <a href="#">Github</a>
          <a href="#">Twitter</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<style>
:root {
  --primary: #6366f1;
  --primary-glow: rgba(99, 102, 241, 0.2);
  --primary-dark: #4f46e5;
  --bg: #ffffff;
  --surface: #f8fafc;
  --text-main: #0f172a;
  --text-muted: #64748b;
  --border: #e2e8f0;
  --max-width: 1100px;
  --radius: 16px;
  --transition-smooth: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

html {
  scroll-behavior: smooth;
}

body {
  margin: 0;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
  color: var(--text-main);
  background-color: var(--bg);
  line-height: 1.6;
  -webkit-font-smoothing: antialiased;
  overflow-x: hidden;
}

.container {
  max-width: var(--max-width);
  margin: 0 auto;
  padding: 0 2rem;
}

/* Header & Nav */
.header {
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-bottom: 1px solid var(--border);
  position: sticky;
  top: 0;
  z-index: 1000;
  transition: var(--transition-smooth);
}

nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 80px;
}

.logo {
  font-weight: 800;
  font-size: 1.6rem;
  letter-spacing: -0.05em;
  cursor: pointer;
}

.logo span {
  color: var(--primary);
  display: inline-block;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.5); opacity: 0.5; }
}

nav ul {
  display: flex;
  list-style: none;
  gap: 2rem;
  margin: 0;
  padding: 0;
  align-items: center;
}

nav a {
  text-decoration: none;
  color: var(--text-muted);
  font-weight: 500;
  font-size: 0.95rem;
  position: relative;
  padding: 0.5rem 0;
  transition: var(--transition-smooth);
}

nav a:hover:not(.nav-cta), 
nav a.active:not(.nav-cta) {
  color: var(--text-main);
}

nav a.active::after {
  content: "";
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background: var(--primary);
  border-radius: 2px;
  animation: slideIn 0.3s ease;
}

@keyframes slideIn {
  from { width: 0; }
  to { width: 100%; }
}

.nav-cta {
  background: var(--text-main);
  color: white !important;
  padding: 0.7rem 1.4rem;
  border-radius: 10px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.1);
}

.nav-cta:hover {
  background: var(--primary);
  transform: translateY(-2px);
  box-shadow: 0 8px 25px var(--primary-glow);
}

/* Hero Section */
.hero {
  padding: 10rem 0 6rem;
  min-height: 80vh;
  display: flex;
  align-items: center;
}

.hero-content {
  max-width: 800px;
}

.hero h1 {
  font-size: clamp(3.5rem, 10vw, 5.5rem);
  line-height: 1;
  font-weight: 900;
  letter-spacing: -0.05em;
  margin-bottom: 2rem;
}

.gradient-text {
  background: linear-gradient(135deg, var(--primary) 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: 1.35rem;
  color: var(--text-muted);
  max-width: 600px;
  margin-bottom: 3rem;
  font-weight: 400;
}

.hero-actions {
  display: flex;
  gap: 1.5rem;
}

.btn-primary {
  background: var(--primary);
  color: white;
  padding: 1rem 2rem;
  border-radius: 12px;
  text-decoration: none;
  font-weight: 600;
  transition: var(--transition-smooth);
  box-shadow: 0 10px 30px var(--primary-glow);
}

.btn-primary:hover {
  transform: translateY(-3px);
  box-shadow: 0 15px 40px var(--primary-glow);
  background: var(--primary-dark);
}

.btn-secondary {
  border: 1px solid var(--border);
  color: var(--text-main);
  padding: 1rem 2rem;
  border-radius: 12px;
  text-decoration: none;
  font-weight: 600;
  transition: var(--transition-smooth);
}

.btn-secondary:hover {
  background: var(--surface);
  border-color: var(--text-main);
}

/* Common Animations */
.reveal-text {
  animation: reveal 1s cubic-bezier(0.77, 0, 0.175, 1);
}

@keyframes reveal {
  0% { transform: translateY(30px); opacity: 0; }
  100% { transform: translateY(0); opacity: 1; }
}

.fade-in-up {
  animation: fadeInUp 1s cubic-bezier(0.77, 0, 0.175, 1) 0.2s both;
}

.fade-in-up-delay {
  animation: fadeInUp 1s cubic-bezier(0.77, 0, 0.175, 1) 0.4s both;
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

/* Sections */
.section {
  padding: 8rem 0;
}

.section h1 {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 4rem;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  letter-spacing: -0.03em;
}

.section h1::after {
  content: "";
  height: 2px;
  flex: 1;
  background: linear-gradient(to right, var(--border), transparent);
}

/* Footer */
.footer {
  background: var(--surface);
  padding: 5rem 0;
  margin-top: 8rem;
  border-top: 1px solid var(--border);
}

.footer .container {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.footer-info {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer p {
  color: var(--text-muted);
  font-size: 1rem;
  margin: 0;
  font-weight: 500;
}

.icp-link {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.85rem;
  opacity: 0.7;
  transition: var(--transition-smooth);
}

.icp-link:hover {
  color: var(--primary);
  opacity: 1;
}

.footer-links {
  display: flex;
  gap: 3rem;
}

.footer-links a {
  color: var(--text-main);
  text-decoration: none;
  font-size: 1rem;
  font-weight: 600;
  transition: var(--transition-smooth);
}

.footer-links a:hover {
  color: var(--primary);
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  nav ul { display: none; }
  .hero h1 { font-size: 3rem; }
  .hero-actions { flex-direction: column; }
  .footer .container { flex-direction: column; gap: 3rem; }
}
</style>
