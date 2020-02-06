<template>
    <div class="row">

        <div class="col border">
            <form @submit.prevent>
                <div class="form-group">
                    <label for="login">Login</label>
                    <input type="text" class="form-control" id="login" v-model="signInData.login" placeholder="Login">
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" class="form-control" id="password" v-model="signInData.password"
                           placeholder="Password">
                </div>
                <button class="btn btn-block btn-primary" @click="signIn()">Login</button>
            </form>
        </div>

        <div class="col border ml-2">
            <form @submit.prevent>
                <div class="form-group">
                    <label for="signUp">SignUp</label>
                    <input type="text" class="form-control" id="signUp" v-model="signUpData.login" placeholder="Login">
                </div>
                <div class="form-group">
                    <label for="passwordUp">Password</label>
                    <input type="password" class="form-control" id="passwordUp" v-model="signUpData.password"
                           placeholder="Password">
                </div>
                <button class="btn btn-block btn-info" @click="signUp()">SignUp</button>
            </form>
        </div>

    </div>
</template>

<script>
    import RegistrationService from '../services/Login.js'
    import { mapMutations } from 'vuex';

    export default {
        name: "login-component",
        data() {
            return {
                signInData: {
                    login: '',
                    password: ''
                },
                signUpData: {
                    login: '',
                    password: ''
                }
            }
        },
        methods: {
            ...mapMutations([
                'loginUser'
            ]),
            async signIn() {
                const res = await RegistrationService.signIn(this.signInData);
                if (res['message'] === 'ok') {
                    this.loginUser(this.signInData.login);
                    this.$toasted.success('Success')
                } else {
                    this.$toasted.error('Wrong login or password')
                }
            },
            async signUp() {
                try {
                    const res = await RegistrationService.signUp(this.signUpData);
                    console.log('mes',res['message']);
                    if (res['message'] === 'ok') {
                        this.$toasted.success('Success. You can login now');
                    } else {
                        this.$toasted.error('User already exists')
                    }
                } catch (e) {
                    console.log("error", e)
                }
            }
        }
    }
</script>

<style scoped>

</style>