import React, {PropTypes} from 'react';
import marked from 'marked'

const Widget = ({body}) => (<p dangerouslySetInnerHTML={{__html: marked(body)}}></p>)

Widget.propTypes = {
  body: PropTypes.string.isRequired
}
export default Widget
