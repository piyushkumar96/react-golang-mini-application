import React, { Component } from "react";
import "./Home.css";
export default class Home extends Component {
  render() {
    return (
      <div className="text-center">
        <h1>This is the home page</h1>
        <hr />
        <div className="tickets"></div>
      </div>
    );
  }
}
