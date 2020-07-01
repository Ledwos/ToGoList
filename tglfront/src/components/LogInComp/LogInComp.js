import React from 'react';
import './LogInComp.css';

const LogInComp = (props) => {
    // const [errTxt, seterrTxt] = useState(false);
    return (
        <div id='loginDiv'>
            <form onSubmit={props.handleLogin} method='POST' id='loginForm'>
                <input 
                    type='text' 
                    id='u_email' 
                    placeholder='someone@organised.com'
                    pattern="[a-zA-Z0-9!#$%&amp;'*+\/=?^_`{|}~.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*"
                    required
                    ></input>
                <input 
                    type='password' 
                    id='u_pass'
                    placeholder='password'
                    required
                    ></input>
                {props.errTxt ? <p className='errTxt'>Please check your login details</p> : null}
                <button type='submit' id='loginBtn'>Log in</button>
            </form>
            <br />
            <br />
            <br />
            <p>Don't have an account? <span onClick={props.toggleSignup} id='clickHere'>click here</span> to sign up</p>
        </div>
    );
};

export default LogInComp;