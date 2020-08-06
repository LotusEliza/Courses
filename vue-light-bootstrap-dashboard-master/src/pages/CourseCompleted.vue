<template>
  <!------------------------------------- IF COURSE COMPLETED TEMPLATE------------------------------->
<!--  <div v-if="current_video.Completed && !$route.params.setCurrentId">-->
  <div >
    <p>Вы успешно окончили данный курс и получили сертификат!</p>
    <router-link  :to="{ name: 'CourseListOfVideos',  params: { course_id: course_id }}">Все видео курса</router-link>
  </div>
</template>

<script>
    export default {
        name: "CourseCompleted",
      data() {
        return {
          course_id: this.$route.params.course_id
        }
      },
      mounted() {
        this.$store.dispatch('videos/getVideosByCourse', this.$route.params.id_course);
        if(this.$route.params.course_id){
          localStorage.setItem('CourseID', this.$route.params.course_id)
          this.current_video = this.$store.getters['videos/current_video_by_course_id'](this.$route.params.course_id)
          this.course_id = Number(this.$route.params.course_id)
        }else {
          this.current_video = this.$store.getters['videos/current_video_by_course_id'](localStorage.getItem('CourseID'))
          this.course_id = Number(localStorage.getItem('CourseID'))
        }
      },
      computed:{
        current_video_by_course_id() {
          if(this.$route.params.course_id){
            return this.$store.getters['videos/current_video_by_course_id'](this.$route.params.course_id)
          }else {
            return this.$store.getters['videos/current_video_by_course_id'](localStorage.getItem('CourseID'))
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
      }
    }
</script>

<style scoped>

</style>
