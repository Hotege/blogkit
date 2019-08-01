function showReplyComment(id, hId, rId, type) {
    document.getElementById(hId).style.display = "block";
    document.getElementById(hId).style.height = document.getElementById(id).parentNode.clientHeight;
    document.getElementById(id).style.display = "block";
    document.getElementById(id).style.top = (document.getElementById(id).parentNode.clientHeight - document.getElementById(id).clientHeight) / 2 + "px";
    document.getElementById('reply_id').value = rId;
    document.getElementById('reply_type').value = type;
    document.getElementById('reply_comment_content').focus();
}

function hideReplyComment(id, hId) {
    document.getElementById(hId).style.display = "none";
    document.getElementById(id).style.display = "none";
}

function submitReplyComment(id, hId) {
    hideReplyComment(id, hId);
    document.getElementById("reply_form").submit();
}

function deleteComment(id) {
    var msg = confirm("Will you delete this comment?");
    if (msg) {
        document.getElementById("delete_id").value = id;
        document.getElementById("delete_form").submit();
    }
}
