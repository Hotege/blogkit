function showModuleEditor(id, hId, mId, pId, value, type) {
    document.getElementById(hId).style.display = "block";
    document.getElementById(hId).style.height = document.getElementById(id).parentNode.clientHeight + "px";
    document.getElementById(id).style.display = "block";
    document.getElementById(id).style.top = (document.getElementById(id).parentNode.clientHeight - document.getElementById(id).clientHeight) / 2 + "px";
    document.getElementById('module_edit_id').value = mId;
    document.getElementById('module_edit_pid').value = pId;
    document.getElementById('module_edit_name').value = value;
    document.getElementById('module_edit_type').value = type;
    var options = document.getElementById('select_previous');
    for (i = 0; i < options.length; i++) {
        if (options[i].id == pId) {
            options[i].selected = true;
            break;
        }
    }
    document.getElementById('module_edit_name').focus();
}

function hideModuleEditor(id, hId) {
    document.getElementById(hId).style.display = "none";
    document.getElementById(id).style.display = "none";
}

function submitModuleEditor(id, hId) {
    hideModuleEditor(id, hId);
    document.getElementById("module_form").submit();
}

function deleteModule(mId) {
    var msg = confirm("Will you delete this module?");
    if (msg) {
        document.getElementById("delete_id").value = mId;
        document.getElementById("delete_form").submit();
    }
}

function deleteArticle(aId) {
    var msg = confirm("Will you delete this article?");
    if (msg) {
        document.getElementById("delete_article_id").value = aId;
        document.getElementById("delete_article_form").submit();
    }
}
