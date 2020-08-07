/*!

 =========================================================
 * Vue Light Bootstrap Dashboard - v2.0.0 (Bootstrap 4)
 =========================================================

 * Product Page: http://www.creative-tim.com/product/light-bootstrap-dashboard
 * Copyright 2019 Creative Tim (http://www.creative-tim.com)
 * Licensed under MIT (https://github.com/creativetimofficial/light-bootstrap-dashboard/blob/master/LICENSE.md)

 =========================================================

 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

 */
import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import JQuery from 'jquery'
import {store} from './store/store'
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'
// import vuePlayer  from  '@algoz098/vue-player'
// import VueCryptojs from 'vue-cryptojs'
import cookie from 'vue-cookies'
//
import VueCookies from 'vue-cookies'
Vue.use(VueCookies)

// Vue.component(vuePlayer)
Vue.use(Buefy)

// LightBootstrap plugin
import LightBootstrap from './light-bootstrap-main'

// router setup
import routes from './routes/routes'

import './registerServiceWorker'
// plugin setup
Vue.use(VueRouter)
Vue.use(LightBootstrap)

//vuex setup
import Vuex from 'vuex'
//VeeValidate

import VueLodash from 'vue-lodash'
import lodash from 'lodash'

//vue-tel-input
import VueTelInput from 'vue-tel-input'

Vue.use(VueTelInput)

// import ModalSessionOver from 'components/Modals/ModalSessionOver.vue';
// Vue.component('ModalSessionOver', ModalSessionOver);

Vue.component('Alert',
    () => import('./components/Alerts/Alert.vue')
)

// name is optional
Vue.use(VueLodash, { name: 'custom' , lodash: lodash })

Vue.use(Vuex)



//axios setup
window.$ = JQuery;
window.axios = require("axios");
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
// window.axios.defaults.headers.common['Origin'] = 'http://localhost:3000';
window.axios.defaults.baseURL = "http://localhost:3000/";
// window.axios.defaults.headers.Authorization = "Bearer NOTOKEN";
// window.axios.defaults.headers.common['Authorization'] = store.getters['login/isLoggedIn'];

// window.axios.defaults.withCredentials = true
// window.axios.defaults.headers.Authorization = "Bearer NOTOKEN";

// axios.get(url, { crossdomain: true })
// window.axios.get('http://localhost:3000/', {withCredentials: true}, { crossdomain: true }); //for GET
// window.axios.post('http://localhost:3000/', data, {withCredentials: true}, { crossdomain: true }); //for POST
// window.axios.put('http://localhost:3000/', data, {withCredentials: true}, { crossdomain: true }); //for PUT
// window.axios.delete('http://localhost:3000/', data, {withCredentials: true}, { crossdomain: true }); //for Delete


// const myBody = window.getElementsByTagName('body')[0];
//chunk setup
window.chunk = require("chunk");

// configure router
export const router = new VueRouter({
  routes, // short for routes: routes
  linkActiveClass: 'nav-item active',
  scrollBehavior: (to) => {
    if (to.hash) {
      return {selector: to.hash}
    } else {
      return { x: 0, y: 0 }
    }
  }
})





// Vue.mixin({
//   methods: {
//     notEmptyObject(someObject){
//       return Object.keys(someObject).length
//     }
//   }
// })

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store: store,
  render: h => h(App),
  router,
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)){
    // if(store.getters['login/isLoggedIn']){
    if(cookie.get('token')){
      console.log("is login")
      // store.dispatch('login/refresh');
      console.log("from")
      console.log(from)
      console.log("to")
      console.log(to)
      if(from.name === "LogIn" || from.name === "Restore"){
        console.log("from LogIn")
      }else {
        console.log("Not from LogIn")
        store.dispatch('login/refresh');
      }
      Vue.prototype.$prevRoute = from
      next()
      return
    }else{
      store.commit('login/AUTH_LOGOUT');
      console.log("not login")
      next('/log-in')
      //clean the vuex state
      location.reload()
    }
  } else {
    // store.dispatch('login/refresh');
    console.log("not protected")
    next()
  }
})
