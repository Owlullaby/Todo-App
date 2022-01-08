window.onload = function() {
    var lists = document.getElementById("lists");

    lists.addEventListener("click", function(e){
        if (e.target && e.target.matches("div > label")){
            if (e.target.parentElement.querySelector("div > label").style.display  === "block"){
                e.target.parentElement.querySelector("div > label").style.display  = "none";
                e.target.parentElement.querySelector("form > input").style.display = "inline-block";
                e.target.parentElement.querySelector("form > button.bg-green-300").style.display = "inline-block";
                e.target.parentElement.querySelector("form > button.bg-red-300").style.display = "inline-block";
                e.target.parentElement.querySelector("form > input").value = e.target.parentElement.querySelector("div > label").innerHTML;
            }
            
        }
    });

    lists.addEventListener("click", function(e){
        if (e.target && e.target.matches("form > button.bg-red-300")){
            e.target.parentElement.parentElement.querySelector("div > label").style.display = "block";
            e.target.parentElement.querySelector("form > input").style.display = "none";
            e.target.parentElement.querySelector("form > button.bg-green-300").style.display = "none";
            e.target.parentElement.querySelector("form > button.bg-red-300").style.display = "none";
            }
        });

}