import React, { Component, PropTypes } from 'react';
import { connect } from 'react-redux'
import {Link, browserHistory} from 'react-router'
import { Thumbnail, FormGroup,ControlLabel, FormControl,  HelpBlock, Button } from 'react-bootstrap';
import i18next from 'i18next';

import {post, get} from '../../ajax'
import {signIn} from '../../actions'
import {TOKEN, DASHBOARD} from '../../constants'

export class Index extends Component{
  constructor(props){
    super(props)
    this.state = {users:[]}
  }
  componentDidMount() {
    get('/users').then(
      function(rst){
        this.setState({users:rst})
      }.bind(this)
    );
  }
  render(){
    return (<div className="row">
    <h3>{i18next.t('auth.users.index.title')}</h3>
    <hr/>
    {this.state.users.map((u,i)=>(<div key={i} className="col-md-3">
      <Thumbnail src={u.logo} alt="242x200">
        <h3>{u.name}</h3>
        <p></p>
      </Thumbnail>
    </div>))}
    </div>)
  }
}

class SignInW extends Component{
  constructor(props){
    super(props)
    this.state = {
      email:'',
      password:'',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    const {signIn} = this.props

    var data = new FormData()
    data.append('email', this.state.email)
    data.append('password', this.state.password)
    post('/users/sign-in', data)
      .then(function(rst){
        sessionStorage.setItem(TOKEN, rst.token)
        signIn(rst.token)
        browserHistory.push(DASHBOARD)
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div>
      <h3>{i18next.t('auth.users.sign-in.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="email">
          <ControlLabel>{i18next.t('attributes.email')}</ControlLabel>
          <FormControl
            type="email"
            value={this.state.email}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="password">
          <ControlLabel>{i18next.t('attributes.password')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.password}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}


SignInW.propTypes = {
  signIn: PropTypes.func.isRequired
}

export const SignIn = connect(
  state => ({}),
  {signIn},
)(SignInW);


export class SignUp extends Component{
  constructor(props){
    super(props)
    this.state = {
      name:'',
      email:'',
      password:'',
      passwordConfirmation:'',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('name', this.state.name)
    data.append('email', this.state.email)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    post('/users/sign-up', data)
      .then(function(rst){
        alert(i18next.t('auth.messages.email-for-confirm'))
        browserHistory.push('/users/sign-in')
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div>
      <h3>{i18next.t('auth.users.sign-up.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="name">
          <ControlLabel>{i18next.t('attributes.fullName')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.name}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="email">
          <ControlLabel>{i18next.t('attributes.email')}</ControlLabel>
          <FormControl
            type="email"
            value={this.state.email}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="password">
          <ControlLabel>{i18next.t('attributes.password')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.password}
            onChange={this.handleChange}
          />
          <HelpBlock>{i18next.t('helps.password')}</HelpBlock>
        </FormGroup>
        <FormGroup controlId="passwordConfirmation">
          <ControlLabel>{i18next.t('attributes.passwordConfirmation')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.passwordConfirmation}
            onChange={this.handleChange}
          />
          <HelpBlock>{i18next.t('helps.passwordConfirmation')}</HelpBlock>
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}


export class EmailForm extends Component{
  constructor(props){
    super(props)
    this.state = {
      email:'',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    const {action} = this.props
    var data = new FormData()
    data.append('email', this.state.email)
    post(`/users/${action}`, data)
      .then(function(rst){
        alert(i18next.t(`auth.messages.email-for-${action}`))
        browserHistory.push('/users/sign-in')
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    const {action} = this.props
    return (<div>
      <h3>{i18next.t(`auth.users.${action}.title`)}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="email">
          <ControlLabel>{i18next.t('attributes.email')}</ControlLabel>
          <FormControl
            type="email"
            value={this.state.email}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}


EmailForm.propTypes = {
  action: PropTypes.string.isRequired
}

export const Confirm = () => (
  <EmailForm action="confirm" />
)

export const Unlock = () => (
  <EmailForm action="unlock" />
)

export const ForgotPassword = () => (
  <EmailForm action="forgot-password" />
)


export class ResetPassword extends Component{
  constructor(props){
    super(props)
    this.state = {
      password:'',
      passwordConfirmation:'',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    post(`/users/reset-password/${this.props.params.token}`, data)
      .then(function(rst){
        alert(rst.message)
        browserHistory.push('/users/sign-in')
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div>
      <h3>{i18next.t('auth.users.reset-password.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="password">
          <ControlLabel>{i18next.t('attributes.password')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.password}
            onChange={this.handleChange}
          />
          <HelpBlock>{i18next.t('helps.password')}</HelpBlock>
        </FormGroup>
        <FormGroup controlId="passwordConfirmation">
          <ControlLabel>{i18next.t('attributes.passwordConfirmation')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.passwordConfirmation}
            onChange={this.handleChange}
          />
          <HelpBlock>{i18next.t('helps.passwordConfirmation')}</HelpBlock>
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}


export const Layout = ({children}) => (
  <div className="row">
    <div className="col-md-offset-1 col-md-9">
    {children}
    <br/>
    <ul>
      {['sign-in', 'sign-up', 'confirm', 'unlock', 'forgot-password'].map((k,i)=>(
        <li key={i}>
          <Link to={`/users/${k}`}>
            {i18next.t(`auth.users.${k}.title`)}
          </Link>
        </li>
      ))}
      <li><Link to="/leave-words/new">{i18next.t("site.leave-words.new.title")}</Link></li>
    </ul>
    </div>
  </div>
)
