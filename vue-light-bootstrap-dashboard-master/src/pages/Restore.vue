<template>
  <div class="content">
    <template v-if="tokenForPass || codeForPass">
    <div class="container-fluid">
      <card>
          <div class="row justify-content-center ">
            <div class="col-md-3 col-md-offset-1">
            </div>
            <div class="col-md-3 col-md-offset-1 pb-20" >
              <div class="typo-line  has-text-centered">
                <h2>Reset Password<br>
                  <small>Please enter your new password below.</small>
                </h2>
              </div>
              <br>
              <ValidationObserver ref="observer">
                <section slot-scope="{ validate }">

<!--                  <div v-if="tokenForPass || codeForPass">-->
                  <b-field label="Password">
                    <BInputWithValidation rules="required|min:5"
                                          placeholder="Your password"
                                          v-model="info.password"
                                          :refer="refer"
                                          :password_reveal="password_reveal"
                                          type="password"
                                          name="password"
                    />
                  </b-field>

                  <b-field label="Password Repeat">
                    <BInputWithValidation rules="required|confirmed:password"
                                          placeholder="Your confirm password"
                                          v-model="info.repeatPassword"
                                          :password_reveal="password_reveal"
                                          type="password"
                                          name="password_confirmation"
                                          data-vv-as="password"
                    />
                  </b-field>

                  <div class="buttons">
                    <b-button type="is-info" v-if="codeForPass" @click="validate().then(setNewPassword)">Send codeForPass Logged In</b-button>
                    <b-button type="is-info" v-if="tokenForPass" @click="validate().then(updatePassword)">Send tokenForPass Not Logged In</b-button>
                  </div>

<!--                </div>-->
                </section>
              </ValidationObserver>
<!--              <div v-if="passwordUpdated">-->
<!--                Your password has been updated!-->
<!--              </div>-->
<!--              <div v-if="!tokenForPass">-->
<!--                No Form Token invalid!!!-->
<!--              </div>-->

          </div>
            <div class="col-md-3 col-md-offset-1">
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
    import {
      ValidationObserver
    } from 'vee-validate/dist/vee-validate.full';
    // import { ValidationObserver } from 'vee-validate';
    import BInputWithValidation from "../components/InputsValidation/BInputWithValidation";

    import { extend } from 'vee-validate';
    import { required, confirmed, min } from 'vee-validate/dist/rules';


    // Add the required rule
    extend('required', {
      ...required,
      message: 'This field is required'
    });
    // Add the required rule
    extend('confirmed', {
      ...confirmed,
      message: 'Not matching'
    });
    extend('min', {
      validate(value, { length }) {
        return value.length >= length;
      },
      params: ['length'],
      message: 'The field must have at least {length} characters'
    });

    export default {
      name: "Restore.vue",
      data () {
        return {
          password_reveal: '',
          refer: "password",
          test: '',
          info: {
            password: '',
            repeatPassword: '',
          }
        }
      },
      components: {
        ValidationObserver,
        BInputWithValidation,
      },
      computed: {
        tokenForPass () {
          return this.$store.getters['login/tokenForPass']
        },
        codeForPass () {
          return this.$store.getters['login/codeForPass']
        },
        passwordUpdated(){
          return this.$store.getters['login/passwordUpdated']
        }
      },
      mounted() {
        this.test = this.$route.query.test
        if (this.test){
          this.$store.dispatch('login/restorePassword', this.test);
        }
      },
      methods: {
        //not loggedin
        updatePassword(){
            this.$store.dispatch('login/updatePassword', this.info.password);
        },
        //if user is logged in
        setNewPassword(){
          if(this.info.password === this.info.repeatPassword && this.info.password !== ''){
            this.$store.dispatch('login/setNewPassword', this.info.password);
          }
        },
        redPage(){
          console.log('No data!')
          this.$router.push({path: '/admin/user'});
        },
      }
    }
</script>

<style>
/*div.sidebar{*/
/*  display: none;*/
/*}*/
  .pb-20{
    padding-bottom: 20px;
  }
</style>
