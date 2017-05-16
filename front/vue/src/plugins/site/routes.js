import Home from './Home'
import Dashboard from './Dashboard'
import Install from './Install'

export default [
  { path: '/', name: 'site.home', component: Home },
  { path: '/dashboard', name: 'site.dashboard', component: Dashboard },
  { path: '/leave-words/new', name: 'site.leave-words.new', component: Dashboard },
  { path: '/install', name: 'site.install', component: Install }
]
