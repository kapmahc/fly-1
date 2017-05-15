import React from 'react';
import { Route } from 'react-router'

import Dashboard from './Dashboard'
import Show from './Show'

export default {
  navLinks: [
    {href:'/blogs/', label:'blog.index.title'},
  ],
  dashboard: <Dashboard key="blog.dashboard"/>,
  routes: [
    (<Route key="blog.show" path="blogs/*" component={Show} />)
  ],
}
