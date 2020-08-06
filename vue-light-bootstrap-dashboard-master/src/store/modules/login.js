import cookie from 'vue-cookies'
import {router} from "../../main.js"
import Vue from 'vue'
import {ToastProgrammatic as Toast} from "buefy";



export default {
  namespaced: true,
  state: {
    status: '',
    token: cookie.get('token') || '',
    admin: localStorage.getItem('UserID') || '',
    user: {},
    tokenForPass: "",
    codeForPass: "",
    updPass: "",
    sessionOver: "",
    enterCodeActive: false,
  },
  getters: {
    isAuthenticated: state => state.token,
    isAdmin: state => state.admin,
    authStatus: state => state.status,
    user: state => state.user,
    tokenForPass: state => state.tokenForPass,
    updatePassResult: state => state.updPass,
    sessionOver: state => state.sessionOver,
    enterCodeActive: state => state.enterCodeActive,
    codeForPass: state => state.codeForPass,
    // authStatus: (state, getters, commit) => {
    //   if (state.status && state.token) {
    //     commit('AUTH_SUCCESS')
    //     return state.status
    //   } else {
    //     return state.status
    //   }
    // },
  },
  mutations: {
    AUTH_REQUEST (state) {
      state.status = 'loading'
    },
    AUTH_SUCCESS (state){
      state.status = 'success'
      state.token = cookie.get('token')
    },
    AUTH_ERROR (state) {
      state.status = 'error'
    },
    SET_INFO(state, user) {
      console.log(user)
      state.user = user
    },
    AUTH_LOGOUT(state){
      state.token = ''
      state.status = ''
    },
    SET_PASS_TOKEN(state, token){
      state.tokenForPass = token
      console.log("state.showForm")
      console.log(state.showForm)
    },
    REMOVE_PASS_TOKEN(state){
      state.tokenForPass = ""
    },
    UPDATE_PASS_RESULT(state, result){
      state.updPass = result
    },
    SET_SESSION_OVER(state){
      state.sessionOver = "yes"
    },
    ENTER_CODE_ACTIVE(state){
      state.enterCodeActive = true
    },
    REMOVE_CODE_ACTIVE(state){
      state.enterCodeActive = false
    },
    SET_PASS_CODE(state, code){
      state.codeForPass = code
      // console.log("state.showForm")
      // console.log(state.showForm)
    },
    REMOVE_PASS_CODE(state){
      state.codeForPass = ""
      state.enterCodeActive = false
    },
  },
  actions: {

    /**
     * Sends info (password, email, name), method - POST
     * @return | gets token http_only OR error
     */
    async loginSend({commit, state}, data) {
      commit('AUTH_REQUEST')
      let response = await window.axios.post('/user_exist',
        {
        Password: data.password,
        Email: data.email,
        Name: data.name,
      },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {
        if (response.data.Error === "ERROR_NO_SUCH_USER"){
          // localStorage.removeItem('UserID');
          cookie.remove('token')
          console.log("No such user in db vue js!!!")
          Toast.open({
            message: "Wrong email or name!",
            position: "is-top-right",
            type:"is-danger",
          });
          commit('AUTH_ERROR')
        } else if (response.data.Error === "ERROR_WRONG_PASSWORD"){
          cookie.remove('token')
          console.log("Wrong Password!!!")
          Toast.open({
            message: "Wrong Password!",
            position: "is-top-right",
            type:"is-danger",
          });
          commit('AUTH_ERROR')
        } else {
          let myResult = cookie.get('token').split(".");
          let decodedString = JSON.parse(atob(myResult[1]));
          localStorage.setItem('UserID', decodedString.userid)
          localStorage.setItem('Admin', decodedString.admin)
          commit('AUTH_SUCCESS')
          await router.push("/");
        }
      }
    },

    /**
     * Sends info from RestoreForm (email, name) and checks if such user exist,
     * Sends email with restore link that contains restore_password_token, method - POST
     * @return | success OR error
     */
    async restoreStart({commit, state, dispatch}, data) {
      // commit('SET_DATA', data);
      let response = await window.axios.post('/restore_start',
        {
          Email: data.email,
          Name: data.name,
        },
        {withCredentials: true});
      if (response.status === 200 || response.status === 204) {
        if (response.data.Error === "ERROR_NO_SUCH_USER"){
          commit('AUTH_ERROR', true);
        }else {
        console.log("router.history.current.path")
        console.log(router.history.current.path)
          await router.push("/email_sent");
          console.log("Email Sent!")
        }
      }
    },


    // async restoreStart({commit, state, dispatch}, data) {
    //   commit('SET_DATA', data);
    //   let response = await window.axios.post('/restore_start',
    //     {
    //       Email: data.email,
    //       Name: data.name,
    //     },
    //     {withCredentials: true});
    //   if (response.status === 200 || response.status === 204) {
    //     if (response.data.Error === "ERROR_NO_SUCH_USER"){
    //       commit('AUTH_ERROR', true);
    //     }else {
    //       console.log("router.history.current.path")
    //       console.log(router.history.current.path)
    //       if (router.history.current.path === '/admin/user') {
    //         console.log("From admin/user!")
    //       }else {
    //         await router.push("/email_sent");
    //       }
    //       console.log("Email Sent!")
    //     }
    //   }
    // },


    /**
     * Refreshes token if his life is less then 30 sec, method - POST
     * @return | success OR error
     */
    async refresh({commit}) {
      console.log("refresh it is!")
      let response = await window.axios.post('/refresh',
        {},
        { withCredentials: true });
      if (response.status == 200 || response.status == 204) {
        if (response.data.Error === "ERROR_USER_NOT_CONFIRMED") {
          commit('AUTH_LOGOUT');
          router.push("/log-in").then(() => console.log("Logged out!!!"))
        }
        // console.log("response")
      }
      // else if (response.status ==425){
      // }
    },

    /**
     * Logout user from admin panel
     */
    logOutClick({commit}){
      console.log("Log Out click")
      cookie.remove('token')
      localStorage.removeItem('UserID');
      commit('AUTH_LOGOUT');
      Toast.open({
        message: "Logged out!",
        position: "is-top-right",
        type:"is-warning",
      });


      router.push("/log-in").then(() => console.log("Logged out!!!"))
      //cleaning the vuex state after logout
      location.reload()

    },

    /**
     * Sends userId with http_only token, method - POST
     * @return | email, name, tel etc OR error
     */
    async info({commit}){
      let response = await window.axios.post('/info',
        {
        },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {
        if (response.data.Error === "ERROR_TOKEN_NOT_FOUND_OR_EXPIRED") {
          commit('AUTH_LOGOUT');
          router.push("/log-in").then(() => console.log("Logged out!!!"))
        }else {
          console.log(response.data)
          commit('SET_INFO', response.data);
        }
      }
    },

    /**
     * Updates user info, method - POST
     * @return |  OR error
     */
    async updateInfo({commit}, data){
      let response = await window.axios.post('/info_update',
        {
          ID: data.ID,
          Email: data.Email,
          Name: data.Name,
          Tel: data.Tel,
        },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {
        console.log("updated info of user")
        commit('SET_INFO', data);
      }
    },

    async recovPassLoggedIn({commit, state}, data){
      console.log("userID")
      console.log(state.user.ID)
      let response = await window.axios.post('/recover_password',
        {
          ID: state.user.ID,
          Method: data,
          Tel: state.user.Tel,
          Email: state.user.Email,
        },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {
        // await router.push("/");
        // console.log("response.data.Error")
        // console.log(response.data.Error)
        // if(response.data.Error === "ERROR_NO_ACCESS_TOKEN"){
        //   console.log("Session is over...")
        //
        //   setTimeout(()=>{
        //     router.push("/");
        //   },1000);
        //
        // }
        // console.log(response.data)
        // console.log("updated info of user")
        commit('ENTER_CODE_ACTIVE', data);


      }
      // if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
      //   // if you ever get an unauthorized, logout the user
      //   this.$store.dispatch(AUTH_LOGOUT)
      //   // you can also redirect to /login if needed !
      // }
    },

    async sendVerCode({commit, state}, code){
      console.log("userID")
      console.log(state.user.ID)
      let response = await window.axios.post('/check_ver_code',
        {
          IdCustomer: state.user.ID,
          Code: Number(code),
        },
        {withCredentials: true});
      if (response.status == 200 || response.status == 204) {

        if (response.data.Error === "ERROR_EXPIRED_TOKEN") {
          commit('REMOVE_CODE_ACTIVE');
          // commit('REMOVE_PASS_CODE');
          Toast.open({
            message: "Too late... Code is expired! Try again!",
            position: "is-top-right",
            type:"is-danger",
          })
          await router.push("/admin/user");
        }else if (response.data.Error === "ERROR_NOT_FOUND"){
          Toast.open({
            message: "Wrong code!",
            position: "is-top-right",
            type:"is-danger",
          })
          // await router.push("/admin/user");
          // commit('SET_PASS_CODE', code);
          // commit('REMOVE_CODE_ACTIVE', code);
          // await router.push("/restore");
        } else {
          console.log("Code is fine")
          commit('SET_PASS_CODE', code);
          commit('REMOVE_CODE_ACTIVE', code);
          await router.push("/restore");
          ///////////////////
        }


      }
      // if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
      //   // if you ever get an unauthorized, logout the user
      //   this.$store.dispatch(AUTH_LOGOUT)
      //   // you can also redirect to /login if needed !
      // }
    },

    /**
     * Checks if restore_password_token exists in database and if yes, - shows the form (new_password), method - POST
     * @return | success OR error
     */
    async restorePassword({commit}, token){
      let response = await window.axios.post('/confirmation',
        {
          Token: token,
        });
      if (response.status == 200 || response.status == 204) {
        console.log("Token change_password exist!")
        commit('SET_PASS_TOKEN', token);
      }
    },

    /**
     * Sends info from RestorePassword form (restore_password_token, new_password), method - POST
     * if everything ok server sets http_only jwt
     * @return | gets token http_only OR error
     */
    async updatePassword({commit, state, dispatch}, password){
      let response = await window.axios.post('/update_password',
        {
          Password: password,
          Token: state.tokenForPass,
        },
        { withCredentials: true });
      if (response.status == 200 || response.status == 204) {
        let myResult = cookie.get('token').split(".");
        let decodedString = JSON.parse(atob(myResult[1]));
        localStorage.setItem('UserID', decodedString.userid)
        commit('REMOVE_PASS_TOKEN');
        commit('AUTH_SUCCESS');
        await router.push("/");
      }
    },
    async setNewPassword({commit, state, dispatch}, password){
      let response = await window.axios.post('/set_new_password',
        {
          Password: password,
          ID: state.user.ID,
        },
        { withCredentials: true });
      if (response.status == 200 || response.status == 204) {
        console.log("Updated pass when user logged in!")
        await router.push("/admin/user")
        commit('REMOVE_PASS_CODE');
        Toast.open({
          message: "Password Updated",
          position: "is-top-right",
          type:"is-success",
        })
      }
    },
    async savePassword({commit, state, dispatch}, data){
      let response = await window.axios.post('/save_password',
        {
          ID: Number(localStorage.getItem('UserID')),
          CurrentPassword: data.currentPass,
          NewPassword: data.newPass,
        },
        { withCredentials: true });
      if (response.status == 200 || response.status == 204) {
        // console.log("Password Changed!")
        // console.log(response.data.Error)
        if(response.data.Error === "ERROR_NO_SUCH_USER"){
          Toast.open({
            message: "Wrong current password!",
            position: "is-top-right",
            type:"is-danger",
          });
          console.log("Password wrong!!!")
          // UPDATE_PASS_RESULT
          commit('UPDATE_PASS_RESULT', "error");
        }else if (response.data.Error === "ERROR_NO_ACCESS_TOKEN"){
            console.log("Session is over...")
          commit('SET_SESSION_OVER');
            // setTimeout(()=>{
            //   router.push("/");
            // },1000);
        }
        else {
          Toast.open({
            message: "Password Updated",
            position: "is-top-right",
            type:"is-success",
          });
          console.log("Password updated!!!")
          commit('UPDATE_PASS_RESULT', "success");
        }
      }
    },
  }
}









// async welcome({commit}) {
//   let response = await window.axios.post('/welcome',
//     {},
//     { withCredentials: true });
//   if (response.status == 200 || response.status == 204) {
//     console.log(response.data)
//     commit('VALID');
//   }
// },


// async loginSendAfterRestore({commit, state}, obj) {
//   let response = await window.axios.post('/user_exist',
//     {
//       Token: obj.token,
//       Password: obj.password,
//     },
//     {withCredentials: true});
//   if (response.status == 200 || response.status == 204) {
//     let myResult = cookie.get('token').split(".");
//     let decodedString = JSON.parse(atob(myResult[1]));
//     localStorage.setItem('UserID', decodedString.userid)
//     commit('SET_TOKEN');
//     await router.push("/");
//   }
// },

