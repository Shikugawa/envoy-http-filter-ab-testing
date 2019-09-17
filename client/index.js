import ReactDOM from "react-dom";
import React from "react";
import { CookiesProvider } from "react-cookie";

class LoginForm extends React.Component {
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
      '' ,{
        method: 'post',
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

class Color extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      color: 'none'
    }
  }

  componentDidMount() {
    const {cookies} = this.props;
    if (!cookies.get("session_id")) {
      return;
    }
    fetch(
      '' ,{
        method: 'post',
        credentials: 'same-origin',
        headers: {
          "Content-Type": "application/json",
          "x-user-session-id-mod": parseInt(cookies.get("user_id"))
        }
      }
    ).then(resp => {
      if (resp.status >= 400) {
        throw new Error(resp.status)
      } else {
        return resp.json()
      }
    }).then(resp =>
      this.setState('color', resp.color)
    ).catch(err => console.error(err))
  }

  render() {
    return (
      <div>
      </div>
    )
  } 
}

class MainComponent extends React.Component {
  render() {
    return (        
      <BrowserRouter>
        <Switch>
          <Route 
            exact={true}
            path="/"
            render={ () => {
              return (
                <div>
                  <Color />
                </div>
              );
            }}
          />
          <Route
            exact={true}
            path="/signin"
            render={ (props) => {
              return (
                <div>
                  <LoginForm />
                </div>
              );
            }}
          />
        </Switch>
      </BrowserRouter>
    );
} 

ReactDOM.render(
  <CookiesProvider>
    <MainComponent />
  </CookiesProvider>, document.querySelector('#container')
);