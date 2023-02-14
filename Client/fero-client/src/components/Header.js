import React from 'react';

function Header({prop}) {
    return(
        <h1 className = "text-5xl text-headerColor">
            {prop}'s Homework
        </h1>
    );
}

export default Header;