import React,{ PropTypes } from 'react'
import { connect } from 'react-redux'
import {NavDropdown, MenuItem} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'

const Widget = ({user}) =>{
  return user.isAdmin ?
  (<NavDropdown title={i18next.t('site.dashboard.title')} id="site-dashboard-dropdown">
    {['info', 'author', 'seo', 'smtp', 'status'].map((k, i)=>(<LinkContainer key={i}to={`/admin/site/${k}`}>
      <MenuItem>{i18next.t(`site.admin.${k}.title`)}</MenuItem>
    </LinkContainer>))}
    <MenuItem divider />
    <LinkContainer to="/admin/locales">
      <MenuItem>{i18next.t('site.admin.locales.index.title')}</MenuItem>
    </LinkContainer>
    <LinkContainer to="/admin/users">
      <MenuItem>{i18next.t('site.admin.users.index.title')}</MenuItem>
    </LinkContainer>
    <LinkContainer to="/admin/notices">
      <MenuItem>{i18next.t('site.notices.index.title')}</MenuItem>
    </LinkContainer>
    <LinkContainer to="/admin/leave-words">
      <MenuItem>{i18next.t('site.leave-words.index.title')}</MenuItem>
    </LinkContainer>
  </NavDropdown>) : null
}


Widget.propTypes = {
  user: PropTypes.object.isRequired
}

export default connect(
  state => ({user: state.currentUser})
)(Widget);
