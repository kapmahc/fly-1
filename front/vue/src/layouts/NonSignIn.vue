<template>
  <Layout>
    <div class="row">
      <div class="offset-2 col-8">
        <h2>{{$t(title)}}</h2>
        <hr />
        <slot />
        <div class="list-group block">
          <router-link class="list-group-item list-group-item-action" v-for="l in links" v-bind:key="l" :to="{name: l.href}">
            <material-icon :name="l.icon" />&nbsp;{{$t(`${l.href}.title`)}}
          </router-link>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from '@/layouts/Application'
import {nonSignInLinks as links} from '@/constants'
import Icon from '@/components/Icon'

export default {
  props: ['title'],
  data () {
    return {
      links
    }
  },
  components: {
    Layout,
    Icon
  }
}
</script>

<style scope>
.block {
  padding-top: 3rem;
  padding-bottom: 3rem;
}
</style>
