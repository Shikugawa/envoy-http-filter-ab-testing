import React from "react";
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Color from "./color";
import { LoginForm } from "./login";

export class MainComponent extends React.Component {
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
                  <p>ああああああ</p>
                  <LoginForm />
                </div>
              );
            }}
          />
        </Switch>
      </BrowserRouter>
    );
  }
}