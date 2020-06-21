import React, { useState, useEffect } from 'react';

const SignUpComp = (props) => {
    const [pass1, setpass1] = useState('i');
    const [pass2, setpass2] = useState('o');

    useEffect(() => {
        checkmatch();
    }, [pass1, pass2]);

    const checkmatch = () => {
        if (pass1 === pass2) {
            return ':)';
        } else {
            return ':(';
        };
    };
    
    const checknum = () => {
        if (/\d/.test(pass1)) {
            return ':)';
        } else {
            return ':(';
        };
    };

    const setp1 = (e) => {
        setpass1(e.target.value);
    };

    const setp2 = (e) => {
        setpass2(e.target.value);
    };

    const checksubmit = (e) => {
        e.preventDefault();
        const name = document.getElementById('username').value;
        const email = document.getElementById('email').value;
        const pass = document.getElementById('pass').value;
        fetch('http://www.localhost:8080/api/newacc', {
            mode: 'cors',
            method: 'post',
            headers: {
                'Accept': 'application/json, text/plain, */*',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "name": name,
                "pass": pass,
                "email": email
            })
        })
        .then((response) => {
            if (response.status === 200) {
                console.log('congrats');
                document.getElementById('signupForm').reset();
                props.resetHome();
            } else {
                console.log('error: ' + response.status);
            }
        });
    }

    return (
        <div>
            <p>I'm the sign up component</p>
            <form id='signupForm' onSubmit={checksubmit} method='POST'>
                <label for='username'>Name:</label>
                <input type='text' id='username' placeholder='Jessica' required></input>
                <label for='email'>Email:</label>
                <input type='text' 
                        id='email' 
                        placeholder='someone@organised.com'
                        pattern="[a-zA-Z0-9!#$%&amp;'*+\/=?^_`{|}~.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*"
                        required></input>
                <label for='pass'>Password:</label>
                <input minLength='8' type='password' id='pass' placeholder='password' required onChange={setp1}></input>
                <label for='rpass'>Re-enter Password:</label>
                <input type='password' minLength='8' id='rpass' placeholder='password' required onChange={setp2}></input>
                <button type='submit'>Sign Up</button>
            </form>
            <p>passwords match: {checkmatch()}</p>
            <p>password contains number(s): {checknum()}</p>
            <br />
            <p onClick={props.toggleLogin}>back to log in</p>
            <p onClick={props.resetHome}>X</p>
        </div>
    );
};

export default SignUpComp;