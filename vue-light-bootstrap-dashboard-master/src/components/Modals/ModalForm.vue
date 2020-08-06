<template>
  <section>
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
              <p class="modal-card-title">Login</p>
        </header>

        <ValidationObserver ref="observer">
          <section slot-scope="{ validate }">
            <section class="modal-card-body">
              <b-field label="Email">
              <BInputWithValidation rules="required|email"
                                    placeholder="Your email"
                                    v-model="email"
              />

              </b-field>

              <b-field label="Name">
                <BInputWithValidation rules="required"
                                      placeholder="Your name"
                                      v-model="name"
                />

              </b-field>

              <b-field label="Password">
                <BInputWithValidation rules="required|min:5"
                                      placeholder="Your password"
                                      v-model="password"
                                      :password_reveal="password_reveal"
                                      type="password"
                />
              </b-field>

            </section>

            <footer class="modal-card-foot">
              <b-button class="button" type="button" @click="$parent.close(); clearFields()">Close</b-button>
              <b-button @click="validate().then(login)" class="button is-primary" type="is-info">Login</b-button>
              <b-button type="is-text" @click="remind()">Forgot password</b-button>
            </footer>

          </section>
        </ValidationObserver>

      </div>
    </form>
  </section>


</template>

<script>
  import {
    ValidationObserver
  } from 'vee-validate/dist/vee-validate.full';
  // import { ValidationObserver } from 'vee-validate';
  import BInputWithValidation from "../InputsValidation/BInputWithValidation";

  import { extend } from 'vee-validate';
  import { required, email } from 'vee-validate/dist/rules';

  // Add the required rule
  extend('required', {
    ...required,
    message: 'This field is required'
  });

  extend('min', {
    validate(value, { length }) {
      return value.length >= length;
    },
    params: ['length'],
    message: 'The field must have at least {length} characters'
  });

  // Add the email rule
  extend('email', {
    ...email,
    message: 'This field must be a valid email'
  });

// TODO make validation for all fields
  export default {
    components: {
      ValidationObserver,
      BInputWithValidation,
    },
    props: {
    //   email: {
    //     type: String,
    //   },
    //   password: {
    //     type: String,
    //   },
    },
    data() {
      return {
        email: '',
        password: '',
        name: '',
        password_reveal: '',
        // show: false,
        // emailRestore: '',
        // nameRestore: '',
      }
    },

    methods:{
      login(){
        this.$emit('login', {
          email: this.email,
          password: this.password,
          name: this.name,
        })
      },
      clearFields(){
        this.name = ""
        this.password = ""
        this.email = ""
      },
      restore(){
        this.$emit('restore', {
          email: this.email,
          name: this.name,
        })
      },

      remind(){
        this.$emit('remind')
      }
      // switchTaps(){
        // this.show = !this.show
        // this.name = ""
        // this.email = ""
        // this.password = ""
      // }
    }
  }
</script>
