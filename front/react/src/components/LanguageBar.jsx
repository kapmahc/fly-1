import React,{PropTypes} from 'react';
import { connect } from 'react-redux'
import { NavDropdown, MenuItem } from 'react-bootstrap';
import {LinkContainer} from 'react-router-bootstrap'
import i18next from 'i18next';

const Widget = ({info}) => (
  <NavDropdown title={i18next.t('language-bar.switch')} id="header-language-bar">
    {info.languages.map((lng, i) => (
      <LinkContainer target="_blank" key={i} to={{ pathname: '/', query: { locale: lng } }}>
        <MenuItem disabled={i18next.language === lng}>
          {i18next.t(`languages.${lng}`)}
        </MenuItem>
      </LinkContainer>
    ))}
  </NavDropdown>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired
}

export default connect(
  state => ({info: state.siteInfo})
)(Widget)
