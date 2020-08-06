<template>
  <ValidationProvider :vid="vid" :placeholder="placeholder" :name="$attrs.name || $attrs.label" :rules="rules" :password_reveal="password_reveal" :type="type" :refer="refer">
      <b-field
        slot-scope="{ errors, valid }"
        v-bind="$attrs"
        :type="{ 'is-danger': errors[0], 'is-success': valid }"
        :message="errors"
      >
        <b-input v-model="innerValue"
                 v-bind="$attrs"
                 :placeholder="placeholder"
                 :password-reveal="password_reveal"
                 :type="type"
                 :ref="refer"
        ></b-input>
      </b-field>
  </ValidationProvider>
</template>

<script>
import { ValidationProvider } from "vee-validate";

export default {
  components: {
    ValidationProvider
  },
  props: {
    refer: {
      type: String
    },
    type: {
      type: String
    },
    password_reveal: {
      type: Boolean
    },
    placeholder: {
      type: String
    },
    vid: {
      type: String
    },
    rules: {
      type: [Object, String],
      default: ''
    },
    // must be included in props
    value: {
      type: null
    }
  },
  data: () => ({
    innerValue: ''
  }),
  watch: {
    // Handles internal model changes.
    innerValue (newVal) {
      this.$emit('input', newVal);
    },
    // Handles external model changes.
    value (newVal) {
      this.innerValue = newVal;
    }
  },
  created () {
    if (this.value) {
      this.innerValue = this.value;
    }
  }
};
</script>
