import Vue from 'vue'
import Vuex from 'vuex'
import courses from './modules/courses'
import videos from './modules/videos'
import quiz from './modules/quiz'
import login from './modules/login'
import customers from './modules/customers'
import orders from './modules/orders'

Vue.use(Vuex);

export const store = new Vuex.Store({
  strict: true,
  modules:{
    courses,
    videos,
    quiz,
    login,
    customers,
    orders,
  },
  state: {
  },
  mutations: {
  },
  getters : {
  }
});
