<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { marked } from 'marked';

interface Blog {
  id: number;
  title: string;
  summary: string;
  date: string;
}

const blogs = ref<Blog[]>([]);
const selectedPost = ref<{ title: string; content: string } | null>(null);

onMounted(async () => {
  try {
    const response = await fetch('/api/blog');
    blogs.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch blogs:', error);
  }
});

const readPost = async (blog: Blog) => {
  try {
    const response = await fetch(`/api/blog/content?title=${encodeURIComponent(blog.title)}&date=${blog.date}`);
    const markdown = await response.text();
    selectedPost.value = {
      title: blog.title,
      content: await marked.parse(markdown)
    };
    // Scroll to top of content
    window.scrollTo({ top: document.getElementById('blog')?.offsetTop, behavior: 'smooth' });
  } catch (error) {
    console.error('Failed to fetch blog content:', error);
  }
};

const closePost = () => {
  selectedPost.value = null;
};
</script>

<template>
  <section id="blog" class="section">
    <h1>Blog</h1>
    
    <!-- Blog List View -->
    <div v-if="!selectedPost" class="blog-list">
      <article v-for="blog in blogs" :key="blog.id" class="blog-post">
        <div class="post-meta">
          <time>{{ blog.date }}</time>
          <span class="dot"></span>
          <span class="read-time">Read more</span>
        </div>
        <h3>{{ blog.title }}</h3>
        <p>{{ blog.summary }}</p>
        <button @click="readPost(blog)" class="read-more-btn">Read article &rarr;</button>
      </article>
      <p v-if="blogs.length === 0" class="empty-msg">No posts found in data/posts/ yet.</p>
    </div>

    <!-- Blog Content View -->
    <div v-else class="blog-content-viewer">
      <button @click="closePost" class="back-btn">&larr; Back to list</button>
      <div class="markdown-body" v-html="selectedPost.content"></div>
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
}
p {
  color: var(--text-muted);
  margin-bottom: 1.25rem;
  font-size: 1.05rem;
}
.read-more-btn {
  background: none;
  border: none;
  color: var(--text-main);
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  padding: 0;
  transition: color 0.2s;
}
.read-more-btn:hover {
  color: var(--primary);
}

.blog-content-viewer {
  background: white;
  padding: 2rem;
  border-radius: var(--radius);
  border: 1px solid var(--border);
}
.back-btn {
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  margin-bottom: 2rem;
  transition: all 0.2s;
}
.back-btn:hover {
  background: var(--border);
}

/* Markdown Styling */
.markdown-body :deep(h1), .markdown-body :deep(h2) {
  border-bottom: 1px solid var(--border);
  padding-bottom: 0.5rem;
  margin-top: 2rem;
}
.markdown-body :deep(p) {
  font-size: 1.1rem;
  line-height: 1.8;
}
.markdown-body :deep(code) {
  background: var(--surface);
  padding: 0.2rem 0.4rem;
  border-radius: 4px;
}

.empty-msg {
  color: var(--text-muted);
  font-style: italic;
}
</style>
