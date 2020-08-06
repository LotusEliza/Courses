<template>
  <div>
  <card>
    <div slot="header">
      <h4 class="card-title bold">Add customer Form</h4>
      <br>
    </div>

    <form>
      userInfo::{{userInfo}}
      <div class="row">
        <div class="col-md-6">
          <base-input type="text"
                    label="Name"
                    placeholder="Name"
                    v-model="userInfo.Name">
          </base-input>
        </div>
        <div class="col-md-6">
          <base-input type="email"
                    label="Email"
                    placeholder="Email"
                    v-model="userInfo.Email">
          </base-input>
        </div>
      </div>

      <div class="row">
        <div class="col-md-6">
          <base-input type="text"
                      label="Tel"
                      placeholder="Tel"
                      v-model="userInfo.Tel">
          </base-input>
        </div>
      </div>

      <div class="text-center">
        <b-button type="submit" class="btn btn-info btn-fill float-right" @click.prevent="updateProfile">
          Update Profile
        </b-button>

      </div>
      <div class="clearfix"></div>
    </form>
  </card>


<!--------------------------------------------    PASSWORD CHANGE FORM------------------------------------------------>
  <card v-bind:style=" show ? 'background-color: #f2f2f2;' : 'background-color: white;'">
    <div slot="header">
      <h3  class="card-title bold">Вход</h3>
      <br>
    </div>

      <h4 class="card-title">Смена пароля</h4>
      Рекомендуем установить надежный пароль, который вы больше нигде не используете.
      <b-button class="btn btn-info btn-fill float-right" @click.prevent="show=!show" v-if="!show" >
        Change Password
      </b-button>
      <b-button class="btn btn-info btn-fill float-right" @click.prevent="show=!show" v-if="show">
        Close
      </b-button>

      <form v-if="show">

        <br>
        <hr>
        userPassword::{{userPassword}}
        <ValidationObserver ref="observer">
          <section slot-scope="{ validate }">

            <div class="row">
              <div class="col-md-4">

                <b-field label="Current Password">
                  <BInputWithValidation rules="required|min:5"
                                        placeholder="Your current password"
                                        v-model="userPassword.currentPass"
                                        type="password"
                                        :password_reveal="password_reveal"
                  />
                </b-field>

              </div>
              <div class="col-md-4">

                <b-field label="New Password">
                  <BInputWithValidation rules="required|min:5"
                                        placeholder="Your new password"
                                        v-model="userPassword.newPass"
                                        :password_reveal="password_reveal"
                                        :refer="refer"
                                        type="password"
                                        name="password"

                  />
                </b-field>

              </div>
              <div class="col-md-4">

                <b-field label="Confirm Password">
                  <BInputWithValidation rules="required|confirmed:password"
                                        placeholder="Your confirmation password"
                                        v-model="userPassword.confirmNewPass"
                                        :password_reveal="password_reveal"
                                        type="password"
                                        data-vv-as="password"
                  />
                </b-field>

              </div>
            </div>
            <div class="row">
              <b-button type="is-text" @click="remindPassword()">Forgot password</b-button>
            </div>
            <div class="row">
              <b-button class="btn btn-info btn-fill float-right" @click="validate().then(savePassword)" :disabled='!isComplete'>
<!--              <b-button class="btn btn-info btn-fill float-right" @click.prevent="validate().then(savePassword)" disabled>-->
                Save changes
              </b-button>
            </div>
          </section>
        </ValidationObserver>
      </form>
  </card>

  </div>
<!--  <div v-else>{{redirectPage()}}</div>-->
</template>
<script>
  import Card from 'src/components/Cards/Card.vue'
  import {mapGetters} from "vuex";
  import BInputWithValidation from "../../../components/InputsValidation/BInputWithValidation";

  import { extend } from 'vee-validate';
  import { required, confirmed } from 'vee-validate/dist/rules';
  import {ValidationObserver} from "vee-validate/dist/vee-validate.full";
  import cookie from "vue-cookies";
  import clone from 'lodash/clone'

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

  export default {
    components: {
      Card,
      ValidationObserver,
      BInputWithValidation,
    },
    data () {
      return {
        userInfo: {},
        userPassword: {
          currentPass: '',
          newPass: '',
          confirmNewPass: '',
        },
        show: false,
        password_reveal: '',
        refer: "password",
        isComponentModalRestoreActive: false,
      }
    },
    mounted() {
      if(Object.keys(this.$store.getters['login/user']).length === 0){
        console.log("Empty obj")
        this.$store.dispatch('login/info').
        then(() => {
          const user =  this.$store.getters['login/user'];
          this.userInfo = clone(user)
        });
        // then(() => this.userInfo = this._.cloneDeep(this.$store.getters['login/user']));

      }else {
        console.log("Not Empty obj")
        // this.userInfo = this._.cloneDeep(this.$store.getters['login/user'])
        const user =  this.$store.getters['login/user'];
        this.userInfo = clone(user)
      }
    },
    computed: {
      isComplete () {
        return this.userPassword.confirmNewPass && this.userPassword.currentPass && this.userPassword.newPass;
      },
      ...mapGetters("login", [
        'user', 'updatePassResult'
      ]),
    },
    //TODO clean the form if changed successfuly
    watch: {
      updatePassResult (newVal, oldVal) {
        if (newVal == 'error'){
          console.log("Password current  is wrong!!!")
          // this.$router.push({ name: 'LogIn' })
        }else if(newVal == 'success'){
          console.log("Password current  is right!!!")
          this.show = false
          this.cleanPassForm()
        }
      }
    },
    methods: {
      updateProfile() {
        alert('Your data: ' + JSON.stringify(this.user))
        this.$store.dispatch('login/updateInfo', this.userInfo).then(()=>{
          const user = this.$store.getters['login/user'];
          this.userInfo = clone(user);
        })
        // store.commit('changeProgress', progress)
      },
      savePassword(){
        // if (this.validate())
        this.$store.dispatch('login/savePassword', this.userPassword)
        console.log("Dispatch save new password")
      },
      cleanPassForm(){
        this.userPassword.newPass = ""
        this.userPassword.currentPass = ""
        this.userPassword.confirmNewPass = ""
      },
      remindPassword(){
          this.$router.push({ name: 'RecovPassLoggedIn' })
      },
      // redirectPage(){
      //   console.log('No cookie!')
      //   this.$router.push({ name: 'LogIn' })
      // }
    }
  }

</script>
<style>
.bold{
  font-weight: bold !important;
}
  .grey{
    background-color: lightgrey;
  }
  /*.card-body{*/
  /*  background-color: pink;*/
  /*}*/
</style>
