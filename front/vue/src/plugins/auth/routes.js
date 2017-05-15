import UsersSignIn from './users/SignIn'
import UsersSignUp from './users/SignUp'

export default [
  { path: '/users/sign-in', name: 'auth.users.sign-in', component: UsersSignIn },
  { path: '/users/sign-up', name: 'auth.users.sign-up', component: UsersSignUp }
]
