import React, { Component, PropTypes } from 'react'
import { connect } from 'react-redux'

import Header from './components/Header'
import Footer from './components/Footer'

import {refresh, signIn} from './actions'
import {get} from './ajax'
import {TOKEN} from './constants'

class Widget extends Component {
  componentDidMount() {
    const { refresh, signIn } = this.props
    var token = sessionStorage.getItem(TOKEN)
    if (token){
      signIn(token)
    }
    get('/site/info').then(
      rst => {
        document.title = rst.title;
        refresh(rst);
      }
    );
  }
  render() {
    const {children} = this.props;
    return (
      <div>
        <Header />
        <div className="container">
          {children}
          <Footer/>
        </div>
      </div>
    );
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  refresh: PropTypes.func.isRequired,
  signIn: PropTypes.func.isRequired
}

export default connect(
  state => ({}),
  {refresh, signIn},
)(Widget);
