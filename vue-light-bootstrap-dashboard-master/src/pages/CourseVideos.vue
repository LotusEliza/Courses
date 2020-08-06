<template>
  <div class="content">
    videos::{{videos}}
    <div class="container-fluid">
      <div class="row">
        <div class="col-12">
          <card class="strpied-tabled-with-hover pd-20" body-classes="table-full-width table-responsive">
            <div v-for="(item, index) in videos">
              <template v-if="index+1 == 1 && $route.params.free_video">
                this is free first video
                <vue-player  :src="require(`../../public/videos/${item.Name}`)" :poster="require(`../../public/img/courses/${item.Poster}`)" :title="item.Title"></vue-player>
                <br>
                videos::{{videos}}
                item.ID::{{item.ID}}
                <div v-if="quiz && quiz.length !== 0">
                  <!----------------------------------- Quiz ---------------------------------------->
                  <quiz v-if="quiz && quiz.length !== 0" @scrollToTop="scrollToTop" :free_video="$route.params.free_video" :setCurrentId="$route.params.setCurrentId" :sequence="item.Sequence" :id_course="course_id" :id_video="item.ID" :key="componentKey" @update="updateKey"></quiz>
                </div>
              </template>

<!--              ------------------------Continue watching from video that user stopped on-->
              <div v-if="index+1 == current.CurrentVideo && !$route.params.free_video">
                current_video.Completed:: {{current_video}}

                <!------------------------------------- IF COURSE COMPLETED TEMPLATE------------------------------->
                <template v-if="current_video.Completed && !$route.params.setCurrentId">
                  <p>Вы успешно окончили данный курс и получили сертификат!</p>
                  <router-link  :to="{ name: 'CourseListOfVideos',  params: { course_id: course_id }}">Все видео курса</router-link>
                </template>
                <!------------------------------------- IF COURSE NOT COMPLETED TEMPLATE------------------------------->
                <template v-if="!current_video.Completed || $route.params.setCurrentId">
                  <!----------------------------------- Player ---------------------------------------->
                  <vue-player  :src="require(`../../public/videos/${item.Name}`)" :poster="require(`../../public/img/courses/${item.Poster}`)" :title="item.Title"></vue-player>
                  <br>
                  <div v-if="quiz && quiz.length !== 0">
                    <!----------------------------------- Quiz ---------------------------------------->
                    <quiz v-if="quiz && quiz.length !== 0" @scrollToTop="scrollToTop" :setCurrentId="$route.params.setCurrentId" :sequence="item.Sequence" :id_course="course_id" :id_video="current.IdVideo" :key="componentKey" @update="updateKey"></quiz>
                  </div>
                  <router-link  :to="{ name: 'CourseListOfVideos',  params: { course_id: course_id }}">Список всех видео этого курса</router-link>
                </template>
              </div>



            </div>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import {mapGetters} from "vuex";
  import vuePlayer  from  '@algoz098/vue-player'
  import quiz from 'src/components/Quiz.vue'


  export default {
    name: "CourseVideos.vue",
    components: {
      vuePlayer,
      quiz
    },
    data() {
      return {
        course_id: this.$route.params.id_course,
        setCurrentId: this.$route.params.setCurrentId,
        current: '',
        componentKey: 0,
        autoplay: true,
        preview_on_mouse: true,
      }
    },
    computed: {
      ...mapGetters("videos", [
        'videos',
      ]),
      ...mapGetters("quiz", [
        'quiz',
      ]),
      current_video() {
        return this.$store.getters['videos/current_video_by_course_id'](this.$route.params.id_course)
      },
    },
    mounted() {
      let currentVideo = this.$store.getters['videos/current_video_by_course_id'](this.$route.params.id_course)
      console.log("currentVideo::"+currentVideo)
      if(!this.current_video) {
        this.current = {"CurrentVideo": 0};
      }else if(this.$route.params.setCurrentId && this.$route.params.setCurrentId !== currentVideo.CurrentVideo){
        //------------click from ListOfVideos component-----------
        this.current_video_clone = Object.assign({}, this.current_video_clone, this.current_video);
        this.current_video_clone.CurrentVideo = this.$route.params.setCurrentId
        this.current_video_clone.IdVideo = this.$store.getters['videos/video_id_by_sequence'](this.$route.params.setCurrentId)
        this.$store.commit('videos/SET_CURRENT_WATCHING_VIDEO', this.current_video_clone);
        this.current = this.$store.getters['videos/current_watching_video'];
      }else {
        //-----------click nextVideo from Quiz component-------------
        this.current = this.current_video;
      }
      this.$store.dispatch('videos/getVideosByCourse', this.$route.params.id_course);
    },
    methods: {
      scrollToTop(){
        window.scrollTo({
          top: 80,
          behavior: "smooth"
        });
      },
      forceRerender() {
        this.componentKey += 1;
      },
      updateKey(){
        this.forceRerender()
      },
    },

  }

</script>
