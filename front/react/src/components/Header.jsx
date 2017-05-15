import React, {PropTypes} from 'react';
import { Navbar, Nav, NavItem } from 'react-bootstrap';
import { connect } from 'react-redux'
import { Link } from 'react-router'
import i18next from 'i18next';
import {IndexLinkContainer, LinkContainer} from 'react-router-bootstrap'

import PersonalBar from './PersonalBar'
import LanguageBar from './LanguageBar'
import {DASHBOARD} from '../constants'
import engines from '../engines'

const Widget = ({info, user}) => (
  <Navbar fixedTop fluid inverse collapseOnSelect>
    <Navbar.Header>
      <Navbar.Brand>
        <Link to="/">{info.subTitle}</Link>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        <IndexLinkContainer to='/'>
          <NavItem>{i18next.t('header.home')}</NavItem>
        </IndexLinkContainer>
        {
          engines.navLinks.map((l,i)=>(<LinkContainer key={i} to={l.href}>
            <NavItem>{i18next.t(l.label)}</NavItem>
          </LinkContainer>))
        }
        {
          user.uid ?
          (<LinkContainer to={DASHBOARD}>
            <NavItem>{i18next.t("header.dashboard")}</NavItem>
          </LinkContainer>) :
          "&nbsp;"
        }
      </Nav>
      <Nav pullRight>
        <LanguageBar/>
        <PersonalBar/>
      </Nav>
    </Navbar.Collapse>
  </Navbar>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired,
  user: PropTypes.object.isRequired,
}

export default connect(
   state => ({info: state.siteInfo, user: state.currentUser})
)(Widget);
