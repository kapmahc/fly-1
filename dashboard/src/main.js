// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
Vue.config.productionTip = false

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)

import App from './App'
import router from './router'
import './main.css'
import {i18n, detect as detectLocale, load as loadLocaleMessage} from './i18n'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  template: '<App/>',
  components: {
    App
  }
})

loadLocaleMessage(detectLocale())
