<template>
    <div class="content">
      authStatus::{{authStatus}}
      isAuthenticated::{{isAuthenticated}}
      <div class="container-fluid">
        <div class="row">
          <div class="col-12" v-if="isAuthenticated">
            courses::{{courses}}
            <card class="strpied-tabled-with-hover"
                  body-classes="table-full-width table-responsive"
            >
              <div class="columns" v-for="courses in chunkedCourses">
                <div class="column" v-for="course in courses">
                  <CardBuefy :name="course.Name"
                             :description="course.Description"
                             :id_course="course.ID"
                             :image="course.Img"
                             :activated="course.Activated"
                             :completed="course.Completed"
                  ></CardBuefy>
                </div>
              </div>
            </card>
          </div>
        </div>
       <div>
       </div>
      </div>
    </div>
</template>

<script>
  import ModalBuefy from 'src/components/Modals/ModalBuefy.vue'
  import CardBuefy from 'src/components/Cards/CardBuefy.vue'
  import { mapGetters } from 'vuex';
  import cookie from 'vue-cookies';

  export default {
    components: {
      CardBuefy, ModalBuefy
    },
    data () {
      return {
        description: 'Hello 1st',
      }
    },
    computed: {
      ...mapGetters("courses", [
        'courses', 'activated',
      ]),
      ...mapGetters("login", [
        'isAuthenticated', 'authStatus'
      ]),
      chunkedCourses() {
        return chunk(this.courses, 2)
      },
      token() {
        return cookie.get('token')
      },

    },
    watch: {
      isLoggedIn: function (newVal, oldVal) {
        console.log("oldVal"+oldVal)
        if(newVal===''){
          this.$router.push({ name: 'LogIn' })
        }
      }
    },
    created() {
      // this.$store.dispatch('login/welcome').
      // then(() => this.checkIfLogIn());


      // this.checkIfLogIn()radio
    },
    mounted(){

      // current_watching_video
      // if(this.$prevRoute.name === "LogIn" || this.$prevRoute.name === 0){
      //   this.notifyVue('top', 'left', `<span>Welcome to <b>Online Courses</b> - our dear user!</span>`)
      // }
      this.$store.commit('videos/RESET_CURRENT_WATCHING_VIDEO')
      this.checkIfLogIn()
      ///-------------TODO Check if jwt valid-------------////
      this.$store.dispatch('courses/getActiveCourses');
      this.$store.dispatch('courses/getCourses');
      /////////////////////////////////////////////////
      // localStorage.setItem('CustomerId', "1")
      this.$store.dispatch('videos/getCurrentVideos',  localStorage.getItem('UserID'));
    },
    methods:{
      notifyVue (verticalAlign, horizontalAlign, msg) {
        const color = Math.floor((Math.random() * 4) + 1)
        this.$notifications.notify(
          {
            message: msg,
            icon: 'nc-icon nc-app',
            horizontalAlign: horizontalAlign,
            verticalAlign: verticalAlign,
            type: this.type[color]
          })
      },
      checkIfLogIn(){
        if(!cookie.get('token')){
          console.log("not valide or login")
          this.$router.push({ name: 'LogIn' })
        }
      },
      welcome(){
        this.$store.dispatch('login/welcome');
      },
    }

  }
</script>
<style scoped>

.column{
  padding-left: 5% !important;
}
</style>
