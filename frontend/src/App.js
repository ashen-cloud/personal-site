import React, { Component } from 'react';
import { get } from './api'
// import logo from './logo.svg';
import './App.css';

const URL_BASE = "http://ashencloud.xyz:8080/"

const promisify = (fn, ...args) => {
  return new Promise(resolve => {
    fn(...args, r => {
      resolve(r)
    })
  })
}

class App extends Component {
  constructor(props) {
    super()
    this.state = {
      data: null,
    }
  }

  componentDidMount() {
    promisify(get, URL_BASE).then(r => this.setState({ fills: r }))
    promisify(get, URL_BASE + "posts").then(r => {
      console.log(r)
      this.setState({ posts: r })
    })
  }

  render() {
    const fills = this.state.fills ? this.state.fills : { links: [] }
    const postsRaw = this.state.posts ? this.state.posts : []

    const links = fills.links.map(l => {
      return <a href={ l[1] }>{ l[0] }</a> // eslint-disable-line
    })
    const posts = postsRaw.map(p => {
      const n = p.Name
      return <li className="posts-item"><a href="/posts/{ n }"></a><p>{ n }</p></li> // eslint-disable-line
    })
    return (
      <body className="body">
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
              <p className="fills-description">{ fills.title }</p>
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
      </body>
    );
  }
}


export default App;
