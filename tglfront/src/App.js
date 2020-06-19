import React, { useState, useEffect } from 'react';
import './App.css';
import TaskComp from './components/TaskComp';
import LogInComp from './components/LogInComp';


const App = () => {
  const [loggedIn, setloggedIn] = useState('false');
  const [uId, setuId] = useState(0);

  useEffect(() => {
    logState();
  }, [loggedIn])

  const logState = () => {
    const logStatus = localStorage.getItem('loggedIn');
    const userId = localStorage.getItem('userId');
    setloggedIn(logStatus);
    setuId(userId);
  };

  const logOut = () => {
    localStorage.setItem('loggedIn', 'false');
    localStorage.setItem('userid', null);
    logState();
  }

  const handleLogin = (e) => {
    e.preventDefault();
    const email = document.getElementById('u_email').value;
    const pass = document.getElementById('u_pass').value;
    console.dir({
      "email": email,
      "e-type": typeof(email),
      "pass": pass,
      "p-type": typeof(pass),
    });
    let resStatus;
    fetch('http://www.localhost:8080/api/login', {
            mode: 'cors',
            method: 'post',
            headers: {
                'Accept': 'application/json, text/plain, */*',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "email": email,
                "pass": pass,
            })
    })
    .then(response => {
      resStatus = response.status;
      return response.json();
    })
    .then(response => {
      if (resStatus === 200) {
        console.log(response);
        localStorage.setItem('loggedIn', 'true');
        localStorage.setItem('userId', response.userid);
        logState();
      }
    })
  }

  if (loggedIn === 'true') {
    return <TaskComp user={uId} logOut={logOut}/>
  } else {
    return <LogInComp handleLogin={handleLogin}/>
  }
}

export default App;
