<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface Blog {
  id: number;
  title: string;
  summary: string;
  date: string;
}

const blogs = ref<Blog[]>([]);

onMounted(async () => {
  try {
    const response = await fetch('/api/blog');
    blogs.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch blogs:', error);
  }
});
</script>

<template>
  <section id="blog" class="section">
    <h1>Blog</h1>
    <div class="blog-list">
      <article v-for="blog in blogs" :key="blog.id" class="blog-post">
        <div class="post-meta">
          <time>{{ blog.date }}</time>
          <span class="dot"></span>
          <span class="read-time">5 min read</span>
        </div>
        <h3>{{ blog.title }}</h3>
        <p>{{ blog.summary }}</p>
        <a href="#" class="read-more">Read article &rarr;</a>
      </article>
    </div>
  </section>
</template>

<style scoped>
.blog-list {
  display: flex;
  flex-direction: column;
  gap: 3rem;
}
.blog-post {
  max-width: 800px;
}
.post-meta {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
  font-size: 0.85rem;
  color: var(--text-muted);
}
.dot {
  width: 3px;
  height: 3px;
  background: var(--border);
  border-radius: 50%;
}
h3 {
  font-size: 1.5rem;
  margin: 0 0 0.75rem;
  font-weight: 700;
  transition: color 0.2s;
}
.blog-post:hover h3 {
  color: var(--primary);
}
p {
  color: var(--text-muted);
  margin-bottom: 1.25rem;
  font-size: 1.05rem;
}
.read-more {
  color: var(--text-main);
  text-decoration: none;
  font-weight: 600;
  font-size: 0.9rem;
}
.read-more:hover {
  color: var(--primary);
}
</style>
