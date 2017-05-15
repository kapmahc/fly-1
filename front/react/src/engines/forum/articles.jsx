import React, {Component, PropTypes } from 'react'
import { connect } from 'react-redux'
import {Table, FormGroup,ControlLabel, FormControl, Checkbox,
  Thumbnail,
  HelpBlock,Pagination,
  ButtonGroup, ButtonToolbar, Button} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'
import {browserHistory, Link} from 'react-router'

import {get, post, _delete} from '../../ajax'
import {List as Comments}  from './comments'
import Markdown from '../../components/Markdown'
import {PAGE_SIZE} from '../../constants'

export const List = ({items}) => (<div className="row">
  {items.map((a,i)=>(<div className="col-md-4" key={i}>
      <Thumbnail>
       <h3>{a.title}</h3>
       <p>{a.summary}</p>
       <p>
         <LinkContainer target="_blank" to={`/forum/articles/${a.id}`}>
           <Button bsStyle="primary">{i18next.t('buttons.view')}</Button>
         </LinkContainer>
       </p>
     </Thumbnail>
  </div>))}
</div>)

List.propTypes = {
  items: PropTypes.array.isRequired
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
    get(`/forum/articles?page=${page}&size=${size}`).then(
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
      <h3>{i18next.t('forum.articles.index.title')}</h3>
      <hr/>
      {pager}
      <List items={this.state.items}/>
      {pager}
    </div>)
  }
}

class Form extends Component{
  constructor(props){
    super(props)
    this.state = {
      id: props.id,
      title:'',
      summary:'',
      type:'',
      body: '',
      tags: [],
      tagBoxes: [],
    }
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  componentDidMount() {
    get('/forum/tags').then(
      function(rst){
        this.setState({tagBoxes:rst})
      }.bind(this)
    );
    var id = this.state.id;
    if(id){
      get(`/forum/articles/${id}`).then(
        function(rst){
          this.setState(rst)
          rst.tags.forEach((t)=>{
            var tmp = {}
            tmp[`tag-${t.id}`] = true
            this.setState(tmp)
          })
        }.bind(this)
      );
    }
  }
  handleChange(e) {
    var data = {};
    var t = e.target
    data[t.id] = t.type === 'checkbox' ? t.checked : t.value;
    this.setState(data);
  }
  handleSubmit(e) {
    e.preventDefault();
    var data = new FormData()
    data.append('title', this.state.title)
    data.append('summary', this.state.summary)
    data.append('body', this.state.body)
    data.append('type', 'markdown')
    this.state.tagBoxes.filter(function(t){
      return this.state[`tag-${t.id}`] === true
    }.bind(this)).forEach((t)=>data.append('tags', t.id))

    var id = this.state.id
    post(id ?`/forum/articles/${id}`: '/forum/articles', data)
      .then(function(rst){
        alert(i18next.t('success'))
        browserHistory.push(`/forum/articles/${rst.id}`)
      })
      .catch((err) => {
        alert(err)
      })
  }
  render() {
    return (<div className="row">
      <h3>{this.state.id ? `${i18next.t('buttons.edit')}[${this.state.id}]` : i18next.t('forum.articles.new.title')}</h3>
      <hr/>
      <form onSubmit={this.handleSubmit}>
        <FormGroup controlId="title">
          <ControlLabel>{i18next.t('attributes.title')}</ControlLabel>
          <FormControl
            type="text"
            value={this.state.title}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup controlId="summary">
          <ControlLabel>{i18next.t('attributes.summary')}</ControlLabel>
          <FormControl
            value={this.state.summary}
            onChange={this.handleChange}
            componentClass="textarea" />
        </FormGroup>
        <FormGroup>
          {this.state.tagBoxes.map(function(t, i){
            var id = `tag-${t.id}`
            return (<Checkbox checked={this.state[id] === true} id={id} key={i} onChange={this.handleChange} inline>{t.name}</Checkbox>)
          }.bind(this))}
        </FormGroup>
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


class ShowW extends Component{
  constructor(props){
    super(props)
    this.state = {
      title:'',
      summary:'',
      body:'',
      comments: [],
      tags:[],
    }
  }
  componentDidMount() {
    const {params} = this.props
    get(`/forum/articles/${params.id}`).then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render() {
    const {user} = this.props
    return (<div className="row">
      <h3>{this.state.title}</h3>
      <hr/>
      <b>{this.state.summary}</b>
      <br/>
      {this.state.tags.map((t,i)=>(<Link key={i} to={`/forum/tags/${t.id}`} className="block">{t.name}</Link>))}
      <br/>
      <Markdown body={this.state.body}/>
      <h4>
        {i18next.t('forum.comments.index.title')}
        {user.uid ? (<Link className="block" to={`/forum/comments/new?articleId=${this.state.id}`}>{i18next.t("buttons.new")}</Link>): <br/>}
      </h4>
      <hr/>
      <Comments items={this.state.comments}/>
    </div>)
  }
}


ShowW.propTypes = {
  user: PropTypes.object.isRequired
}

export const Show = connect(
  state => ({user: state.currentUser})
)(ShowW)

export const New = ()=>(<Form />)
export const Edit = ({params})=>(<Form id={params.id} />)

export class Dashboard extends Component{
  constructor(props){
    super(props)
    this.state = {
      items: []
    }
    this.handleRemove = this.handleRemove.bind(this);
  }
  componentDidMount() {
    get('/forum/my/articles').then(
      function(rst){
        this.setState({items:rst})
      }.bind(this)
    );
  }
  handleRemove(id) {
    if(confirm(i18next.t('are-you-sure'))){
      _delete(`/forum/articles/${id}`)
        .then(function(rst){
          alert(i18next.t('success'))
          this.setState({
            items: this.state.items.filter((a, _) => a.id !== id)
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
      <h3>{i18next.t('forum.dashboard.articles.title')}</h3>
      <hr/>
      <Table striped bordered condensed hover>
        <thead>
          <tr>
            <th>{i18next.t('attributes.updatedAt')}</th>
            <th>{i18next.t('attributes.title')}</th>
            <th>
              {i18next.t('buttons.manage')}
              <LinkContainer to="/forum/articles/new">
                <Button bsStyle="success" bsSize="sm">{i18next.t('buttons.new')}</Button>
              </LinkContainer>
            </th>
          </tr>
        </thead>
        <tbody>
          {this.state.items.map((a,i)=>(<tr key={i}>
            <td>{a.updatedAt}</td>
            <td>{a.title}</td>
            <td>
              <ButtonToolbar><ButtonGroup bsSize="sm">
                <LinkContainer to={`/forum/articles/${a.id}/edit`} target="_blank">
                  <Button bsStyle="warning">{i18next.t('buttons.edit')}</Button>
                </LinkContainer>
                <Button bsStyle="danger" onClick={()=>this.handleRemove(a.id)}>{i18next.t('buttons.remove')}</Button>
              </ButtonGroup></ButtonToolbar>
          </td>
          </tr>))}
        </tbody>
      </Table>
      </div>
    </div>)
  }
}
