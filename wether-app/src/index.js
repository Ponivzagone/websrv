import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min"
import React from 'react';
import ReactDOM from 'react-dom';
import {Router} from "react-router";
import { createBrowserHistory } from 'history';
import App from './component/App';

const history = createBrowserHistory();

ReactDOM.render(
    <Router history={history}>
        <App />
    </Router>,
    document.getElementById('root')
);

