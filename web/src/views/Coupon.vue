<template>
  <v-container class="text-xs-center">
    <v-layout row wrap>
      <v-flex xs12>
        <div class="available d-flex" style="flex-direction: column;">
          <div class="description headline mt-5 grey--text">Coupons Available:</div>
          <div class="coupon primary--text">{{ count }}</div>
        </div>
      </v-flex>
      <v-flex xs12>
        <div class="request">
          <v-btn color="primary" large>Give Me A Coupon</v-btn>
        </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  data () {
    return {
      count: 5566
    }
  }
}
</script>

<style lang="stylus" scoped>
.available
  .coupon
    font-size 200px
</style>
