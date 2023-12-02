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
           if( resGuest["key"] ){
               header.set("session-key" , resGuest["key"]);
           }
           return resGuest
        })
}

async function _enter(slotId){
    let body = JSON.stringify({
        index: 1,
    });
    return await fetch(`/api/game/${slotId}/enter`,
        {
            method: "POST",
            headers: header,
            body:body
        })
        .then(res => res.json())
        .then(res => res["gameInfo"]);
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
    Guest : _guest,
    Enter : _enter
}