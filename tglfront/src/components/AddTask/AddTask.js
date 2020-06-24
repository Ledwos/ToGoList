import React from 'react';
import './AddTask.css';

const AddTask = (props) => {
    let year = new Date().getFullYear();
    return (
        <div id='addTaskDiv'>
            <form onSubmit={props.handleSubmit} method='POST' id='taskForm'>
                <label for='taskname' className='txtLabels'>Task Name:</label>
                <input type='text' id='taskname' placeholder='Clean garage' required></input>
                <label for='taskdesc' className='txtLabels'>Description:</label>
                <textarea id='taskdesc' placeholder='(optional)'></textarea>
                <label for='taskdate' className='txtLabels'>Date:</label>
                <div id='taskdate'>
                    <input type='number' min="1" max="31" placeholder="DD" id="dateD"></input>/
                    <input type='number' min="1" max="12" placeholder="MM" id="dateM"></input>/
                    <input type='number' min={year} max="2099" placeholder="YYYY" id="dateY"></input>
                </div>
                <label for='tasktime' className='txtLabels'>Time:</label>
                <div id='tasktime'>
                    <input type='number' min="0" max="23" placeholder="HH" id="timeH"></input>: 
                    <input type='number' min="0" max="59" placeholder="MM" id="timeM"></input>
                </div>
                <button type='submit' id='addBtn'>Add!</button>
            </form>
        </div>
    );
};

export default AddTask;