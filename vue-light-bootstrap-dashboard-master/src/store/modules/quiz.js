
export default {
  namespaced: true,
  state: {
    quiz: {},
  },
  getters: {
    quiz: state => state.quiz,
  },
  mutations: {
    SET_QUIZ(state, quiz){
      console.log("Set quiz vuex")
      console.log(quiz)
      state.quiz = {
        user: "User",
        questions: quiz,
      }
    },
    RESET_QUIZ(state){
     state.quiz = {}
    },
  },
  actions: {
    async getQuiz({commit}, id_video) {
      let response = await window.axios.post('/quiz', {
        Id: id_video,
      });
      if (response.status == 200 || response.status == 204) {
        commit('SET_QUIZ', response.data.Items);
      }
    },
  }
}
