import Home from './Home'
import Dashboard from './Dashboard'
import Install from './Install'
import NewLeaveWord from './leave-words/New'

export default [
  { path: '/', name: 'site.home', component: Home },
  { path: '/dashboard', name: 'site.dashboard', component: Dashboard },
  { path: '/leave-words/new', name: 'site.leave-words.new', component: NewLeaveWord },
  { path: '/install', name: 'site.install', component: Install }
]
