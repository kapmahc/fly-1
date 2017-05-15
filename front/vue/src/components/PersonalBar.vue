<template>
  <b-nav-item-dropdown text="User" right-alignment right>
    <b-dropdown-item to="/users/info">Profile</b-dropdown-item>
    <b-dropdown-item to="/users/sign-out">Signout</b-dropdown-item>
  </b-nav-item-dropdown>
</template>

<script>
export default {
  name: 'personal-bar',
  data () {
    return {
    }
  }
}
</script>
