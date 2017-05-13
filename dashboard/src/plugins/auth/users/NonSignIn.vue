<template>
  <Application>
    <div class="row">
      <div class="col-12">
        <slot />
      </div>
      <div class="col-12">non sign in </div>
    </div>
  </Application>
</template>

<script>
import Application from '@/layouts/Application'

export default {
  name: 'auth-non-sign-in',
  data () {
    return {
    }
  },
  components: {
    Application
  }
}
</script>

<style scope>
body {
  padding-top: 4rem;
}
</style>
