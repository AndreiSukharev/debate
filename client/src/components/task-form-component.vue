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
            <label for="dueDate">Due Date</label>
            <input type="date" class="form-control" id="dueDate" v-model="task.dueDate" placeholder="Enter your deadline">
        </div>

    </form>

</template>

<script>
    export default {
        name: "task-form-component",
        data() {
            return {
                task: {
                    id: '',
                    title: '',
                    goal: '',
                    dueDate: ''
                }
            }
        },
        mounted() {
            this.$root.$on('getModalEditTask', (task) => {
                this.task = task;
            });
        },
        watch: {
            task: {
                immediate: true,
                deep: true,
                handler() {
                    this.$emit('updateTask', this.task)
                }
            }
        }
    }
</script>

<style scoped>

</style>