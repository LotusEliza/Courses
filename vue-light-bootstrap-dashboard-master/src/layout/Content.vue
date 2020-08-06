<template>
  <transition name="fade" mode="out-in">
    <router-view>
<!--      <modal-session-over></modal-session-over>-->
    </router-view>
  </transition>
</template>
<script>
  // import ModalSessionOver from '../components/Modals/ModalSessionOver.vue'
  export default {
    // components: {
    //   ModalSessionOver
    // },
  }
</script>
<style>
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity .1s
  }

  .fade-enter,
  .fade-leave-to
    /* .fade-leave-active in <2.1.8 */

  {
    opacity: 0
  }
</style>
