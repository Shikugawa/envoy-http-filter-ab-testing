import React from "react";
import { BrowserRouter, Switch, Route } from 'react-router-dom';

const API_ENDPOINT = "http://localhost:8080"

export class LoginForm extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      username: '',
      password: ''
    }
  }
  
  changeHandler(e) {
    const key = e.target.id;
    this.setState({[key]: e.target.value});
  }

  loginHandler(e) {
    fetch(
      API_ENDPOINT + "/login" ,{
        method: 'post',
        mode: 'no-cors',
        credentials: 'same-origin',
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          username: this.state.username,
          password: this.state.password
        })
      }
    ).then(resp => {
      if (resp.status >= 400) {
        alert(resp.status)
      }
    }).catch(err => console.error(err))
  }

  render() {
    return (
      <div>
        <form>
          <div><input ref="username" placeholder="username" onChange={this.changeHandler.bind(this)} /></div>
          <div><input ref="password" placeholder="password" onChange={this.changeHandler.bind(this)} /></div>
          <button onClick={this.loginHandler.bind(this)}>ログイン</button>
        </form>
      </div>
    )
  }
}