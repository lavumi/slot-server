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
const initialData = {
    us : [11, 12, 13, 14, 10],
    s :  [[12, 14, 11], [14, 13, 20], [20, 12, 12], [20, 13, 10], [21, 14, 22],],
    ds : [11, 12, 13, 14, 10]
};
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

function waitSec(time) {
    return new Promise(
        (resolve) => setTimeout(resolve, time)
    )
}
function requestSpin() {
    loopingLinePay = false;
    updateWallet(-1.0 * 25);
    stopLineWinAnimation();
    startSpin()
        .then(() => {

        return Network.Spin(0, 1.0, 25)
    })
        .then(setSpinResult)
        .then(changeGrid)
        .then(stopSpin)
        .then(setWinAmount)
        .then(showLineWins)
}
//region [ UI ]
async function startSpin() {
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
    let up = spinOutput.us;
    let grid = spinOutput.s;
    let down = spinOutput.ds;

    function changeReel(f, s, e) {
        this[0].innerHTML = symbols[f];
        this[1].innerHTML = symbols[s[0]];
        this[2].innerHTML = symbols[s[1]];
        this[3].innerHTML = symbols[s[2]];
        this[4].innerHTML = symbols[e];
    }
    for (let i = 0; i < reels.length; i++) {
        changeReel.call(reels[i], up[i], grid[i], down[i]);
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
    return res.SpinOutput;
}

function setWinAmount( spinOutput ){
    winAmountTxt.innerHTML = `$ ${spinOutput.tw}`;

    updateWallet(spinOutput.tw);
    return spinOutput;
}

async function showLineWins(spinOutput){


    let linePayList = spinOutput.lp;
    if (linePayList === null )return;
    loopingLinePay = true;

    async function linePay( winLines ){

        for (let i = 0; i < reels.length; i++) {
            if (winLines[i] !== undefined){
                for (let j = 0; j < winLines[i].length; j++) {
                    reels[i][winLines[i][j]+1].style.animation = 'highlight 0.6s ease infinite';
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
    for (let i = 0; loopingLinePay === true; i++) {
        if ( i === linePayList.length )
            i = 0;
        await linePay(linePayList[i].pos);
    }
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
    return initialData
}

//region

//



function InitSlotUI() {
    updateWallet(1000000);
    Network
        .Load()
        .then(res => {
            symbols = res.symbols;
            return res;
        })
        .then(initPayTable)
        .then(changeGrid)
}
InitSlotUI();
