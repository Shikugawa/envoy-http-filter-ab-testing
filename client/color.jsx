import React from "react";
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { withCookies } from 'react-cookie';

const API_ENDPOINT = "http://localhost:8080"

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
      API_ENDPOINT + "/welcome" ,{
        method: 'post',
        credentials: 'same-origin',
        mode: 'no-cors',
        headers: {
          "Content-Type": "application/json",
          "x-user-session-id-mod": parseInt(cookies.get("user_id")) % 10 // 10の剰余をヘッダに乗せて、EnvoyのHTTP Filterでパースしてリクエスト先を決定する
        },
        body: JSON.stringify({
          session_id: cookies.get("session_id")
        })
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
        <p>{this.state.color}</p>
      </div>
    )
  } 
}

export default withCookies(Color)