<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="personalBar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t('personal-bar.sign-in-or-up')}}
    </a>
    <div class="dropdown-menu" aria-labelledby="personalBar">
      <router-link class="dropdown-item" :to="{name: 'auth.users.sign-in'}">
        sign in
      </router-link>
      <router-link class="dropdown-item" :to="{name: 'auth.users.sign-in'}">
        sign up
      </router-link>
    </div>
  </li>
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
