import React from 'react';
import {Alert} from 'react-bootstrap';
import i18next from 'i18next';

export const Forbidden = () => (<Alert bsStyle="danger">
  <strong>{i18next.t('flashs.error')}</strong> {i18next.t('errors.forbidden')}
</Alert>)
