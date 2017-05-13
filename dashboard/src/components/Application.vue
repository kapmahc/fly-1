<template>
  <div>
    <app-header />
    <h2>application</h2>
    <slot />
    <app-footer />
  </div>
</template>

<script>
import Header from './Header'
import Footer from './Footer'

export default {
  name: 'layout-application',
  data () {
    return {
    }
  },
  components: {
    'app-header': Header,
    'app-footer': Footer
  }
}
</script>
