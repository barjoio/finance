const $ = x => document.getElementById(x);
const $body = document.body;

//
//  Popout
//

const $popout = $("popout");
const $popoutOpen = $("popout_open");
const $popoutClose = $("popout_close");

$popoutOpen.onclick = () => {
    $popout.classList.add("open");
    $body.classList.add("covered");
};
$popoutClose.onclick = () => {
    $popout.classList.remove("open");
    $body.classList.remove("covered");
};