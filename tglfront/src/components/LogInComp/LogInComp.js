import React from 'react';
import './LogInComp.css';

const LogInComp = (props) => {
    return (
        <div id='loginDiv'>
            <form onSubmit={props.handleLogin} method='POST' id='loginForm'>
                {/* <label for='u_email'>email: </label> */}
                <input 
                    type='text' 
                    id='u_email' 
                    placeholder='someone@organised.com'
                    pattern="[a-zA-Z0-9!#$%&amp;'*+\/=?^_`{|}~.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*"
                    required
                    ></input>
                {/* <label for='u_pass'>password: </label> */}
                <input 
                    type='password' 
                    id='u_pass'
                    placeholder='password'
                    required
                    ></input>
                <button type='submit' id='loginBtn'>Log in</button>
            </form>
            <br />
            <br />
            <br />
            <p>Don't have an account? <span onClick={props.toggleSignup} id='clickHere'>click here</span> to sign up</p>
            {/* <p onClick={props.resetHome}>X</p> */}
        </div>
    );
};

export default LogInComp;