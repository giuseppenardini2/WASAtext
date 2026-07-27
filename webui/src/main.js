import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMessage from './components/ErrorMessage.vue'
import UserSearchBar from './components/UserSearchBar.vue'
import MessageMenu from './components/MessageMenu.vue'
import ConversationModal from './components/ConversationModal.vue'

import './assets/css/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMessage", ErrorMessage);
app.component("UserSearchBar", UserSearchBar);
app.component("MessageMenu", MessageMenu);
app.component("ConversationModal", ConversationModal);
app.use(router)
app.mount('#app')
