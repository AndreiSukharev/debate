<template>

    <div>
        <nav class="navbar navbar-expand-lg navbar-light" style="background-color: #e3f2fd;">
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item">
                        <a class="navbar-brand" href="#">Task Manager</a>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" to="/" exact @click.native="logout">Log out</router-link>
                    </li>
                </ul>
                <button class="btn btn-outline-success my-2 my-sm-0" @click="openModal">Add task</button>
            </div>
        </nav>

        <b-modal ref="modal-add-task" id="modal-add-task" @close="closeModal()">
            <div slot="modal-header">
                <h2 class="text-center">
                    Task Manager
                </h2>
            </div>
            <task-form-component @updateTask="task = $event"></task-form-component>
            <div slot="modal-footer" class="text-center form-actions">
                <button class="btn btn-danger" @click="closeModal">
                    Cancel
                </button>
                <button class="btn btn-success pull-right mr-2" @click="addTask()">
                    Add task
                </button>
            </div>

        </b-modal>

    </div>


</template>

<script>
    import TaskFormComponent from "./task-form-component.vue"
    import TaskService from '../services/Task.js'
    import { mapMutations } from 'vuex';

    export default {
        name: "header-component",
        components: {
            TaskFormComponent
        },

        data() {
            return {
                task: {
                    title: '',
                    goal: '',
                    dueDate: ''
                }
            }
        },
        methods: {
            ...mapMutations([
                'logoutUser'
            ]),
            async addTask() {
                await TaskService.addTask(this.task);
                this.$toasted.success("Task is added");
                this.$root.$emit('updateTasks')
            },
            openModal() {
                this.$refs['modal-add-task'].show();
            },
            closeModal() {
                this.clean();
                this.$refs['modal-add-task'].hide();
            },
            clean() {
                this.task = {
                    title: '',
                    goal: '',
                    dueDate: ''
                }
            },
             logout() {
                this.logoutUser();
            }
        }
    }
</script>

<style scoped>

</style>