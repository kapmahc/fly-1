import React, {Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {Table, FormGroup, FormControl,ControlLabel,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';
import Dropzone from 'react-dropzone'
import Clipboard from 'clipboard'
import octicons from 'octicons'

import {get, post, _delete} from '../../ajax'
import {Forbidden} from '../../components/alerts'

class DashboardW extends Component{
  constructor(props){
    super(props)
    this.state = {
      items:[],
      id:null,
      title:null,
    }
    this.handleDrop = this.handleDrop.bind(this);
    this.handleRemove = this.handleRemove.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
  }
  componentDidMount() {
    get(`/attachments`).then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
    new Clipboard('.clipboard');
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
    var id = this.state.id
    post(`/attachments/${id}`, data)
      .then(function(rst){
        alert(i18next.t('success'))
        var items = this.state.items.filter((a, _) => a.id !== id)
        items.unshift(rst)
        this.setState({
          title: null,
          id:null,
          items: items
        })
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  }
  handleEdit(a){
    this.setState({id:a.id, title:a.title})
  }
  handleRemove(id) {
    if(confirm(i18next.t('are-you-sure'))){
      _delete(`/attachments/${id}`)
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
  handleDrop(files) {
    var data = new FormData()
    files.forEach((file)=> {
      data.append('files', file)
    });
    post('/attachments', data)
      .then(function(rst){
        alert(i18next.t('success'))
        var items=this.state.items
        this.setState({items:rst.concat(items)})
      }.bind(this))
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    const {user} = this.props
    var cb = function(a){
      if (a.mediaType.startsWith('image')){
        return `![${a.title}](${a.url})`
      }
      return `[${a.title}](${a.url})`
    }
    return user.uid ? (<div className="row">
    <h3>{i18next.t('auth.attachments.index.title')}</h3>
    <hr/>
    <Dropzone onDrop={this.handleDrop}>
      <div>{i18next.t('hints.uploader')}</div>
    </Dropzone>
    <br/>
    { this.state.id ?
        (<div className="col-md-10 col-md-offset-1">
          <h3>{i18next.t('buttons.edit')}[{this.state.id}]</h3>
          <hr/>
          <form onSubmit={this.handleSubmit}>
            <FormGroup controlId="title">
              <ControlLabel>{i18next.t('attributes.title')}</ControlLabel>
              <FormControl
                value={this.state.title}
                onChange={this.handleChange} />
            </FormGroup>
            <Button type="submit" bsStyle="primary">
              {i18next.t('buttons.submit')}
            </Button>
          </form>
        </div>) : <br/>
      }
      <br/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.updatedAt')}</th>
            <th>
              {i18next.t('attributes.name')}
            </th>
            <th>
              {i18next.t('buttons.manage')}
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((a,i)=>(<tr key={i}>
            <td>{a.updatedAt}</td>
            <td>
              <a href={a.url} target="_blank">{a.title}</a>
              <button alt={i18next.t('auth.attachments.index.copy-toclip-board')} dangerouslySetInnerHTML={{__html:octicons.clippy.toSVG()}} className="clipboard" data-clipboard-text={cb(a)}></button>
            </td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <Button bsStyle="warning" onClick={()=>this.handleEdit(a)}>{i18next.t('buttons.edit')}</Button>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(a.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
          </td>
          </tr>))}
        </tbody>
      </Table>
    </div>):<Forbidden/>
  }
}

DashboardW.propTypes = {
  user: PropTypes.object.isRequired
}

export const Dashboard = connect(
  state => ({user:state.currentUser})
)(DashboardW);
