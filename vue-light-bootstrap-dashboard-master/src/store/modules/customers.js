import {ToastProgrammatic as Toast} from "buefy";
import {router} from "@/main";

export default {
  namespaced: true,
  state: {
    customers: [],
    // quiz: {},
  },
  getters: {
    customers: state => state.customers,
  },
  mutations: {
    SET_CUSTOMERS(state, items){
     state.customers = items
    },
    REMOVE_CUSTOMER(state, id){
      let c = state.customers.filter(s => s.ID != id);
      state.customers = c;
    }
  },
  actions: {
    async getCustomers({commit}) {
      let response = await window.axios.post('/get_customers', {});
      if (response.status == 200 || response.status == 204) {
        commit('SET_CUSTOMERS', response.data.Items);
      }
    },
    async removeCustomer({commit}, id) {
      let response = await window.axios.post('/customer_remove', {
        ID: id,
      });
      if (response.status == 200 || response.status == 204) {
        commit('REMOVE_CUSTOMER', id);
        return 'OK';
        // console.log("Removed user")
      }
    },
    async saveCustomer({commit, state}, user){
      let response = await window.axios.post('/save_customer',
        {
          Name: user.name,
          Tel: user.tel,
          Email: user.email,
          Password: user.password,
        },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {
        if (response.data.Error === "ERROR_REQUEST_DATA") {
          console.log("Mistake data request!")
          Toast.open({
            message: "Some problem occurred. please try again!",
            position: "is-top-right",
            type:"is-danger",
          })
          // commit('REMOVE_CODE_ACTIVE');
        }else{
          Toast.open({
            message: "Customer inserted!",
            position: "is-top-right",
            type:"is-success",
          });
          console.log("Inserted new customer")
        }
      }

    },
  }
}
