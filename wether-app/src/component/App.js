import React from "react";
import {Switch, Route} from "react-router";

import Login from "./authorization/login";
import Registration from "./authorization/register";
import Wrapper from "./weather/wrapper"

import "./App.css"

class App extends React.Component {
  constructor(props) {
    super(props);

    const token = localStorage.getItem("bearer");
    this.state = {
        token: token
    }

  }

  handleAuthentication = (token) => {
      localStorage.setItem("bearer", token);
      this.setState({ token });
  };

  handleLogout = () => {
      localStorage.removeItem("bearer");
      this.setState({ token: null });
  };

  render() {
    return (

      <div data-test-component="App" className="AppContent">
        <Switch>
          <Route 
            exact path="/"
            render={(props) =>
                <Wrapper
                    {...props}
                    token={this.state.token}
                    onLogout={this.handleLogout}
                />
            }
          />
          <Route
            path="/login"
            render={(props) =>
                <Login
                    {...props}
                    onAuthenticated={this.handleAuthentication}
                />
            }
          />
          <Route
            path="/register"
            render={(props) =>
                <Registration
                    {...props}
                    onRegistered={this.handleAuthentication}
                />
            }
          />
        </Switch>
      </div>
    );
  }
}

export default App;

