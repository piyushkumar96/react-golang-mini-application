import React, { Component, Fragment } from "react";
import { Link } from "react-router-dom";

export default class Genres extends Component {
  state = {
    genres: [],
    isLoaded: false,
    error: null,
  };
  componentDidMount() {
    fetch("http://localhost:4000/v1/genres")
      .then((response) => {
        if (response.status !== 200) {
          let err = Error;
          err.message = "Invalid repsonse code: " + response.status;
          this.setState({ error: err });
        }
        return response.json();
      })
      .then((json) => {
        this.setState(
          {
            genres: json.genres,
            isLoaded: true,
          },
          (error) => {
            this.setState({
              isLoaded: true,
              error,
            });
          }
        );
      });
  }

  render() {
    const { genres, error, isLoaded } = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
      return <p>Loading....</p>;
    } else {
      return (
        <Fragment>
          <h2>Genres</h2>
          <ul>
            {genres.map((m) => (
              <li key={m.id}>
                <Link
                  to={{ pathname: `/genre/${m.id}`, genreName: m.genre_name }}
                >
                  {m.genre_name}
                </Link>
              </li>
            ))}
          </ul>
        </Fragment>
      );
    }
  }
}
