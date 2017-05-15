import 'jquery'
import 'bootstrap'

import Vue from 'vue'
Vue.config.productionTip = false

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
