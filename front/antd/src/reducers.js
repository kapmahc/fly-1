import {USERS_SIGN_IN, USERS_SIGN_OUT, REFRESH_SITE_INFO} from './actions'
import jwtDecode from 'jwt-decode'


const currentUser = (state={}, action) => {
  switch(action.type){
    case USERS_SIGN_IN:
      try{
        return jwtDecode(action.token)
      }catch(e){
        console.log(e)
      }
      return {}
    case USERS_SIGN_OUT:
      return {}
    default:
      return state
  }
}

const siteInfo = (state={languages:[]}, action) => {
  switch(action.type){
    case REFRESH_SITE_INFO:
      return Object.assign({}, action.info)
    default:
      return state;
  }
}

export default {
  currentUser,
  siteInfo
}
