<template>
  <section>
    customers::{{customers[0]}}
    <Table :data="customers"
           :columns="columns"
           @update="updateCustomer"
           @add="addCustomer"
           @delete="deleteCustomer">
      <!--      <p>I'm a slot title</p>-->


    </Table>
  </section>
</template>

<script>
import Table from "@/pages/Admin/Table";
import {mapGetters} from "vuex";
import {ToastProgrammatic as Toast} from "buefy";

export default {
  name: "ListOfCustomers",
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
          field: 'Email',
          label: 'Email',
        },
        {
          field: 'Tel',
          label: 'Tel',
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
    this.$store.dispatch('customers/getCustomers');
  },
  computed: {
    ...mapGetters("customers", [
      'customers'
    ]),
  },
  methods:{
    updateCustomer(value){
      console.log("Update customer from child in parent!!!")
      console.log(value)
    },
    async deleteCustomer(values){
      console.log("Delete! customer from child in parent!!!")
      let count = null;
      for (let i = 0; i < values.length; i++) {
        const response = await this.$store.dispatch('customers/removeCustomer', values[i].ID);
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
    addCustomer(value){
      console.log("Add! customer from child in parent!!!")
      console.log(value)
    }
  }
}

</script>

