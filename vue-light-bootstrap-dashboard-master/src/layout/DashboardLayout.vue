<template>
  <div class="wrapper">
    <side-bar >
      <mobile-menu slot="content"></mobile-menu>
<!--     For simple user-->
      <div v-if="isAuthenticated && !isAdmin">
        <sidebar-link to="/admin/courses">
          <i class="nc-icon nc-chart-pie-35"></i>
          <p>Курсы</p>
        </sidebar-link>
        <sidebar-link to="/admin/user">
          <i class="nc-icon nc-circle-09"></i>
          <p>User Profile</p>
        </sidebar-link>
      </div>
<!--      For Admin-->

      <div v-if="isAuthenticated && isAdmin">
        <sidebar-link to="/list_of_courses">
          <i class="nc-icon nc-chart-pie-35"></i>
          <p>Курсы</p>
        </sidebar-link>
        <sidebar-link to="/list_of_customers">
          <i class="nc-icon nc-chart-pie-35"></i>
          <p>Пользователи</p>
        </sidebar-link>
        <sidebar-link to="/list_of_orders">
          <i class="nc-icon nc-chart-pie-35"></i>
          <p>Заказы</p>
        </sidebar-link>
        <sidebar-link to="/add_course">
          <i class="nc-icon nc-chart-pie-35"></i>
          <p>SMS рассылки</p>
        </sidebar-link>
<!--        <sidebar-link to="/admin/user">-->
<!--          <i class="nc-icon nc-circle-09"></i>-->
<!--          <p>User Profile</p>-->
<!--        </sidebar-link>-->
      </div>
    </side-bar>
    <div class="main-panel">
      <top-navbar></top-navbar>

      <dashboard-content @click="toggleSidebar">

      </dashboard-content>


      <content-footer></content-footer>
    </div>
  </div>
</template>
<style lang="scss">

</style>
<script>
  import TopNavbar from './TopNavbar.vue'

  import ContentFooter from './ContentFooter.vue'
  import DashboardContent from './Content.vue'
  import MobileMenu from './MobileMenu.vue'
  import {mapGetters} from "vuex";
  import cookie from 'vue-cookies'

  export default {
    data() {
      return {
        cookie: cookie.get('token')
      }
    },
    components: {
      TopNavbar,
      ContentFooter,
      DashboardContent,
      MobileMenu,

    },
    methods: {
      toggleSidebar () {
        if (this.$sidebar.showSidebar) {
          this.$sidebar.displaySidebar(false)
        }
      }
    },
    computed:{
      ...mapGetters("login", [
        'isAuthenticated', 'authStatus', 'isAdmin'
      ]),
    },
  }

</script>


