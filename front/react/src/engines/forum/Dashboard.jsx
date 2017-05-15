import React,{ PropTypes } from 'react'
import { connect } from 'react-redux'
import {NavDropdown, MenuItem} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'

const Widget = ({user}) =>{
  return (<NavDropdown title={i18next.t('forum.dashboard.title')} id="forum-dashboard-dropdown">
    <LinkContainer to="/forum/articles/new">
      <MenuItem>{i18next.t('forum.articles.new.title')}</MenuItem>
    </LinkContainer>
    <LinkContainer to="/forum/dashboard/articles">
      <MenuItem>{i18next.t('forum.dashboard.articles.title')}</MenuItem>
    </LinkContainer>
    <LinkContainer to="/forum/dashboard/comments">
      <MenuItem>{i18next.t('forum.dashboard.comments.title')}</MenuItem>
    </LinkContainer>
    {user.isAdmin ? (<LinkContainer to="/forum/dashboard/tags">
      <MenuItem>{i18next.t('forum.dashboard.tags.title')}</MenuItem>
    </LinkContainer>) : null}
  </NavDropdown>)
}


Widget.propTypes = {
  user: PropTypes.object.isRequired
}

export default connect(
  state => ({user: state.currentUser})
)(Widget);
