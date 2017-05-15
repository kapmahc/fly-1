import Vue from 'vue'
import Router from 'vue-router'

import site from './plugins/site/routes'
import auth from './plugins/auth/routes'

Vue.use(Router)

export default new Router({
  routes: [].concat(site).concat(auth)
})
