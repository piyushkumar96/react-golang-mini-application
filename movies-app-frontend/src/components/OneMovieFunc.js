import React, { useState, useEffect, Fragment } from "react";

function OneMovieFunc(props) {
  const [movie, setMovie] = useState({});
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("http://localhost:4000/v1/movie/" + props.match.params.id)
      .then((response) => {
        if (response.status !== 200) {
          setError("Invalid response code: ", response.status);
        } else {
          setError(null);
        }

        return response.json();
      })
      .then((json) => {
        setMovie(json.movie);
      });
  }, [props.match.params.id]);

  if (movie.genres) {
    movie.genres = Object.values(movie.genres);
  } else {
    movie.genres = [];
  }

  if (error !== null) {
    return <div>Error: {error.message}</div>;
  } else {
    return (
      <Fragment>
        <h2>
          Movie: {movie.title} ({movie.year})
        </h2>
        <div className="float-start">
          <small>Rating: {movie.cbfc_rating}</small>
        </div>
        <div className="float-end">
          {movie.genres.map((m, index) => (
            <span className="badge bg-secondary me-1" key={index}>
              {m}
            </span>
          ))}
        </div>
        <div className="clearfix"></div>

        <hr />
        <table className="table table-compact table-striped">
          <thead></thead>
          <tbody>
            <tr>
              <td>
                <strong>Title:</strong>
              </td>
              <td>{movie.title}</td>
            </tr>
            <tr>
              <td>
                <strong>Description:</strong>
              </td>
              <td>{movie.description}</td>
            </tr>
            <tr>
              <td>
                <strong>Runtime:</strong>
              </td>
              <td>{movie.runtime} minutes</td>
            </tr>
          </tbody>
        </table>
      </Fragment>
    );
  }
}

export default OneMovieFunc;
