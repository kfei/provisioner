<template>
    <v-tooltip left>
      <div class="dot-indicator" :class="wsStatus === 1 ? 'green' : 'red'" slot="activator" />
      <span>WebSocket {{ wsStatus === 1 ? 'connected' : 'disconnected' }}</span>
    </v-tooltip>
</template>

<script>
export default {
  computed: {
    wsStatus () {
      return this.$store.state.wsStatus
    }
  }
}
</script>

<style lang="stylus" scoped>
.dot-indicator
  width 20px
  height 20px
  border-radius 100%
</style>
