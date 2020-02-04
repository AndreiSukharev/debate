<template>
    <div id="app">
        <header>
            <header-component/>
        </header>
        <main>
            <div class="container mt-3">
                <router-view v-if="loggedUser.length > 0"/>
                <login-component v-else/>
            </div>
        </main>
        <footer>
        </footer>
    </div>
</template>

<script>
    import HeaderComponent from "./components/header-component.vue"
    import LoginComponent from  "./components/login-component.vue"
    import {mapState} from 'vuex';

    export default {
        name: 'App',
        components: {
            HeaderComponent,
            LoginComponent
        },
        data() {
            return {
                checkLogged: false,
            };
        },
        computed: {
            ...mapState([
                'loggedUser',
            ]),
        },
    };
</script>


<style>
</style>