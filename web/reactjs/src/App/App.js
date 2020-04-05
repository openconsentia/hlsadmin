// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react';
import { Router } from 'react-router-dom';

import { Provider } from 'react-redux';

import { ThemeProvider } from '@material-ui/styles';

import { createBrowserHistory } from 'history';
const browserHistory = createBrowserHistory();

import { Routes } from '../Routes';
import theme from './theme';
import {store} from '../modules/store';

export default class App extends React.Component {
    render() {
        return (
            <Provider store={store}>
                <ThemeProvider theme={theme}>
                    <Router history={browserHistory}>
                        <Routes />
                    </Router>
                </ThemeProvider>
            </Provider>
        );
    }
}
