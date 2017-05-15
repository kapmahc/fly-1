import React, {Component, PropTypes } from 'react'
import {Table, FormGroup,ControlLabel, FormControl,
  HelpBlock,Pagination,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'
import {browserHistory, Link} from 'react-router'

import {get, post, _delete} from '../../ajax'
import {timeago} from '../../utils'
import Markdown from '../../components/Markdown'
import {PAGE_SIZE} from '../../constants'

export const List = ({items, more}) => (<div className="col-md-12">
{items.map((c,i)=>(<blockquote key={i}>
  <Markdown body={c.body}/>
  <footer>
    {more ? (<Link to={`/forum/articles/${c.articleId}`}>{i18next.t('buttons.view')}</Link>):<span/>}
    <cite>{timeago(c.updatedAt)}</cite>
  </footer>
</blockquote>))}
</div>)

List.propTypes = {
  items: PropTypes.array.isRequired,
  more: PropTypes.bool
}


export class Index extends Component{
  constructor(props){
    super(props)
    this.state = {
      page:1,
      size:PAGE_SIZE,
      total:0,
      count:1,
      items:[]
    }
    this.handleSelect = this.handleSelect.bind(this);
  }
  componentDidMount() {
    this.loadArticles(1, PAGE_SIZE)
  }
  handleSelect(page){
    this.loadArticles(page, PAGE_SIZE)
  }
  loadArticles(page, size){
    get(`/forum/comments?page=${page}&size=${size}`).then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render() {
    var pager = (<div className="col-md-12">
      <Pagination
          prev
          next
          first
          last
          ellipsis
          boundaryLinks
          items={this.state.count}
          maxButtons={5}
          activePage={this.state.page}
          onSelect={this.handleSelect} />
      </div>)
    return (<div className="row">
      <h3>{i18next.t('forum.comments.index.title')}</h3>
      <hr/>
      {pager}
      <List items={this.state.items} more/>
      {pager}
    </div>)
  }
}

export class New extends Component{
  constructor(props){
    super(props)
    this.state = {
      body: '',
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
    data.append('articleId', this.props.location.query.articleId)
    data.append('body', this.state.body)
    data.append('type', 'markdown')
    post('/forum/comments', data)
      .then(function(rst){
        alert(i18next.t('success'))
        browserHistory.push(`/forum/articles/${rst.articleId}`)
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="row">
      <h3>{i18next.t('buttons.new')}</h3>
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
    </div>)
  }
}



export class Edit extends Component{
  constructor(props){
    super(props)
    this.state = {
      body: '',
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get(`/forum/comments/${this.props.params.id}`).then(
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
    data.append('body', this.state.body)
    data.append('type', 'markdown')

    post(`/forum/comments/${this.state.id}`, data)
      .then(function(rst){
        alert(i18next.t('success'))
        browserHistory.push(`/forum/articles/${rst.articleId}`)
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="row">
      <h3>{i18next.t('buttons.edit')}[{this.state.id}]</h3>
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
    </div>)
  }
}

Edit.propTypes = {
  id: PropTypes.number.isRequired
}


export class Dashboard extends Component{
  constructor(props){
    super(props)
    this.state = {
      items: []
    }
    this.handleRemove = this.handleRemove.bind(this);
  }
  componentDidMount() {
    get('/forum/my/comments').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  handleRemove(id) {
    if(confirm(i18next.t('are-you-sure'))){
      _delete(`/forum/comments/${id}`)
        .then(function(rst){
          alert(i18next.t('success'))
          this.setState({
            items: this.state.items.filter((c, _) => c.id !== id)
          });
        }.bind(this))
        .catch((err) => {
          alert(err)
        })
    }
  }
  render() {

    return (<div className="row">
      <div className="col-md-10 col-md-offset-1">
      <h3>{i18next.t('forum.dashboard.comments.title')}</h3>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.updatedAt')}</th>
            <th>{i18next.t('attributes.body')}</th>
            <th>
              {i18next.t('buttons.manage')}
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((c,i)=>(<tr key={i}>
            <td>{c.updatedAt}</td>
            <td><pre><code>{c.body}</code></pre></td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <LinkContainer to={`/forum/comments/${c.id}/edit`} target="_blank">
                  <Button bsStyle="warning">{i18next.t('buttons.edit')}</Button>
                </LinkContainer>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(c.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
          </td>
          </tr>))}
        </tbody>
      </Table>
      </div>
    </div>)
  }
}
