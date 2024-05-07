import * as api from './api.js'
const root = document.querySelector('main')
const homeRoot = document.getElementById('home_root')
const meetupInfoRoot = document.getElementById('meetup_info_root')

function clearBase() {
  homeRoot.style.display = 'none'
  meetupInfoRoot.style.display = 'none'
  // for (let node of root.children) {
  //   node.style.display = 'none'
  //   node.innerHTML = ''
  // }
}

function viewMoreRoute(e) {
  e.preventDefault()
  history.pushState(null, '', e.target.href)
  route()
}

async function homepage() {
  clearBase()
  homeRoot.style.display = 'block'
  homeRoot.innerText = "Loading Content";
  const contents = await api.getMeetups()
  console.log(contents)
  homeRoot.innerText = "";
  for (let i of contents) {
    const meetup = document.createElement('div')
    meetup.classList.add('meetup')
    meetup.innerHTML = `
      <h2> ${i.name} </h2>
      <p>Starts at: <b>${i.date}</b> </p>
      <p>Address: <b>${i.address}</b> </p>
      <a class="view_more" href="/meetup?id=${i.id}">View More</a>
    `
    meetup.querySelector('a').addEventListener('click', (e) => viewMoreRoute(e))
    homeRoot.appendChild(meetup)
  }
}

async function meetupInfo(id) {
  clearBase()
  meetupInfoRoot.style.display = 'block'
  console.log(meetupInfoRoot.style.display)
  meetupInfoRoot.innerText = 'Loading content'
  const contents = await api.getMeetup(id)
  meetupInfoRoot.innerText = ''
  meetupInfoRoot.innerHTML = `
    <h1> ${contents.name} </h1>
    <p> ${contents.description} </p>
    <div class="metadata">
      <div>Starts at: <b>${contents.date}</b></div>
      <div>Address: <b>${contents.address}</b></div>
    </div>
  `
}

function route() {
  const regex = new RegExp("https?://[^/]+")
  const urlPath = location.href.slice(location.href.match(regex)[0].length)
  console.log(urlPath)

  switch (true) {
    case /\/meetup[?]id=\d+/.test(urlPath):
      console.log("here")
      meetupInfo(Number(urlPath.slice(urlPath.indexOf('=') + 1)));
      break
    case /\//.test(urlPath): homepage(); break
  }
}

function main() {
  window.addEventListener('load', route)
  window.addEventListener('popstate', route)
} main()

// console.log(history.state)
// history.pushState({ data: "here" }, '', "/newstate")
// console.log(history.state)
// console.log(location.href)
