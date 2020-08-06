<template>
  <div class="content">
    <div class="container-fluid">
      <ModalBuefy></ModalBuefy>
      <!--    <b-button v-on:click="welcome">welcome</b-button>-->
    </div>
  </div>

</template>

<script>
    import ModalBuefy from "../components/Modals/ModalBuefy";

    export default {
      name: "LogIn.vue",
      components: {
        ModalBuefy
      }
    }
</script>

<style scoped>
div.sidebar{
  display: none;
}
</style>
