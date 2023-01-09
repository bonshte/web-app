const button = document.getElementById('menu-button');

button.addEventListener('click', (event) => {
  var menu = document.getElementById("menu-box");
  var status = window.getComputedStyle(menu).display;
  if (status == "none") {
    menu.style.display = "flex";
  } else {
    menu.style.display = "none";
  }
});

var queryForLargeScreen = "(min-width: 701px)";
var mql = window.matchMedia(queryForLargeScreen);
function checkForWideScreen() {
  var menu = document.getElementById("menu-box");
  if(mql.matches) {
    menu.style.display = "flex";
  } else {
    menu.style.display = "none";
  }
}
window.addEventListener('resize', checkForWideScreen);
