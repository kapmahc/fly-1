<template>
  <footer>
    <p class="float-right">
      <a href="#">Back to top</a>
    </p>
    <p>
      &copy; {{ $t("site.copyright") }}
      &middot; <a href="#">Privacy</a>
      &middot; <a href="#">Terms</a>
    </p>
  </footer>
</template>

<script>
export default {
  name: 'footer',
  data () {
    return {
    }
  }
}
</script>
