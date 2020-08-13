import cookie from "vue-cookies";
import {ToastProgrammatic as Toast} from "buefy";
import {router} from "@/main";

export default {
  namespaced: true,
  state: {
    orders: [],
    // quiz: {},
  },
  getters: {
    orders: state => state.orders,
  },
  mutations: {
    SET_ORDERS(state, items){
     state.orders = items
    },
    REMOVE_ORDER(state, id){
      let o = state.orders.filter(s => s.ID != id);
      state.orders = o;
    },
  },
  actions: {
    async getOrders({commit}) {
      let response = await window.axios.post('/get_orders', {});
      if (response.status == 200 || response.status == 204) {
        commit('SET_ORDERS', response.data.Items);
      }
    },
    async removeOrder({commit}, id) {
      let response = await window.axios.post('/order_remove', {
        ID: id,
      });
      if (response.status == 200 || response.status == 204) {
        commit('REMOVE_ORDER', id);
        return 'OK';
        // console.log("Removed user")
      }
    },
    async saveCustomer({commit, state}, data) {
      // commit('AUTH_REQUEST')
      let response = await window.axios.post('/save_customer',
        {
          Password: data.password,
          Email: data.email,
          Name: data.name,
          Tel: data.tel,
        },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {
        if (response.data.Error === "ERROR_NO_SUCH_USER"){
          // localStorage.removeItem('UserID');
          cookie.remove('token')
          console.log("No such user in db vue js!!!")
          Toast.open({
            message: "Wrong email or name!",
            position: "is-top-right",
            type:"is-danger",
          });
          commit('AUTH_ERROR')
        } else if (response.data.Error === "ERROR_WRONG_PASSWORD"){
          cookie.remove('token')
          console.log("Wrong Password!!!")
          Toast.open({
            message: "Wrong Password!",
            position: "is-top-right",
            type:"is-danger",
          });
          commit('AUTH_ERROR')
        } else {
          console.log("Customer inserted!")
          // let myResult = cookie.get('token').split(".");
          // let decodedString = JSON.parse(atob(myResult[1]));
          // localStorage.setItem('UserID', decodedString.userid)
          // localStorage.setItem('Admin', decodedString.admin)
          // commit('AUTH_SUCCESS')
          // await router.push("/");
        }
      }
    },
  }
}
