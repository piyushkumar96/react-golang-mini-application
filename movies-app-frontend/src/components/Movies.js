import React, { Component, Fragment } from "react";
import { Link } from "react-router-dom";

export default class Movies extends Component {
  state = { movies: [] };

  componentDidMount() {
    this.setState({
      movies: [
        { id: 1, title: "The Shawshank Redemption", runtime: 142 },
        { id: 2, title: "Theory of Everything", runtime: 122 },
      ],
    });
  }

  render() {
    return (
      <Fragment>
        <h1>Choose a movie</h1>
        <ul>
          {this.state.movies.map((m) => (
            <li key={m.id}>
              <Link to={`/movies/${m.id}`}>{m.title}</Link>
            </li>
          ))}
        </ul>
      </Fragment>
    );
  }
}
