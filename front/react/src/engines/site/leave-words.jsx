import React, {Component } from 'react'
import {Table, HelpBlock, FormGroup,ControlLabel, FormControl, Button} from 'react-bootstrap';
import i18next from 'i18next';

import {get, post, _delete} from '../../ajax'


export class Admin extends Component{
  constructor(props){
    super(props)
    this.state = {
      items:[]
    }
    this.handleRemove = this.handleRemove.bind(this);
  }
  componentDidMount() {
    get('/leave-words').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  handleRemove(id) {
    if(confirm(i18next.t('are-you-sure'))){
      _delete(`/leave-words/${id}`)
        .then(function(rst){
          alert(i18next.t('success'))
          this.setState({
            items: this.state.items.filter((lw, _) => lw.id !== id)
          });
        }.bind(this))
        .catch((err) => {
          alert(err)
        })
      }
  }
  render() {
    return (<div className="row">
      <h3>{i18next.t('site.leave-words.index.title')}</h3>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.createdAt')}</th>
            <th>{i18next.t('attributes.body')}</th>
            <th>{i18next.t('buttons.manage')}</th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((lw,i)=>(<tr key={i}>
            <td>{lw.createdAt}</td>
            <td>
              <pre><code>
                {lw.body}
              </code></pre>
            </td>
            <td>
              <Button bsStyle="danger" bsSize="sm" onClick={()=>this.handleRemove(lw.id)}>{i18next.t('buttons.remove')}</Button>
            </td>
          </tr>))}
        </tbody>
      </Table>
    </div>)
  }
}


export class New extends Component{
  constructor(props){
    super(props)
    this.state = {
      body:''
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
    data.append('type', 'markdown')
    data.append('body', this.state.body)
    post('/leave-words', data)
      .then(function(rst){
        alert(i18next.t('success'))
        this.setState({body:''})
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="row">
      <h3>{i18next.t('site.leave-words.new.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="body">
          <ControlLabel>{i18next.t('attributes.body')}</ControlLabel>
          <FormControl
            value={this.state.body}
            onChange={this.handleChange}
            rows={8}
            componentClass="textarea" />
          <HelpBlock>{i18next.t('site.helps.leave-word.body')}</HelpBlock>
        </FormGroup>
        <Button type="submit" bsStyle="primary">
          {i18next.t('buttons.submit')}
        </Button>
      </form>
    </div>)
  }
}
