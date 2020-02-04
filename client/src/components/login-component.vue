<template>
    <div class="row">

        <div class="col border">
            <form>
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
            <form>
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
                if (res === 'ok') {
                    this.loginUser(this.signInData.login)
                } else {
                    this.$toasted.error('Wrong login or password')
                }
            },
            async signUp() {
                const res = await RegistrationService.signUp(this.signUpData);
            }
        }
    }
</script>

<style scoped>

</style>