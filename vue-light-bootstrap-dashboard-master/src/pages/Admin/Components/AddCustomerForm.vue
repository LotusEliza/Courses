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
        user::{{user}}
        <ValidationObserver ref="observer">
          <section slot-scope="{ validate }">

            <div class="row">
              <div class="col-md-4">
                <b-field label="Name">
                  <BInputWithValidation rules="required"
                                        placeholder="Name"
                                        v-model="user.name"
                  />
                </b-field>
              </div>



              <div class="col-md-4">
                <b-field label="Email">
                  <BInputWithValidation rules="required"
                                        placeholder="Email"
                                        v-model="user.email"
                  />
                </b-field>
              </div>

              <div class="col-md-4">
                <vue-tel-input v-model="user.tel"
                               v-bind="bindProps"
                               @validate="yourValidationMethod"

                ></vue-tel-input>
                <!--                <b-field label="Tel">-->
                <!--                  <BInputWithValidation rules="required|min:13"-->
                <!--                                        placeholder=""-->
                <!--                                        v-model="user.tel"-->
                <!--                  />-->
                <!--                </b-field>-->
              </div>
            </div>

            <div class="row">
<!--              <div class="col-md-4">-->

<!--                <b-field label="Current Password">-->
<!--                  <BInputWithValidation rules="required|min:5"-->
<!--                                        placeholder="Your current password"-->
<!--                                        v-model="userPassword.currentPass"-->
<!--                                        type="password"-->
<!--                                        :password_reveal="password_reveal"-->
<!--                  />-->
<!--                </b-field>-->

<!--              </div>-->
              <div class="col-md-4">

                <b-field label="New Password">
                  <BInputWithValidation rules="required|min:5"
                                        placeholder="Your new password"
                                        v-model="user.password"
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
                                        v-model="user.confirmPass"
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
import { VueTelInput } from 'vue-tel-input'

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
    VueTelInput,
  },
  data () {
    return {
        show: true,
        user: {
          name: '',
          tel: '',
          email: '',
          password: '',
          confirmPass: '',
        },
        bindProps: {
          mode: "international",
          defaultCountry: "FR",
          disabledFetchingCountry: false,
          disabled: false,
          disabledFormatting: false,
          placeholder: "Enter a phone number",
          required: true,
          enabledCountryCode: false,
          enabledFlags: true,
          preferredCountries: ["AU", "BR"],
          onlyCountries: [],
          ignoredCountries: [],
          autocomplete: "off",
          name: "telephone",
          maxLen: 11,
          wrapperClasses: "",
          inputClasses: "",
          // dropdownOptions: {
          //   disabledDialCode: false
          // },
          // inputOptions: {
          //   showDialCode: false
          // }
          // customValidate: true,
        },
      password_reveal: true,
      refer: "password",
      // isComponentModalRestoreActive: false,
      }
  },
  computed: {
    isComplete() {
      return this.user.password && this.user.confirmPass;
    },
  },
  methods: {
    yourValidationMethod: function ({ number, isValid, country }) {
      console.log("yourValidationMethod")
      console.log(number)
      console.log(isValid)
      // console.log(country)
      return false
      // Do stuff with the arguments passed by the vue-tel-input component
    },
  }
}
</script>

<style scoped>

.vue-tel-input {
  margin-top: 24px !important;
  height: 36px !important;
}

.vue-tel-input:focus-within {
  box-shadow: none !important;
  border-color: #66afe9;
  border-top-color: rgb(102, 175, 233);
  border-right-color: rgb(102, 175, 233);
  border-bottom-color: rgb(102, 175, 233);
  border-left-color: rgb(102, 175, 233);
}
</style>
