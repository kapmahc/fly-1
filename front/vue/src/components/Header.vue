<template>
  <b-navbar toggleable type="inverse" variant="inverse" fixed="top">

    <b-nav-toggle target="nav_collapse"></b-nav-toggle>

    <b-link class="navbar-brand" to="#">
      <span>BootstrapVue</span>
    </b-link>

    <b-collapse is-nav id="nav_collapse">

      <b-nav is-nav-bar>
        <b-nav-item>Support</b-nav-item>
        <b-nav-item>Docs</b-nav-item>
        <b-nav-item>Contact Us</b-nav-item>
      </b-nav>

      <b-nav is-nav-bar class="ml-auto">
        <LanguageBar />
        <PersonalBar />
      </b-nav>
    </b-collapse>
  </b-navbar>
</template>

<script>
import PersonalBar from './PersonalBar'
import LanguageBar from './LanguageBar'

export default {
  name: 'header',
  data () {
    return {
    }
  },
  components: {
    PersonalBar,
    LanguageBar
  }
}
</script>
