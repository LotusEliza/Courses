<template>
  <section id="app" class="hero is-warning is-fullheight" v-if="showData">
    sequence::{{sequence}}
    amount::{{amount}}
    current_watching_video::{{current_watching_video}}

    <!--heroBody-->
    <div class="hero-body">

      <!--container-->
      <div class="container is-fluid">

        <!--columns-->
        <div class="columns is-centered is-marginless is-paddingless box">

          <!--sidebar-->
          <div class="column is-5 sidebarQuiz">
            <div class="sidebarContainer">

              <div class="tags is-marginless">
                <span class="tag is-medium is-rounded">#vuejs</span>
                <span class="tag is-medium is-rounded">первое занятие</span>
              </div>

              <!-- title -->
              <h1 class="title is-2">Тесты</h1>
              <h2 class="subtitle is-5">Необходимо ответить правильно на все вопросы.</h2>
              <br>
              <br>
              <br>
              <br>
              <!-- description -->
              <!--              <p class="subtitle is-6 has-text-justified">Created using <em>VueJs</em>, <em>Bulma</em>, <em>animate.css</em>, and <em>Font Awesome 4</em>.</p>-->
            </div>

            <div class="sidebarFooter">
              <p>Photo by <a href="https://unsplash.com/@debidiemski">Deborah Diem</a></p>
            </div>
          </div>
          <!--/sidebar-->

          <!--questionBox-->
          <div class="column is-7 questionBox is-paddingless is-marginless">

            <!-- transition -->
            <transition enter-active-class="animated jackInTheBox" leave-active-class="animated zoomOut" mode="out-in">

              <div v-if="!quizStarted" v-bind:key="quizStarted" class="quizForm">
                <div class="titleContainer">
                  <h2 class="title is-4">Getting Started</h2>
                </div>

                <div class="quizFormContainer">
                  <div class="field">
                    <div class="field-label is-normal is-size-5">
                      <label class="label">Name</label>
                    </div>
                    <div class="control">
                      <input class="input is-rounded" type="text" v-model="quiz.user" placeholder="Enter your name" required>
                    </div>
                  </div>
                  <div class="field">
                    <div class="field-label is-normal is-size-5">
                      <label class="label">Difficulty</label>
                    </div>

                    <div class="control">
                      <label class="radio">
                        <input type="radio" name="0" checked>
                        Easy
                      </label>
                      <label class="radio">
                        <input type="radio" name="1">
                        Medium
                      </label>
                      <label class="radio">
                        <input type="radio" name="2">
                        Hard
                      </label>
                    </div>
                  </div>

                  <a class="pagination-previous button is-medium is-info is-rounded is-outlined is-pulled-right" v-on:click="quizLength" style="margin-bottom:12px">
                    Start
                  </a>
                </div>
              </div>

              <!--qusetionContainer-->
              <div class="questionContainer" v-if="questionIndex<quiz.questions.length && quizStarted" v-bind:key="questionIndex">

                <!-- questionTitle -->
                <div class="titleContainer">
                  <h2 class="title is-5">{{questionIndex+1}}. {{ quiz.questions[questionIndex].Text }}</h2>
                </div>

                <!-- quizOptions -->
                <div class="optionContainer">
                  <div class="option" v-for="(response, index) in quiz.questions[questionIndex].Responses" @click="selectOption(index)" :class="{ 'is-selected': userResponses[questionIndex] == index}" :key="index">
                    {{ index | charIndex }}. {{ response.Text }}
                  </div>
                </div>

                <!--quizFooter: navigation and progress-->
                <div class="questionFooter">

                  <!--pagination-->
                  <nav class="pagination is-centered" role="navigation" aria-label="pagination">

                    <!-- prevButton -->
                    <a class="pagination-previous button is-info is-rounded is-outlined" v-on:click="prev()" :disabled="questionIndex < 1">
                      Previous Question
                    </a>

                    <!-- prevButton -->
                    <a class="pagination-next button is-info is-rounded" v-on:click="next()" :disabled="questionIndex>=quiz.questions.length">
                      {{ (userResponses[questionIndex]==null)?'Skip':'Next' }} Question
                    </a>

                  </nav>
                  <!--/pagination-->

                  <!--progress-->
                  <div class="progressContainer">
                    <progress class="progress is-success is-small" :value="(questionIndex/quiz.questions.length)*100" max="100">{{(questionIndex/quiz.questions.length)*100}}%</progress>
                  </div>
                  <!--/progress-->

                </div>
                <!--/quizFooter-->

              </div>
              <!--/questionContainer-->

              <!--quizCompletedResult-->
              <div v-if="questionIndex >= quiz.questions.length && quizStarted" v-bind:key="questionIndex" class="quizCompleted has-text-centered">

                <!-- quizCompletedIcon: Achievement Icon -->
                <span class="icon is-large has-text-success">
                <i class="fa fa-check-circle-o fa-3x"></i>
              </span>

                <!--resultTitleBlock-->
                <h2 class="title">
                  Хорошая работа!
                </h2>
                <p class="subtitle">
                  Total score: {{ score() }} / {{ quiz.questions.length }}
                </p>
