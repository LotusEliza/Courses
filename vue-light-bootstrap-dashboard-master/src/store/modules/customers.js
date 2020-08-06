
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
  }
}
