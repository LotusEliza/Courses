<template>
  <section>
    courses::{{courses[0]}}
    <Table :data="courses"
           :columns="columns"
           @update="updateCourse"
           @add="addCourse"
           @delete="deleteCourse">
<!--      <p>I'm a slot title</p>-->


    </Table>
  </section>
</template>

<script>
  import Table from "@/pages/Admin/Table";
  import {mapGetters} from "vuex";
  import {ToastProgrammatic as Toast} from "buefy";

  export default {
    name: "ListOfCourses",
    components: {
      Table
    },
    data () {
      return {
        columns: [
          {
            field: 'ID',
            label: 'ID',
            width: '40',
            numeric: true
          },
          {
            field: 'Name',
            label: 'Name',
          },
          {
            field: 'Description',
            label: 'Description',
          },
          {
            field: 'Price',
            label: 'Price',
            centered: true
          },
          {
            field: 'Duration',
            label: 'Duration',
          }
        ],
      }
    },
    mounted() {
      this.$store.dispatch('courses/getCourses');
    },
    // columns: [
    //   {
    //     field: 'ID',
    //     label: 'ID',
    //     width: '40',
    //     numeric: true
    //   },
    //   {
    //     field: 'Name',
    //     label: 'Name',
    //   },
    //   {
    //     field: 'Description',
    //     label: 'Description',
    //   },
    //   {
    //     field: 'Price',
    //     label: 'Price',
    //     centered: true
    //   },
    //   {
    //     field: 'Duration',
    //     label: 'Duration',
    //   }
    // ]
    computed: {
      ...mapGetters("courses", [
        'courses'
      ]),
    },
    methods:{

      updateCourse(value){
        console.log("Update course from child in parent!!!")
        console.log(value)
      },
      // deleteCourse(value){
      //
      //   console.log("Delete! course from child in parent!!!")
      //   console.log(value)
      // },
      async deleteCourse(values){
        console.log("Delete! Course from child in parent!!!")
        let count = null;
        for (let i = 0; i < values.length; i++) {
          const response = await this.$store.dispatch('courses/removeCourse', values[i].ID);
          if(response === 'OK'){
            console.log(response);
            count++;
          }
        }
        if(values.length === count){
          Toast.open({
            message: "removed!",
            type: 'is-danger'
          });
        }
      },
      addCourse(value){
        console.log("Add! course from child in parent!!!")
        console.log(value)
      }
    }
  }

</script>

