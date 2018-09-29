import axios from 'axios'

export function getTagPlaytimeData(username, callback) {
  let url = '/api/data/' + username;
  axios.get(url)
    .then((resp) => callback(resp.data))
    .catch((resp) => console.log(resp))
}
