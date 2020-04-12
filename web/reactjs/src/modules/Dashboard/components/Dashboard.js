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
import {ServiceItem} from './ServiceItem';
import Grid from '@material-ui/core/Grid';

import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({
    root: {
        padding: theme.spacing(4)
    }
}));

const addServiceItems = () => {
    var items = [];
    for (var index=0; index < 3; index++){
        items.push(
            <ServiceItem key={index}/>
        );
    }
    return items;
};

const Dashboard = () => {

    const classes = useStyles();

    return (
        <div className={classes.root}>
            <Grid container spacing={10}>
                {
                    addServiceItems()
                }
            </Grid>
        </div>
    );
};

export default Dashboard;
