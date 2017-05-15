import React, { Component, PropTypes } from 'react';
import {Link} from 'react-router'
import i18next from 'i18next';

import {get} from '../../ajax'
import Markdown from '../../components/Markdown'

const Show = ({item}) => (<div className="row">
  <h2>{item.title}</h2>
  <hr/>
  <Markdown body={item.body}/>
</div>)

Show.propTypes = {
  item: PropTypes.object.isRequired
}

class Widget extends Component{
  constructor(props){
    super(props)
    this.state = {
      cur: null,
      items:[]
    }
  }
  componentDidMount() {
    const {params} = this.props
    get('/blogs').then(
      function(rst){
        this.setState({items:rst})
        this.setCur(params.splat)
      }.bind(this)
    );
  }
  componentWillReceiveProps(nextProps){
    const {params} = nextProps
    this.setCur(params.splat)
  }
  setCur(href){
    if(href === ""){
      this.setState({cur:null})
    }else{
      this.state.items.forEach(function(b){
        if(b.href === href){
          this.setState({cur:b})
        }
      }.bind(this))
    }
  }
  render(){
    const {cur, items} = this.state
    var body = cur ?
    (<Show item={cur} />) : items.map((b,i)=>(<div className="row" key={i}>
      <Show item={b} />
    </div>))

    return (<div className="row">
    <div className="col-md-9">
      {body}
    </div>
    <div className="col-md-3">
      <h4>{i18next.t('blog.list.title')}</h4>
      <ol className="list-unstyled">
        {items.map((b,i)=><li key={i}>
          <Link to={`/blogs/${b.href}`}>{b.title}</Link>
        </li>)}
      </ol>
    </div>
  </div>)
  }
}

export default Widget
