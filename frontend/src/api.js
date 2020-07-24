function call(url, cb, m, h, b = {}, p = {}) {
  const pLen = Object.keys(p).length

  if (pLen) {
    url += "?"
    let i = 1
    for (const key in p) {
      url += `${key}=${p[key]}`
      if (i !== pLen) {
        url += "&"
      }
      i++
    }
  }
 
  return window.fetch(url, {
    "method": m,
    "headers": h,
    "body": Object.keys(b).length ? JSON.stringify(b) : null
  })
  .then(r => r.json())
  .then(r => cb(r))
  .catch(err => {
    console.error(err)
  })
}

function post(url, b = {}) {
  const h = {}
  return call(url, () => {}, "POST", h, b)
}

function get(url, cb, p = {}) {
  const h = {}
  return call(url, cb, "GET", h, {}, p)
}

export { post, get }