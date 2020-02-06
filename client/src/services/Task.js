import api from './api.js';

class TaskService {

    static getTasks(data) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await api().get('tasks');
                resolve(res.data)
            } catch (err) {
                reject(err);
            }
        })
    }

    static addTask(data) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await api().post('tasks', data);
                resolve(res.data)
            } catch (err) {
                reject(err);
            }
        })
    }

    static deleteTask(id) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await api().delete('tasks', id);
                resolve(res.data)
            } catch (err) {
                reject(err);
            }
        })
    }

    static updateTask(data) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await api().put('tasks', data);
                resolve(res.data)
            } catch (err) {
                reject(err);
            }
        })
    }


}

export default TaskService;