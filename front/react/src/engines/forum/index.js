import React from 'react';
import { Route } from 'react-router'

import {Index as IndexTags, Show as ShowTag, Dashboard as DsTags} from './tags'
import {Index as IndexArticles, New as NewArticle, Show as ShowArticle, Edit as EditArticle, Dashboard as DsArticles} from './articles'
import {Index as IndexComments, New as NewComment, Edit as EditComment, Dashboard as DsComments} from './comments'
import Dashboard from './Dashboard'
import Root from '../../Dashboard'

export default {
  navLinks: [
    {href:'/forum/articles', label:'forum.articles.index.title'},
    {href:'/forum/tags', label:'forum.tags.index.title'},
    {href:'/forum/comments', label:'forum.comments.index.title'},
  ],
  dashboard: <Dashboard key='forum.dashboard'/>,
  routes: [
    (<Route key="forum.engine" path="forum">
      <Route path="articles" component={IndexArticles}/>
      <Route path="articles/new" component={NewArticle}/>
      <Route path="articles/:id" component={ShowArticle}/>
      <Route path="articles/:id/edit" component={EditArticle}/>
      <Route path="tags" component={IndexTags}/>
      <Route path="tags/:id" component={ShowTag}/>
      <Route path="comments" component={IndexComments}/>
      <Route path="comments/new" component={NewComment}/>
      <Route path="comments/:id/edit" component={EditComment}/>
      <Route path="dashboard" component={Root}>
        <Route path="comments" component={DsComments}/>
        <Route path="articles" component={DsArticles}/>
        <Route path="tags" component={DsTags}/>
      </Route>
    </Route>)
  ]
}
