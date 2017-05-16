<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="personalBar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t('personal-bar.sign-in-or-up')}}
    </a>
    <div v-if="user" class="dropdown-menu" aria-labelledby="personalBar">
      <router-link class="dropdown-item" :to="{name: 'auth.users.sign-in'}">
        sign in
      </router-link>
      <router-link class="dropdown-item" :to="{name: 'auth.users.sign-in'}">
        sign out
      </router-link>
    </div>
    <div v-else class="dropdown-menu" aria-labelledby="personalBar">
      <router-link class="dropdown-item" v-for="l in links" v-bind:key="l" :to="{name: l.href}">
        {{$t(`${l.href}.title`)}}
      </router-link>
    </div>
  </li>
</template>

<script>
import {nonSignInLinks as links} from '@/constants'

export default {
  data () {
    return {
      links
    }
  },
  computed: {
    user () {
      return this.$store.state.currentUser
    }
  }
}
</script>
