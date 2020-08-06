// import Vue from "vue";
import router from "../../main.js"



export default {
  namespaced: true,
  state: {
    videos: [],
    current_videos: [],
    current_watching_video: [],
    amount: 0,
  },
  getters: {
    videos: state => state.videos,
    amount: state => state.amount,
    current_videos: state => state.current_videos,
    video_id_by_sequence: (state)=> (sequence) => {
      let video = state.videos.find(v => v.Sequence === sequence);
      return video.ID
    },
    current_video_by_course_id: (state)=> (courseId) => {
      return state.current_videos.find(current_video => current_video.IdCourse === courseId)
    },
    current_watching_video: state => state.current_watching_video,
    // get_video_by_course_id: (state)=> (courseId) => {
    //   return state.current_videos.find(current_video => current_video.IdCourse === courseId)
    // }
  },
  mutations: {
    SET_VIDEOS(state, videos){
      state.videos = videos;
      state.amount = videos.length
    },
    SET_CURRENT_VIDEOS(state, current_videos){
      state.current_videos = current_videos;
    },
    UPDATE_CURRENT_VIDEO(state, array){
      let cv = state.current_videos.find(cv =>  cv.IdCourse ==  array.IdCourse);
      cv.CurrentVideo = array.CurrentVideo;
      let video = state.videos.find(v => v.Sequence === array.CurrentVideo);
      console.log("state.videos")
      console.log(state.videos)
      cv.IdVideo = video.ID;
    },
    SET_CURRENT_WATCHING_VIDEO(state, current_video){
      state.current_watching_video = current_video;
    },
    INCR_CURRENT_WATCHING_VIDEO(state){
      console.log('INCR_CURRENT_WATCHING_VIDEO::')
      // console.log(state.current_watching_video.CurrentVideo)
      state.current_watching_video.CurrentVideo++;
      // console.log('state.videos.find(v => v.Sequence === state.current_watching_video)')
      // console.log(state.videos.find(v => v.Sequence === state.current_watching_video))
      let video = state.videos.find(v => v.Sequence === state.current_watching_video.CurrentVideo)
      state.current_watching_video.IdVideo = video.ID;
      // console.log(state.current_watching_video.CurrentVideo)
    },
    RESET_CURRENT_WATCHING_VIDEO(state){
      state.current_watching_video = []
    }
  },
  actions: {
    async getVideos({commit, state}) {
      let response = await window.axios.post('/courses');
      for (let i = 0; i < response.data.Items.length; i++) {
        Vue.set(response.data.Items[i], "Activated", false);
      }
      for (let i = 0; i < state.activated.length; i++) {
        for (let j = 0; j < response.data.Items.length; j++) {
          if (response.data.Items[j].ID === state.activated[i].ID){
            response.data.Items[i].Activated = true
          }
        }
      }
      commit('SET_COURSES', response.data.Items);
    },
    async getVideosByCourse({commit}, id) {
      let response = await window.axios.post('/videos', {
        Id: id,
      });
      if (response.status == 200 || response.status == 204) {
        commit('SET_VIDEOS', response.data.Items);
      }
    },
    async getCurrentVideos({commit}, id) {
      let response = await window.axios.post('/customer_course_video', {
        Id: Number(id),
      });
      if (response.status == 200 || response.status == 204) {
        commit('SET_CURRENT_VIDEOS', response.data.Items);
      }
    },
    async updateCurrentVideo({commit, getters}, id_course){
      let current = getters.current_video_by_course_id(id_course)
      console.log(current)

      let completed, current_video
      if(current.CurrentVideo === getters.amount){
        completed = 1
        current_video = Number(current.CurrentVideo)
      }else{
        current_video = Number(current.CurrentVideo+1)
      }
      // let current= localStorage.getItem('CustomerId')
      let response = await window.axios.post('/current_video_update', {
        IdCourse: Number(id_course),
        IdCustomer: current.IdCustomer,
        CurrentVideo: current_video,
        Completed: completed,
      });
      if (response.status == 200 || response.status == 204) {
        let array = {'IdCourse': id_course, 'CurrentVideo': current.CurrentVideo+1}
        commit('UPDATE_CURRENT_VIDEO', array);
        // router.push({ name: 'CourseVideos', params: { id_course: id_course }});

      }
    }
  }
}
