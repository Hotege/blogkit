function showLogin(id, hId, uId, pId) {
    document.getElementById(hId).style.display = "block";
    document.getElementById(hId).style.height = document.body.clientHeight + "px";
    document.getElementById(id).style.display = "block";
    document.getElementById(id).style.top = (document.body.clientHeight - document.getElementById(id).clientHeight) / 2 + "px";
    document.getElementById(uId).value = "";
    document.getElementById(pId).value = "";
    document.getElementById(uId).focus();
}

function hideLogin(id, hId) {
    document.getElementById(hId).style.display = "none";
    document.getElementById(id).style.display = "none";
}
