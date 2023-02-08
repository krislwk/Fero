import React from 'react';
import axios from 'axios'

function AddButton() {

    function clickHandler() {
        axios.get("http://127.0.0.1:4000/healthcheck").then(response => {console.log(response)})
    }

    return(
        <button onClick={clickHandler}>Add</button>
    )
}

export default AddButton;