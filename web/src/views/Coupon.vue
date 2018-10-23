<template>
  <v-container class="text-xs-center">
    <v-layout row wrap>
      <v-flex xs12>
        <div class="available d-flex" style="flex-direction: column;">
          <div class="description headline mt-5 grey--text">Coupons Available:</div>
          <div class="coupon primary--text" :class="{ increased: increased, decreased: decreased }">{{ available }}</div>
        </div>
      </v-flex>
      <v-flex xs12>
        <div class="request">
          <v-btn color="primary" large @click="requestCoupon">Give Me A Coupon</v-btn>
        </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapState } from 'vuex'
import ws from '@/ws'

export default {
  computed: {
    ...mapState({
      'available': state => state.available,
      'increased': state => state.delta > 0,
      'decreased': state => state.delta < 0
    })
  },

  methods: {
    requestCoupon () {
      ws.send(JSON.stringify({
        command: 'request-coupon',
        requestCoupon: {
          count: 1
        }
      }))
    }
  }
}
</script>

<style lang="stylus" scoped>
.available
  .coupon
    font-size 200px

    &.increased
      animation color-increased 1s infinite
    &.decreased
      animation color-decreased 1s infinite

primaryColor = '#3498db'
@keyframes color-increased
  0%
    color primaryColor
  50%
    color green
  100%
    color primaryColor
@keyframes color-decreased
  0%
    color primaryColor
  50%
    color red
  100%
    color primaryColor
</style>
