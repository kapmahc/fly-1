import React from 'react';
import { Route, IndexRoute } from 'react-router'

import Home from './Home'
import {Index as IndexNotices, Admin as AdminNotices} from './notices'
import {New as NewLeaveWord, Admin as AdminLeaveWords} from './leave-words'
import Dashboard from './Dashboard'
import Root from '../../Dashboard'
import {
  Info as SiteInfo,
  Author as SiteAuthor,
  Seo as SiteSeo,
  Smtp as SiteSmtp,
  Status as SiteStatus,
  Locales, Users
} from './admin'

export default {
  navLinks: [
    {href:'/notices', label:'site.notices.index.title'},
  ],
  dashboard: <Dashboard key='site.dashboard'/>,
  routes: [
    (<IndexRoute key="site.home" component={Home} />),
    (<Route key="site.notices" path="/notices" component={IndexNotices} />),
    (<Route key="site.leave-words.new" path="/leave-words/new" component={NewLeaveWord} />),
    (<Route key="site.admin" path="/admin" component={Root}>
      <Route path="site/info" component={SiteInfo}/>
      <Route path="site/author" component={SiteAuthor}/>
      <Route path="site/seo" component={SiteSeo}/>
      <Route path="site/smtp" component={SiteSmtp}/>
      <Route path="site/status" component={SiteStatus}/>
      <Route path="locales" component={Locales}/>
      <Route path="users" component={Users}/>
      <Route path="leave-words" component={AdminLeaveWords}/>
      <Route path="notices" component={AdminNotices}/>
    </Route>)
  ]
}
