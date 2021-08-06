import React, { useState, useEffect, Fragment } from "react";
import { Link } from "react-router-dom";
import { confirmAlert } from "react-confirm-alert"; // Import
import "react-confirm-alert/src/react-confirm-alert.css";

import "./EditMovie.css";
import Input from "./form-components/Input";
import TextArea from "./form-components/TextArea";
import Select from "./form-components/Select";
import Alert from "./ui-components/Alert";

function EditMovieFunc(props) {
  const [movie, setMovie] = useState({
    id: 0,
    title: "",
    release_date: "",
    runtime: "",
    cbfc_rating: "",
    rating: "",
    description: "",
  });
  const [error, setError] = useState(null);
  const [valErrors, setValErrors] = useState([]);
  const [alert, setAlert] = useState({ type: "d-none", message: "" });
  const cbfcOptions = [
    { id: "U", value: "U" },
    { id: "U/A", value: "U/A" },
    { id: "A", value: "A" },
    { id: "S", value: "S" },
  ];

  useEffect(() => {
    if (props.jwt === "") {
      props.history.push({
        pathname: "/login",
      });
      return;
    }

    const id = props.match.params.id;

    if (id > 0) {
      fetch("http://localhost:4000/v1/movie/" + id)
        .then((response) => {
          if (response.status !== 200) {
            setError("Invalid response : ", response.status);
          } else {
            setError(null);
          }
          return response.json();
        })
        .then((json) => {
          const releaseDate = new Date(json.movie.release_date);

          json.movie.release_date = releaseDate.toISOString().split("T")[0];
          setMovie(json.movie);
        });
    }
  }, [props.history, props.jwt, props.match.params.id]);

  const handleChange = () => (evt) => {
    let name = evt.target.name;
    let value = evt.target.value;
    setMovie({
      ...movie,
      [name]: value,
    });
  };

  const handleSubmit = (evt) => {
    evt.preventDefault();

    // client side input validation
    let valErrors = [];
    if (movie.title === "") {
      valErrors.push("title");
    }
    if (movie.release_date === "") {
      valErrors.push("release_date");
    }
    if (movie.runtime === "") {
      valErrors.push("runtime");
    }
    if (movie.cbfc_rating === "") {
      valErrors.push("cbfc_rating");
    }
    if (movie.rating === "") {
      valErrors.push("rating");
    }
    if (movie.description === "") {
      valErrors.push("description");
    }

    setValErrors(valErrors);

    if (valErrors.length > 0) {
      return false;
    }

    const data = new FormData(evt.target);
    const payload = Object.fromEntries(data.entries());
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + props.jwt);

    const requestOptions = {
      method: "POST",
      body: JSON.stringify(payload),
      headers: headers,
    };

    fetch("http://localhost:4000/v1/admin/editmovie", requestOptions)
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          setAlert({
            type: "alert-danger",
            message: data.error.message,
          });
        } else {
          setAlert({
            type: "alert-success",
            message: "Changes saved!",
          });
          props.history.push({
            pathname: "/admin",
          });
        }
      });
  };

  const confirmDelete = (e) => {
    confirmAlert({
      title: "Delete Movie!",
      message: "Are you sure?",
      buttons: [
        {
          label: "Yes",
          onClick: () => {
            const headers = new Headers();
            headers.append("Content-Type", "application/json");
            headers.append("Authorization", "Bearer " + props.jwt);

            fetch("http://localhost:4000/v1/admin/movie/" + movie.id, {
              method: "DELETE",
              headers: headers,
            })
              .then((response) => response.json)
              .then((data) => {
                if (data.error) {
                  setAlert({
                    type: "alert-danger",
                    message: data.error.message,
                  });
                } else {
                  setAlert({
                    type: "alert-success",
                    message: "Movie deleted",
                  });
                  props.history.push({
                    pathname: "/admin",
                  });
                }
              });
          },
        },
        {
          label: "No",
          onClick: () => {},
        },
      ],
    });
  };

  function hasValError(key) {
    return valErrors.indexOf(key) !== -1;
  }

  if (error !== null) {
    return <div>Error: {error.message}</div>;
  } else {
    return (
      <Fragment>
        <h2>Add/Edit Movie</h2>
        <Alert alertType={alert.type} alertMessage={alert.message} />
        <hr />
        <form onSubmit={handleSubmit}>
          <input
            type="hidden"
            name="id"
            id="id"
            value={movie.id}
            onChange={handleChange(movie.id)}
          />
          <Input
            title={"Title"}
            className={hasValError("title") ? "is-invalid" : ""}
            type={"text"}
            name={"title"}
            value={movie.title}
            handleChange={handleChange("title")}
            errorDiv={hasValError("title") ? "text-danger" : "d-none"}
            errorMsg={"Please enter the title"}
          />
          <Input
            title={"Release"}
            className={hasValError("release_date") ? "is-invalid" : ""}
            type={"date"}
            name={"release_date"}
            value={movie.release_date}
            handleChange={handleChange("release_date")}
            errorDiv={hasValError("release_date") ? "text-danger" : "d-none"}
            errorMsg={"Please select release date"}
          />
          <Input
            title={"Runtime"}
            className={hasValError("runtime") ? "is-invalid" : ""}
            type={"text"}
            name={"runtime"}
            value={movie.runtime}
            handleChange={handleChange("runtime")}
            errorDiv={hasValError("runtime") ? "text-danger" : "d-none"}
            errorMsg={"Please enter the runtime"}
          />
          <Select
            title={"CBFC Rating"}
            className={hasValError("cbfc_rating") ? "is-invalid" : ""}
            name={"cbfc_rating"}
            options={cbfcOptions}
            value={movie.cbfc_rating}
            handleChange={handleChange("cbfc_rating")}
            placeholder={"Choose..."}
            errorDiv={hasValError("cbfc_rating") ? "text-danger" : "d-none"}
            errorMsg={"Please select cbfc rating"}
          />
          <Input
            title={"Rating"}
            className={hasValError("rating") ? "is-invalid" : ""}
            type={"text"}
            name={"rating"}
            value={movie.rating}
            handleChange={handleChange("rating")}
            errorDiv={hasValError("rating") ? "text-danger" : "d-none"}
            errorMsg={"Please enter rating"}
          />
          <TextArea
            title={"Description"}
            className={hasValError("description") ? "is-invalid" : ""}
            type={"textarea"}
            name={"description"}
            value={movie.description}
            handleChange={handleChange("description")}
            rows="3"
            errorDiv={hasValError("description") ? "text-danger" : "d-none"}
            errorMsg={"Please write description"}
          />
          <hr />
          <button className="btn btn-primary">Save</button>
          <Link to="/admin" className="btn btn-warning ms-1">
            Cancel
          </Link>
          {movie.id > 0 && (
            <a
              href="#!"
              /* if u don't use () => then it will just fire up alert on page load*/
              onClick={() => confirmDelete()}
              className="btn btn-danger ms-1"
            >
              Delete
            </a>
          )}
        </form>
      </Fragment>
    );
  }
}

export default EditMovieFunc;
