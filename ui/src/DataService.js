import axios from 'axios'

export function getTestData(callback) {
  axios.get('/api/test')
    .then((resp) => callback(resp.data))
    .catch((resp) => console.log(resp))
}