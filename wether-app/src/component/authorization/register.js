import React from "react";
import {fetchRegistration} from "../../api/authorization";

import "./style.css"

export default class Registration extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            user: "",
            password: ""
        };
    }

    handleUsernameChange = (event) => {
        this.setState({user: event.target.value});
    };

    handlePasswordChange = (event) => {
        this.setState({password: event.target.value});
    };

    register = async (e) => {
        e.preventDefault();
        try {
            const token = await fetchRegistration(this.state.user, this.state.password);

            //this.props.onRegistered(token);
            this.props.history.push("/login");
        } catch (e) {
            console.log(e);
        }
    };

    render() {
        return (
            <div className="container h-100">
                <div className="row">
                    <div className="col-md-12">
                        <h3>Fill Registration form</h3>
                    </div>
                </div>
                <form >
                    <div className="row">
                        <div className="col-md-3 offset-md-2">
                            <input 
                                type="text"
                                className="form-control"
                                value={this.state.user}
                                onChange={this.handleUsernameChange}
                                placeholder="Username"
                                name="username"
                                autoComplete="off"
                            />
                        </div>
                        <div className="col-md-3">
                            <input 
                                type="password"
                                className="form-control"
                                value={this.state.password}
                                onChange={this.handlePasswordChange}
                                placeholder="Password"
                                name="password"
                                autoComplete="off"
                            />
                        </div>
                        <div className="col-md-3 mt-md-0 mt-2 text-md-left">
                            <button
                                type="submit"
                                onClick={this.register}
                                className="btn btn-warning"
                            >
                                Sign Up
                            </button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}