<template>

    <div>

        <div class="row">
            <div class="col-4" v-for="task in tasks">
                <task-component :task="task"></task-component>
            </div>
        </div>

        <b-modal ref="modal-edit-task" id="modal-edit-task" @close="closeModal()">
            <div slot="modal-header">
                <h2 class="text-center">
                    Task Manager
                </h2>
            </div>
            <task-form-component @updateTask="updatedTask=$event"></task-form-component>
            <div slot="modal-footer" class="text-center form-actions">
                <button class="btn btn-danger" @click="closeModal()">
                    Cancel
                </button>
                <button class="btn btn-success pull-right mr-2" @click="editTask()">
                    Edit task
                </button>
            </div>
        </b-modal>

    </div>


</template>

<script>
    import TaskComponent from "./task-component.vue"
    import TaskFormComponent from "./task-form-component.vue"

    export default {
        name: "manager-component",
        components: {
            TaskComponent,
            TaskFormComponent
        },
        data() {
            return {
                updatedTask: {},
                tasks: [
                    {title: "do1", goal: "to do homework", dueDate: "01.02.2020"},
                    {title: "do2", goal: "work", dueDate: "03.12.2020"}
                ],
            };
        },
        methods: {
            editTask() {

            },
            closeModal() {
                this.$refs['modal-edit-task'].hide();
                this.clean();
            },
            clean() {
                this.updatedTask = {}
            }
        },


    }
</script>

<style scoped>

</style>