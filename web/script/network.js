'use strict'


let header = new Headers();

async function _guest() {
    return await fetch(
        `/api/auth/guest`,
        {
            method: "POST",
        })
        .then(res => res.json())
        .then(resGuest =>{
            console.log(resGuest);
           if( resGuest["key"] ){
               header.set("session-key" , resGuest["key"]);
           }
           return resGuest
        })
}



async function _loadSheet(){
    return await fetch(`/config/0.json`)
        .then(res => res.json())
}

async function _spin(slotId, bet) {

    let body = JSON.stringify({
        index: 1,
        counter: 1,
        bet: bet
    });
    return await fetch(
        `/api/game/${slotId}/spin`,
        {
            method: "POST",
            headers: header,
            body:body
        })
        .then(res => res.json())
}


let Network = {
    Spin : _spin,
    Load : _loadSheet,
    Guest : _guest
}