import Vue from 'vue'
import {sync} from 'vuex-router-sync'
import VueResource from 'vue-resource'
import App from './components/App'
import router from './router'
import store from './store'

Vue.use(VueResource)
Vue.http.options.xhr = {withCredentials: true}
Vue.config.devtools = true

sync(store, router)

const app = new Vue({
  router,
  store,
  ...App
})

export {app, router, store}