<!--                -----------------------------IF TEST CORRECT-------------------------------------->
                current_watching_video.length:::{{current_watching_video.length}} === 0
                sequence ::: {{sequence}} !== amount :::{{amount}} && free_video {{free_video}}
                <template v-if="score()/quiz.questions.length === 1">
                  <b-button v-if="current_watching_video.length === 0 && sequence !== amount && free_video" size="is-medium" @click="pay()">
                    Следующее видео (Оплата)
                  </b-button>
                  <b-button v-if="current_watching_video.length === 0 && sequence !== amount && !free_video" size="is-medium" @click="nextVideo() + scrollToTop()">
                    Следующее видео
                  </b-button>
                  <b-button v-if="current_watching_video.length !== 0 && sequence !== amount" @click="incrementWatchingVideo()">
                    Cледующее видео Watching
                  </b-button>
<!--                  -----------------------------if last video-------------------------------->
                  <b-button v-if="sequence === amount && !current_watching_video.Completed" size="is-medium" @click="getCertificate()">
                    Получить сертификат
                  </b-button>
                  <p v-if="current_watching_video.Completed && sequence === amount">Курс окончен!</p>
                </template>

<!--                ------------------------------IF TEST INCORRECT------------------------------------>
                <div v-if="score()/quiz.questions.length < 1">
                    <b-button  size="is-medium" @click="update">Пройти тест еще раз</b-button>
                    <br>
                    <p>или</p>
                    <b-button size="is-medium" @click="scrollToTop">Смотреть видео еще раз</b-button>
                </div>

                <!--/resultTitleBlock-->

              </div>
              <!--/quizCompetedResult-->

            </transition>

          </div>
          <!--/questionBox-->

        </div>
        <!--/columns-->

      </div>
      <!--/container-->

    </div>
    <!--/heroBody-->

  </section>
</template>

<script>
  import {mapGetters} from "vuex";

  export default {
    name: "Quiz",
    props: ['id_video', 'id_course', 'sequence', 'setCurrentId', 'free_video'],
    data() {
      return {
        questionIndex: 0,
        userResponses: [],
        quizStarted: true,
        isActive: false,
        showData: false,
      }
    },
    computed: {
      ...mapGetters("quiz", [
        'quiz',
      ]),
      ...mapGetters("videos", [
        'current_watching_video', 'amount'
      ]),
    },
    mounted() {
      console.log("mounted")
      if(Object.keys(this.$store.getters['quiz/quiz']).length === 0){
        console.log("If")
        console.log(this.id_video)
        this.getQuiz(this.id_video)
      }else {
        console.log("Else")
        this.getQuiz(this.$store.getters['videos/current_video_by_course_id'](this.id_course).IdVideo)
      }
    },
    filters: {
      charIndex: function(i) {
        return String.fromCharCode(97 + i);
      }
    },
    methods: {
      async getQuiz(id_video){
        try {
          await this.$store.dispatch('quiz/getQuiz', id_video);
        } catch (error) {
          console.log('error')
        } finally {
          this.userResponses = Array(this.quiz.questions.length).fill(null)
          this.showData = true
        }
      },
      scrollToTop(){
        this.$emit('scrollToTop')
      },
      quizLength(){
        if(quiz.user.length>0){quizStarted=!quizStarted}
      },
      selectOption: function(index) {
        this.$set(this.userResponses, this.questionIndex, index);
      },
      next: function() {
        if (this.questionIndex < this.quiz.questions.length)
          this.questionIndex++;
      },
      prev: function() {
        if (this.quiz.questions.length > 0) this.questionIndex--;
      },
      // Return "true" count in userResponses
      score: function() {
        var score = 0;
        for (let i = 0; i < this.userResponses.length; i++) {
          if (
            typeof this.quiz.questions[i].Responses[
              this.userResponses[i]
              ] !== "undefined" &&
            this.quiz.questions[i].Responses[this.userResponses[i]].Correct
          ) {
            score = score + 1;
          }
        }
        return score;
      },
      async update(){
        try {
          this.userResponses = Array(quiz.questions.length).fill(null);
          console.log("After Try: "+this.userResponses)
        } catch (error) {
          console.log('error')
        } finally {
          this.$emit('update')
        }
      },
      async nextVideo(){
        let currentVideo = this.$store.getters['videos/current_video_by_course_id']((this.id_course))
          if(this.sequence === currentVideo.CurrentVideo){
            try {
              await this.$store.dispatch('videos/updateCurrentVideo', this.id_course);
            } catch (error) {
              console.log('error')
            } finally {
              await this.update()
            }
          }else {
            console.log('Already watched video')
          }
        },
      async incrementWatchingVideo(){
        if(this.current_watching_video.CurrentVideo < this.$store.getters['videos/current_video_by_course_id'](this.current_watching_video.IdCourse).CurrentVideo){
          try {
            await this.$store.commit('videos/INCR_CURRENT_WATCHING_VIDEO');
            console.log("incrementWatchingVideo vuex")
            let current_watching_video = this.$store.getters['videos/current_watching_video']
            // await this.getQuiz(this.$store.getters['videos/video_id_by_sequence'](current_watching_video.CurrentVideo))
          } catch (error) {
            console.log('error')
          } finally {
            await this.update()
          }
        }else {
          console.log("Is empty watching video[]")
            try {
              await this.$store.commit('videos/INCR_CURRENT_WATCHING_VIDEO');
              await this.$store.dispatch('videos/updateCurrentVideo', this.id_course);
            } catch (error) {
              console.log('error')
            } finally {
              await this.update()
            }
        }

      },
      getCertificate(){
        console.log('Certificate!')
        this.$store.dispatch('videos/updateCurrentVideo', this.id_course);
      },
      pay(){
        console.log("Payment page redirect")
        this.$router.push({ name: 'BeforePayment', params: { id_course: this.id_course }});
        this.$store.dispatch('courses/buyCourse', this.id_course);
      }

    },
    beforeDestroy(){
      this.$store.commit('quiz/RESET_QUIZ');
      // this.$store.commit('videos/RESET_CURRENT_WATCHING_VIDEO');
    }
  }
