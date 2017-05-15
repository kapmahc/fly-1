import React from 'react';
import PropTypes from 'prop-types';

import Layout from '../../../layouts/Application'

const Widget = ({children}) => (
  <Layout>
    <div>non sign in</div>
    {children}
  </Layout>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
