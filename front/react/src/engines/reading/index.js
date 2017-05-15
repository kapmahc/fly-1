import React from 'react';
import { Route } from 'react-router'

import Dashboard from './Dashboard'
import {Show as ShowBook, Index as IndexBooks} from './books'

export default {
  navLinks: [
    {href:'/reading/books', label:'reading.books.index.title'},
  ],
  dashboard: <Dashboard key="reading.dashboard"/>,
  routes: [
    (<Route key="reading.engine" path="reading">
      <Route path="books" component={IndexBooks}/>,
      <Route path="books/:id" component={ShowBook}/>
    </Route>)
  ],
}
