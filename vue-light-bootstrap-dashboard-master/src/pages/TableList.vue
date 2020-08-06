<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-12">
          <card class="strpied-tabled-with-hover"
                body-classes="table-full-width table-responsive"
          >
            <Quiz></Quiz>
          </card>

        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import Card from 'src/components/Cards/Card.vue'
  import Quiz from 'src/components/Quiz.vue'
  export default {
    components: {
      Card,
      Quiz
    },
    data () {
      return {
      }
    }
  }
</script>
<style>
</style>
