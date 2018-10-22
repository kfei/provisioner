<template>
  <v-app>
    <v-navigation-drawer
      permanent
      :mini-variant="miniVariant"
      enable-resize-watcher
      fixed
      app
      width="240"
    >
      <v-toolbar flat class="transparent">
        <v-list class="pa-0">
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <img :src="require('./assets/logo.svg')" alt="logo">
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title class="font-weight-bold">{{ appName }}</v-list-tile-title>
              <v-list-tile-sub-title>for some coupons</v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-toolbar>

      <v-divider></v-divider>

      <v-list>
        <v-list-tile
          active-class="grey lighten-3 primary--text"
          exact
          v-for="(link, i) in links"
          :key="i"
          :to="{ name: link.route }"
        >
          <v-list-tile-action>
            <v-icon v-html="link.icon"></v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title v-text="link.title"></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar
      app
      class="white elevation-1"
    >
      <v-btn icon @click.stop="miniVariant = !miniVariant">
        <v-icon v-html="miniVariant ? 'chevron_right' : 'chevron_left'"></v-icon>
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>menu</v-icon>
      </v-btn>
    </v-toolbar>

    <v-content class="grey lighten-5">
      <router-view/>
    </v-content>

  </v-app>
</template>

<script>

export default {
  name: 'App',
  data () {
    return {
      appName: require('../package.json').name,
      links: [{
        icon: 'attach_money',
        title: 'Coupon',
        route: 'coupon'
      }, {
        icon: 'info',
        title: 'About',
        route: 'about'
      }],
      miniVariant: true
    }
  }
}
</script>

<style lang="stylus">
body
  margin 0
#app
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
</style>
