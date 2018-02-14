const mq = window.matchMedia("(min-width: 601px)");

function openNav() {
  if (mq.matches) {
    document.getElementById('sidenav').style.width = "25%";
  } else {
    document.getElementById('sidenav').style.width = "50%";
  }
}

function closeNav() {
  document.getElementById('sidenav').style.width = "0%";
}