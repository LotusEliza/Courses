<template>
  <div class="content">
  <template v-if="user.Email">
    user::{{user}}
    radio::{{radio}}
    <!--    _______________________________Form for choosing method that will send ver code_________________________-->
    <b-loading :is-full-page="isFullPage" :active.sync="isLoading" :can-cancel="true"></b-loading>
    <div class="container-fluid" v-if="!enterCodeActive">
      <card>
        <div class="row justify-content-center ">
          <div class="col-md-4">
          </div>
          <div class="col-md-4 pb-20" >
            <div class="typo-line  has-text-centered">
              <h2>Reset Your Password<br>
                <hr>
                <small>How do you want to recieve the code to reset your password?</small>
              </h2>

              <section>
                <div class="field level-left">
                  <b-radio v-model="radio"
                           native-value="email"
                           type="is-info">
                    <p>
                      <i class="nc-icon nc-email-83"></i> Send code via email
                    </p>
                    <small>{{user.Email}}</small>
                  </b-radio>


                </div>

                <br>
                <div class="field level-left">
                  <b-radio v-model="radio"
                           native-value="tel"
                           type="is-info">
                    <p>
                      <i class="nc-icon nc-send"></i>Send code via tel
                    </p>
                    <small>{{user.Tel}}</small>
                  </b-radio>
                </div>
                <hr>
                <div class="buttons level-right" >
                  <b-button type="is-info" @click="contin">Continue</b-button>
                </div>
              </section>

            </div>
          </div>
          <div class="col-md-4">
          </div>
        </div>
      </card>
    </div>

    <!--    _______________________________Form for entering ver code__________________________________________________-->
    <div class="container-fluid" v-if="enterCodeActive">
      <card>
        verCode::{{verCode}}
        Enter Ver Code now:
        <div class="row justify-content-center ">
          <div class="col-md-4">
          </div>
          <div class="col-md-4 pb-20" >
            <div class="typo-line  has-text-centered">
              <h2>Verification Code<br>
                <hr>
                <small>Please enter verification code that you've got:</small>
              </h2>
              <ValidationObserver ref="observer">
                <section slot-scope="{ validate }" class="level-left">
                  <b-field label="Password">
                    <BInputWithValidation rules="required|numeric|length:4"
                                          placeholder="Your password"
                                          v-model="verCode"
                    />
                  </b-field>
                  <hr>
                  <div class="buttons level-right" >
                    <b-button type="is-info" @click="validate().then(sendVerCode)">Verify</b-button>
                  </div>
                </section>
              </ValidationObserver>
            </div>
          </div>
          <div class="col-md-4">
          </div>
        </div>
      </card>
    </div>
  </template>
  <template v-else>{{redPage()}}</template>>

  </div>
</template>

<script>
    import {mapGetters} from "vuex";
    import BInputWithValidation from "../components/InputsValidation/BInputWithValidation";

    import { extend } from 'vee-validate';
    import { required, numeric, length} from 'vee-validate/dist/rules';
    import {ValidationObserver} from "vee-validate/dist/vee-validate.full";

    // Add the required rule
    extend('required', {
      ...required,
      message: 'This field is required'
    });
    // Add the numeric rule
    extend('numeric', {
      ...numeric,
      message: 'This field should contain only digits'
    });
    // Add the numeric rule
    extend('length', {
      ...length,
      message: 'The code should be 4 digits'
    });

    // extend('max', {
    //   validate(value, { length }) {
    //     return value.length >= length;
    //   },
    //   params: ['length'],
    //   message: 'The field must have not more then {length} characters'
    // });

    export default {
        name: "RecovPassLoggedIn.vue",
      data () {
        return{
          radio: 'email',
          verCode: '',
          isLoading: false,
          isFullPage: true
        }
        },
      components: {
        ValidationObserver,
        BInputWithValidation,
      },
      computed: {
        ...mapGetters("login", [
          'user', 'enterCodeActive'
        ]),
      },
      watch: {
        enterCodeActive: function (newVal, oldVal) {
          if(newVal){
            this.closeLoading()
          }
        }
      },

      // TODO redirect to enter restore code page
      // Tell user that email has been sent
      methods:{
        contin(){
          console.log("continue")
          this.openLoading()
            this.$store.dispatch('login/recovPassLoggedIn', this.radio);
        },
        sendVerCode(){
          console.log("sendVerCode")
          if(this.verCode && this.verCode.length == 4){
            this.$store.dispatch('login/sendVerCode', this.verCode);
          }
        },
        openLoading() {
          this.isLoading = true
          setTimeout(() => {
            this.isLoading = false
          }, 10 * 1000)
        },
        closeLoading(){
          this.isLoading = false
        },
        redPage(){
          console.log('No data!')
          this.$router.push({path: '/admin/user'});
        },
      },

      beforeDestroy() {
        this.$store.commit('login/REMOVE_CODE_ACTIVE')
      }
    }
</script>

<style scoped>
  .field p {
    padding-top: 35px;
  }

</style>
