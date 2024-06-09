var head = document.getElementById('head');
var musicLogo = document.getElementById('musicImage');
var headerText = document.getElementsByClassName('headerText');
var headerTextAmount = headerText.length;
var musicButton = document.getElementById('musicButton');
var imgListenMusic = document.getElementById('imgListenMusic');
var titleListenMusic = document.getElementById('titleListenMusic');
var musicTextBackground = document.getElementById('musicTextBackground');

var filterButton = document.getElementById('filterButton');
var filters = document.getElementsByClassName('filter')[0];




document.addEventListener('scroll', function () {
    if (window.scrollY >= 5) {
        head.classList.add('fixed');
        headerText.item(0).classList.add('textHeaderFixed')
        for (let index = 0; index < headerTextAmount; index++) {
            headerText.item(index).classList.add('textHeaderFixed')
        }
    } else {
        head.classList.remove('fixed');
        for (let index = 0; index < headerTextAmount; index++) {
            headerText.item(index).classList.remove('textHeaderFixed')
        }
    }
});

function expandButton() {

    musicButton.id = 'musicButtonExpanded';
    musicLogo.id = 'musicImage2'
    setTimeout(() => { 
        musicButton.id = 'musicButtonExpanded2';
        musicButton.disabled = true;
    }, 200);
    setTimeout(() => {
        imgListenMusic.id = 'imgListenMusic2';
        titleListenMusic.id = 'titleListenMusicAfter';
    }, 500);
}

function expandTextBackground() {
    console.log("hover");
    musicTextBackground.id = 'musicTextBackgroundAfter';
}

function retractTextBackground() {
    musicTextBackground.id = 'musicTextBackground';
}

musicButton.onclick = expandButton;

musicButton.addEventListener('mouseover', expandTextBackground);
musicButton.addEventListener('mouseout', retractTextBackground);

filterButton.onclick = filterBand;

function filterBand() {
    filters.id = "filterAfter";
}



/* See More Button */


var seeMoreButton = document.getElementsByClassName('cardButton');
var popUps = document.getElementsByClassName('popUP');
var closePopUpButton = document.getElementsByClassName('closePopUp');

for (let index = 0; index < seeMoreButton.length; index++) {
    seeMoreButton[index].addEventListener('click', function () {
        console.log('clicked');
        seeMorePopUp(index);
    });
}

for (let index = 0; index < closePopUpButton.length; index++) {
    closePopUpButton[index].addEventListener('click', function () {
        console.log('clicked');
        closePopUp(index);
    })
}

function seeMorePopUp(index) {
    popUps[index].style.display = 'block';
}

function closePopUp(index) {
    popUps[index].style.display = 'none';
}