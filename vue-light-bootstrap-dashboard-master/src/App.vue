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
  .btn {
    padding: 5px 10px 5px 10px !important;
  }

  .b-tabs .tabs li.is-active a {
    border-bottom-color: #43d0ed;
    color: #43d0ed;
  }

  /*table tbody .b-checkbox.checkbox input[type=checkbox]:checked + .check {*/
  /*  background: #43d0ed url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 1 1'%3E%3Cpath style='fill:white' d='M 0.04038059,0.6267767 0.14644661,0.52071068 0.42928932,0.80355339 0.3232233,0.90961941 z M 0.21715729,0.80355339 0.85355339,0.16715729 0.95961941,0.2732233 0.3232233,0.90961941 z'%3E%3C/path%3E%3C/svg%3E") no-repeat center center;*/
  /*  border-color: #43d0ed;*/
  /*}*/

  /*table tbody .b-checkbox.checkbox:hover input[type=checkbox]:not(:disabled) + .check {*/
  /*  border-color: #43d0ed;*/
  /*}*/

  /*#__lpform_input_idx_4 > img, #__lpform_input_idx_8 > img {*/
  /*!*#__lpform_input_idx_4_icon img {*!*/
  /*  display: none;*/
  /*}*/
  .icon.has-text-primary {
    color: #339db2 !important;
  }
</style>
<style lang="scss">
@import "../node_modules/bulma/sass/utilities/_all.sass";
@import "../node_modules/buefy/src/scss/utils/_all.scss";
/*@import "../node_modules/buefy/dist/buefy.css";*/
/*/home/lotus/go/src/courses/vue-light-bootstrap-dashboard-master/node_modules/buefy/src/scss/utils/_all.scss*/
/*/home/lotus/go/src/courses/vue-light-bootstrap-dashboard-master/node_modules/buefy/dist/buefy.css*/

/*$color: red;*/
/*$green: green;*/




/*// Set your colors*/
$primary: #37d822;
$primary-invert: findColorInvert($primary);
$twitter: #4099FF;
$twitter-invert: findColorInvert($twitter);

/*// Setup $colors to use as bulma classes (e.g. 'is-twitter')*/
$colors: (
"white": ($white, $black),
"black": ($black, $white),
"light": ($light, $light-invert),
"dark": ($dark, $dark-invert),
"primary": ($primary, $primary-invert),
"info": ($info, $info-invert),
"success": ($success, $success-invert),
"warning": ($warning, $warning-invert),
"danger": ($danger, $danger-invert),
"twitter": ($twitter, $twitter-invert)
);

$has-text-primary: #37d822;
 /*{*/
/*  color: orangered !important;*/
/*}*/
/*@import "../node_modules/buefy/src/scss/buefy.scss";*/
/*table tbody .b-checkbox.checkbox:hover input[type=checkbox]:not(:disabled) + .check {*/
/*  border-color: $color;*/
/*}*/
</style>

<!--<style lang="scss">-->
<!--/*@import "../styles/_variables.scss";*/-->
<!--@import "../../node_modules/bulma/sass/utilities/_all.sass";-->
<!--table td:nth-child(1) {-->
<!--  background-color: $violetpale;-->
<!--}-->
<!--/*.table thead th {*/-->
<!--/*    background-color: $pink;*/-->
<!--/*}*/-->
<!--$table-row-hover-background-color: $hover;-->
<!--$narbar-hover-background-color: $hover;-->
<!--.b-table .table {-->
<!--  background-color: $table;-->
<!--};-->
<!--/*// Import Bulma and Buefy styles*/-->
<!--$colors: (-->
<!--"light": ($pinkstrong, white),-->
<!--"primary": ($violet, white),-->
<!--"info": ($peachy, $white),-->
<!--"success": ($green, $success-invert),-->
<!--"warning": ($darkyellow, $white),-->
<!--"danger": ($pinkstrong, white),-->
<!--);-->
<!--.title{-->
<!--  color: $violet !important;-->
<!--}-->
<!--.navbar{-->
<!--  background-color: $darkgrey !important;-->
<!--}-->
<!--.navbar-item, .navbar-link{-->
<!--  color: $green !important;-->
<!--  font-size: 20px;-->
<!--}-->
<!--$navbar-item-hover-background-color:  $violetpale;-->
<!--.navbar-burger{-->
<!--  color: $green !important;-->
<!--}-->
<!--$pagination-margin: 0.25rem !default;-->
<!--$pagination-color: $green !default;-->
<!--$pagination-border-color: $green !default;-->
<!--$pagination-active-color:  $peachy !default;-->
<!--$pagination-current-background-color: $violet !default;-->
<!--.navbar-menu{-->
<!--  background-color: $darkgrey !important;-->
<!--}-->
<!--@import "../../node_modules/bulma";-->
<!--@import "../../node_modules/buefy/src/scss/buefy.scss";-->
<!--</style>-->
