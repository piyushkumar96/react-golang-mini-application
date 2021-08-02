import React, { Component, Fragment } from "react";

import "./EditMovie.css";
import Input from "./form-components/Input";
import TextArea from "./form-components/TextArea";
import Select from "./form-components/Select";

export default class EditMovie extends Component {
  constructor(props) {
    super(props);
    this.state = {
      movie: {
        id: 0,
        title: "",
        release_date: "",
        runtime: "",
        cbfc_rating: "",
        rating: "",
        description: "",
      },
      cbfcOptions: [
        { id: "U", value: "U" },
        { id: "U/A", value: "U/A" },
        { id: "A", value: "A" },
        { id: "S", value: "S" },
      ],
      isLoaded: false,
      error: null,
      valErrors: [],
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit = (evt) => {
    evt.preventDefault();

    // client side input validation
    let valErrors = [];
    if (this.state.movie.title === "") {
      valErrors.push("title");
    }
    if (this.state.movie.release_date === "") {
      valErrors.push("release_date");
    }
    if (this.state.movie.runtime === "") {
      valErrors.push("runtime");
    }
    if (this.state.movie.cbfc_rating === "") {
      valErrors.push("cbfc_rating");
    }
    if (this.state.movie.rating === "") {
      valErrors.push("rating");
    }
    if (this.state.movie.description === "") {
      valErrors.push("description");
    }

    this.setState({ valErrors: valErrors });

    if (valErrors.length > 0) {
      return false;
    }

    const data = new FormData(evt.target);
    const payload = Object.fromEntries(data.entries());
    const requestOptions = {
      method: "POST",
      body: JSON.stringify(payload),
    };

    fetch("http://localhost:4000/v1/admin/editmovie", requestOptions)
      .then((response) => response.json())
      .then((data) => {
        alert(JSON.stringify(data));
      });
  };

  handleChange = (evt) => {
    let name = evt.target.name;
    let value = evt.target.value;
    this.setState((prevState) => ({
      movie: {
        ...prevState.movie,
        [name]: value,
      },
    }));
  };

  hasValError(key) {
    return this.state.valErrors.indexOf(key) !== -1
  }

  componentDidMount() {
    const id = this.props.match.params.id;
    if (id > 0) {
      fetch("http://localhost:4000/v1/movie/" + id)
        .then((response) => {
          if (response.status !== 200) {
            let err = Error;
            err.message = "Invalid reponse code: " + response.status;
            this.setState({ error: err });
          }
          return response.json();
        })
        .then((json) => {
          const releaseDate = new Date(json.movie.release_date);

          this.setState(
            {
              movie: {
                id: id,
                title: json.movie.title,
                release_date: releaseDate.toISOString().split("T")[0],
                runtime: json.movie.runtime,
                cbfc_rating: json.movie.cbfc_rating,
                rating: json.movie.rating,
                description: json.movie.description,
              },
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
    } else {
      this.setState({
        isLoaded: true,
      });
    }
  }

  render() {
    let { movie, isLoaded, error } = this.state;

    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
      return <p>Loading...</p>;
    } else {
      return (
        <Fragment>
          <h2>Add/Edit Movie</h2>
          <hr />
          <form onSubmit={this.handleSubmit}>
            <input
              type="hidden"
              id="id"
              value={movie.id}
              onChange={this.handleChange}
            />
            <Input
              title={"Title"}
              className={this.hasValError("title")? "is-invalid": ""}
              type={"text"}
              name={"title"}
              value={movie.title}
              handleChange={this.handleChange}
              errorDiv={this.hasValError("title") ? "text-danger": "d-none"}
              errorMsg={"Please enter the title"}
            />
            <Input
              title={"Release"}
              className={this.hasValError("release_date")? "is-invalid": ""}
              type={"date"}
              name={"release_date"}
              value={movie.release_date}
              handleChange={this.handleChange}
              errorDiv={this.hasValError("release_date") ? "text-danger": "d-none"}
              errorMsg={"Please select release date"}
            />
            <Input
              title={"Runtime"}
              className={this.hasValError("runtime")? "is-invalid": ""}
              type={"text"}
              name={"runtime"}
              value={movie.runtime}
              handleChange={this.handleChange}
              errorDiv={this.hasValError("runtime") ? "text-danger": "d-none"}
              errorMsg={"Please enter the runtime"}
            />
            <Select
              title={"CBFC Rating"}
              className={this.hasValError("cbfc_rating")? "is-invalid": ""}
              name={"cbfc_rating"}
              options={this.state.cbfcOptions}
              value={movie.cbfc_rating}
              handleChange={this.handleChange}
              placeholder={"Choose..."}
              errorDiv={this.hasValError("cbfc_rating") ? "text-danger": "d-none"}
              errorMsg={"Please select cbfc rating"}
            />
            <Input
              title={"Rating"}
              className={this.hasValError("rating")? "is-invalid": ""}
              type={"text"}
              name={"rating"}
              value={movie.rating}
              handleChange={this.handleChange}
              errorDiv={this.hasValError("rating") ? "text-danger": "d-none"}
              errorMsg={"Please enter rating"}
            />
            <TextArea
              title={"Description"}
              className={this.hasValError("description")? "is-invalid": ""}
              type={"textarea"}
              name={"description"}
              value={movie.description}
              handleChange={this.handleChange}
              rows="3"
              errorDiv={this.hasValError("description") ? "text-danger": "d-none"}
              errorMsg={"Please write description"}
            />
            <hr />
            <button className="btn btn-primary">Save</button>
          </form>
          <div className="mt-3">
            <pre>{JSON.stringify(this.state, null, 3)}</pre>
          </div>
        </Fragment>
      );
    }
  }
}
