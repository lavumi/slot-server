'use strict'


let header = new Headers({
    'Authorization': 'Bearer ' + "mytoken",
    'uid': 123
});

function _login(userid) {
    header = new Headers({
        'Authorization': 'Bearer ' + "mytoken",
        'uid': userid
    });
}

async function _loadSheet(){
    return await fetch(`/config/0.json`)
        .then(res => res.json())
}

async function _spin(slotId, bet, line) {

    let body = JSON.stringify({
        index: 1,
        counter: 1,
        bet: bet,
        line: line
    });
    console.log(body);
    return await fetch(
        `/api/game/${slotId}/spin`,
        {
            method: "POST",
            // headers: header,
            body:body
        })
        .then(res => res.json())
}




let Network = {
    Spin : _spin,
    Load : _loadSheet
}