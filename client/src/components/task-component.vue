<template>

    <div>
        <div class="card text-center" style="width: 18rem;">
            <div class="card-body">
                <h5 class="card-title">{{task.title}}</h5>
                <p class="card-text">Goal: {{task.goal}}</p>
                <p class="card-text">Due date: {{task.duedate}}</p>
                <button class="btn btn-primary btn-block" @click="openModal()">Edit</button>
                <button class="btn btn-danger btn-block" @click="deleteTask">Delete</button>
            </div>
        </div>


    </div>

</template>

<script>
    import TaskService from '../services/Task.js'

    export default {
        name: "task-component",
        components: {
        },
        props: {
            task: Object
        },
        data() {
            return {
            }
        },
        methods: {
            openModal() {
                this.$emit('openModal', this.task);
            },
            async deleteTask() {
                await TaskService.deleteTask(this.task.id);
                this.$emit('updateTasks');
                this.$toasted.info('Task is deleted')

            }

        },

    }
</script>

<style scoped>

</style>