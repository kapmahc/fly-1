<template>
  <footer>
    <p class="float-right">
      {{$t('footer.other-languages')}}:
      <b-link v-bind:key="l" v-on:click="setLocale(l)" v-for="l in languages">
        {{$t(`languages.${l}`)}}
      </b-link>
    </p>
    <p>
      &copy; {{ $t("site.copyright") }}
      &middot; <a href="#">Privacy</a>
      &middot; <a href="#">Terms</a>
    </p>
  </footer>
</template>

<script>
import {LANGUAGES, load as setLocale} from '@/i18n'

export default {
  name: 'footer',
  data () {
    return {
      languages: LANGUAGES
    }
  },
  methods: {
    setLocale
  }
}
</script>
