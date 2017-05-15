import React, { Component } from 'react';
import {ListGroup, ListGroupItem, FormGroup,ControlLabel, FormControl,  HelpBlock, Button} from 'react-bootstrap';
import i18next from 'i18next';

import {get, post} from '../../ajax'


export class Info extends Component{
  constructor(props){
    super(props)
    this.state = {
      name:'',
      home:'',
      logo:'',
      email: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get('/users/info').then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
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
    data.append('home', this.state.home)
    data.append('logo', this.state.logo)
    post('/users/info', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('auth.users.info.title')}</h3>
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
            disabled
          />
        </FormGroup>
        <FormGroup controlId="home">
          <ControlLabel>{i18next.t('auth.attributes.user.home')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.home}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="logo">
          <ControlLabel>{i18next.t('auth.attributes.user.logo')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.logo}
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



export class Logs extends Component {
  constructor(props){
    super(props)
    this.state = {logs:[]}
    get('/users/logs').then(
      function(rst){
        this.setState({logs:rst})
      }.bind(this)
    );
  }
  render(){
    return (<div className="col-md-offset-1 col-md-10">
    <h3>{i18next.t('auth.users.logs.title')}</h3>
    <hr/>
    <ListGroup>
    {
      this.state.logs.map((l,i)=>(
        <ListGroupItem key={i}>[{l.ip}]{l.createdAt}: {l.message}</ListGroupItem>
      ))
    }
    </ListGroup>
    </div>)
  }
}



export class ChangePassword extends Component{
  constructor(props){
    super(props)
    this.state = {
      currentPassword:'',
      newPassword:'',
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
    data.append('currentPassword', this.state.currentPassword)
    data.append('newPassword', this.state.newPassword)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    post('/users/change-password', data)
      .then(function(rst){
        alert(i18next.t('success'))
        this.setState({currentPassword:'', newPassword:'', passwordConfirmation:''})
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div>
      <h3>{i18next.t('auth.users.change-password.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="currentPassword">
          <ControlLabel>{i18next.t('attributes.currentPassword')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.currentPassword}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="newPassword">
          <ControlLabel>{i18next.t('attributes.newPassword')}</ControlLabel>
          <FormControl
            type="password"
            value={this.state.newPassword}
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
