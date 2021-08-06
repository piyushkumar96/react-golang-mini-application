import React, { useState, useEffect, Fragment } from "react";
import { Link } from "react-router-dom";

function OneGenreFunc(props) {
  let [movies, setMovies] = useState([]);
  const [error, setError] = useState(null);
  let [genreName, setGenreName] = useState("");

  useEffect(() => {
    fetch("http://localhost:4000/v1/movies/genre/" + props.match.params.id)
      .then((response) => {
        if (response.status !== 200) {
          setError("Invalid response : ", response.status);
        } else {
          setError(null);
        }
        return response.json();
      })
      .then((json) => {
        setGenreName(props.location.genreName);
        setMovies(json.movies);
      });
  }, [props.match.params.id, props.location.genreName]);

  if (!movies) {
    movies = [];
  }

  if (error !== null) {
    return <div>Error: {error.message}</div>;
  } else {
    return (
      <Fragment>
        <h2>Genre: {genreName}</h2>
        <div className="list-group">
          {movies.map((m) => (
            <Link
              to={`/movies/${m.id}`}
              className="list-group-item list-group-item-action"
            >
              {m.title}
            </Link>
          ))}
        </div>
      </Fragment>
    );
  }
}

export default OneGenreFunc;
