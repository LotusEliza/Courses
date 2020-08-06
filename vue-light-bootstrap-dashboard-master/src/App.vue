<template>

  <div :class="{'nav-open': $sidebar.showSidebar}">
<!--    sessionOver:::{{sessionOver}}-->
    <notifications></notifications>

    <alert ></alert>
<!--    <alert ref="confirm"></alert>-->
<!--    <ModalSessionOver :isComponentModalActive="isComponentModalActive"></ModalSessionOver>-->
    <router-view></router-view>
  </div>
</template>

<script>
  import {mapGetters} from "vuex";
  // import ModalSessionOver from "./components/Alerts/Alert.vue";

  export default {
    components: {
      // Alert
    },

    data() {
      return {
      }
    },
    computed: {
      ...mapGetters("login", [
        'sessionOver',
      ]),
    },
    watch: {
      sessionOver: function (newVal, oldVal) {
        console.log("oldVal"+oldVal)
        if(newVal==="yes"){
          console.log("From APP vue watcher")
          // this.$refs.confirm.confirm()
          // this.$children.confirm()
          // this.$router.push({ name: 'LogIn' })
        }
      },
    },
  }
</script>
<style lang="scss">
  .vue-notifyjs.notifications{
    .list-move {
      transition: transform 0.3s, opacity 0.4s;
    }
    .list-item {
      display: inline-block;
      margin-right: 10px;

    }
    .list-enter-active {
      transition: transform 0.2s ease-in, opacity 0.4s ease-in;
    }
    .list-leave-active {
      transition: transform 1s ease-out, opacity 0.4s ease-out;
    }

    .list-enter {
      opacity: 0;
      transform: scale(1.1);

    }
    .list-leave-to {
      opacity: 0;
      transform: scale(1.2, 0.7);
    }
  }

</style>
