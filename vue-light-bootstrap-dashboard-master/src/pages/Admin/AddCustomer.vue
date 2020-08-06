<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-8">
          <add-customer-form>
          </add-customer-form>
        </div>
        <div class="col-md-4">
          <user-card>
          </user-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import AddCustomerForm from './Forms/AddCustomerForm.vue'
import UserCard from '../UserProfile/UserCard.vue'
import AddCustomerForm from "@/pages/Admin/Forms/AddCustomerForm";


export default {
name: "AddCustomer",
  components: {
    AddCustomerForm,
    UserCard
  },
}
</script>

<style scoped>

</style>
