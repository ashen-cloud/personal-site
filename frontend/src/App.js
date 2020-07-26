import React, { Component } from 'react';
import { get } from './api'
// import logo from './logo.svg';
import './App.css';

// const URL_BASE = "https://ashencloud.xyz:8080/"
const URL_BASE = "http://192.168.88.211:8080/"

const promisify = (fn, ...args) => {
  return new Promise(resolve => {
    fn(...args, r => {
      resolve(r)
    })
  })
}

class App extends Component {
  constructor(_) {
    super()
    this.state = {
      data: null,
    }
  }

  componentDidMount() {
    promisify(get, URL_BASE).then(r => r && this.setState({ fills: r }))
    promisify(get, URL_BASE + "posts").then(r => r && this.setState({ posts: r }))
  }

  showPost(name) {
    const posts = this.state.posts
    if (posts && posts.length) {
      const post = posts.reverse().find(x => x.Name === name)
      promisify(get, URL_BASE + "posts/" + post.Id).then(r => {
        if (r) {
          this.setState({ postRes: r })
          this.render()
        }
      })
    }
  }

  render() {
    const fills = this.state.fills ? this.state.fills : { links: [] }
    const postsRaw = this.state.posts ? this.state.posts : []
    const postText = this.state.postRes ? this.state.postRes.Content : ""

    const genRandKey = _ => (Math.random() + 1) + ""
    
    const links = fills.links.map(l => {
      return <a key={ genRandKey() } href={ l[1] }> { l[0] } </a>
    })
    const posts = postsRaw.map(p => {
      const n = p.Name
      return <li key={ genRandKey() } className="posts-item"><a onClick={ _ => this.showPost(n) }></a><p>{ n }</p></li> // eslint-disable-line
    })
    return (
      <div className="body">
        <div className="main">

          <div className="main_title">
            <span>
              { fills.name }
            </span>
          </div>

          <div className="links">
            { links }
          </div>

          <div className="site-content">
            <div className="content">
              <h1>{ fills.title }</h1>
              <p className="fills-description">{ fills.description }</p>
            </div>

            <div className="post-text">
              <p> { postText } </p>
            </div>

            <div className="side-bar">
              <div className="posts-title">
                POSTS&nbsp;
              </div>
              <ul>
                { posts }
              </ul>
            </div>
          </div>
        </div>
      </div>
    );
  }
}


export default App;
