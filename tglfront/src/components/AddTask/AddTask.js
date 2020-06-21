import React from 'react';

const AddTask = (props) => {
    let year = new Date().getFullYear();
    return (
        <div>
            <form onSubmit={props.handleSubmit} method='POST' id='taskForm'>
                <label for='taskname'>Title:</label>
                <input type='text' id='taskname' required></input>
                <label for='taskdesc'>Description:</label>
                <textarea id='taskdesc' placeholder='(optional)'></textarea>
                <label for='taskdate'>Date:</label>
                <div id='taskdate'>
                    <label for='dateD'>DD:</label>
                    <input type='number' min="1" max="31" placeholder="01" id="dateD"></input> -
                    <label for='dateM'>MM:</label> 
                    <input type='number' min="1" max="12" placeholder="01" id="dateM"></input> -
                    <label for='dateD'>YY:</label>
                    <input type='number' min={year} max="2099" placeholder={year} id="dateY"></input>
                </div>
                <label for='tasktime'>Time:</label>
                <div id='tasktime'>
                    <input type='number' min="0" max="23" placeholder="00" id="timeH"></input> :
                    <input type='number' min="0" max="59" placeholder="00" id="timeM"></input>
                </div>
                <button type='submit'>Add!</button>
            </form>
        </div>
    );
};

export default AddTask;