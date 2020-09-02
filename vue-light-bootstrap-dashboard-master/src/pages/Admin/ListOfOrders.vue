<template>
  <section>
    orders::{{orders[0]}}
    <Table :data="orders"
           :columns="columns"
           @update="updateOrder"
           @add="addOrder"
           @delete="deleteOrder">
      <!--      <p>I'm a slot title</p>-->


    </Table>
  </section>
</template>

<script>
import Table from "@/pages/Admin/Table";
import {mapGetters} from "vuex";
import {ToastProgrammatic as Toast} from "buefy";

export default {
  name: "ListOfOrders",
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
          numeric: true,
          centered: true
        },
        {
          field: 'IdCustomer',
          label: 'IdCustomer',
          centered: true
        },
        {
          field: 'IdCourse',
          label: 'IdCourse',
          centered: true
        },
        {
          field: 'Status',
          label: 'Status',
          centered: true
        },
        {
          field: 'DateCreate',
          label: 'Created',
        }
      ],
    }
  },
  mounted() {
    this.$store.dispatch('orders/getOrders');
  },
  computed: {
    ...mapGetters("orders", [
      'orders'
    ]),
  },
  methods:{
    updateOrder(orders){
      console.log("Update Order from child in parent!!!")
      console.log(orders)
      this.$router.push({ name: 'UpdateOrder',  params: {id: 2 }})
    },
    // deleteOrder(value){
    //   console.log("Delete! Order from child in parent!!!")
    //   console.log(value)
    // },
    addOrder(value){
      console.log("Add! Order from child in parent!!!")
      console.log(value)
    },
    async deleteOrder(values){
      console.log("Delete! customer from child in parent!!!")
      let count = null;
      for (let i = 0; i < values.length; i++) {
        const response = await this.$store.dispatch('orders/removeOrder', values[i].ID);
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
  }
}

</script>

