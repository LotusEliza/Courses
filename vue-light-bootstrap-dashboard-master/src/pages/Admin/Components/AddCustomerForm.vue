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
    <b-button class="btn btn-info btn-fill float-right" @click.prevent="cleanFields" v-if="show">
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
                  <BInputWithValidation rules="required|email"
                                        placeholder="Email"
                                        v-model="user.email"
                  />
                </b-field>
              </div>

              <div class="col-md-4">
                <vue-tel-input v-model="user.tel"
                               v-bind="bindProps"
                               @validate="yourValidationMethod"
                               :class="dynamicClass"

                ></vue-tel-input>
                <p :class="help, is_danger" v-if="dynamicClass === 'is-danger'">
                  {{ errorStr }}
                </p>

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

            <div class="row buttons pt-3">
              <b-button class="btn btn-info btn-fill float-right" @click="validate().then(addCustomer)" :disabled='!isComplete'>
                <!--              <b-button class="btn btn-info btn-fill float-right" @click.prevent="validate().then(savePassword)" disabled>-->
                Add Customer
              </b-button>
<!--              <b-button class="btn btn-info btn-fill float-right" @click.prevent="cleanFields" v-if="show">-->
<!--                Cancel-->
<!--              </b-button>-->
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
import {required, confirmed, email} from 'vee-validate/dist/rules';
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
// Add the email rule
extend('email', {
  ...email,
  message: 'This field must be a valid email'
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
        show: false,
        user: {
          name: '',
          tel: '',
          email: '',
          password: '',
          confirmPass: '',
        },
        bindProps: {
          // onlyCountries: ['UK', 'USA'],
          validCharactersOnly: true,
          mode: "international",
          defaultCountry: "UA",
          disabledFetchingCountry: true,
          disabled: false,
          disabledFormatting: false,
          placeholder: "Enter a phone number",
          required: true,
          enabledCountryCode: true,
          enabledFlags: true,
          // preferredCountries: ["AU", "BR"],
          onlyCountries: ["UA", "US", "TR"],
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
      // isValid: false,
      dynamicClass: '',
      is_danger: 'is-danger',
      help: 'help',
      errorStr: ""
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

      if (isValid === true){
        this.dynamicClass = 'is-success'
      }else if(isValid === false && number.input !== ""){
        this.errorStr = "This field is required"
        this.dynamicClass = 'is-danger'
      }
      // else if (isValid === false && number.input !== "" && !this.phoneValidator()){
      //   this.errorStr = "This field must contain only numbers"
      //   this.dynamicClass = 'is-danger'
      // }
      // console.log(country)
      return false
      // Do stuff with the arguments passed by the vue-tel-input component
    },
    async addCustomer(event){
      this.user.tel = this.user.tel.replace(/[\s\/]/g, '')
      await this.$store.dispatch('customers/saveCustomer', this.user).then(() => this.cleanFields());
      // this.$store.dispatch('customers/saveCustomer', this.user);
      // this.phoneValidator()
      // if(this.user.tel)
        // this.$emit('add')
    },
    cleanFields(){
      this.show=!this.show
      this.user.name = ''
      this.user.tel = ''
      this.user.email = ''
      this.user.password = ''
      this.user.confirmPass = ''
    }
    // phoneValidator () {
    //   let phoneno = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
    //   this.user.tel = this.user.tel.replace(/[\s\/]/g, '')
    //   if (this.user.tel.match(phoneno)) {
    //     console.log('Phone is true')
    //     return true
    //   }
    //   else {
    //     console.log('Phone is false')
    //     return false
    //   }
    // }
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

.is-danger {
  border-color: #ff3860;
}

.is-success {
  border-color: #23d160;
}
.vue-tel-input[data-v-31b4a26e]:focus-within {
  border-color: #7957d5 !important;
  box-shadow: 0 0 0 0.125em rgba(255, 56, 96, 0.25) !important;
}
.help {
  display: block;
  font-size: 0.75rem;
  margin-top: 0.25rem;
  color: #ff3860;
}
.btn {
  padding: 5px 10px 5px 10px !important;
}

section div.buttons{
  padding-left: 15px;
}


</style>
