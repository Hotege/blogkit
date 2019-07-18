function addStep() {
    id = document.getElementById('steps_last_id').value;
    var node = document.getElementById('steps').cloneNode(true);
    node.id = node.id + id;
    node.style.display = 'block';
    var select = node.getElementsByTagName("select")[0];
    select.id = 'step_s' + id;
    select.onchange = function() {
        var oId = this.options[this.options.selectedIndex].id;
        var sId = select.id.split('step_s')[1];
        var ch = oId[6];
        document.getElementById('step_tt' + sId).style.display = 'none';
        document.getElementById('step_ii' + sId).style.display = 'none';
        document.getElementById('step_ff' + sId).style.display = 'none';
        document.getElementById('step_cc' + sId).style.display = 'none';
        document.getElementById('step_' + ch + ch + sId).style.display = 'block';
    };
    var divs = node.getElementsByTagName("div");
    for (var i = 0; i < divs.length; i++) {
        if (typeof(divs[i].id) != "undefined") {
            divs[i].id = divs[i].id + id;
        }
        if (divs[i].id == 'step_tt' + id) {
            divs[i].style.display = 'block';
        }
        if (divs[i].id == 'step_ii' + id) {
            divs[i].style.display = 'none';
            var ifs = divs[i].getElementsByTagName("iframe")[0];
            ifs.id = ifs.name + id;
            ifs.name = ifs.name + id;
            var iform = divs[i].getElementsByTagName("form")[0];
            iform.id = 'step_i_form' + id;
            iform.target = ifs.name;
            var ia = divs[i].getElementsByTagName("a")[0];
            ia.href = 'javascript:submitFile("step_i_form' + id + '", "step_i_f' + id + '", "step_i_name' + id + '");';
            var iname = divs[i].getElementsByTagName("input")[1];
            iname.id = iname.id + id;
        }
        if (divs[i].id == 'step_ff' + id) {
            divs[i].style.display = 'none';
            var cfs = divs[i].getElementsByTagName("iframe")[0];
            cfs.id = cfs.name + id;
            cfs.name = cfs.name + id;
            var cform = divs[i].getElementsByTagName("form")[0];
            cform.id = 'step_f_form' + id;
            cform.target = cfs.name;
            var ca = divs[i].getElementsByTagName("a")[0];
            ca.href = 'javascript:submitFile("step_f_form' + id + '", "step_f_f' + id + '", "step_f_name' + id + '");';
            var fname = divs[i].getElementsByTagName("input")[1];
            fname.id = fname.id + id;
        }
        if (divs[i].id == 'step_cc' + id) {
            divs[i].style.display = 'none';
        }
    }
    var aa = node.getElementsByTagName("a");
    for (var i = 0; i < aa.length; i++) {
        if (aa[i].id == 'step_r') {
            aa[i].id = aa[i].id + id;
            aa[i].href = 'javascript:removeStep("' + node.id + '")';
        }
    }
    document.getElementById('steps_group').appendChild(node);
    document.getElementById('steps_last_id').value = parseInt(id) + 1;
}
function removeStep(id) {
    document.getElementById('steps_group').removeChild(document.getElementById(id));
}
function submitFile(id, fId, nId) {
    document.getElementById(id).submit();
    document.getElementById(fId).onload = function() {
        var inputs = document.getElementById(fId).contentWindow.document.getElementsByTagName("input");
        var s = new Array(inputs.length);
        for (var i = 0; i < inputs.length; i++) {
            var layers = inputs[i].value.split('/');
            s[i] = layers[layers.length - 1];
        }
        document.getElementById(fId).parentNode.getElementsByTagName("input")[1].value = s.join('/');
    };
}
