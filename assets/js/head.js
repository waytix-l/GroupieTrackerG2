var head = document.getElementById('head');
var docHeight = document.documentElement.scrollHeight;
var halfDocHeight = docHeight / 10;

document.addEventListener('scroll', function () {
    if (window.scrollY >= halfDocHeight) {
        head.classList.add('fixed');
    } else {
        head.classList.remove('fixed');
    }
});