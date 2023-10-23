'use strict'

// JavaScript 코드 추가

const reels = [
    document.querySelectorAll('.reel0'),
    document.querySelectorAll('.reel1'),
    document.querySelectorAll('.reel2'),
    document.querySelectorAll('.reel3'),
    document.querySelectorAll('.reel4')
];
const spinButton = document.querySelector('.spin-btn');
const winAmountTxt = document.querySelector('.win-amount');
const payTable = document.querySelector('#pay-table');
const walletText = document.querySelector('.wallet span');
const spinTimeInterval = 100;
let wallet = 0;

let symbols = null;
let isSpinning = false;
let loopingLinePay = false
spinButton.addEventListener('click', () => {
    if (!isSpinning) {
        isSpinning = true;
        requestSpin();
    }
});


function InitSlotUI() {
    Network
        .Load()
        .then(res => {
            symbols = res.symbols;
            return res;
        })
        .then(initPayTable)
        .then(changeGrid)
        .then(()=>{
            return Network.Guest()
        })
        .then(guestRes=>{
            updateWallet(guestRes["cash"]);
        })
}
function requestSpin() {
    let bet = 1.0;
    let line = 50;
    loopingLinePay = false;
    updateWallet(-bet * line);
    stopLineWinAnimation();
    SpinReels()
        .then(() => {
        return Network.Spin(0, bet)
    })
        .then(setSpinResult)
        .then(changeGrid)
        .then(stopSpin)
        .then(setWinAmount)
        .then(showLineWins)
}


//region [ UI ]

function initPayTable(slotConfig){
    let payouts = slotConfig.payout;

    for (const symbolId in payouts) {
        let symbol = symbols[symbolId];
        for (let i = 0; i < payouts[symbolId].length; i++) {
            if ( payouts[symbolId][i] > 0 ){
                for (let j = 0; j < i+1; j++) {
                    payTable.innerHTML += symbol;
                }
                payTable.innerHTML += `&nbsp x ${payouts[symbolId][i]}<br>`;
            }
        }
    }

    return slotConfig["initialData"]["spin"]
}

async function SpinReels() {
    function _spinReel() {
        for (let i = 0; i < this.length; i++) {
            this[i].style.animation = 'moveSlots 0.1s linear infinite';
        }
    }
    for (let i = 0; i < reels.length; i++) {
        _spinReel.call(reels[i])
        await waitSec(spinTimeInterval)
    }
}

function changeGrid(spinOutput) {

    let baseRes = spinOutput["res"];

    let up = baseRes.up;
    let grid = baseRes.reel;
    let down = baseRes.dn;

    function changeReel(f, s, e) {
        this[0].innerHTML = symbols[f];
        this[1].innerHTML = symbols[s[0]];
        this[2].innerHTML = symbols[s[1]];
        this[3].innerHTML = symbols[s[2]];
        this[4].innerHTML = symbols[e];
    }
    for (let i = 0; i < reels.length; i++) {
        changeReel.call(reels[i], up[i], grid[i].strip, down[i]);
    }
    return spinOutput
}

async function stopSpin(spinOutput) {

    function _stopReel() {
        for (let i = 0; i < this.length; i++) {
            this[i].style.animation = 'none';
        }
    }


    for (let i = 0; i < reels.length; i++) {
        _stopReel.call(reels[i])
        await waitSec(spinTimeInterval)
    }
    isSpinning = false;

    return spinOutput;
}

function setSpinResult(res) {
    return res["spin"];
}

function setWinAmount( spinOutput ){

    let winAmount = spinOutput.res.win ?? 0;
    winAmountTxt.innerHTML = `$ ${winAmount}`;
    updateWallet(winAmount);
    return spinOutput;
}

async function showLineWins(spinOutput){
    let linePayList = spinOutput.res.lineWins;
    console.log(linePayList);
    if (!!linePayList === false ) return;
    loopingLinePay = true;
    async function linePay( winLines ){
        for (let i = 0; i < reels.length; i++) {
            if (winLines[i] !== undefined){
                let bit = ConvertToBitArray( winLines[i]);//.toString(2);
                for (let j = 0; j < bit.length; j++) {
                    if ( bit[j] === 1)
                        reels[i][j+1].style.animation = 'highlight 0.6s ease infinite';
                }
            }
        }
        await waitSec(600);
        for (let i = 0; i < reels.length; i++) {
            if (winLines[i] !== undefined){
                for (let j = 0; j < winLines[i].length; j++) {
                    reels[i][winLines[i][j]+1].style.animation = 'none';
                }
            }
        }
        await waitSec(100);
    }


    // for (let i = 0; loopingLinePay === true; i++) {
    //     if ( i === linePayList.length )
    //         i = 0;
        await linePay(linePayList[0].position);
    // }
}

function stopLineWinAnimation(){
    for (let i = 0; i < reels.length; i++) {
        for (let j = 0; j < reels[i].length; j++) {
            reels[i][j].style.animation = 'none';
            reels[i][j].style.backgroundColor= '#f1f1f1';
            reels[i][j].style.fontSize= `3rem`;
        }
    }
}

function updateWallet(amount){
    wallet += amount;
    walletText.innerHTML = `$ ${wallet.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')}`
}




//region [ UTIL ]
function waitSec(time) {
    return new Promise(
        (resolve) => setTimeout(resolve, time)
    )
}

function ConvertToBitArray(intNumber) {
    const bitArray = [];

    let bitString = intNumber.toString(2);
    for (let i =bitString.length -1; i >=0 ; i--) {
         bitArray.push( parseInt(bitString[i]));
    }
    return bitArray;
}

//endregion

InitSlotUI();