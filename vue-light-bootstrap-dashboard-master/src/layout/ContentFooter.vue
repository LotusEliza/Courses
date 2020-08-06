<template>
  <footer class="footer">
    <div class="container-fluid">
      <nav v-if="isAuthenticated && authStatus == 'success'">
        <ul class="footer-menu">
          <li>
            <router-link :to="{path:'/admin/courses'}" >Dashboard</router-link>
          </li>
        </ul>
      </nav>
      <div class="copyright text-center">
        Copyright &copy; 2020 Online courses. All rights reserved
        <i class="fa fa-heart heart"></i>
<!--        by-->
<!--        <a href="https://binarcode.com" target="_blank">BinarCode</a>.-->
<!--        Designed by <a href="https://www.creative-tim.com/?ref=pdf-vuejs" target="_blank">Creative Tim</a>.-->
      </div>
    </div>
  </footer>
</template>
<script>
  import {mapGetters} from "vuex";

  export default {
    computed: {
      ...mapGetters("login", [
        'isAuthenticated', 'authStatus'
      ])
    },
  }


</script>
<style>

</style>
