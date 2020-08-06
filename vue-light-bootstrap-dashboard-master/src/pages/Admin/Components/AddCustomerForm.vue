<template>
  <!--------------------------------------------    PASSWORD CHANGE FORM------------------------------------------------>


<!--    <h4 class="card-title">Смена пароля</h4>-->
<!--    Рекомендуем установить надежный пароль, который вы больше нигде не используете.-->
<!--    <b-button class="btn btn-info btn-fill float-right" @click.prevent="show=!show" v-if="!show" >-->
<!--      Change Password-->
<!--    </b-button>-->


  <card v-bind:style=" show ? 'background-color: #f2f2f2;' : 'background-color: white;'">
    <div slot="header">
      <h3  class="card-title bold">Add Customer</h3>
      <br>
    </div>

<!--    <h4 class="card-title">Please add a customer</h4>-->
    Рекомендуем установить надежный пароль.
    <b-button class="btn btn-info btn-fill float-right" @click.prevent="show=!show" v-if="!show" >
      Add Customer
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
                <b-field label="Customer Name">
                  <BInputWithValidation rules="required|min:5"
                                        placeholder="Customer Name"
                                        v-model="userPassword.currentPass"
                  />
                </b-field>
              </div>

              <div class="col-md-4">
                <b-field label="Customer Name">
                  <BInputWithValidation rules="required|min:5"
                                        placeholder="Customer Name"
                                        v-model="userPassword.currentPass"
                  />
                </b-field>
              </div>

              <div class="col-md-4">
                <b-field label="Customer Name">
                  <BInputWithValidation rules="required|min:5"
                                        placeholder="Customer Name"
                                        v-model="userPassword.currentPass"
                  />
                </b-field>
              </div>
            </div>

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
  <!--          <div class="row">-->
  <!--            <b-button type="is-text" @click="remindPassword()">Forgot password</b-button>-->
  <!--          </div>-->

            <div class="row">
              <b-button class="btn btn-info btn-fill float-right" @click="validate().then(savePassword)" :disabled='!isComplete'>
                <!--              <b-button class="btn btn-info btn-fill float-right" @click.prevent="validate().then(savePassword)" disabled>-->
                Add Customer
              </b-button>
              <b-button class="btn btn-info btn-fill float-right" @click.prevent="show=!show" v-if="show">
                Cancel
              </b-button>
            </div>

          </section>
        </ValidationObserver>

    </form>
  </card>

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
  name: "AddCustomerForm",
  components: {
    Card,
    ValidationObserver,
    BInputWithValidation,
  },
  data() {
      return{
        show: true,
        userPassword: {
          currentPass: '',
          newPass: '',
          confirmNewPass: '',
        },
      }
  }
}
</script>

<style scoped>

</style>
