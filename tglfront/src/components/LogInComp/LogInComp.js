import React from 'react';

const LogInComp = (props) => {
    return (
        <div>
            <p>I'm the log in component</p>
            <form onSubmit={props.handleLogin} method='POST'>
                <label for='u_email'>email: </label>
                <input 
                    type='text' 
                    id='u_email' 
                    placeholder='someone@organised.com'
                    pattern="[a-zA-Z0-9!#$%&amp;'*+\/=?^_`{|}~.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*"
                    required
                    ></input>
                <label for='u_pass'>password: </label>
                <input 
                    type='password' 
                    id='u_pass'
                    required
                    ></input>
                <button type='submit'>Log in</button>
            </form>
            <br />
            <p>Don't have an account? <span onClick={props.toggleSignup}>click here</span> to sign up</p>
            <p onClick={props.resetHome}>X</p>
        </div>
    );
};

export default LogInComp;