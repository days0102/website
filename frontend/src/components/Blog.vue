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
const isVisible = ref(false)
const sectionRef = ref<HTMLElement | null>(null)

onMounted(async () => {
  try {
    const response = await fetch('/api/blog');
    blogs.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch blogs:', error);
  }

  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      isVisible.value = true
    }
  }, { threshold: 0.1 })

  if (sectionRef.value) observer.observe(sectionRef.value)
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
    const section = document.getElementById('blog');
    if (section) {
      const offset = section.offsetTop - 80;
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
  <section id="blog" ref="sectionRef" class="section">
    <h1 :class="{ 'reveal-text': isVisible }">Blog</h1>
    
    <div class="blog-wrapper">
      <Transition name="fade-slide" mode="out-in">
        <!-- Blog List View -->
        <div v-if="!selectedPost" key="list" class="blog-list">
          <article 
            v-for="(blog, index) in blogs" 
            :key="blog.id" 
            class="blog-post"
            :style="{ animationDelay: `${index * 150 + 200}ms` }"
            :class="{ 'fade-in-up': isVisible }"
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
  gap: 4rem;
}

.blog-post {
  max-width: 850px;
  opacity: 0;
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
  font-size: 1.8rem;
  margin: 0 0 1rem;
  font-weight: 800;
  letter-spacing: -0.02em;
  transition: color 0.3s;
  line-height: 1.3;
}

.blog-post:hover h3 {
  color: var(--primary);
}

.read-more-btn {
  background: none;
  border: none;
  color: var(--text-main);
  font-weight: 700;
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
  padding: 4rem;
  border-radius: var(--radius);
  border: 1px solid var(--border);
  box-shadow: 0 10px 40px rgba(0,0,0,0.02);
}

.content-header {
  margin-bottom: 3.5rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid var(--border);
}

.back-btn {
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 0.7rem 1.4rem;
  border-radius: 12px;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
  font-weight: 700;
  font-size: 0.9rem;
  margin-bottom: 2.5rem;
  transition: var(--transition-smooth);
}

.back-btn:hover {
  background: var(--border);
  transform: translateX(-6px);
}

.post-display-title {
  font-size: 2.8rem;
  font-weight: 900;
  letter-spacing: -0.04em;
  line-height: 1.1;
  margin: 0;
  color: var(--text-main);
}

/* Transitions */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(20px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

/* Markdown Body Styling */
.markdown-body :deep(h2) {
  font-size: 2rem;
  margin-top: 3rem;
  margin-bottom: 1.5rem;
  font-weight: 800;
  letter-spacing: -0.02em;
}

.markdown-body :deep(p) {
  font-size: 1.15rem;
  line-height: 1.9;
  margin-bottom: 1.75rem;
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
  font-family: 'JetBrains Mono', monospace;
}

.markdown-body :deep(ul) {
  padding-left: 1.5rem;
  margin-bottom: 2rem;
}

.markdown-body :deep(li) {
  margin-bottom: 1rem;
  font-size: 1.15rem;
}

@media (max-width: 768px) {
  .blog-content-viewer {
    padding: 2rem;
  }
  .post-display-title {
    font-size: 2rem;
  }
}
</style>
