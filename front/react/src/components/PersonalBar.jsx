import React,{Component, PropTypes} from 'react';
import { connect } from 'react-redux'
import {browserHistory} from 'react-router'
import { NavDropdown, MenuItem } from 'react-bootstrap';
import {LinkContainer} from 'react-router-bootstrap'
import i18next from 'i18next';

import {signOut} from '../actions'
import {DASHBOARD} from '../constants'

class Widget extends Component{
  constructor(props){
    super(props)
    this.state = {}
    this.handleSignOut = this.handleSignOut.bind(this);
  }
  handleSignOut(e) {
    if(confirm(i18next.t('are-you-sure'))){
      browserHistory.push('/')
      sessionStorage.clear();
      const {signOut} = this.props
      signOut()
    }
  }
  render(){
    const {user} = this.props
    return (
      user.uid ?
      <NavDropdown title={i18next.t("personal-bar.welcome", {name:user.name})} id="header-personal-bar">
        <LinkContainer to={DASHBOARD}>
          <MenuItem>{i18next.t("personal-bar.dashboard")}</MenuItem>
        </LinkContainer>
        <MenuItem divider />
        <MenuItem onClick={this.handleSignOut}>{i18next.t("personal-bar.sign-out")}</MenuItem>
      </NavDropdown> :
      <NavDropdown title={i18next.t("personal-bar.sign-in-or-up")} id="header-personal-bar">
        {['sign-in', 'sign-up', 'confirm', 'unlock', 'forgot-password'].map((k,i)=>(
          <LinkContainer key={i} to={`/users/${k}`}>
            <MenuItem>{i18next.t(`auth.users.${k}.title`)}</MenuItem>
          </LinkContainer>
        ))}
      </NavDropdown>
    )
  }
}

Widget.propTypes = {
  user: PropTypes.object.isRequired,
  signOut: PropTypes.func.isRequired
}

export default connect(
  state => ({user: state.currentUser}),
  {signOut},
)(Widget)
