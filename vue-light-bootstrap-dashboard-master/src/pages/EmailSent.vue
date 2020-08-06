<template>
  <div class="content">
    <div class="container-fluid">
      <card>
          <div class="row justify-content-center ">
            <div class="col-md-3 col-md-offset-1">
            </div>
            <div class="col-md-3 col-md-offset-1 pb-20" >
              <div class="typo-line  has-text-centered">
                <h2>A verification link has been sent to your email account<br>
                  <small>Please click on the link that has been sent to your email account to verify your email
                    and continue the restoring process.</small>
                </h2>
                <b-button type="is-info" class="is-large"
                          @click="$router.go(-1)">
                  Back to Login page
                </b-button >
              </div>
            </div>
            <div class="col-md-3 col-md-offset-1">
            </div>
        </div>
      </card>
    </div>
  </div>
</template>

<script>
    import {mapGetters} from "vuex";

    export default {
      name: "EmailSent.vue",
      data () {
        return {
        }
      },

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
