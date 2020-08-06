<template>

      <card>
        <div class="row justify-content-center ">
          <div class="col-md-2">
          </div>
          <div class="col-md-7 col-md-offset-1 has-text-centered">
              <h1 >Welcome to our online courses!</h1>
              <h3>To continue please login to out system:</h3>
              <br>
<!-------------------------------------------Loading----------------------------------------------------------------->
              <b-field>
<!--                <button class="button is-primary is-medium" @click="openLoading">-->
<!--                  Launch loading-->
<!--                </button>-->
              </b-field>
                <b-loading :is-full-page="isFullPage" :active.sync="isLoading" :can-cancel="true"></b-loading>

<!------------------------------------------------------------------------------------------------------------------->

              <b-button type="is-info" class="is-large"
                        @click="isComponentModalActive = true">
                LogIn
              </b-button >

              <b-modal :active.sync="isComponentModalActive"
                       has-modal-card
                       trap-focus
                       :destroy-on-hide="true"
                       aria-role="dialog"
                       aria-modal>
                <modal-form @login='onLogin' @remind="remindPassword"></modal-form>
              </b-modal>

              <b-modal :active.sync="isComponentModalRestoreActive"
                       has-modal-card
                       trap-focus
                       :destroy-on-hide="true"
                       aria-role="dialog"
                       aria-modal>
                <modal-form-restore @restore="OnRestore" @closeRestore="closeRestore"></modal-form-restore>
              </b-modal>

          </div>
          <div class="col-md-2">
          </div>
        </div>
      </card>
<!--    </div>-->
<!--  </div>-->
</template>

<script>
  import ModalForm from './ModalForm.vue'
  import ModalFormRestore from './ModalFormRestore.vue'
  import {ToastProgrammatic as Toast} from "buefy";

  export default {
    components: {
      ModalForm,
      ModalFormRestore,
    },
    data() {
      return {
        isComponentModalActive: false,
        isComponentModalRestoreActive: false,
        isLoading: false,
        isFullPage: true
      }
    },

    methods: {
      openLoading() {
        this.isLoading = true
        setTimeout(() => {
          this.isLoading = false
        }, 10 * 1000)
      },
      onLogin(data){
        this.$store.dispatch('login/loginSend', { email: data.email, name: data.name, password: data.password }).then(() => {
          if(this.$store.getters['login/authStatus'] == 'error'){
            this.close()
            // Toast.open({
            //   message: "No such User in DB! Please Try again!",
            //   position: "is-top-right",
            //   type:"is-danger",
            //   duration: 4000
            // });
          }else {
            this.close()
          }
        })

      },
      OnRestore(data){
        this.openLoading()
        this.$store.dispatch('login/restoreStart', { email: data.email, name: data.name}).then(() => {
          if(this.$store.getters['login/authStatus'] == 'error'){
            this.isLoading = false
            this.close()
            Toast.open({
              message: "No such User in DB! Please Try again!",
              position: "is-top-right",
              type:"is-danger",
              duration: 4000
            });
          }else {
            this.close()
            Toast.open({
              message: "The email with a link has been sent to your email address!",
              position: "is-top-right",
              type:"is-success",
              duration: 4000
            });
          }
        })

        this.closeRestore()
      },
      close(){
        this.isComponentModalActive = false
      },
      remindPassword(){
        this.close()
        this.isComponentModalActive = false
        this.isComponentModalRestoreActive = true
      },
      closeRestore(){
        this.isComponentModalRestoreActive = false
      },
    }
  }
</script>
<style>
  /*.center{*/
  /*  text-align: center;*/
  /*}*/
</style>
