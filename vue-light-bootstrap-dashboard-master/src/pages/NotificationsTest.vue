<template>
  <div class="content">
    <div class="container-fluid">
      <card>
        <div class="places-buttons">
          <div class="row justify-content-center">
            <div class="col-md-3">
              <button class="btn btn-default btn-block" @click="notifyVue('bottom', 'right')">Bottom Right</button>
            </div>
          </div>
        </div>
      </card>
    </div>
  </div>
</template>
<script>
  import Card from 'src/components/Cards/Card.vue'

  export default {
    components: {
      Card
    },
    data () {
      return {
        type: ['', 'info', 'success', 'warning', 'danger'],
        notifications: {
          topCenter: false
        }
      }
    },
    methods: {
      notifyVue (verticalAlign, horizontalAlign) {
        const color = Math.floor((Math.random() * 4) + 1)
        this.$notifications.notify(
          {
            message: `<span>Welcome to <b>Light Bootstrap Dashboard</b> - a beautiful freebie for every web developer.</span>`,
            icon: 'nc-icon nc-app',
            horizontalAlign: horizontalAlign,
            verticalAlign: verticalAlign,
            type: this.type[color]
          })
      }
    }
  }

</script>
<style lang="scss">

</style>
