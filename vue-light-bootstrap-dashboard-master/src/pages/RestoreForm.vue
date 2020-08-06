<template>
  <div class="col-12">
    Restore Form page
<!--    Test :: {{test}}-->
<!--    <ModalBuefy></ModalBuefy>-->
<!--    <b-button v-on:click="welcome">welcome</b-button>-->
  </div>
</template>

<script>
    export default {
      name: "LogIn.vue",
      data () {
        return {
          // test: ''
        }
      },
      components: {
      },
      mounted() {
        // this.test = this.$route.query.test
        // this.$store.dispatch('login/restorePassword', this.test);
        // console.log(this.$route.query.test)
      }
    }
</script>

<style scoped>
div.sidebar{
  display: none;
}
</style>
