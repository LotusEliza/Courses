
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
  }
}
