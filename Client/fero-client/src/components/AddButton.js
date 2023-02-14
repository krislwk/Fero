import React from 'react';
import axios from 'axios'

function AddButton() {

    function clickHandler() {
        axios.get("http://127.0.0.1:4000/healthcheck").then(response => {console.log(response)})
    }

    return(
        <div class = "grid place-items-center h-screen flex flex-col items-center justify-center">
            <button
                className="text-white font-semibold w-24 h-9 border-2 border-[#447dbd] rounded-md bg-[#447dbd]"
                onClick={clickHandler}>
                Add Task
            </button>
        </div>
    )
}

export default AddButton;