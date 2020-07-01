import React from 'react';
import LogInComp from '../LogInComp/LogInComp';
import SignUpComp from '../SignUpComp/SignUpComp';
import './HomeComp.css';


const HomeComp = (props) => {
    if (props.loginState) {
        return <LogInComp {...props} />
    } else if (props.signupState) {
        return <SignUpComp {...props} />
    } else {
        return (
            <div id='homeDiv'>
                <div id='homeTxt'>
                    <p>Welcome to ToGoList!</p> <br/>
                    <p>A task manager app to help you keep track of your work!</p><br/>
                    {props.loggedIn ? null : <p>Are you ready to become more organised? Click below to get started:</p>}<br/>
                </div>
                {props.loggedIn ? null : <button onClick={props.toggleSignup} id='signupBtn'>Organise Me!</button> }
            </div>
        );
    }
};

export default HomeComp;