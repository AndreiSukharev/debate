import api from './api.js';

class RegistrationService {

    static signUp(data) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await api().get('signup');
                // const res = await api().post('signup', data);
                resolve(res.data)
            } catch (err) {
                reject(err);
            }
        })
    }

    static signIn(data) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await api().post('signin', data);
                resolve(res.data)
            } catch (err) {
                reject(err);
            }
        })
    }
}

export default RegistrationService;