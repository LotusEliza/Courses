<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-12">
          current_video:{{current_video}}
          <card class="strpied-tabled-with-hover" body-classes="table-full-width table-responsive">
            <div class="columns" v-for="videos in chunkedVideos">
              <div class="column" v-for="video in videos">
              <!----------------------------------- CardBuefy ---------------------------------------->
                <CardBuefy :name="video.Title"
                           :video_sequence="video.Sequence"
                           :id_course="course_id"
                           :image="video.Poster"
                           :show="true"
                           :isCurrent="isCurrent(video.Sequence, current_video.CurrentVideo)"
                           :isActive="isActiveFunc(video.Sequence, current_video.CurrentVideo)"
                           :isNotActive="isActiveFunc( current_video.CurrentVideo, video.Sequence-1)">
                </CardBuefy>

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
  import CardBuefy from "../components/Cards/CardBuefy";

  export default {
    name: "CourseListOfVideos.vue",
    components: {
      CardBuefy
    },
    data() {
      return {
        // TODO remove from localstorage after
        course_id: 0,
        // course_id: this.$route.params.course_id || localStorage.getItem('CourseID'),
        // current_video: [],
      }
    },
    mounted() {
      this.$store.dispatch('videos/getVideosByCourse', this.$route.params.id_course);
      if(this.$route.params.course_id){
        localStorage.setItem('CourseID', this.$route.params.course_id)
        this.current_video = this.$store.getters['videos/current_video_by_course_id'](this.$route.params.course_id)
        this.course_id = Number(this.$route.params.course_id)
      }else {
        this.current_video = this.$store.getters['videos/current_video_by_course_id'](Number(localStorage.getItem('CourseID')))
        this.course_id = Number(localStorage.getItem('CourseID'))
      }

    },
    computed: {
      ...mapGetters("videos", [
        'videos',
      ]),

      chunkedVideos() {
        return chunk(this.videos, 3)
      },
      current_video_by_course_id() {
        if(this.$route.params.course_id){
          return this.$store.getters['videos/current_video_by_course_id'](this.$route.params.course_id)
        }else {
          return this.$store.getters['videos/current_video_by_course_id'](Number(localStorage.getItem('CourseID')))
        }

      },
      current_video: {
        get(){
          return this.current_video_by_course_id
        },
        set(newName){
          return newName
        }
      }
    },
    methods: {
      isActiveFunc(a,b){
        if (a <= b){
          return true
        }else{
          return false
        }
      },
      isCurrent(a,b){
        if (this.current_video.Completed || a !== b){
          return false
        }else if(a === b){
            return true
        }
      }
    },
  }
</script>

<style>

</style>
