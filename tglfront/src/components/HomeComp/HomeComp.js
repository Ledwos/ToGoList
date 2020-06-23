import React from 'react';
import LogInComp from '../LogInComp/LogInComp';
import SignUpComp from '../SignUpComp/SignUpComp';


const HomeComp = (props) => {
    if (props.loginState) {
        return <LogInComp {...props} />
    } else if (props.signupState) {
        return <SignUpComp {...props} />
    } else {
        return (
            <div>
                <p>Welcome to ToGoList!</p>
                <p>A task manager app to help you keep track of your work!</p>
                <p>Are you ready to become more organised? Click below to get started:</p>
                <button onClick={props.toggleSignup}>Organise Me!</button>
            </div>
        );
    }
};

export default HomeComp;