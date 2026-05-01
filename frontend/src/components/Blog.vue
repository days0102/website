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
const isLoading = ref(false);

onMounted(async () => {
  try {
    const response = await fetch('/api/blog');
    blogs.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch blogs:', error);
  }
});

const readPost = async (blog: Blog) => {
  isLoading.value = true;
  try {
    const response = await fetch(`/api/blog/content?title=${encodeURIComponent(blog.title)}&date=${blog.date}`);
    const markdown = await response.text();
    selectedPost.value = {
      title: blog.title,
      content: await marked.parse(markdown)
    };
    // Scroll to the top of the blog section instead of the entire page
    const section = document.getElementById('blog');
    if (section) {
      const offset = section.offsetTop - 80; // 80px is the header height
      window.scrollTo({ top: offset, behavior: 'smooth' });
    }
  } catch (error) {
    console.error('Failed to fetch blog content:', error);
  } finally {
    isLoading.value = false;
  }
};

const closePost = () => {
  selectedPost.value = null;
};
</script>

<template>
  <section id="blog" class="section">
    <h1>Blog</h1>
    
    <div class="blog-wrapper">
      <Transition name="fade-slide" mode="out-in">
        <!-- Blog List View -->
        <div v-if="!selectedPost" key="list" class="blog-list">
          <article 
            v-for="(blog, index) in blogs" 
            :key="blog.id" 
            class="blog-post"
            :style="{ transitionDelay: `${index * 100}ms` }"
          >
            <div class="post-meta">
              <time>{{ blog.date }}</time>
              <span class="dot"></span>
              <span class="read-time">Tech Notes</span>
            </div>
            <h3>{{ blog.title }}</h3>
            <p>{{ blog.summary }}</p>
            <button @click="readPost(blog)" class="read-more-btn" :disabled="isLoading">
              {{ isLoading ? 'Loading...' : 'Read article →' }}
            </button>
          </article>
          <p v-if="blogs.length === 0" class="empty-msg">No posts found yet.</p>
        </div>

        <!-- Blog Content View -->
        <div v-else key="content" class="blog-content-viewer">
          <header class="content-header">
            <button @click="closePost" class="back-btn">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="12" x2="5" y2="12"></line><polyline points="12 19 5 12 12 5"></polyline></svg>
              Back to list
            </button>
            <div class="post-title-area">
              <h2 class="post-display-title">{{ selectedPost.title }}</h2>
            </div>
          </header>
          <div class="markdown-body" v-html="selectedPost.content"></div>
        </div>
      </Transition>
    </div>
  </section>
</template>

<style scoped>
.blog-wrapper {
  position: relative;
  min-height: 400px;
}

.blog-list {
  display: flex;
  flex-direction: column;
  gap: 3.5rem;
}

.blog-post {
  max-width: 850px;
  animation: slideUp 0.6s ease forwards;
  opacity: 0;
}

@keyframes slideUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.post-meta {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
  font-size: 0.85rem;
  color: var(--text-muted);
  font-weight: 500;
}

.dot {
  width: 4px;
  height: 4px;
  background: var(--primary);
  border-radius: 50%;
  opacity: 0.6;
}

h3 {
  font-size: 1.75rem;
  margin: 0 0 1rem;
  font-weight: 800;
  letter-spacing: -0.02em;
  transition: color 0.3s;
}

.blog-post:hover h3 {
  color: var(--primary);
}

.read-more-btn {
  background: none;
  border: none;
  color: var(--text-main);
  font-weight: 600;
  font-size: 0.95rem;
  cursor: pointer;
  padding: 0;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.2s;
}

.read-more-btn:hover {
  color: var(--primary);
  gap: 0.8rem;
}

/* Content Viewer */
.blog-content-viewer {
  background: white;
  padding: 3.5rem;
  border-radius: var(--radius);
  border: 1px solid var(--border);
  box-shadow: 0 4px 20px rgba(0,0,0,0.03);
}

.content-header {
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid var(--border);
}

.back-btn {
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
  font-weight: 600;
  font-size: 0.9rem;
  margin-bottom: 2rem;
  transition: all 0.2s;
}

.back-btn:hover {
  background: var(--border);
  transform: translateX(-4px);
}

.post-display-title {
  font-size: 2.5rem;
  font-weight: 800;
  letter-spacing: -0.04em;
  line-height: 1.2;
  margin: 0;
}

/* Transitions */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.4s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Markdown Body Styling */
.markdown-body :deep(h2) {
  font-size: 1.8rem;
  margin-top: 2.5rem;
  margin-bottom: 1.25rem;
  font-weight: 700;
}

.markdown-body :deep(p) {
  font-size: 1.15rem;
  line-height: 1.8;
  margin-bottom: 1.5rem;
  color: #334155;
}

.markdown-body :deep(strong) {
  color: var(--text-main);
  font-weight: 700;
}

.markdown-body :deep(code) {
  background: #f1f5f9;
  padding: 0.2rem 0.5rem;
  border-radius: 6px;
  font-size: 0.9em;
  color: var(--primary-dark);
}

.markdown-body :deep(ul) {
  padding-left: 1.5rem;
  margin-bottom: 1.5rem;
}

.markdown-body :deep(li) {
  margin-bottom: 0.75rem;
  font-size: 1.1rem;
}

@media (max-width: 768px) {
  .blog-content-viewer {
    padding: 1.5rem;
  }
  .post-display-title {
    font-size: 1.8rem;
  }
}
</style>
