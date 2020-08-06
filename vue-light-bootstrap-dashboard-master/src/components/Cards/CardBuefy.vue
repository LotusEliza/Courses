<template>
  <div class="card"
       :class="{ active_green: isActive || activated && completed, not_active_grey: isNotActive, current_green: isCurrent || activated && !completed, active_hover: hover }"
       @mouseover="hover = true"
       @mouseleave="hover = false"
       @click="show ? clickMethodVideos() : clickMethodCourses()"
  >
    id_course:{{id_course}}
    <div class="card-image">
      <figure class="image is-4by3">
        {{image}}
        <img :src="require(`../../../public/img/courses/${image}`)" alt="Placeholder image">
      </figure>
    </div>
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title is-4">{{name}} <i v-if="completed" class="nc-icon nc-check-2"></i>
          </p>
        </div>
      </div>
      <div class="content">
        <p v-if="!show">
          {{description}}.
        </p>
        <template v-if="!show">
          <router-link v-if="activated && !completed" :to="{ name: 'CourseVideos', params: { id_course: id_course } }"><b-button type="is-info" expanded outlined>Перейти к обучению...<i class="nc-icon nc-ruler-pencil"></i></b-button></router-link>
          <router-link v-if="activated && completed" :to="{ name: 'Course', params: { id_course: id_course } }"><b-button type="is-info" expanded outlined>Курс завершен! Подробнее...<i class="nc-icon nc-refresh-02"></i></b-button></router-link>
          <router-link  v-if="!activated" :to="{ name: 'Course', params: { id_course: id_course } }"><b-button type="is-info" expanded outlined>Узнать подробнее... <i class="nc-icon nc-tap-01"></i></b-button></router-link>
        </template>
<!--        <template v-if="show && isActive">-->
<!--&lt;!&ndash;          <router-link  v-if="!activated" :to="{ name: 'CourseVideos', params: { id_course: id_course } }">Посмотреть видео еще раз</router-link>&ndash;&gt;-->
<!--        </template>-->
<!--          <slot></slot>-->
      </div>
    </div>
  </div>
</template>

<script>

  // import { bus } from "../../main";

  export default {
    name: 'card-buefy',
    data() {
      return {
        hover: false,
      };
    },
    props: {
      description: {
        type: String,
      },
      name: {
        type: String,
      },
      id_course: {
        type: Number,
      },
      image: {
        type: String,
      },
      activated: {
        type: Boolean,
      },
      show: {
        type: Boolean,
      },
      isActive: {
        type: Boolean,
      },
      isNotActive: {
        type: Boolean,
      },
      isCurrent: {
        type: Boolean,
      },
      video_sequence: {
        type: Number,
      },
      completed: {
        type: Boolean,
      }
    },
    mounted() {
      // bus.$on('router_link', (item) => {
      //   console.log(item)
      //   if(this.show) {
      //     this.clickMethodVideosTest(item)
      //   }
        // console.log("Bus bus")
        // this.$router.push({ name: 'CourseVideos', params: { id_course: item.id_course, setCurrentId: item.video_sequence }});
      // })

    },
    methods: {
      clickMethodVideos(){
        console.log('Click from videos')
        if(this.show && this.isActive || this.isCurrent){
          this.$router.push({ name: 'CourseVideos', params: { id_course: this.id_course, setCurrentId: this.video_sequence }});
        }else {
          console.log("Closed video!")
        }
      },

      clickMethodCourses(){
        console.log('Click from courses')
      }
    }
  }

</script>
<style scoped>
  .card .card-image {
    height: 180px !important;
  }
  .active_green{
    background: #bce6fe ;
  }
  .not_active_grey{
    background: lightgrey;
    filter: grayscale(100%);
  }
  .current_green{
    background: #87cdf6 ;
  }
  .active_hover{
    background: rgba(188, 230, 254, 0.7) ;
  }
</style>
