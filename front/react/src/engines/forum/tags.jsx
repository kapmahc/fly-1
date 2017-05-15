import React, {Component } from 'react'
import {Table, FormGroup,ControlLabel, FormControl, Thumbnail,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'

import {get, post, _delete} from '../../ajax'
import {List as Articles} from './articles'

export class Index extends Component{
  constructor(props){
    super(props)
    this.state = {
      items: []
    }
  }
  componentDidMount() {
    get('/forum/tags').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  render() {
    return (<div className="row">
      <h3>{i18next.t('forum.tags.index.title')}</h3>
      <hr/>
      {this.state.items.map((t,i)=>(<div key={i} className="col-md-2">
      <Thumbnail>
        <h3>{t.name}</h3>
        <p>
          <LinkContainer target="_blank" to={`/forum/tags/${t.id}`}>
            <Button bsStyle="primary">{i18next.t("buttons.view")}</Button>
          </LinkContainer>
        </p>
      </Thumbnail>
      </div>))}
    </div>)
  }
}

export class Show extends Component{
  constructor(props){
    super(props)
    this.state = {
      name: '',
      articles:[]
    }
  }
  componentDidMount() {
    const {params} = this.props
    get(`/forum/tags/${params.id}`).then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render() {
    return (<div className="row">
      <h3>{this.state.name}</h3>
      <hr/>
      <Articles items={this.state.articles || []}/>
    </div>)
  }
}

export class Dashboard extends Component{
  constructor(props){
    super(props)
    this.state = {
      id: null,
      name: '',
      title: null,
      items: []
    }
    this.handleRemove = this.handleRemove.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
    this.handleNew = this.handleNew.bind(this);
  }
  componentDidMount() {
    get('/forum/tags').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  handleChange(e) {
    var data = {};
    data[e.target.id] = e.target.value;
    this.setState(data);
  }
  handleNew(e) {
    this.setState({id:null, title:i18next.t('buttons.new')});
  }
  handleEdit(t) {
    this.setState({id: t.id, name: t.name, title:`${i18next.t('buttons.edit')} [${t.id}]`});
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('name', this.state.name)
    var id = this.state.id
    post(id ? `/forum/tags/${id}`:'/forum/tags', data)
      .then(function(rst){
        alert(i18next.t('success'))
        var items = this.state.items.filter((t, _) => t.id !== id)
        items.unshift(rst)
        this.setState({
          title: null,
          name:'',
          id:null,
          items: items
        })
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  }
  handleRemove(id) {
    if(confirm(i18next.t('are-you-sure'))){
      _delete(`/forum/tags/${id}`)
        .then(function(rst){
          alert(i18next.t('success'))
          this.setState({
            items: this.state.items.filter((t, _) => t.id !== id)
          });
        }.bind(this))
        .catch((err) => {
          alert(err)
        })
    }
  }
  render() {

    return (<div className="row">
    { this.state.title ?
        (<div className="col-md-10 col-md-offset-1">
          <h3>{this.state.title}</h3>
          <hr/>
          <form onSubmit={this.handleSubmit}>
            <FormGroup controlId="name">
              <ControlLabel>{i18next.t('attributes.name')}</ControlLabel>
              <FormControl
                value={this.state.name}
                onChange={this.handleChange} />
            </FormGroup>
            <Button type="submit" bsStyle="primary">
              {i18next.t('buttons.submit')}
            </Button>
          </form>
        </div>) : <br/>
      }
      <div className="col-md-10 col-md-offset-1">
      <h3>{i18next.t('forum.dashboard.tags.title')}</h3>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.updatedAt')}</th>
            <th>{i18next.t('attributes.name')}</th>
            <th>
              {i18next.t('buttons.manage')}
              <Button bsStyle="success" bsSize="sm" onClick={this.handleNew}>{i18next.t('buttons.new')}</Button>
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((t,i)=>(<tr key={i}>
            <td>{t.updatedAt}</td>
            <td>{t.name}</td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <Button bsStyle="warning" onClick={()=>this.handleEdit(t)}>{i18next.t('buttons.edit')}</Button>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(t.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
          </td>
          </tr>))}
        </tbody>
      </Table>
      </div>
    </div>)
  }
}
