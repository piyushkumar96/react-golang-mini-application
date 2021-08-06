import React, { useEffect, useState, Fragment } from "react";
import { Link } from "react-router-dom";

function MoviesFunc(props) {
  const [movies, setMovies] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("http://localhost:4000/v1/movies")
      .then((response) => {
        if (response.status !== 200) {
          setError("Invalid response code: ", response.status);
        } else {
          setError(null);
        }

        return response.json();
      })
      .then((json) => {
        setMovies(json.movies);
      });
  }, []); // *** must set to default value
  if (error !== null) {
    return <div>Error: {error.message}</div>;
  } else {
    return (
      <Fragment>
        <h1>Choose a movie</h1>
        <div className="list-group">
          {movies.map((m) => (
            <Link
              key={m.id}
              className="list-group-item list-group-item-action"
              to={`/movies/${m.id}`}
            >
              {m.title}
            </Link>
          ))}
        </div>
      </Fragment>
    );
  }
}

export default MoviesFunc;
