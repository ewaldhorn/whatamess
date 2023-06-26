// waits for document to load, then triggers
$(function () {
    // the JS way
    // var x = document.getElementById("demo");
    // x.onclick = function() {
    //     document.body.innerHTML = Date();
    // };

    // the jquery way
    $("#demo").click(function(){
        $("body").html(Date());
    });
});
