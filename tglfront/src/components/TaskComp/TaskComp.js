import React, { useState, useEffect } from 'react';
import AddTask from '../AddTask/AddTask';
import './TaskComp.css';

const TaskComp = (props) => {
    
    const [tasks, setTasks] = useState([]);
    const [taskForm, setTaskForm] = useState(false);

    useEffect(() => {
        const getTasks = () => {
            fetch(`api/tasks/${props.user}`)
            .then(res => res.json())
            .then(data => setTasks(data));
        };
        getTasks();
    }, [props.user]);

    const getTasks = () => {
        fetch(`api/tasks/${props.user}`)
        .then(res => res.json())
        .then(data => setTasks(data));
    };

    const toggleForm = () => {
        setTaskForm(!taskForm);
    };

    const compTask = (e) => {
        let taskId = parseInt(e.target.id);
        fetch('api/task/u/complete', {
            mode: 'cors',
            method: 'post',
            headers: {
                'Accept': 'application/json, text/plain, */*',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({"taskid": taskId})
        }).then((response) => {
            if (response.status === 200) {
                getTasks();
            } else {
                console.log("error: " + response.status)
            }
        })
    }

    //Delete task function
    // const delTask = (e) => {
    //     let taskId = parseInt(e.target.id);
    //     fetch('api/task/del', {
    //         mode: 'cors',
    //         method: 'post',
    //         headers: {
    //             'Accept': 'application/json, text/plain, */*',
    //             'Content-Type': 'application/json'
    //         },
    //         body: JSON.stringify({"taskid": taskId})
    //     }).then((response) => {
    //         if (response.status === 200) {
    //             getTasks();
    //         } else {
    //             console.log("error: " + response.status)
    //         }
    //     })
    // };

    const timeString = () => {
        let hr;
        let min;
        let formhr = document.getElementById('timeH').value;
        console.log(typeof(formhr));
        let formmin = document.getElementById('timeM').value;
        formhr.length === 1 ? hr = "0" + formhr : hr = formhr;
        formmin.length === 1 ? min = "0" + formmin : min = formmin;
        formhr === '' ? hr = "00" : hr = hr;
        formmin === '' ? min = "00" : min = min;
        let tstring = hr + min + "00";
        if (parseInt(tstring) === 0) {
            return 'none';
        } else {
            return tstring;
        }
    };

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
        let dstring = yy + mm + dd;
        if (dd === '' || mm === '' || yy === '' ) {
            return 'none';
        } else {
            return dstring;
        }
    };

    const descString = () => {
        let desc = document.getElementById('taskdesc').value;
        if (desc === '') {
            return 'none';
        } else {
            return desc;
        };
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        const uid = props.user;
        const tname = document.getElementById('taskname').value;
        const tdesc = descString();
        const tdate = dateString();
        const ttime = timeString();
        fetch('api/newtask', {
            mode: 'cors',
            method: 'post',
            headers: {
                'Accept': 'application/json, text/plain, */*',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "userid": uid,
                "taskname": tname,
                "taskdesc": tdesc,
                "taskdate": tdate,
                "tasktime": ttime
            })
        })
        .then((response) => {
            if (response.status === 200) {
                document.getElementById('taskForm').reset();
                setTaskForm(false);
                getTasks();
            } else {
                console.log("error: " + response.status)
            }
        });
    };

    return (
        <div id='taskDiv'>
            <h4>Welcome back, {props.uname}</h4>
            <div id='taskContainer'>
                {tasks.length > 0 ? [
                    tasks.map(task => (
                    <div key={task.Tid} className='task'>
                        <div className='taskHeader'>
                            <h4 className='taskTitle'>{task.Tname}</h4>
                            <p id={task.Tid} className='compBtn' onClick={compTask}>X</p>
                        </div>
                        <div className='taskDesc'>
                            {task.Tdesc}
                        </div>
                    </div>
                ))] : <p>No tasks set</p>
                 }
                 <p id='addTaskBtn' onClick={toggleForm}>New Task</p>
                {taskForm ? <AddTask handleSubmit={handleSubmit}/> : null}
            </div>
        </div>
    );
};

export default TaskComp;