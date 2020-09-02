<template>
  <div>
    This is UpdateOrder!!!
    orders:::{{orders}}
  </div>

</template>
<script>
export default {
name: "UpdateOrder",
  props: {
    orders: {
      type: Array
    }
  },
  mounted() {
  console.log(this.orders)
  }
}
</script>

<style scoped>

</style>
