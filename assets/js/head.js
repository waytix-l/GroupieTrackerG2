var head = document.getElementById('head');
var musicLogo = document.getElementById('musicImage');
var headerText = document.getElementsByClassName('headerText')
var headerTextAmount = headerText.length;
var musicButton = document.getElementById('musicButton');
var imgListenMusic = document.getElementById('imgListenMusic')

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
    }, 500);
}

musicButton.onclick = expandButton;
