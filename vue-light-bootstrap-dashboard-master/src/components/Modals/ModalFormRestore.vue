<template>
  <section>
    <form action="">
      <div class="modal-card" style="width: auto">

        <header class="modal-card-head">
              <p class="modal-card-title">Restore Password</p>
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
            </section>

            <footer class="modal-card-foot">
              <b-button class="button" type="button" @click="closeRestore(); clearFields()">Close</b-button>
              <b-button @click="validate().then(restore)" class="button is-primary" type="is-info">Restore</b-button>
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

  export default {
    components: {
      ValidationObserver,
      BInputWithValidation,
    },
    props: {
      emailProp: {
        type: String
      },
      nameProp: {
        type: String
      },
    },
    data() {
      return {
        email: this.emailProp || '',
        name: this.nameProp|| '',
      }
    },
    methods:{
      clearFields(){
        this.name = ""
        this.email = ""
      },
      restore(){
        this.$emit('restore', {
          email: this.email,
          name: this.name,
        })
      },

      closeRestore(){
        this.$emit('closeRestore')
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
