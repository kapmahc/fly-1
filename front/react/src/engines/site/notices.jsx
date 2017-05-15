import React, {Component } from 'react'
import {Table, HelpBlock, FormGroup,ControlLabel, FormControl,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';

import {get, post, _delete} from '../../ajax'
import {timeago} from '../../utils'
import Markdown from '../../components/Markdown'

export class Index extends Component{
  constructor(props){
    super(props)
    this.state = {
      items:[]
    }
  }
  componentDidMount() {
    get('/notices').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  render() {
    return (<div className="row">
      <h3>{i18next.t('site.notices.index.title')}</h3>
      <hr/>
          {this.state.items.map((n,i)=>(<blockquote key={i}>
            <Markdown body={n.body}/>
            <footer><cite>{timeago(n.updatedAt)}</cite></footer>
          </blockquote>))}
    </div>)
  }
}

export class Admin extends Component{
  constructor(props){
    super(props)
    this.state = {
      id: null,
      body: '',
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
    get('/notices').then(
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
  handleEdit(n) {
    this.setState({id: n.id, body: n.body, title:`${i18next.t('buttons.edit')} [${n.id}]`});
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('type', 'markdown')
    data.append('body', this.state.body)
    var id = this.state.id
    post(id ? `/notices/${id}`:'/notices', data)
      .then(function(rst){
        alert(i18next.t('success'))
        var items = this.state.items.filter((n, _) => n.id !== id)
        items.unshift(rst)
        this.setState({
          title: null,
          body:'',
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
      _delete(`/notices/${id}`)
        .then(function(rst){
          alert(i18next.t('success'))
          this.setState({
            items: this.state.items.filter((n, _) => n.id !== id)
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
            <FormGroup controlId="body">
              <ControlLabel>{i18next.t('attributes.body')}</ControlLabel>
              <FormControl
                value={this.state.body}
                onChange={this.handleChange}
                rows={8}
                componentClass="textarea" />
              <HelpBlock>{i18next.t('helps.markdown')}</HelpBlock>
            </FormGroup>
            <Button type="submit" bsStyle="primary">
              {i18next.t('buttons.submit')}
            </Button>
          </form>
        </div>) : <br/>
      }
      <div className="col-md-10 col-md-offset-1">
      <h3>{i18next.t('site.notices.index.title')}</h3>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.updatedAt')}</th>
            <th>{i18next.t('attributes.body')}</th>
            <th>
              {i18next.t('buttons.manage')}
              <Button bsStyle="success" bsSize="sm" onClick={this.handleNew}>{i18next.t('buttons.new')}</Button>
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((n,i)=>(<tr key={i}>
            <td>{n.updatedAt}</td>
            <td>
              <pre><code>
                {n.body}
              </code></pre>
            </td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <Button bsStyle="warning" onClick={()=>this.handleEdit(n)}>{i18next.t('buttons.edit')}</Button>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(n.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
          </td>
          </tr>))}
        </tbody>
      </Table>
      </div>
    </div>)
  }
}
