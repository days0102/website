<script setup lang="ts">
import { ref } from 'vue';

const email = ref('');
const message = ref('');
const isSubmitting = ref(false);
const statusMsg = ref('');

const submitForm = async () => {
  isSubmitting.value = true;
  statusMsg.value = '';
  
  try {
    const response = await fetch('/api/contact', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: email.value, message: message.value })
    });

    if (response.ok) {
      statusMsg.value = 'Message sent successfully! I will get back to you soon.';
      email.value = '';
      message.value = '';
    } else {
      statusMsg.value = 'Failed to send message. Please try again later.';
    }
  } catch (error) {
    console.error('Submission error:', error);
    statusMsg.value = 'An error occurred. Please check your connection.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>

<template>
  <section id="contact" class="section">
    <h1>Contact</h1>
    <div class="contact-container">
      <div class="contact-info">
        <h2>Let's collaborate.</h2>
        <p>I'm always open to discussing new projects, creative ideas or opportunities to be part of your visions.</p>
        <div class="email-direct">
          <span>Email me at</span>
          <a href="mailto:hello@lwz.me">hello@lwz.me</a>
        </div>
      </div>
      
      <form @submit.prevent="submitForm" class="contact-form">
        <div class="form-group">
          <label>Your Email</label>
          <input v-model="email" type="email" placeholder="email@example.com" required />
        </div>
        <div class="form-group">
          <label>Message</label>
          <textarea v-model="message" rows="5" placeholder="Tell me about your project..." required></textarea>
        </div>
        <button type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? 'Sending...' : 'Send Message' }}
        </button>
        <p v-if="statusMsg" :class="['status-msg', { error: statusMsg.includes('Failed') || statusMsg.includes('error') }]">
          {{ statusMsg }}
        </p>
      </form>
    </div>
  </section>
</template>

<style scoped>
.contact-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: start;
}
.contact-info h2 {
  font-size: 2.5rem;
  margin-top: 0;
  line-height: 1.2;
}
.contact-info p {
  color: var(--text-muted);
  font-size: 1.1rem;
  margin-bottom: 2.5rem;
}
.email-direct span {
  display: block;
  font-size: 0.85rem;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  color: var(--text-muted);
  margin-bottom: 0.5rem;
}
.email-direct a {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--primary);
  text-decoration: none;
}
.contact-form {
  background: var(--surface);
  padding: 3rem;
  border-radius: var(--radius);
  border: 1px solid var(--border);
}
.form-group {
  margin-bottom: 1.5rem;
}
label {
  display: block;
  font-size: 0.9rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}
input, textarea {
  width: 100%;
  padding: 1rem;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: white;
  font-family: inherit;
  font-size: 1rem;
  transition: border-color 0.2s;
  box-sizing: border-box;
}
input:focus, textarea:focus {
  outline: none;
  border-color: var(--primary);
}
button {
  width: 100%;
  padding: 1rem;
  background-color: var(--text-main);
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.2s;
}
button:hover:not(:disabled) {
  background-color: var(--primary);
  transform: translateY(-2px);
}
button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
.status-msg {
  margin-top: 1rem;
  font-size: 0.9rem;
  color: var(--primary);
}
.status-msg.error {
  color: #ef4444;
}

@media (max-width: 900px) {
  .contact-container {
    grid-template-columns: 1fr;
    gap: 3rem;
  }
  .contact-form {
    padding: 2rem;
  }
}
</style>
