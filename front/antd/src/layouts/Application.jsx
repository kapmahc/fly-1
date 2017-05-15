import React from 'react';
import PropTypes from 'prop-types';
import Header from '../components/Header'
import Footer from '../components/Footer'

const Widget = ({children}) => (
  <div>
    <Header />
    {children}
    <Footer />
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
