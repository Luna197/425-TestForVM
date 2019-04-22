import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import '../node_modules/bootstrap/dist/css/bootstrap.css';
import '../node_modules/semantic-ui-css/semantic.min.css';

import {HashRouter as Router, Route, Link, hashHistory } from 'react-router-dom'


// Include your new Components here
import Contest from './components/Contest/Contest.js';
import Submission from './components/Submission/Submission.js';


 ReactDOM.render(
<div>
 
 {/* <Contest />, document.getElementById('root') */}
 
 <Router  >
 <div>
 <Route exact path='/' component={Contest}/>
 <Route exact path='/submission' component={Submission}/>
 </div>
</Router>
 </div>,
 document.getElementById('root')
 
 );
// ReactDOM.render(<Submission />, document.getElementById('root'));

serviceWorker.unregister();
