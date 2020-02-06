<template>

    <form>
        <div class="form-group">
            <label for="title">Title</label>
            <input type="text" class="form-control" id="title" v-model="task.title" placeholder="">
        </div>

        <div class="form-group">
            <label for="goal">Goal</label>
            <input type="text" class="form-control" id="goal" v-model="task.goal" placeholder="Enter your goal">
        </div>

        <div class="form-group">
            <label for="duedate">Due Date</label>
            <input type="date" class="form-control" id="duedate" v-model="task.duedate"
                   placeholder="Enter your deadline">
        </div>

    </form>

</template>

<script>
    export default {
        name: "task-form-component",
        props: {
            taskForForm: Object
        },
        data() {
            return {
                task: {
                    id: '',
                    title: '',
                    goal: '',
                    duedate: ''
                }
            }
        },
        mounted() {
        },
        watch: {
            task: {
                immediate: true,
                deep: true,
                handler() {
                    this.$emit('updateTask', this.task)
                }
            },
            taskForForm: {
                immediate: true,
                deep: true,
                handler(task) {
                    this.task = task;
                }
            },
        }
    }
</script>

<style scoped>

</style>