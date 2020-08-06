import Vue from "vue";

export default {
  namespaced: true,
  state: {
    courses: [],
    activated: [],
    data:[],
  },
  getters: {
    courses: state => state.courses,
    courseById: (state) => (id) => {
      return state.courses.find(course => course.ID === id)
    },
    activated: state => state.activated,
    data: state => state.data
  },
  mutations: {
    SET_COURSES(state, courses){
      state.courses = courses;
    },
    SET_ACTIVATED_COURSES(state, activated){
      state.activated = activated;
    },
    SET_DATA(state, data){
      state.data = data
    },
    REMOVE_COURSE(state, id){
      let c = state.courses.filter(s => s.ID != id);
      state.courses = c;
    }
  },
  actions: {
    async getCourses({commit, state}) {
      let response = await window.axios.post('/courses');
      for (let i = 0; i < response.data.Items.length; i++) {
        Vue.set(response.data.Items[i], "Activated", false);
        Vue.set(response.data.Items[i], "Completed", false);
      }
      for (let i = 0; i < state.activated.length; i++) {
      for (let j = 0; j < response.data.Items.length; j++) {
        if (response.data.Items[j].ID === state.activated[i].ID){
          response.data.Items[i].Activated = true
          if (response.data.Items[j].ID === state.activated[i].ID && state.activated[i].Completed){
            response.data.Items[i].Completed = true
          }
        }
      }
      }
      commit('SET_COURSES', response.data.Items);
    },
    async getActiveCourses({commit}) {
      let response = await window.axios.post('/courses/activated', {
        Id: Number(localStorage.getItem('UserID')),
      });
      if (response.status == 200 || response.status == 204) {
        commit('SET_ACTIVATED_COURSES', response.data.Items);
      }
    },
    async buyCourse({commit}, id_course){
      let id_customer = 1
      let response = await window.axios.post('/place_order', {
        IdCourse: id_course,
        IdCustomer: id_customer
      });
      if (response.status == 200 || response.status == 204) {
        commit('SET_DATA', response.data);
      }
    },
    async removeCourse({commit}, id) {
      let response = await window.axios.post('/course_remove', {
        ID: id,
      });
      if (response.status == 200 || response.status == 204) {
        commit('REMOVE_COURSE', id);
        return 'OK';
        // console.log("Removed user")
      }
    },
  }
}
