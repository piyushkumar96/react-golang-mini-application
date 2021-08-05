import React, {Component, Fragment} from 'react';

export default class GraphQL extends Component {
    constructor(props) {
        super(props);

        this.state = {
            movies: [],
            isLoaded: false,
            error: null,
            alert: {
                type: "d-none",
                message: "",
            }
        }
    }

    componentDidMount() {
        const payload = `
        {   
            list {
                id
                title
                runtime
                year
                description
            }
        }
        `

        const headers = new Headers();
        headers.append("Content-Type", "application/json")

        const requestOptions = {
            method: "POST",
            body: payload,
            headers: headers
        }

        fetch("http://localhost:4000/v1/graphql", requestOptions)
            .then((response) => response.json())
            .then((data) => {
                let theList = Object.values(data.data.list)
                return theList
            })
            .then((theList) => {
                this.setState({
                    movies: theList
                })
            })
    }

    render() {
        let { movies } = this.state
        return (
            <Fragment>
                <h2>GraphQL</h2>
                <hr />
                <div className="list-group">
                    {movies.map((m) => (
                        <a
                        key={m.id}
                        className="list-group-item list-group-item-action"
                        href="#!">
                            <strong>{m.title}</strong><br />
                            <small className="text-muted">
                                ({m.year}) - {m.runtime} minutes
                            </small>
                            <br />
                            {m.description.slice(0, 100)}...

                        </a>
                    ))}
                </div>
            </Fragment>

        )
    }
}