import React, { useState, useEffect } from 'react';

const TaskComp = (props) => {
    // const [tasks, setTasks] = useState([{
    //     Tid: 6,
    //     Tname: "testing no desc",
    //     Tdesc: "",
    //     Tdate: "2020-06-01T00:00:00Z",
    //     Ttime: "0000-01-01T17:00:00Z",
    //     Tcomp: false
    //     },
    //     {
    //     Tid: 8,
    //     Tname: "should have false automatically",
    //     Tdesc: "",
    //     Tdate: "2020-06-01T00:00:00Z",
    //     Ttime: "0000-01-01T17:00:00Z",
    //     Tcomp: false
    //     }]);
    
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        fetch(`http://www.localhost:8080/api/tasks/${props.user}`)
        .then(res => res.json())
        .then(data => setTasks(data));
    }, []);

    return (
        <div>
            <p>I'm the task component</p>
            <p>the user id is {props.user}</p>
            <ul>
                {tasks.map(task => (
                    <li key={task.Tid}>
                        {task.Tname}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TaskComp;