import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMessage from './components/ErrorMessage.vue'
import UserSearchBar from './components/UserSearchBar.vue'

import './assets/css/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMessage", ErrorMessage);
app.use(router)
app.mount('#app')
