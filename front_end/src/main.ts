import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'
import Tags from './components/Tags.vue'
import Upload from './components/Upload.vue'
const routerHistory = createWebHistory()
const router = createRouter({
  history: routerHistory,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/tags',
      name: 'tags',
      component: Tags
    },
    {
      path: '/upload',
      name: 'upload',
      component: Upload
    }
  ]
})

const app = createApp(App)
app.use(router)
app.use(ElementPlus)
app.mount('#app')