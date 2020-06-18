import React, { useState, useEffect } from 'react';
import './App.css';
import TaskComp from './components/TaskComp';
import LogInComp from './components/LogInComp';


const App = () => {
  const [loggedIn, setloggedIn] = useState(true);
  const [uId, setuId] = useState(5);

  if (loggedIn) {
    return <TaskComp user={uId}/>
  } else {
    return <LogInComp />
  }
}

export default App;
