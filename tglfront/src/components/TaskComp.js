import React, { useState, useEffect } from 'react';
import AddTask from './AddTask';

const TaskComp = (props) => {
    
    const [tasks, setTasks] = useState([]);
    const [taskForm, setTaskForm] = useState(false);

    useEffect(() => {
        getTasks();
    }, []);

    const getTasks = () => {
        fetch(`http://www.localhost:8080/api/tasks/${props.user}`)
        .then(res => res.json())
        .then(data => setTasks(data));
    }

    const toggleForm = () => {
        setTaskForm(!taskForm);
    }

    const delTask = (e) => {
        let taskId = e.target.id;
        console.log("I delete you! number: " + taskId)
        // fetch('http://www.localhost:8080/api/task/del', {
        //     method: 'DELETE',
        //     body: JSON.stringify({taskid: tId})
        // })
    }

    const timeString = () => {
        let hr;
        let min;
        let formhr = document.getElementById('timeH').value;
        let formmin = document.getElementById('timeM').value;
        formhr.length === 1 ? hr = "0" + formhr : hr = formhr;
        formmin.length === 1 ? min = "0" + formmin : min = formmin;
        return hr + min + "00";
    }

    const dateString = () => {
        let dd;
        let mm;
        let yy;
        let formdd = document.getElementById('dateD').value;
        let formmm = document.getElementById('dateM').value;
        let formyy = document.getElementById('dateY').value;
        formdd.length === 1 ? dd = "0" + formdd : dd = formdd;
        formmm.length === 1 ? mm = "0" + formmm : mm = formmm;
        formyy.length === 1 ? yy = "0" + formyy : yy = formyy;
        return yy + mm + dd;
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        const uid = props.user;
        const tname = document.getElementById('taskname').value;
        const tdesc = document.getElementById('taskdesc').value;
        const tdate = dateString();
        const ttime = timeString();
        console.dir({
            "uid": uid,
            "tName": tname,
            "tDesc": tdesc,
            "tdate": tdate,
            "tTime": ttime
        });
    }

    return (
        <div>
            <p>I'm the task component</p>
            <p>the user id is {props.user}</p>
            <div id='addTaskBtn' onClick={toggleForm}>New Task</div>
            {tasks.map(task => (
                <div key={task.Tid}>
                    <div className='taskHeader'>
                        <h4>{task.Tname}</h4>
                        <p id={task.Tid} onClick={delTask}>X</p>
                    </div>
                    <div className='taskDesc'>
                        {task.Tdesc}
                    </div>
                </div>
            ))}
            {taskForm ? <AddTask handleSubmit={handleSubmit}/> : null}
        </div>
    );
};

export default TaskComp;