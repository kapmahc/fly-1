import React, {Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {ListGroup, Table, ListGroupItem,
  Checkbox, HelpBlock, FormGroup,ControlLabel, FormControl,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';

import {get, post, _delete} from '../../ajax'


class InfoW extends Component{
  constructor(props){
    super(props)
    this.state = {
      title:'',
      subTitle:'',
      keywords:'',
      description: '',
      copyright: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentWillReceiveProps(nextProps) {
    const {info} = nextProps
    this.setState(info)
  }
  componentDidMount(){
    const {info} = this.props
    if(info){
      this.setState(info)
    }
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('title', this.state.title)
    data.append('subTitle', this.state.subTitle)
    data.append('keywords', this.state.keywords)
    data.append('description', this.state.description)
    data.append('copyright', this.state.copyright)
    post('/admin/site/info', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.info.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="title">
          <ControlLabel>{i18next.t('site.attributes.title')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.title}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="subTitle">
          <ControlLabel>{i18next.t('site.attributes.subTitle')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.subTitle}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="keywords">
          <ControlLabel>{i18next.t('site.attributes.keywords')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.keywords}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="description">
          <ControlLabel>{i18next.t('site.attributes.description')}</ControlLabel>
          <FormControl
            value={this.state.description}
            onChange={this.handleChange}
            componentClass="textarea" />
        </FormGroup>
        <FormGroup controlId="copyright">
          <ControlLabel>{i18next.t('site.attributes.copyright')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.copyright}
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

InfoW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Info = connect(
  state => ({info:state.siteInfo})
)(InfoW);

class AuthorW extends Component{
  constructor(props){
    super(props)
    this.state = {
      name:'',
      email: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentWillReceiveProps(nextProps) {
    const {info} = nextProps
    this.setState(info.author)
  }
  componentDidMount(){
    const {info} = this.props
    if(info.author){
      this.setState(info.author)
    }
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
    post('/admin/site/author', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.author.title')}</h3>
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
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}

AuthorW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Author = connect(
  state => ({info:state.siteInfo})
)(AuthorW);

class SeoW extends Component{
  constructor(props){
    super(props)
    this.state = {
      googleVerifyCode:'',
      baiduVerifyCode: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get('/admin/site/seo').then(
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
    data.append('googleVerifyCode', this.state.googleVerifyCode)
    data.append('baiduVerifyCode', this.state.baiduVerifyCode)
    post('/admin/site/seo', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    const {info} = this.props
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.seo.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="googleVerifyCode">
          <ControlLabel>{i18next.t('site.attributes.googleVerifyCode')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.googleVerifyCode}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="baiduVerifyCode">
          <ControlLabel>{i18next.t('site.attributes.baiduVerifyCode')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.baiduVerifyCode}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
      <br/>
      <ListGroup>
        {info.languages.map((l)=>(`rss-${l}.atom`)).concat(["sitemap.xml.gz", "robots.txt", `google${this.state.googleVerifyCode}.html`, `baidu_verify_${this.state.baiduVerifyCode}.html`]).map((f,i)=>(
          <ListGroupItem key={i} target="_blank" href={`/${f}`}>{f}</ListGroupItem>
        ))}
      </ListGroup>
    </div>)
  }
}

SeoW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Seo = connect(
  state => ({info:state.siteInfo})
)(SeoW);

export class Smtp extends Component{
  constructor(props){
    super(props)
    this.state = {
      host:'',
      port: 465,
      ssl:true,
      username: '',
      password: '',
      passwordConfirmation:'',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get('/admin/site/smtp').then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  handleChange(e) {
    var data = {};
    var t = e.target;
    data[t.id] = t.type === 'checkbox' ? t.checked : t.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('host', this.state.host)
    data.append('port', this.state.port)
    data.append('ssl', this.state.ssl)
    data.append('username', this.state.username)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)
    post('/admin/site/smtp', data)
      .then(function(rst){
        alert(i18next.t('success'))
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.smtp.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="host">
          <ControlLabel>{i18next.t('attributes.host')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.host}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="port">
          <ControlLabel>{i18next.t('attributes.port')}</ControlLabel>
          <FormControl value={this.state.port} onChange={this.handleChange} componentClass="select" placeholder="select">
            {[25,465,587].map((p,i)=>(<option value={p} key={i}>{p}</option>))}
          </FormControl>
        </FormGroup>
        <Checkbox id="ssl" checked={this.state.ssl} onChange={this.handleChange}>
          {i18next.t('attributes.ssl')}
        </Checkbox>
        <FormGroup controlId="username">
          <ControlLabel>{i18next.t('site.admin.smtp.sender')}</ControlLabel>
          <FormControl
            type="email"
            value={this.state.username}
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


export class Status extends Component{
  constructor(props){
    super(props)
    this.state = {
      host:'',
      port: 465,
      ssl:true,
      username: '',
      password: '',
      passwordConfirmation:'',
    }
  }
  componentDidMount() {
    get('/admin/site/status').then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render() {
    return (<div className="row">
      {["os", "database", "jobs"].map((k,i)=>(<div key={i} className="col-md-6">
        <h4>{i18next.t(`site.admin.status.${k}`)}</h4>
        <hr/>
        <pre><code>{JSON.stringify(this.state[k], null, 2)}</code></pre>
      </div>))}
      <div className="col-md-6">
        <h4>{i18next.t('site.admin.status.cache')}</h4>
        <hr/>
        <pre><code>{this.state.cache}</code></pre>
      </div>
    </div>)
  }
}


export class Locales extends Component{
  constructor(props){
    super(props)
    this.state = {
      code:'',
      message: '',
      items:[]
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
    this.handleNew = this.handleNew.bind(this);
    this.handleRemove = this.handleRemove.bind(this);
  }
  componentDidMount() {
    get('/admin/locales').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  handleChange(e) {
    var data = {};
    var t = e.target;
    data[t.id] = t.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('code', this.state.code)
    data.append('message', this.state.message)
    post('/admin/locales', data)
      .then(function(rst){
        alert(i18next.t('success'))
        var items = this.state.items.filter((l, _) => l.code !== rst.code)
        items.unshift(rst)
        this.setState({code:'', message:'', items:items})
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  }
  handleNew() {
    this.setState({code:'', message:''})
  }
  handleEdit(l) {
    this.setState({code:l.code, message:l.message})
  }
  handleRemove(id) {
    if(confirm(i18next.t('are-you-sure'))){
      _delete(`/admin/locales/${id}`)
        .then(function(rst){
          alert(i18next.t('success'))
          this.setState({
            items: this.state.items.filter((l, _) => l.id !== id)
          });
        }.bind(this))
        .catch((err) => {
          alert(err)
        })
    }
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.locales.index.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="code">
          <ControlLabel>{i18next.t('site.attributes.locale.code')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.code}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="message">
          <ControlLabel>{i18next.t('site.attributes.locale.message')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.message}
            onChange={this.handleChange}
          />
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
      <br/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('site.attributes.locale.code')}</th>
            <th>{i18next.t('site.attributes.locale.message')}</th>
            <th width="12%">
              {i18next.t('buttons.manage')}
              <Button bsStyle="success" bsSize="sm" onClick={this.handleNew}>{i18next.t('buttons.new')}</Button>
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((l,i)=>(<tr key={i}>
            <td>{l.code}</td>
            <td>
              <pre><code>
                {l.message}
              </code></pre>
            </td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <Button bsStyle="warning" onClick={()=>this.handleEdit(l)}>{i18next.t('buttons.edit')}</Button>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(l.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
            </td>
          </tr>))}
        </tbody>
      </Table>
    </div>)
  }
}



export class Users extends Component{
  constructor(props){
    super(props)
    this.state = {
      users:[]
    }
  }
  componentDidMount() {
    get('/admin/users').then(
      function(rst){
        this.setState({users:rst})
      }.bind(this)
    );
  }
  render() {
    return (<div className="col-md-offset-1 col-md-10">
      <h3>{i18next.t('site.admin.users.index.title')}</h3>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.user')}</th>
            <th>{i18next.t('auth.attributes.user.lastSignIn')}</th>
            <th>{i18next.t('auth.attributes.user.currentSignIn')}</th>
          </tr>
        </thead>
        <tbody>
          {this.state.users.map((u,i)=>(<tr key={i}>
            <td>{u.name}&lt;{u.email}&gt;[{u.signInCount}]</td>
            <td>{u.lastSignInIP} {u.lastSignInAt}</td>
            <td>{u.currentSignInIP} {u.currentSignInAt}</td>
          </tr>))}
        </tbody>
      </Table>
    </div>)
  }
}
