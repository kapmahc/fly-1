import React, { Component } from 'react';
import i18next from 'i18next';
import {Pagination, Thumbnail, Button} from 'react-bootstrap';
import {LinkContainer} from 'react-router-bootstrap'


import {get} from '../../ajax'
import {PAGE_SIZE} from '../../constants'

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
      this.loadBooks(1, PAGE_SIZE)
  }
  handleSelect(page){
    this.loadBooks(page, PAGE_SIZE)
  }
  loadBooks(page, size){
    get(`/reading/books?size=${size}&page=${page}`).then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render(){
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
      <h3>{i18next.t('reading.books.index.title')}</h3>
      <hr/>
      {pager}
      {this.state.items.map((b,i)=>(<div className="col-md-4" key={i}>
        <Thumbnail>
          <h3>{b.title}</h3>
          <p>{b.author}</p>
          <p>
            <LinkContainer to={`/reading/books/${b.id}`} target="_blank">
              <Button bsStyle="primary">{i18next.t('buttons.view')}</Button>
            </LinkContainer>
          </p>
        </Thumbnail>
      </div>))}
      {pager}
    </div>)
  }
}

export class Show extends Component{
  constructor(props){
    super(props)
    this.state = {
      book:{},
      home: ''
    }
  }
  componentDidMount() {
    const {params} = this.props
    get(`/reading/books/${params.id}`).then(
      function(rst){
        this.setState(rst)
      }.bind(this)
    );
  }
  render(){
    return (<div className="row">
      <h3>{this.state.book.title}-{this.state.book.author}</h3>
      <hr/>
      <p dangerouslySetInnerHTML={{__html: this.state.home}}></p>
    </div>)
  }
}
