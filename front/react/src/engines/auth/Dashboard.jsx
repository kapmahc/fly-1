import React from 'react';
import {NavDropdown, MenuItem} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'

const Widget = () => (<NavDropdown title={i18next.t('auth.dashboard.title')} id="auth-dashboard-dropdown">
  {['logs', 'info', 'change-password'].map((k, i)=>(<LinkContainer key={i}to={`/users/${k}`}>
    <MenuItem>{i18next.t(`auth.users.${k}.title`)}</MenuItem>
  </LinkContainer>))}
  <LinkContainer to={'/attachments'}>
    <MenuItem>{i18next.t('auth.attachments.index.title')}</MenuItem>
  </LinkContainer>
</NavDropdown>)

export default Widget