</script>

<style lang="scss" >

  $trans_duration: 0.3s;
  $titleBg: #29c5dc;

  @import url("https://fonts.googleapis.com/css?family=Montserrat:400,400i,700");
  @import url("https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700");

  body {
    font-family: "Open Sans", sans-serif;
    font-size: 14px;

    /* mocking native UI */
    cursor: default !important; /* remove text selection cursor */
    user-select: none; /* remove text selection */
    user-drag: none; /* disbale element dragging */
  }

  .button {
    transition: $trans_duration;
  }
  .title,
  .subtitle {
    font-family: Montserrat, sans-serif;
    font-weight: normal;
  }
  .animated {
    transition-duration: $trans_duration/2;
  }

  .sidebarQuiz {
    background: url("https://source.unsplash.com/RVF0ngUujks");
    background-size: cover;

    border-radius: 6px 0px 0px 6px;
    z-index: 10;

  .sidebarContainer {
    padding: 10px;
    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.25);

  .tag {
    background: rgba(255, 255, 255, 0.25);
    color: rgba(255, 255, 255, 0.8);
  }

  .title {
    color: white !important;
  }
  .subtitle {
    color: rgba(255, 255, 255, 0.8) !important;
  }
  }
  .sidebarFooter {
    display: none;
    color: red;
  }
  }

  .questionBox {
    height: 100%;
    position: relative;
    display: flex;

  .titleContainer {
    background: $titleBg;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
    margin-bottom: 12px;

  h2 {
    color: white;
    padding: 18px;
  }
  }

  .quizForm {
    display: block;
    white-space: normal;

    height: 100%;
    width: 100%;

  .quizFormContainer {
    height: 100%;
    margin: 15px 18px;

  .field-label {
    text-align: left;
    margin-bottom: 0.5rem;
  }
  }
  }
  .quizCompleted {
    width: 100%;
    padding: 25px;
  }
  .questionContainer {
    white-space: normal;

    height: 100%;
    width: 100%;

  .optionContainer {
    margin-top: 12px;
    flex-grow: 1;
  .option {
    border-radius: 290486px;
    padding: 9px 18px;
    margin: 0 18px;
    margin-bottom: 12px;
    transition: $trans_duration;
    cursor: pointer;
    background-color: rgba(0, 0, 0, 0.05);
    border: transparent 1px solid;

  &.is-selected {
     border-color: #209cee;
     background-color: white;
   }
  &:hover {
     background-color: rgba(0, 0, 0, 0.1);
   }
  &:active {
     transform: scaleX(0.9);
   }
  }
  }

  .questionFooter {
    width: 100%;
    align-self: flex-end;

  .pagination {
  //padding: 10px 15px;
    margin: 15px 25px;
  }
  .progressContainer {
    margin: 15px 25px;
  }
  }
  }
  }
  @media screen and (min-width: 769px) {
    .container {
      height: 100%;

    .columns {
      height: 100%;

    .column {
      height: 100%;
    }
  }
  }

  .sidebarQuiz {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .questionBox {
    align-items: center;
    justify-content: center;

  .questionContainer {
    display: flex;
    flex-direction: column;
  }
  }
  }

  @media screen and (max-width: 768px) {
    .sidebarQuiz {
      height: auto !important;
      border-radius: 6px 6px 0px 0px;
    }
  }


</style>
