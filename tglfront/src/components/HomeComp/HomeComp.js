import React from 'react';

const HomeComp = (props) => {
    return (
        <div>
            <p>I'm the Home component</p>
            <button onClick={props.taskPage}>Tasks</button>
            <button onClick={props.loginPage}>Log In</button>
        </div>
    );
};

export default HomeComp;