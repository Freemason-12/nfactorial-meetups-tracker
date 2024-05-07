async function getMeetups() {
  // return fetch('/api/meetups').then(response => response.text()).then(text => console.log(text))
  return fetch('/api/meetups').then(response => response.json())
}

async function getMeetup(id) {
  return fetch(`/api/meetup?id=${id}`).then(response => response.json())
}

export { getMeetups, getMeetup }
