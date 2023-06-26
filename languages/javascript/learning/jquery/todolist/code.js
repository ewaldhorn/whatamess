// waits for document to load, then triggers
$(function () {
    $(function () {
        $("#add").on("click", function () {
            var val = $("input").val();
            if (val !== '') {
                var elem = $("<li></li>").text(val);
                $(elem).append("<button class='rem'>X</button>");
                $("#mylist").append(elem);
                $("input").val(""); //clear the input
                $(".rem").on("click", function() {
                    $(this).parent().remove();
                  });
            }
        });
    });

});
