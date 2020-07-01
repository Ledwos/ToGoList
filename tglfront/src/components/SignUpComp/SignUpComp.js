import React, { useState, useEffect } from 'react';
import './SignUpComp.css';

const SignUpComp = (props) => {
    const [pass1, setpass1] = useState('i');
    const [pass2, setpass2] = useState('o');
    const [emErr, setemErr] = useState(false);

    useEffect(() => {
        const checkmatch = () => {
            if (pass1 === pass2) {
                return '😀';
            } else {
                return '🙁';
            };
        };
        checkmatch();
    }, [pass1, pass2]);

    const checkempty = () => {
        if (pass1 === '') {
            setpass1('i');
        }
        if (pass2 === '') {
            setpass2('o');
        }
    }

    const checkmatch = () => {
        checkempty();
        if (pass1 === pass2) {
            return '😀';
        } else {
            return '🙁';
        };
    };
    
    const checknum = () => {
        if (/\d/.test(pass1) | /\d/.test(pass2)) {
            return '😀';
        } else {
            return '🙁';
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
        fetch('api/newacc', {
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
            if (response.status === 409) {
                console.log('email already exists');
                setemErr(true);
            }
            else if (response.status === 200) {
                console.log(response);
                document.getElementById('signupForm').reset();
                props.resetHome();
            } else {
                console.log('error: ' + response.status);
            }
        });
    }

    return (
        <div id='signupDiv'>
            <p>Fill in the fields below</p>
            <form id='signupForm' onSubmit={checksubmit} method='POST'>
                <input type='text' id='username' className='signupItem' placeholder='First Name' required></input>
                <input type='text' 
                        id='email'
                        className='signupItem' 
                        placeholder='Email'
                        pattern="[a-zA-Z0-9!#$%&amp;'*+\/=?^_`{|}~.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*"
                        required></input>
                {emErr ? <p className='errTxt'>email already exists</p> : null}
                <input minLength='8' 
                       type='password' 
                       id='pass'
                       className='signupItem' 
                       placeholder='Password' 
                       required 
                       onChange={setp1}></input>
                <input type='password' 
                       minLength='8' 
                       id='rpass' 
                       className='signupItem'
                       placeholder='Confirm Password' 
                       required 
                       onChange={setp2}></input>
                <button type='submit' id='signupBtn'>Sign Up</button>
                <table id='passTable'>
                    <tbody>
                        <tr>
                            <td className='passTxt'>Passwords match :</td>
                            <td className='passVal'>{checkmatch()}</td>
                        </tr>
                        <tr>
                            <td className='passTxt'>Contain number(s) :</td>
                            <td className='passVal'>{checknum()}</td>
                        </tr>
                    </tbody> 
                </table>
            </form>
            <br />
            <p onClick={props.toggleLogin} id='clickLogin'>back to log in</p>
        </div>
    );
};

export default SignUpComp;