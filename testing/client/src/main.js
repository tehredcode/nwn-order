import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store'

/* TODO: refactor using store and router
         remove unnecessary dependencys and plugins
         clean up ui
         add comments where needed
*/

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
