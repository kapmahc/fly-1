import React from 'react';
import { Route } from 'react-router'

import {Index, SignIn,SignUp,Confirm,Unlock,ForgotPassword,ResetPassword, Layout} from './non-sign-in'
import {Info, Logs, ChangePassword} from './must-sign-in'
import {Dashboard as Attachments} from './attachments'
import Dashboard from './Dashboard'
import Root from '../../Dashboard'

export default {
  navLinks: [
    {href:'/users', label:'auth.users.index.title'},
  ],
  dashboard: <Dashboard key="auth.dashboard"/>,
  routes: [
    (<Route key="auth.users.index" path="/users" component={Index} />),
    (<Route key="auth.non-sign-in" path="/users" component={Layout}>
      <Route path="sign-in" component={SignIn}/>
      <Route path="sign-up" component={SignUp}/>
      <Route path="confirm" component={Confirm}/>
      <Route path="unlock" component={Unlock}/>
      <Route path="forgot-password" component={ForgotPassword}/>
      <Route path="reset-password/:token" component={ResetPassword}/>
    </Route>),
    (<Route key="auth.attachments" path="attachments" component={Attachments}/>),
    (<Route key="auth-must-sign-in" path="/users" component={Root}>
      <Route path="info" component={Info}/>
      <Route path="logs" component={Logs}/>
      <Route path="change-password" component={ChangePassword}/>
    </Route>),
  ],
}
