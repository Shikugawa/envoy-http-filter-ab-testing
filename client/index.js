import ReactDOM from "react-dom";
import React from "react";
import { CookiesProvider } from "react-cookie";
import { MainComponent } from "./main"; 

ReactDOM.render(
  <CookiesProvider>
    <MainComponent />
  </CookiesProvider>, document.querySelector('#container')
);