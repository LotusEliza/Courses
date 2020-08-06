<template>
  <card class="card-user">
    <img slot="image" src="https://www.codingdojo.com/blog/wp-content/uploads/coding-bootcamps-vs-cs-degrees.png" alt="..."/>
    <div class="author">
      <a href="#">
        <img class="avatar border-gray" src="img/faces/face-0.jpg" alt="..."/>

        <h4 class="title">Mike Andrew<br />
          <small>michael24</small>
        </h4>
      </a>
    </div>
<!--    <p class="description text-center"> "Lamborghini Mercy <br>-->
<!--      Your chick she so thirsty <br>-->
<!--      I'm in that two seat Lambo"-->
<!--    </p>-->
<!--    <div slot="footer" class="text-center d-flex justify-content-center">-->
<!--      <button href="#" class="btn btn-simple"><i class="fa fa-facebook-square"></i></button>-->
<!--      <button href="#" class="btn btn-simple"><i class="fa fa-twitter"></i></button>-->
<!--      <button href="#" class="btn btn-simple"><i class="fa fa-google-plus-square"></i></button>-->
<!--    </div>-->
  </card>
</template>
<script>
  import Card from 'src/components/Cards/Card.vue'
  export default {
    components: {
      Card
    },
    data () {
      return {
        details: [
          {
            title: '12',
            subTitle: 'Files'
          },
          {
            title: '2GB',
            subTitle: 'Used'
          },
          {
            title: '24,6$',
            subTitle: 'Spent'
          }
        ]
      }
    },
    methods: {
      getClasses (index) {
        var remainder = index % 3
        if (remainder === 0) {
          return 'col-md-3 col-md-offset-1'
        } else if (remainder === 2) {
          return 'col-md-4'
        } else {
          return 'col-md-3'
        }
      }
    }
  }

</script>
<style>

</style>
