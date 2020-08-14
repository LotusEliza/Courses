<template>
  <div class="content">
    course_id::{{course_id}}
    <div class="container-fluid">
      <div class="row">
        <div class="col-12">
          <card class="strpied-tabled-with-hover pd-20" v-if="!current_video.Completed"
                body-classes="table-full-width table-responsive"
          >
            <h2 class="title is-2">{{course.Name}}</h2>
            <img :src="require(`../../public/img/courses/${course.Img}`)" alt="Placeholder image">
            <h4 class="subtitle is-4">Продолжительность курса: {{course.Duration}} ч</h4>
            <br>
            <div class="couses">
              <h6 class="subtitle is-6">Программа курса:</h6>
              <p v-html="course.Program"></p>
            </div>

            <h6 class="subtitle is-6">Описание курса:</h6>
            <p class="is-size-4	">
              {{course.Description}}
            </p>
            <router-link  :to="{ name: 'CourseVideos',  params: { id_course: $route.params.id_course, free_video: true }}"><b-button type="is-info" expanded outlined>Перейти к обучению!</b-button></router-link>

          </card>
          <!------------------------------------- IF COURSE COMPLETED TEMPLATE------------------------------->
          <div v-if="current_video.Completed && !$route.params.setCurrentId">
            <p>Вы успешно окончили данный курс и получили сертификат!</p>
            <router-link  :to="{ name: 'CourseListOfVideos',  params: { course_id: course_id, id_course: course_id, }}">Все видео курса</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>

    import {mapGetters} from "vuex";

    export default {
      name: "Course.vue",
      data () {
        return {
          course_id: this.$route.params.id_course,
        }
      },
      created() {
        console.log(this.$route.params.id_course)
        // this.course = this.course()
      },
      computed: {
        ...mapGetters("videos", [
          'current_videos'
        ]),
        course() {
          return this.$store.getters['courses/courseById'](this.$route.params.id_course)
        },
        current_video() {
          return this.$store.getters['videos/current_video_by_course_id'](this.$route.params.id_course)
        },
      },
      methods: {

      }

    }

</script>

<style scoped>
.pd-20{
  padding-left: 30px;
}
</style>
