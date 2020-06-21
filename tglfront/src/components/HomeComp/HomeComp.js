import React, { useState } from 'react';
import LogInComp from '../LogInComp/LogInComp';
import SignUpComp from '../SignUpComp/SignUpComp';


const HomeComp = (props) => {
    const [loginState, setloginState] = useState(false);
    const [signupState, setsignupState] = useState(false);


    const toggleLogin = () => {
        setsignupState(false);
        setloginState(!loginState);
    };
    const toggleSignup = () => {
        setloginState(false);
        setsignupState(!signupState);
    };

    const resetHome = () => {
        setsignupState(false);
        setloginState(false);
    }

    return (
        <div>
            <p>I'm the Home component</p>
            {props.loggedIn === 'true' ? <button onClick={props.taskPage}>Tasks</button> : null}
            {props.loggedIn === 'true' ? <button onClick={props.logOut}>Log out</button> : <button onClick={toggleLogin}>Log In</button>}
            {loginState | signupState === true ? null : <p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus.</p>}
            {loginState ? <LogInComp toggleSignup={toggleSignup} resetHome={resetHome} {...props} /> : null}
            {signupState ? <SignUpComp toggleLogin={toggleLogin} resetHome={resetHome} /> : null}
        </div>
    );
};

export default HomeComp;