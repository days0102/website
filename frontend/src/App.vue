<script setup lang="ts">
import About from './components/About.vue'
import Projects from './components/Projects.vue'
import Blog from './components/Blog.vue'
import Contact from './components/Contact.vue'
</script>

<template>
  <header class="header">
    <nav class="container">
      <div class="logo">LWZ <span>.</span></div>
      <ul>
        <li><a href="#about">About</a></li>
        <li><a href="#projects">Work</a></li>
        <li><a href="#blog">Writing</a></li>
        <li><a href="#contact" class="nav-cta">Say Hello</a></li>
      </ul>
    </nav>
  </header>

  <main>
    <div class="hero">
      <div class="container">
        <h1>Exploring the <br/><span class="gradient-text">depths of systems.</span></h1>
        <p class="subtitle">HPC Researcher & Full-stack Engineer. Dedicated to Parallel File Systems, Storage Architectures, and Kernel Internals.</p>
      </div>
    </div>
    <div class="container">
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
</template>

<style>
:root {
  --primary: #6366f1;
  --primary-dark: #4f46e5;
  --bg: #ffffff;
  --surface: #f8fafc;
  --text-main: #0f172a;
  --text-muted: #64748b;
  --border: #e2e8f0;
  --max-width: 1100px;
  --radius: 12px;
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
}

.container {
  max-width: var(--max-width);
  margin: 0 auto;
  padding: 0 2rem;
}

/* Header & Nav */
.header {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  border-bottom: 1px solid var(--border);
  position: sticky;
  top: 0;
  z-index: 100;
}

nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 80px;
}

.logo {
  font-weight: 800;
  font-size: 1.5rem;
  letter-spacing: -0.05em;
}

.logo span {
  color: var(--primary);
}

nav ul {
  display: flex;
  list-style: none;
  gap: 2.5rem;
  margin: 0;
  padding: 0;
  align-items: center;
}

nav a {
  text-decoration: none;
  color: var(--text-main);
  font-weight: 500;
  font-size: 0.95rem;
  transition: all 0.2s;
}

nav a:hover:not(.nav-cta) {
  color: var(--primary);
}

.nav-cta {
  background: var(--text-main);
  color: white !important;
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
}

.nav-cta:hover {
  background: var(--primary);
  transform: translateY(-1px);
}

/* Hero Section */
.hero {
  padding: 8rem 0 4rem;
  text-align: left;
}

.hero h1 {
  font-size: clamp(3rem, 8vw, 5rem);
  line-height: 1.1;
  font-weight: 800;
  letter-spacing: -0.04em;
  margin-bottom: 1.5rem;
  border-left: none;
  padding-left: 0;
}

.gradient-text {
  background: linear-gradient(135deg, var(--primary), #a855f7);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
  max-width: 600px;
}

/* Common Section Styles */
.section {
  padding: 6rem 0;
}

.section h1 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 3rem;
  display: flex;
  align-items: center;
  gap: 1rem;
}

.section h1::after {
  content: "";
  height: 1px;
  flex: 1;
  background: var(--border);
}

/* Footer */
.footer {
  background: var(--surface);
  padding: 4rem 0;
  margin-top: 6rem;
  border-top: 1px solid var(--border);
}

.footer .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.footer-info {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.footer p {
  color: var(--text-muted);
  font-size: 0.9rem;
  margin: 0;
}

.icp-link {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.85rem;
  transition: color 0.2s;
}

.icp-link:hover {
  color: var(--primary);
}

.footer-links {
  display: flex;
  gap: 2rem;
}

.footer-links a {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.9rem;
}

.footer-links a:hover {
  color: var(--primary);
}

@media (max-width: 768px) {
  nav ul {
    display: none;
  }
  .hero {
    padding: 4rem 0;
  }
  .footer .container {
    flex-direction: column;
    gap: 1.5rem;
    text-align: center;
  }
}
</style>
