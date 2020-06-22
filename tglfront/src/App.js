import React, { useState, useEffect } from 'react';
import {
  Switch,
  Route,
  useHistory,
  Redirect
} from 'react-router-dom';
import './App.css';
import TaskComp from './components/TaskComp/TaskComp';
import HomeComp from './components/HomeComp/HomeComp';


const App = () => {
  const [loggedIn, setloggedIn] = useState('false');
  const [uId, setuId] = useState(0);
  const [uName, setuName] = useState('');

  useEffect(() => {
    logState();
  }, [loggedIn]);

  let history = useHistory();
  // route functions
  // home direct
  const homePage = () => {
    history.push('/');
  };
  // task direct
  const taskPage = () => {
    history.push('/tasks');
  };
  // login direct
  const loginPage = () => {
    history.push('/login');
  };
  // signup direct
  // const signupPage = () => {
  //   history.push('/signup');
  // };

  const logState = () => {
    const logStatus = localStorage.getItem('loggedIn');
    const userId = localStorage.getItem('userId');
    const userName = localStorage.getItem('userName');
    setloggedIn(logStatus);
    setuId(userId);
    setuName(userName);
  };

  const logOut = () => {
    localStorage.clear();
    logState();
    homePage();
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
    fetch('api/login', {
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
        localStorage.setItem('userName', response.username);
        logState();
        taskPage();
      }
    })
  };

  return (
    <div id='App'>
      <Switch>
        <Route exact path='/' children={<HomeComp taskPage={taskPage} loginPage={loginPage} logOut={logOut} handleLogin={handleLogin} loggedIn={loggedIn} />} />
        <Route path='/tasks' children={loggedIn ? <TaskComp user={uId} uname={uName} logOut={logOut} homePage={homePage}/> : <Redirect to='/' />} />
      </Switch>
    </div>
  );
}

export default App;
