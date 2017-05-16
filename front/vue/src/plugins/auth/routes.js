import UsersSignIn from './users/SignIn'
import UsersSignUp from './users/SignUp'

export default [
  { path: '/users/sign-in', name: 'auth.users.sign-in', component: UsersSignIn },
  { path: '/users/sign-up', name: 'auth.users.sign-up', component: UsersSignUp },
  { path: '/users/forgot-password', name: 'auth.users.forgot-password', component: UsersSignUp },
  { path: '/users/confirm', name: 'auth.users.confirm', component: UsersSignUp },
  { path: '/users/unlock', name: 'auth.users.unlock', component: UsersSignUp }
]
