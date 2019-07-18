function showReplyComment(id, hId, rId, type) {
    var obj = document.getElementById(id);
    document.getElementById(hId).style.display = "block";
    document.getElementById(hId).style.height = document.body.clientHeight + "px";
    obj.style.display = "block";
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
