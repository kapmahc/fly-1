import React, { Component, PropTypes } from 'react';
import {Nav} from 'react-bootstrap';
import { connect } from 'react-redux'

import engines from './engines'
import {Forbidden} from './components/alerts'

class Widget extends Component {
  render() {
    const {children, user} = this.props
    return user.uid ? (
      <div className="row">
        <Nav bsStyle="tabs">
          {engines.dashboard}
        </Nav>
        {children}
      </div>
    ) :
    <div className="row">
      <br/>
      <Forbidden />
    </div>
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  user: PropTypes.object.isRequired
}


export default connect(
  state => ({user: state.currentUser})
)(Widget);
