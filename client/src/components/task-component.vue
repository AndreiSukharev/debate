<template>

    <div>
        <div class="card text-center" style="width: 18rem;">
            <img class="card-img-top" src="" alt="photo">
            <div class="card-body">
                <h5 class="card-title">{{task.title}}</h5>
                <p class="card-text">Goal: {{task.goal}}</p>
                <p class="card-text">Due date: {{task.dueDate}}</p>
                <button class="btn btn-primary btn-block" @click="openModal()">Edit</button>
                <button class="btn btn-danger btn-block">Delete</button>
            </div>
        </div>


    </div>

</template>

<script>

    export default {
        name: "task-component",
        components: {
        },
        props: {
            task: Object
        },
        data() {
            return {
                // task: {
                //     title: '',
                //     goal: '',
                //     dueDate: ''
                // }
            }
        },
        methods: {

            openModal() {
                this.$root.$emit('getModalEditTask', this.task);
                this.$bvModal.show('modal-edit-task')
            },

        },

    }
</script>

<style scoped>

</style>