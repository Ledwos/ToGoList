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
import NotFound from './components/NotFound/NotFound';


const App = () => {
  const [loggedIn, setloggedIn] = useState('false');
  const [uId, setuId] = useState(0);
  const [uName, setuName] = useState('');
  const [loginState, setloginState] = useState(false);
  const [signupState, setsignupState] = useState(false);
  const [errTxt, seterrTxt] = useState(false);

  useEffect(() => {
    logState();
  }, [loggedIn]);

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
  };

  let history = useHistory();
  // route functions
  // home direct
  const homePage = () => {
    resetHome();
    history.push('/');
  };
  // task direct
  const taskPage = () => {
    history.push('/tasks');
  };
  // login direct
  // const loginPage = () => {
  //   history.push('/login');
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
        resetHome();
      } else if (resStatus === 400) {
        seterrTxt(true);
      }
    })
  };

  return (
    <div id='App'>
      <nav>
          <h3 id='title' onClick={homePage}>ToGoList</h3> 
          {loggedIn ? [
            <ul className='navMain' key='x'>
            <li className='navItem' key='1' onClick={taskPage}>Tasks</li>
            <li className='navItem' key='2' onClick={logOut}>Log Out</li>
            </ul>
          ] : [
                <ul className='navMain' key='y'>
                  <li className='navItem' key='4' onClick={toggleSignup}>Sign Up</li>
                  <li className='navItem' key='3' onClick={toggleLogin}>Log In</li>
                </ul>
              ]}
        </nav>
      <Switch>
        <Route exact path='/' children={<HomeComp handleLogin={handleLogin} loggedIn={loggedIn} resetHome={resetHome} loginState={loginState} signupState={signupState} toggleSignup={toggleSignup} toggleLogin={toggleLogin} errTxt={errTxt} seterrTxt={seterrTxt}/> } />
        <Route path='/tasks' children={loggedIn ? <TaskComp user={uId} uname={uName} logOut={logOut} homePage={homePage}/> : <Redirect to='/' />} />
        <Route children={<NotFound homePage={homePage}/>} />
      </Switch>
    </div>
  );
}

export default App;
