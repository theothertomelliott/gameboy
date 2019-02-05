function scrollToPC() {
    var pc = document.getElementById("PC");
    var targetPos = pc.offsetTop - document.getElementsByClassName('debugspacer')[0].offsetHeight;
    if ('scrollRestoration' in window.history) {
        window.history.scrollRestoration = 'manual'
    }
    window.scrollTo(0, targetPos);
}