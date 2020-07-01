import React from 'react';
import './NotFound.css';

const NotFound = (props) => {
    return (
        <div id='nfDiv'>
            <h1 id='errTitle'>404</h1>
            <p id='errTxt'>Whoops, something<br /> went wrong  X_X</p>
            <button id='errBtn' onClick={props.homePage}>Back to Home</button>
        </div>
    );
};

export default NotFound;